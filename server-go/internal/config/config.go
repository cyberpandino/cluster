/*
 * PandaOS
 * Copyright (C) 2025  Cyberpandino
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3.
 */

// Package config holds static hardware configuration for GPIO, sensors, and ignition.
// Edit this file to change pin assignments, sensor calibration, or power scripts.
package config

// GPIO holds global GPIO timing settings.
var GPIO = struct {
	DebounceTime    int // debounce filter in ms
	PollingInterval int // warning-light read interval in ms
}{
	DebounceTime:    50,
	PollingInterval: 100,
}

// Ignition is the optocoupler-based ignition key detector configuration.
var Ignition = struct {
	Enabled  bool
	Pin      int
	ActiveOn int // 0 = active-low (opto pulls to GND when ON), 1 = active-high
	Scripts  struct {
		LowPower string // executed when ignition turns OFF
		Wake     string // executed when ignition turns ON
	}
}{
	Enabled:  true,
	Pin:      21,
	ActiveOn: 0,
}

func init() {
	Ignition.Scripts.LowPower = "./scripts/low-power.sh"
	Ignition.Scripts.Wake = "./scripts/wake.sh"
}

// Temperature is the DS18B20 1-Wire temperature sensor configuration.
var Temperature = struct {
	Enabled      bool
	SensorID     string // empty = auto-detect first 28-* device
	BasePath     string
	ReadInterval int // ms
}{
	Enabled:      true,
	SensorID:     "",
	BasePath:     "/sys/bus/w1/devices",
	ReadInterval: 5000,
}

// Fuel is the ADS1115 ADC configuration for the resistive fuel-level sender.
var Fuel = struct {
	Enabled      bool
	ReadInterval int // ms
	// ADS1115 I2C settings
	I2CBus  string // e.g. "/dev/i2c-1"
	Address uint16 // default 0x48
	// Voltage divider: V_sensor is measured across R2
	VoltageDivider struct{ R1, R2 float64 }
	// Calibration: map sensor voltage to 0–100 %
	Calibration struct{ VoltageEmpty, VoltageFull float64 }
}{
	Enabled:      true,
	ReadInterval: 500,
	I2CBus:       "/dev/i2c-1",
	Address:      0x48,
}

func init() {
	Fuel.VoltageDivider.R1 = 100_000 // 100 kΩ
	Fuel.VoltageDivider.R2 = 33_000  // 33 kΩ
	Fuel.Calibration.VoltageEmpty = 0.5
	Fuel.Calibration.VoltageFull = 4.0
}

// WarningMapping links a frontend warning key to a physical GPIO pin.
type WarningMapping struct {
	Pin         int
	Name        string
	Description string
}

// Warnings maps frontend state.warnings keys (client/src/store/state.tsx)
// to their GPIO BCM pin numbers via optocouplers.
// HIGH (1) = warning light ON; LOW (0) = OFF.
var Warnings = map[string]WarningMapping{
	"turnSignals":  {17, "Frecce", "Indicatori di direzione"},
	"battery":      {27, "Alternatore", "Carica batteria/alternatore"},
	"engineOil":    {22, "Pressione olio", "Pressione olio motore"},
	"brakeSystem":  {23, "Freni", "Sistema frenante"},
	"injectors":    {24, "Iniettori", "Sistema iniezione carburante"},
	"keyOn":        {25, "KEY", "Chiave inserita/quadro acceso"},
	"highBeam":     {5, "Abbaglianti", "Fari abbaglianti"},
	"lowBeam":      {6, "Anabbaglianti", "Fari anabbaglianti"},
	"hazard":       {12, "Quattro frecce", "Luci di emergenza"},
	"fogLight":     {13, "Fendinebbia", "Fendinebbia posteriore"},
	"engineColant": {16, "Temperatura raffreddamento", "Temperatura liquido refrigerante"},
	"rearDefrost":  {19, "Termoresistenza lunotto", "Sbrinatore lunotto posteriore"},
	"fuel":         {20, "Riserva carburante", "Livello carburante in riserva"},
}

// AllPins returns every configured GPIO BCM pin number.
func AllPins() []int {
	pins := make([]int, 0, len(Warnings))
	for _, m := range Warnings {
		pins = append(pins, m.Pin)
	}
	return pins
}

// PinToKey returns a map from BCM pin number to frontend warning key.
func PinToKey() map[int]string {
	m := make(map[int]string, len(Warnings))
	for key, mapping := range Warnings {
		m[mapping.Pin] = key
	}
	return m
}
