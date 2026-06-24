//go:build linux

/*
 * PandaOS
 * Copyright (C) 2025  Cyberpandino
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3.
 */

// Package gpio implements Raspberry Pi GPIO access for vehicle warning lights
// and ignition detection via optocouplers.
package gpio

import (
	"context"
	"log"
	"sort"
	"time"

	"github.com/warthog618/gpiod"

	"github.com/cyberpandino/cluster/server/internal/config"
	"github.com/cyberpandino/cluster/server/internal/hub"
)

// Service polls all warning-light GPIO pins and broadcasts state changes.
type Service struct {
	hub_         *hub.Hub
	chip         *gpiocdev.Chip
	lines        *gpiocdev.Lines
	pinOrder     []int           // ordered list matching lines request
	pinToKey     map[int]string  // BCM pin → warning key
	warningState map[string]bool // current state per key
	initialized  bool
}

// NewService creates a GPIO service. Returns a non-nil service even when
// hardware is unavailable (it will skip polling gracefully).
func NewService(h *hub.Hub) *Service {
	s := &Service{
		hub_:         h,
		pinToKey:     config.PinToKey(),
		warningState: make(map[string]bool, len(config.Warnings)),
	}
	for key := range config.Warnings {
		s.warningState[key] = false
	}
	s.init()
	return s
}

func (s *Service) init() {
	chip, err := gpiocdev.NewChip("gpiochip0")
	if err != nil {
		log.Printf("GPIO chip non disponibile (non su Raspberry Pi): %v", err)
		return
	}
	s.chip = chip

	pins := config.AllPins()
	// Sort for deterministic ordering; lines.Values uses same order.
	sort.Ints(pins)
	s.pinOrder = pins

	lines, err := chip.RequestLines(pins,
		gpiocdev.AsInput,
		gpiocdev.WithDebounce(time.Duration(config.GPIO.DebounceTime)*time.Millisecond),
	)
	if err != nil {
		log.Printf("Errore configurazione GPIO lines: %v", err)
		chip.Close()
		s.chip = nil
		return
	}
	s.lines = lines
	s.initialized = true
	log.Printf("GPIO inizializzato: %d pin configurati", len(pins))
}

// StartPolling runs the GPIO polling loop until ctx is cancelled.
func (s *Service) StartPolling(ctx context.Context) {
	if !s.initialized {
		log.Println("GPIO polling saltato: hardware non disponibile")
		return
	}

	ticker := time.NewTicker(time.Duration(config.GPIO.PollingInterval) * time.Millisecond)
	defer ticker.Stop()

	log.Printf("GPIO polling avviato (intervallo: %dms)", config.GPIO.PollingInterval)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			s.readAll()
		}
	}
}

func (s *Service) readAll() {
	vals := make([]int, len(s.pinOrder))
	if err := s.lines.Values(vals); err != nil {
		log.Printf("Errore lettura GPIO: %v", err)
		return
	}

	var changes []hub.WarningChange

	for i, pin := range s.pinOrder {
		key, ok := s.pinToKey[pin]
		if !ok {
			continue
		}
		isActive := vals[i] == 1
		if s.warningState[key] != isActive {
			s.warningState[key] = isActive
			name := config.Warnings[key].Name
			log.Printf("Spia %s: %s", name, map[bool]string{true: "ACCESA", false: "SPENTA"}[isActive])
			changes = append(changes, hub.WarningChange{
				Warning: key,
				State:   isActive,
				Pin:     pin,
				Name:    name,
			})
		}
	}

	if len(changes) > 0 {
		// Copy current state snapshot.
		snapshot := make(map[string]bool, len(s.warningState))
		for k, v := range s.warningState {
			snapshot[k] = v
		}
		s.hub_.EmitWarnings(hub.WarningsData{
			Timestamp: time.Now().UTC().Format(time.RFC3339Nano),
			Warnings:  snapshot,
			Changes:   changes,
		})
	}
}

// AllWarningStates returns a copy of the current warning state map.
func (s *Service) AllWarningStates() map[string]bool {
	out := make(map[string]bool, len(s.warningState))
	for k, v := range s.warningState {
		out[k] = v
	}
	return out
}

// Cleanup releases GPIO resources.
func (s *Service) Cleanup() {
	if s.lines != nil {
		s.lines.Close()
		s.lines = nil
	}
	if s.chip != nil {
		s.chip.Close()
		s.chip = nil
	}
	log.Println("GPIO cleanup completato")
}
