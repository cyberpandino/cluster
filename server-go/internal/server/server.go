/*
 * PandaOS
 * Copyright (C) 2025  Cyberpandino
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3.
 */

// Package server orchestrates all PandaOS services: OBD-II, GPIO, sensors, and WebSocket hub.
package server

import (
	"context"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/cyberpandino/cluster/server/internal/gpio"
	"github.com/cyberpandino/cluster/server/internal/hub"
	"github.com/cyberpandino/cluster/server/internal/obd"
	"github.com/cyberpandino/cluster/server/internal/sensors"
)

const (
	maxReconnectAttempts = 20
	reconnectDelay       = 5 * time.Second
	periodicScanInterval = 30 * time.Second
	serverAddr           = ":3001"
)

// Server is the top-level orchestrator. Create one with New, then call Start.
type Server struct {
	hub_     *hub.Hub
	comm     *obd.Comm
	monitor  *obd.Monitor
	gpio_    *gpio.Service
	ignition *gpio.IgnitionService
	temp     *sensors.TemperatureService
	fuel     *sensors.FuelService

	forceRestartCh chan struct{}
}

// New allocates a Server. No I/O happens until Start is called.
func New() *Server {
	return &Server{forceRestartCh: make(chan struct{}, 1)}
}

// Start runs all services, blocking until ctx is cancelled.
// Returns a non-nil error only for unrecoverable failures.
func (s *Server) Start(ctx context.Context) error {
	checkPlatform()

	// Hub (Socket.IO) must come first — everything else writes to it.
	h, err := hub.New(serverAddr, func() {
		select {
		case s.forceRestartCh <- struct{}{}:
		default:
		}
	})
	if err != nil {
		return err
	}
	s.hub_ = h

	// GPIO, ignition, sensors run independently of OBD.
	s.gpio_ = gpio.NewService(h)
	s.ignition = gpio.NewIgnitionService(h)
	s.temp = sensors.NewTemperatureService(h)
	s.fuel = sensors.NewFuelService(h)

	go s.gpio_.StartPolling(ctx)
	go s.temp.StartReading(ctx)
	go s.fuel.StartReading(ctx)

	log.Println("Avvio servizio OBD-II per Fiat Panda 141")

	// OBD connection + monitoring retry loop.
	return s.runOBDLoop(ctx)
}

// runOBDLoop connects to OBD, monitors, and retries on failure.
// Mirrors the startWithRetry + startMonitoring logic from OBDServer.js.
func (s *Server) runOBDLoop(ctx context.Context) error {
	attempts := 0

	for {
		// Check for force-restart signal from client or unrecoverable error.
		select {
		case <-s.forceRestartCh:
			log.Println("Riavvio forzato richiesto — uscita")
			s.stop()
			os.Exit(1)
		case <-ctx.Done():
			s.stop()
			return nil
		default:
		}

		attempts++
		msg := "Tentativo di connessione al connettore OBD..."
		if attempts > 1 {
			msg = "Tentativo di connessione #" + itoa(attempts) + "..."
		}
		log.Println(msg)
		s.hub_.EmitStatus(hub.StatusData{
			Connected:         false,
			Message:           msg,
			ReconnectAttempts: attempts,
		})

		s.comm = obd.NewComm()
		if err := s.comm.Connect(ctx); err != nil {
			log.Printf("Connessione fallita (tentativo %d): %v", attempts, err)
			s.hub_.EmitStatus(hub.StatusData{
				Connected: false,
				Message:   "Connessione fallita (tentativo " + itoa(attempts) + "). Riavvio tra 5 secondi...",
				Error:     err.Error(),
			})

			if attempts >= maxReconnectAttempts {
				log.Println("Troppi tentativi falliti — uscita")
				s.hub_.EmitStatus(hub.StatusData{Message: "Riavvio completo in corso...", Restarting: true})
				s.stop()
				os.Exit(1)
			}

			select {
			case <-ctx.Done():
				s.stop()
				return nil
			case <-time.After(reconnectDelay):
				continue
			}
		}

		log.Println("Connessione OBD stabilita con successo")
		attempts = 0
		s.hub_.EmitStatus(hub.StatusData{Connected: true, Message: "ELM327 connesso"})

		if err := s.comm.Initialize(ctx); err != nil {
			log.Printf("Errore inizializzazione ELM327: %v", err)
			s.comm.Disconnect()
			continue
		}

		workingPIDs := s.scanAllPIDs(ctx)

		if len(workingPIDs) > 0 {
			// Kick off periodic re-scan in background.
			scanCtx, cancelScan := context.WithCancel(ctx)
			go s.periodicScan(scanCtx, workingPIDs)

			// StartMonitoring blocks until OBD disconnects or ctx is cancelled.
			s.monitor = obd.NewMonitor(s.comm, s.hub_)
			s.monitor.StartMonitoring(ctx, workingPIDs)
			cancelScan()
		}

		s.comm.Disconnect()

		select {
		case <-ctx.Done():
			s.stop()
			return nil
		case <-time.After(2 * time.Second):
		}
	}
}

