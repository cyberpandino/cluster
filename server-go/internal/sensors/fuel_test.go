/*
 * PandaOS
 * Copyright (C) 2025  Cyberpandino
 */

package sensors

import (
	"math"
	"testing"

	"github.com/cyberpandino/cluster/server/internal/config"
)

// processReading is private; these tests live in package sensors.

func TestProcessReading_Zero(t *testing.T) {
	svc := &FuelService{}
	// 0 mV → 0V at ADC → 0V sensor → below VoltageEmpty → clamped to 0%
	percent, voltage := svc.processReading(0)
	if percent != 0 {
		t.Errorf("percent: got %.4f, want 0", percent)
	}
	if voltage != 0 {
		t.Errorf("voltage: got %.4f, want 0", voltage)
	}
}

func TestProcessReading_HalfTank(t *testing.T) {
	svc := &FuelService{}
	// Compute mV that corresponds to half-tank voltage.
	// half-tank voltage = (VoltageEmpty + VoltageFull) / 2 = (0.5 + 4.0) / 2 = 2.25V
	// voltage = vAdc * (R1+R2)/R2  →  vAdc = voltage * R2/(R1+R2)
	r1, r2 := config.Fuel.VoltageDivider.R1, config.Fuel.VoltageDivider.R2
	halfV := (config.Fuel.Calibration.VoltageEmpty + config.Fuel.Calibration.VoltageFull) / 2
	vAdc := halfV * r2 / (r1 + r2)
	halfMV := vAdc * 1000

	percent, voltage := svc.processReading(halfMV)

	if math.Abs(percent-50.0) > 0.01 {
		t.Errorf("percent: got %.4f, want ~50.0", percent)
	}
	if math.Abs(voltage-halfV) > 0.01 {
		t.Errorf("voltage: got %.4f, want ~%.4f", voltage, halfV)
	}
}

func TestProcessReading_FullTank(t *testing.T) {
	svc := &FuelService{}
	// Compute mV that corresponds exactly to VoltageFull at the sensor.
	r1, r2 := config.Fuel.VoltageDivider.R1, config.Fuel.VoltageDivider.R2
	fullV := config.Fuel.Calibration.VoltageFull
	vAdc := fullV * r2 / (r1 + r2)
	fullMV := vAdc * 1000

	percent, _ := svc.processReading(fullMV)
	if math.Abs(percent-100.0) > 0.01 {
		t.Errorf("percent: got %.4f, want ~100.0", percent)
	}
}

func TestProcessReading_EmptyTank(t *testing.T) {
	svc := &FuelService{}
	// Voltage at ADC corresponding to VoltageEmpty at the sensor.
	r1, r2 := config.Fuel.VoltageDivider.R1, config.Fuel.VoltageDivider.R2
	emptyV := config.Fuel.Calibration.VoltageEmpty
	vAdc := emptyV * r2 / (r1 + r2)
	emptyMV := vAdc * 1000

	percent, _ := svc.processReading(emptyMV)
	if math.Abs(percent-0.0) > 0.01 {
		t.Errorf("percent: got %.4f, want ~0.0", percent)
	}
}

func TestProcessReading_ClampsBelowZero(t *testing.T) {
	svc := &FuelService{}
	// Negative mV → below empty → clamped to 0.
	percent, _ := svc.processReading(-500)
	if percent != 0 {
		t.Errorf("percent: got %.4f, want 0 (clamped)", percent)
	}
}

func TestProcessReading_ClampsAbove100(t *testing.T) {
	svc := &FuelService{}
	// Very high mV → well above full → clamped to 100.
	percent, _ := svc.processReading(5000)
	if percent != 100 {
		t.Errorf("percent: got %.4f, want 100 (clamped)", percent)
	}
}

func TestProcessReading_VoltageDividerRatio(t *testing.T) {
	svc := &FuelService{}
	// Test that the voltage divider formula is applied correctly.
	// With 1V at ADC and R1=100k, R2=33k: sensorV = 1 * (133000/33000) ≈ 4.03V
	r1, r2 := config.Fuel.VoltageDivider.R1, config.Fuel.VoltageDivider.R2
	expectedV := 1.0 * (r1 + r2) / r2
	_, voltage := svc.processReading(1000) // 1000 mV = 1V at ADC
	if math.Abs(voltage-expectedV) > 0.001 {
		t.Errorf("voltage: got %.4f, want %.4f", voltage, expectedV)
	}
}

// ---- benchmarks ----

func BenchmarkProcessReading(b *testing.B) {
	svc := &FuelService{}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		svc.processReading(558.3)
	}
}
