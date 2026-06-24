/*
 * PandaOS
 * Copyright (C) 2025  Cyberpandino
 */

package obd

import (
	"testing"
)

func TestMonitor_InitialState(t *testing.T) {
	m := &Monitor{}
	if m.IsActive() {
		t.Error("new Monitor should not be active")
	}
}

func TestMonitor_Stop(t *testing.T) {
	m := &Monitor{active: true}
	m.Stop()
	if m.IsActive() {
		t.Error("Stop() should set active=false")
	}
}

func TestMonitor_IsPIDMonitored_Empty(t *testing.T) {
	m := &Monitor{}
	if m.IsPIDMonitored("rpm") {
		t.Error("should not monitor any PID when working set is empty")
	}
}

func TestMonitor_IsPIDMonitored_Present(t *testing.T) {
	m := &Monitor{
		workingPIDs: []PIDDef{
			{PID: "010C", Name: "RPM", Key: "rpm"},
			{PID: "010D", Name: "Speed", Key: "speed"},
		},
	}
	if !m.IsPIDMonitored("rpm") {
		t.Error("expected rpm to be monitored")
	}
	if !m.IsPIDMonitored("speed") {
		t.Error("expected speed to be monitored")
	}
	if m.IsPIDMonitored("coolant_temp") {
		t.Error("coolant_temp should not be in empty-initialized monitor")
	}
}

func TestUpdatePIDsLocked_FiltersMonitorKeys(t *testing.T) {
	m := &Monitor{}
	working := []PIDDef{
		{Key: "rpm"},             // monitor key ✓
		{Key: "speed"},           // monitor key ✓
		{Key: "monitor_status"},  // NOT a monitor key
		{Key: "coolant_temp"},    // monitor key ✓
		{Key: "obd_standards"},   // NOT a monitor key
	}
	m.mu.Lock()
	m.updatePIDsLocked(working)
	m.mu.Unlock()

	if len(m.workingPIDs) != 5 {
		t.Errorf("workingPIDs: got %d, want 5", len(m.workingPIDs))
	}
	wantMonitor := 3 // rpm, speed, coolant_temp
	if len(m.monitorPIDs) != wantMonitor {
		t.Errorf("monitorPIDs: got %d, want %d", len(m.monitorPIDs), wantMonitor)
	}
	for _, p := range m.monitorPIDs {
		if !IsMonitorKey(p.Key) {
			t.Errorf("non-monitor key %q ended up in monitorPIDs", p.Key)
		}
	}
}

func TestUpdateWorkingPIDs_InactiveMonitor(t *testing.T) {
	// When inactive, UpdateWorkingPIDs should be a no-op (no panic on nil hub).
	m := &Monitor{active: false}
	newPIDs := []PIDDef{{Key: "rpm"}}
	// Should not panic even with nil hub_.
	m.UpdateWorkingPIDs(newPIDs)
	if len(m.workingPIDs) != 0 {
		t.Error("inactive monitor should not update working PIDs")
	}
}

func TestMonitor_StopTwice(t *testing.T) {
	m := &Monitor{active: true}
	m.Stop()
	m.Stop() // must not panic
	if m.IsActive() {
		t.Error("should still be inactive after second Stop()")
	}
}

// ---- benchmarks ----

func BenchmarkIsPIDMonitored_Hit(b *testing.B) {
	m := &Monitor{
		workingPIDs: []PIDDef{
			{Key: "rpm"}, {Key: "speed"}, {Key: "coolant_temp"},
		},
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		m.IsPIDMonitored("coolant_temp")
	}
}

func BenchmarkIsPIDMonitored_Miss(b *testing.B) {
	m := &Monitor{
		workingPIDs: []PIDDef{
			{Key: "rpm"}, {Key: "speed"}, {Key: "coolant_temp"},
		},
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		m.IsPIDMonitored("obd_standards")
	}
}
