/*
 * PandaOS
 * Copyright (C) 2025  Cyberpandino
 */

package config

import "testing"

func TestAllPins_Count(t *testing.T) {
	pins := AllPins()
	if len(pins) != len(Warnings) {
		t.Errorf("AllPins len=%d, want %d (one per warning)", len(pins), len(Warnings))
	}
}

func TestAllPins_NoDuplicates(t *testing.T) {
	seen := make(map[int]string)
	for _, p := range AllPins() {
		if prev, ok := seen[p]; ok {
			t.Errorf("pin %d used by at least two warnings (%s and another)", p, prev)
		}
		seen[p] = ""
	}
}

func TestPinToKey_Count(t *testing.T) {
	m := PinToKey()
	if len(m) != len(Warnings) {
		t.Errorf("PinToKey len=%d, want %d", len(m), len(Warnings))
	}
}

func TestPinToKey_Reversible(t *testing.T) {
	m := PinToKey()
	for key, mapping := range Warnings {
		got, ok := m[mapping.Pin]
		if !ok {
			t.Errorf("pin %d (from key %q) not in PinToKey map", mapping.Pin, key)
			continue
		}
		if got != key {
			t.Errorf("PinToKey[%d] = %q, want %q", mapping.Pin, got, key)
		}
	}
}

func TestIgnitionPin_NotInWarnings(t *testing.T) {
	for key, m := range Warnings {
		if m.Pin == Ignition.Pin {
			t.Errorf("ignition pin %d conflicts with warning key %q", Ignition.Pin, key)
		}
	}
}

func TestWarnings_HaveNames(t *testing.T) {
	for key, m := range Warnings {
		if m.Name == "" {
			t.Errorf("warning %q has empty Name", key)
		}
		if m.Pin == 0 {
			t.Errorf("warning %q has zero Pin", key)
		}
	}
}

func TestFuelConfig_VoltageDivider(t *testing.T) {
	if Fuel.VoltageDivider.R1 <= 0 {
		t.Errorf("R1 must be positive, got %f", Fuel.VoltageDivider.R1)
	}
	if Fuel.VoltageDivider.R2 <= 0 {
		t.Errorf("R2 must be positive, got %f", Fuel.VoltageDivider.R2)
	}
	if Fuel.Calibration.VoltageEmpty >= Fuel.Calibration.VoltageFull {
		t.Errorf("VoltageEmpty (%.2f) must be less than VoltageFull (%.2f)",
			Fuel.Calibration.VoltageEmpty, Fuel.Calibration.VoltageFull)
	}
}

func TestIgnitionConfig_Scripts(t *testing.T) {
	if Ignition.Scripts.LowPower == "" {
		t.Errorf("LowPower script path must not be empty")
	}
	if Ignition.Scripts.Wake == "" {
		t.Errorf("Wake script path must not be empty")
	}
}

func TestTemperatureConfig_Paths(t *testing.T) {
	if Temperature.BasePath == "" {
		t.Errorf("BasePath must not be empty")
	}
	if Temperature.ReadInterval <= 0 {
		t.Errorf("ReadInterval must be positive")
	}
}
