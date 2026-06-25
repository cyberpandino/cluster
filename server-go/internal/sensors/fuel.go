/*
 * PandaOS
 * Copyright (C) 2025  Cyberpandino
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3.
 */

package sensors

import (
	"context"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"time"

	"golang.org/x/sys/unix"

	"github.com/cyberpandino/cluster/server/internal/config"
	"github.com/cyberpandino/cluster/server/internal/hub"
)

const (
	i2cSlave    = 0x0703 // ioctl request: set slave address
	ads1115Addr = 0x48   // default I2C address when ADDR pin is GND

	// ADS1115 register pointers.
	regConversion = 0x00
	regConfig     = 0x01

	// Config register value: OS=1 (start), MUX=100 (AIN0 vs GND),
	// PGA=001 (±4.096 V), MODE=1 (single-shot), DR=100 (250 SPS),
	// COMP disabled (11).
	// Binary: 1_100_001_1_1000_0011 = 0xC383
	ads1115Config = 0xC383

	// Full-scale voltage for PGA=001 (±4.096 V) in millivolts.
	fullScaleMillivolts = 4096.0
)

// FuelService reads fuel level from an ADS1115 ADC via the Linux I2C device.
type FuelService struct {
	hub_         *hub.Hub
	fd           int
	lastLevel    float64
	initialized  bool
}

// NewFuelService creates the service and opens the I2C device.
func NewFuelService(h *hub.Hub) *FuelService {
	s := &FuelService{hub_: h, fd: -1}
	if config.Fuel.Enabled {
		s.init()
	} else {
		log.Println("Sensore carburante ADS1115 disabilitato nella configurazione")
	}
	return s
}

func (s *FuelService) init() {
	fd, err := unix.Open(config.Fuel.I2CBus, unix.O_RDWR, 0)
	if err != nil {
		log.Printf("Errore apertura I2C %s: %v — sensore carburante disabilitato", config.Fuel.I2CBus, err)
		return
	}

	// Set ADS1115 as the I2C target.
	if _, _, errno := unix.Syscall(unix.SYS_IOCTL,
		uintptr(fd), i2cSlave, uintptr(ads1115Addr)); errno != 0 {
		log.Printf("Errore ioctl I2C_SLAVE: %v", errno)
		unix.Close(fd)
		return
	}

	s.fd = fd
	s.initialized = true
	log.Println("Sensore carburante ADS1115 inizializzato correttamente")
}

// StartReading runs the periodic ADC read loop until ctx is cancelled.
func (s *FuelService) StartReading(ctx context.Context) {
	if !config.Fuel.Enabled || !s.initialized {
		return
	}

	interval := time.Duration(config.Fuel.ReadInterval) * time.Millisecond
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	log.Printf("Lettura carburante ADS1115 avviata (intervallo: %dms)", config.Fuel.ReadInterval)

	// Read immediately on start.
	s.performReading()

	for {
		select {
		case <-ctx.Done():
			log.Println("Lettura carburante ADS1115 fermata")
			s.close()
			return
		case <-ticker.C:
			s.performReading()
		}
	}
}

func (s *FuelService) performReading() {
	mv, raw, err := s.readADC()
	if err != nil {
		log.Printf("Errore lettura ADC: %v", err)
		return
	}

	level, voltage := s.processReading(mv)

	// Only broadcast when level changed by more than 0.5%.
	if math.Abs(level-s.lastLevel) < 0.5 {
		return
	}
	s.lastLevel = level

	d := hub.FuelData{Timestamp: time.Now().UTC().Format(time.RFC3339Nano)}
	d.Fuel.Level = math.Round(level*10) / 10
	d.Fuel.Unit = "%"
	d.Fuel.Sensor = "ADS1115"
	d.Fuel.Raw.ADC = raw
	d.Fuel.Raw.Voltage = math.Round(voltage*1000) / 1000
	s.hub_.EmitFuelLevel(d)
	log.Printf("Livello carburante: %.1f%%", level)
}

// readADC triggers a single-shot conversion and returns (millivolts, raw_adc, error).
func (s *FuelService) readADC() (float64, int, error) {
	// Write config register to start conversion.
	cfgHigh := byte(ads1115Config >> 8)
	cfgLow := byte(ads1115Config & 0xFF)
	if _, err := unix.Write(s.fd, []byte{regConfig, cfgHigh, cfgLow}); err != nil {
		return 0, 0, fmt.Errorf("writing config register: %w", err)
	}

	// Wait for conversion (250 SPS → ~4 ms; use 6 ms for margin).
	time.Sleep(6 * time.Millisecond)

	// Point at conversion register.
	if _, err := unix.Write(s.fd, []byte{regConversion}); err != nil {
		return 0, 0, fmt.Errorf("setting conversion register: %w", err)
	}

	// Read 2-byte result.
	buf := make([]byte, 2)
	if _, err := unix.Read(s.fd, buf); err != nil {
		return 0, 0, fmt.Errorf("reading conversion: %w", err)
	}

	raw16 := int16(binary.BigEndian.Uint16(buf))
	// Full-scale (32767) = ±4.096 V for PGA=001.
	mv := float64(raw16) * fullScaleMillivolts / 32767.0
	return mv, int(raw16), nil
}

// processReading converts millivolts → fuel percentage using the voltage divider
// and calibration values from config.
func (s *FuelService) processReading(mv float64) (percent, voltage float64) {
	vAdc := mv / 1000 // mv → V at ADC input
	r1, r2 := config.Fuel.VoltageDivider.R1, config.Fuel.VoltageDivider.R2
	voltage = vAdc * ((r1 + r2) / r2) // reverse voltage divider to get sensor voltage

	empty := config.Fuel.Calibration.VoltageEmpty
	full := config.Fuel.Calibration.VoltageFull
	percent = (voltage - empty) / (full - empty) * 100
	percent = math.Max(0, math.Min(100, percent))
	return
}

func (s *FuelService) close() {
	if s.fd >= 0 {
		unix.Close(s.fd)
		s.fd = -1
	}
}

