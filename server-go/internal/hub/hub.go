/*
 * PandaOS
 * Copyright (C) 2025  Cyberpandino
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3.
 */

// Package hub wraps the Socket.IO server and exposes typed broadcast methods
// that match the event names and payloads expected by the React client.
package hub

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/zishang520/socket.io/v2/socket"
)

// Hub is a thread-safe Socket.IO v4 broadcaster.
type Hub struct {
	io             *socket.Server
	mu             sync.RWMutex
	lastOBDData    interface{} // cached for new-client handshake
	onForceRestart func()
}

// ---- event payload types (JSON tags match client expectations) ----

type StatusData struct {
	Timestamp         string `json:"timestamp"`
	Connected         bool   `json:"connected,omitempty"`
	Monitoring        bool   `json:"monitoring,omitempty"`
	Message           string `json:"message,omitempty"`
	Error             string `json:"error,omitempty"`
	Restarting        bool   `json:"restarting,omitempty"`
	ReconnectAttempts int    `json:"reconnectAttempts,omitempty"`
	WorkingPIDs       int    `json:"working_pids,omitempty"`
	TotalPIDs         int    `json:"total_pids,omitempty"`
	AddedPIDs         int    `json:"added_pids,omitempty"`
	RemovedPIDs       int    `json:"removed_pids,omitempty"`
	MonitoringPIDs    int    `json:"monitoring_pids,omitempty"`
}

type PIDValue struct {
	Name      string      `json:"name"`
	Value     interface{} `json:"value"`
	Unit      string      `json:"unit"`
	Success   bool        `json:"success"`
	Error     bool        `json:"error,omitempty"`
	Timestamp string      `json:"timestamp"`
}

type OBDData struct {
	Timestamp  string              `json:"timestamp"`
	Vehicle    map[string]string   `json:"vehicle,omitempty"`
	Parameters map[string]PIDValue `json:"parameters"`
}

type OBDLiveData struct {
	Timestamp  string              `json:"timestamp"`
	Monitoring bool                `json:"monitoring"`
	Parameters map[string]PIDValue `json:"parameters"`
}

type WarningChange struct {
	Warning string `json:"warning"`
	State   bool   `json:"state"`
	Pin     int    `json:"pin"`
	Name    string `json:"name"`
}

type WarningsData struct {
	Timestamp string          `json:"timestamp"`
	Warnings  map[string]bool `json:"warnings"`
	Changes   []WarningChange `json:"changes"`
}

type TemperatureData struct {
	Timestamp   string `json:"timestamp"`
	Temperature struct {
		External float64 `json:"external"`
		Unit     string  `json:"unit"`
		Sensor   string  `json:"sensor"`
	} `json:"temperature"`
}

type FuelData struct {
	Timestamp string `json:"timestamp"`
	Fuel      struct {
		Level  float64 `json:"level"`
		Unit   string  `json:"unit"`
		Sensor string  `json:"sensor"`
		Raw    struct {
			ADC     int     `json:"adc"`
			Voltage float64 `json:"voltage"`
		} `json:"raw"`
	} `json:"fuel"`
}

type IgnitionData struct {
	Timestamp string `json:"timestamp"`
	Ignition  struct {
		On    bool   `json:"on"`
		State string `json:"state"`
	} `json:"ignition"`
}

// New starts a Socket.IO v4 server listening on addr (e.g. ":3001").
// onForceRestart is called when a client sends the "force-restart" event.
func New(addr string, onForceRestart func()) (*Hub, error) {
	io := socket.NewServer(nil, nil)
	h := &Hub{io: io, onForceRestart: onForceRestart}

	io.On("connection", func(clients ...any) {
		s := clients[0].(*socket.Socket)
		log.Printf("Client connesso: %s", s.Id())

		// Send cached OBD data so the new client sees a non-empty gauge immediately.
		h.mu.RLock()
		last := h.lastOBDData
		h.mu.RUnlock()
		if last != nil {
			if err := s.Emit("obd-data", last); err != nil {
				log.Printf("obd-data handshake emit error: %v", err)
			}
		}

		s.On("force-restart", func(...any) {
			log.Println("Riavvio forzato richiesto dal client")
			if h.onForceRestart != nil {
				h.onForceRestart()
			}
		})

		s.On("disconnect", func(reason ...any) {
			log.Printf("Client disconnesso: %s (%v)", s.Id(), reason)
		})
	})

	mux := http.NewServeMux()
	mux.Handle("/socket.io/", corsMiddleware(io.ServeHandler(nil)))
	mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	httpSrv := &http.Server{Addr: addr, Handler: mux}
	go func() {
		log.Printf("Server OBD in ascolto su %s", addr)
		if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("HTTP server error: %v", err)
		}
	}()

	return h, nil
}

// corsMiddleware adds permissive CORS headers required for Socket.IO polling.
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func now() string { return time.Now().UTC().Format(time.RFC3339Nano) }

func (h *Hub) EmitStatus(d StatusData) {
	d.Timestamp = now()
	h.io.Emit("status", d)
}

func (h *Hub) EmitOBDData(d OBDData) {
	h.mu.Lock()
	h.lastOBDData = d
	h.mu.Unlock()
	h.io.Emit("obd-data", d)
}

func (h *Hub) EmitOBDLiveData(d OBDLiveData) {
	h.io.Emit("obd-live", d)
}

func (h *Hub) EmitWarnings(d WarningsData) {
	h.io.Emit("gpio-warnings", d)
}

func (h *Hub) EmitExternalTemperature(d TemperatureData) {
	h.io.Emit("external-temperature", d)
}

func (h *Hub) EmitFuelLevel(d FuelData) {
	h.io.Emit("fuel-level", d)
}

func (h *Hub) EmitIgnitionState(d IgnitionData) {
	h.io.Emit("ignition-state", d)
}

func (h *Hub) EmitError(msg string) {
	h.io.Emit("error", map[string]string{
		"message":   msg,
		"timestamp": now(),
	})
}
