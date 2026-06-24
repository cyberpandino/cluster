/*
 * PandaOS
 * Copyright (C) 2025  Cyberpandino
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3.
 */

package obd

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/cyberpandino/cluster/server/internal/hub"
)

// Monitor continuously polls the working PID set and broadcasts live data.
type Monitor struct {
	comm        *Comm
	hub_        *hub.Hub
	mu          sync.Mutex
	monitorPIDs []PIDDef // subset of working PIDs used for live polling
	workingPIDs []PIDDef // full set of responsive PIDs
	active      bool
}

// NewMonitor creates a Monitor. Call StartMonitoring to begin the polling loop.
func NewMonitor(comm *Comm, h *hub.Hub) *Monitor {
	return &Monitor{comm: comm, hub_: h}
}

// StartMonitoring begins the live polling loop for the given working PID set.
// Blocks until ctx is cancelled, the OBD connection drops, or Stop is called.
func (m *Monitor) StartMonitoring(ctx context.Context, workingPIDs []PIDDef) {
	if len(workingPIDs) == 0 {
		log.Println("Nessun PID funzionante disponibile per il monitoraggio")
		return
	}

	m.mu.Lock()
	m.active = true
	m.updatePIDsLocked(workingPIDs)
	m.mu.Unlock()

	log.Println("Avvio monitoraggio continuo PID")
	m.hub_.EmitStatus(hub.StatusData{Monitoring: true, Message: "Monitoraggio attivo"})

	for {
		select {
		case <-ctx.Done():
			m.mu.Lock()
			m.active = false
			m.mu.Unlock()
			m.hub_.EmitStatus(hub.StatusData{Message: "Monitoraggio terminato"})
			return
		default:
		}

		m.mu.Lock()
		if !m.active || !m.comm.IsConnected() {
			m.mu.Unlock()
			break
		}
		pids := make([]PIDDef, len(m.monitorPIDs))
		copy(pids, m.monitorPIDs)
		m.mu.Unlock()

		data := hub.OBDLiveData{
			Timestamp:  time.Now().UTC().Format(time.RFC3339Nano),
			Monitoring: true,
			Parameters: make(map[string]hub.PIDValue, len(pids)),
		}

		for _, def := range pids {
			result := m.comm.ReadPID(ctx, def)
			data.Parameters[def.Key] = hub.PIDValue{
				Name:      result.Name,
				Value:     result.Value,
				Unit:      result.Unit,
				Success:   result.Success,
				Error:     result.Error,
				Timestamp: result.Timestamp,
			}
			select {
			case <-ctx.Done():
				m.mu.Lock()
				m.active = false
				m.mu.Unlock()
				m.hub_.EmitStatus(hub.StatusData{Message: "Monitoraggio terminato"})
				return
			case <-time.After(50 * time.Millisecond):
			}
		}

		m.hub_.EmitOBDLiveData(data)

		select {
		case <-ctx.Done():
			m.mu.Lock()
			m.active = false
			m.mu.Unlock()
			m.hub_.EmitStatus(hub.StatusData{Message: "Monitoraggio terminato"})
			return
		case <-time.After(300 * time.Millisecond):
		}
	}

	m.hub_.EmitStatus(hub.StatusData{Monitoring: false, Message: "Monitoraggio terminato"})
}

// UpdateWorkingPIDs replaces the current PID set while monitoring is running.
func (m *Monitor) UpdateWorkingPIDs(newPIDs []PIDDef) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if !m.active {
		return
	}
	prev := len(m.monitorPIDs)
	m.updatePIDsLocked(newPIDs)
	next := len(m.monitorPIDs)

	switch {
	case next > prev:
		added := next - prev
		log.Printf("Lista PID aggiornata: +%d PID aggiunti (totale: %d)", added, next)
		m.hub_.EmitStatus(hub.StatusData{
			Message:        fmt.Sprintf("PID aggiornati: %d nuovi PID aggiunti", added),
			MonitoringPIDs: next,
			AddedPIDs:      added,
		})
	case next < prev:
		removed := prev - next
		log.Printf("Lista PID aggiornata: -%d PID rimossi (totale: %d)", removed, next)
		m.hub_.EmitStatus(hub.StatusData{
			Message:        fmt.Sprintf("PID aggiornati: %d PID rimossi", removed),
			MonitoringPIDs: next,
			RemovedPIDs:    removed,
		})
	}
}

// IsPIDMonitored reports whether the given key is in the current working set.
func (m *Monitor) IsPIDMonitored(key string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, p := range m.workingPIDs {
		if p.Key == key {
			return true
		}
	}
	return false
}

// Stop signals the monitoring loop to exit.
func (m *Monitor) Stop() {
	m.mu.Lock()
	m.active = false
	m.mu.Unlock()
	log.Println("Arresto monitoraggio PID")
}

// IsActive reports whether the monitoring loop is running.
func (m *Monitor) IsActive() bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.active
}

// updatePIDsLocked rebuilds monitorPIDs from workingPIDs. Caller must hold m.mu.
func (m *Monitor) updatePIDsLocked(working []PIDDef) {
	m.workingPIDs = working
	m.monitorPIDs = m.monitorPIDs[:0]
	for _, p := range working {
		if IsMonitorKey(p.Key) {
			m.monitorPIDs = append(m.monitorPIDs, p)
		}
	}
}
