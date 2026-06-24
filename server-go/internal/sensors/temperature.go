/*
 * PandaOS
 * Copyright (C) 2025  Cyberpandino
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3.
 */

// Package sensors implements DS18B20 (1-Wire temperature) and ADS1115 (I2C ADC fuel level) readers.
package sensors

import (
	"context"
	"log"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/cyberpandino/cluster/server/internal/config"
	"github.com/cyberpandino/cluster/server/internal/hub"
)

var tempLineRE = regexp.MustCompile(`t=(\d+)`)

// TemperatureService reads the DS18B20 sensor via the Linux 1-Wire sysfs interface.
type TemperatureService struct {
	hub_         *hub.Hub
	sensorPath   string
	lastTemp     float64
	initialized  bool
}

// NewTemperatureService creates the service. It auto-detects the sensor path.
func NewTemperatureService(h *hub.Hub) *TemperatureService {
	s := &TemperatureService{hub_: h}
	if config.Temperature.Enabled {
		s.init()
	} else {
		log.Println("Sensore temperatura DS18B20 disabilitato nella configurazione")
	}
	return s
}

func (s *TemperatureService) init() {
	var path string
	if config.Temperature.SensorID != "" {
		path = filepath.Join(config.Temperature.BasePath, config.Temperature.SensorID, "w1_slave")
	} else {
		path = s.autoDetect()
	}

	if path == "" {
		log.Println("Sensore DS18B20 non trovato, temperatura esterna disabilitata")
		return
	}
	if _, err := os.Stat(path); err != nil {
		log.Printf("Sensore DS18B20 non trovato a %s", path)
		return
	}

	s.sensorPath = path
	s.initialized = true
	log.Printf("Sensore temperatura DS18B20 trovato: %s", path)

	if t, ok := s.read(); ok {
		log.Printf("Temperatura esterna iniziale: %.1f°C", t)
	}
}

func (s *TemperatureService) autoDetect() string {
	entries, err := os.ReadDir(config.Temperature.BasePath)
	if err != nil {
		return ""
	}
	for _, e := range entries {
		if strings.HasPrefix(e.Name(), "28-") {
			return filepath.Join(config.Temperature.BasePath, e.Name(), "w1_slave")
		}
	}
	return ""
}

// StartReading runs the periodic temperature read loop until ctx is cancelled.
func (s *TemperatureService) StartReading(ctx context.Context) {
	if !config.Temperature.Enabled || !s.initialized {
		return
	}

	interval := time.Duration(config.Temperature.ReadInterval) * time.Millisecond
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	log.Printf("Lettura temperatura DS18B20 avviata (intervallo: %dms)",
		config.Temperature.ReadInterval)

	// Emit immediately on start.
	if t, ok := s.read(); ok {
		s.emit(t)
	}

	for {
		select {
		case <-ctx.Done():
			log.Println("Lettura temperatura DS18B20 fermata")
			return
		case <-ticker.C:
			if t, ok := s.read(); ok && t != s.lastTemp {
				s.lastTemp = t
				s.emit(t)
			}
		}
	}
}

// read parses the 1-Wire sysfs file and returns the temperature in °C.
func (s *TemperatureService) read() (float64, bool) {
	data, err := os.ReadFile(s.sensorPath)
	if err != nil {
		log.Printf("Errore lettura sensore DS18B20: %v", err)
		return 0, false
	}

	lines := strings.Split(string(data), "\n")
	if len(lines) < 2 || !strings.Contains(lines[0], "YES") {
		return 0, false // CRC invalid
	}

	m := tempLineRE.FindStringSubmatch(lines[1])
	if m == nil {
		return 0, false
	}
	raw, err := strconv.Atoi(m[1])
	if err != nil {
		return 0, false
	}
	// raw is in thousandths of a degree Celsius.
	t := math.Round(float64(raw)/100) / 10 // one decimal place
	return t, true
}

func (s *TemperatureService) emit(temp float64) {
	d := hub.TemperatureData{Timestamp: time.Now().UTC().Format(time.RFC3339Nano)}
	d.Temperature.External = temp
	d.Temperature.Unit = "°C"
	d.Temperature.Sensor = "DS18B20"
	s.hub_.EmitExternalTemperature(d)
	log.Printf("Temperatura esterna: %.1f°C", temp)
}
