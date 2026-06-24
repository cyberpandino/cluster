//go:build linux

/*
 * PandaOS
 * Copyright (C) 2025  Cyberpandino
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3.
 */

package gpio

import (
	"log"
	"os/exec"
	"sync/atomic"
	"time"

	"github.com/warthog618/gpiod"

	"github.com/cyberpandino/cluster/server/internal/config"
	"github.com/cyberpandino/cluster/server/internal/hub"
)

// IgnitionService watches a dedicated GPIO pin for ignition ON/OFF events
// and runs the configured power-management shell scripts.
type IgnitionService struct {
	hub_         *hub.Hub
	line         *gpiocdev.Line
	chip         *gpiocdev.Chip
	currentState atomic.Int32 // 0 = unknown, 1 = ON, 2 = OFF
	initialized  bool
}

// NewIgnitionService creates and initialises the ignition GPIO watcher.
func NewIgnitionService(h *hub.Hub) *IgnitionService {
	s := &IgnitionService{hub_: h}
	if config.Ignition.Enabled {
		s.init()
	}
	return s
}

func (s *IgnitionService) init() {
	chip, err := gpiocdev.NewChip("gpiochip0")
	if err != nil {
		log.Printf("IgnitionService: GPIO chip non disponibile: %v", err)
		return
	}
	s.chip = chip

	pin := config.Ignition.Pin
	line, err := chip.RequestLine(pin,
		gpiocdev.AsInput,
		gpiocdev.WithBothEdges,
		gpiocdev.WithDebounce(time.Duration(config.GPIO.DebounceTime)*time.Millisecond),
		gpiocdev.WithEventHandler(s.handleEdge),
	)
	if err != nil {
		log.Printf("IgnitionService: errore configurazione GPIO %d: %v", pin, err)
		chip.Close()
		s.chip = nil
		return
	}
	s.line = line

	// Set initial state from current pin value.
	val, err := line.Value()
	if err == nil {
		s.applyValue(val, true)
	}

	s.initialized = true
	log.Printf("IgnitionService inizializzato su GPIO %d (activeOn=%d)",
		pin, config.Ignition.ActiveOn)
}

// handleEdge is called by go-gpiocdev on every detected edge.
func (s *IgnitionService) handleEdge(evt gpiocdev.LineEvent) {
	val := 0
	if evt.Type == gpiocdev.LineEventRisingEdge {
		val = 1
	}
	s.applyValue(val, false)
}

func (s *IgnitionService) applyValue(val int, isInitial bool) {
	isOn := val == config.Ignition.ActiveOn
	newState := int32(2) // OFF
	if isOn {
		newState = 1 // ON
	}

	// Deduplicate: skip if state hasn't changed (except on first read).
	if !isInitial && s.currentState.Swap(newState) == newState {
		return
	}
	s.currentState.Store(newState)

	if isOn {
		log.Println("QUADRO ACCESO — Wake Mode")
		s.runScript(config.Ignition.Scripts.Wake)
	} else {
		log.Println("QUADRO SPENTO — Low Power Mode")
		s.runScript(config.Ignition.Scripts.LowPower)
	}

	d := hub.IgnitionData{Timestamp: time.Now().UTC().Format(time.RFC3339Nano)}
	d.Ignition.On = isOn
	d.Ignition.State = map[bool]string{true: "ON", false: "OFF"}[isOn]
	s.hub_.EmitIgnitionState(d)
}

func (s *IgnitionService) runScript(path string) {
	if path == "" {
		return
	}
	go func() {
		out, err := exec.Command("bash", path).CombinedOutput()
		if err != nil {
			log.Printf("Ignition script %s error: %v", path, err)
			return
		}
		if len(out) > 0 {
			log.Printf("Ignition script %s: %s", path, string(out))
		}
	}()
}

// Cleanup releases GPIO resources.
func (s *IgnitionService) Cleanup() {
	if s.line != nil {
		s.line.Close()
		s.line = nil
	}
	if s.chip != nil {
		s.chip.Close()
		s.chip = nil
	}
}