// scanAllPIDs probes every known PID and returns the ones that responded.
// Mirrors OBDServer.scanAllPIDs.
func (s *Server) scanAllPIDs(ctx context.Context) []obd.PIDDef {
	defs := obd.Definitions()
	log.Println("Scansione PID iniziale avviata")
	s.hub_.EmitStatus(hub.StatusData{Message: "Scansione PID in corso..."})

	data := hub.OBDData{
		Timestamp: time.Now().UTC().Format(time.RFC3339Nano),
		Vehicle: map[string]string{
			"brand": "Fiat",
			"model": "Panda 141",
			"ecu":   "Magneti Marelli IAW 4AF",
		},
		Parameters: make(map[string]hub.PIDValue, len(defs)),
	}

	var working []obd.PIDDef
	for _, def := range defs {
		result := s.comm.ReadPID(ctx, def)
		data.Parameters[def.Key] = hub.PIDValue{
			Name:      result.Name,
			Value:     result.Value,
			Unit:      result.Unit,
			Success:   result.Success,
			Error:     result.Error,
			Timestamp: result.Timestamp,
		}
		if result.Success {
			log.Printf("PID funzionante: %s: %v %s", result.Name, result.Value, result.Unit)
			working = append(working, def)
		} else {
			log.Printf("PID non funzionante: %s", result.Name)
		}

		select {
		case <-ctx.Done():
			return working
		case <-time.After(500 * time.Millisecond):
		}
	}

	s.hub_.EmitOBDData(data)
	s.hub_.EmitStatus(hub.StatusData{
		Message:     "Scansione completata: " + itoa(len(working)) + "/" + itoa(len(defs)) + " PID funzionanti",
		WorkingPIDs: len(working),
		TotalPIDs:   len(defs),
	})
	log.Printf("%d/%d PID funzionanti", len(working), len(defs))
	return working
}

// periodicScan re-probes all PIDs every 30 s and updates the monitor's list.
// Mirrors OBDServer.startPeriodicPIDScan / periodicPIDScan.
func (s *Server) periodicScan(ctx context.Context, initial []obd.PIDDef) {
	ticker := time.NewTicker(periodicScanInterval)
	defer ticker.Stop()
	log.Println("Avvio scansione periodica PID (ogni 30 secondi)")

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if !s.comm.IsConnected() || (s.monitor != nil && !s.monitor.IsActive()) {
				continue
			}
			log.Println("Scansione periodica PID in corso")
			s.hub_.EmitStatus(hub.StatusData{Message: "Scansione periodica PID..."})

			defs := obd.Definitions()
			var found []obd.PIDDef
			newCount := 0
			for _, def := range defs {
				result := s.comm.ReadPID(ctx, def)
				if result.Success {
					if s.monitor != nil && !s.monitor.IsPIDMonitored(def.Key) {
						newCount++
						log.Printf("Nuovo PID disponibile: %s: %v %s", result.Name, result.Value, result.Unit)
					}
					found = append(found, def)
				}
				select {
				case <-ctx.Done():
					return
				case <-time.After(200 * time.Millisecond):
				}
			}

			if s.monitor != nil {
				s.monitor.UpdateWorkingPIDs(found)
			}

			msg := "Scansione periodica: " + itoa(len(found)) + "/" + itoa(len(defs)) + " PID funzionanti"
			if newCount > 0 {
				msg = "Scansione periodica: " + itoa(newCount) + " nuovi PID trovati (" + itoa(len(found)) + "/" + itoa(len(defs)) + " totali)"
			}
			s.hub_.EmitStatus(hub.StatusData{
				Message:     msg,
				WorkingPIDs: len(found),
				TotalPIDs:   len(defs),
			})
			log.Println(msg)
		}
	}
}

func (s *Server) stop() {
	if s.monitor != nil {
		s.monitor.Stop()
	}
	if s.comm != nil {
		s.comm.Disconnect()
	}
	if s.gpio_ != nil {
		s.gpio_.Cleanup()
	}
	if s.ignition != nil {
		s.ignition.Cleanup()
	}
	log.Println("Servizio OBD arrestato correttamente")
}

// checkPlatform replicates the JS checkRaspberryPiDependencies guard.
// Exits the process on non-ARM Linux unless DEV_MODE=true is set.
func checkPlatform() {
	devMode := os.Getenv("DEV_MODE") == "true"
	goos := runtime.GOOS
	goarch := runtime.GOARCH

	if goos != "linux" || (goarch != "arm" && goarch != "arm64") {
		if devMode {
			log.Printf("DEV_MODE attivo — server avviato su %s/%s senza hardware (non funzionale)", goos, goarch)
			return
		}
		log.Printf("Piattaforma non supportata: %s %s — richiesto Linux ARM (Raspberry Pi)", goos, goarch)
		log.Println("Per sviluppo locale usa il client in modalità mock (websocket.mock = true)")
		os.Exit(1)
	}
}

// itoa is a lightweight int-to-string helper to avoid importing strconv in this file.
func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	neg := n < 0
	if neg {
		n = -n
	}
	var buf [20]byte
	pos := len(buf)
	for n > 0 {
		pos--
		buf[pos] = byte('0' + n%10)
		n /= 10
	}
	if neg {
		pos--
		buf[pos] = '-'
	}
	return string(buf[pos:])
}
