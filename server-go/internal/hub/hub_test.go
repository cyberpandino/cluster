/*
 * PandaOS
 * Copyright (C) 2025  Cyberpandino
 */

package hub

import (
	"encoding/json"
	"strings"
	"testing"
)

// These tests verify that struct tags produce the JSON shape the React client expects.
// If a field name changes, these tests catch the regression before a deploy.

func TestStatusData_JSON_OmitEmpty(t *testing.T) {
	d := StatusData{Message: "hello"}
	b, err := json.Marshal(d)
	if err != nil {
		t.Fatal(err)
	}
	s := string(b)

	mustContain(t, s, `"message":"hello"`)
	// omitempty fields with zero values must be absent.
	for _, absent := range []string{`"connected"`, `"monitoring"`, `"restarting"`,
		`"error"`, `"working_pids"`, `"total_pids"`} {
		if strings.Contains(s, absent) {
			t.Errorf("expected %q to be omitted, but found in %s", absent, s)
		}
	}
}

func TestStatusData_JSON_AllFields(t *testing.T) {
	d := StatusData{
		Connected:         true,
		Monitoring:        true,
		Message:           "ok",
		Error:             "oops",
		Restarting:        true,
		ReconnectAttempts: 3,
		WorkingPIDs:       5,
		TotalPIDs:         19,
	}
	b, _ := json.Marshal(d)
	s := string(b)

	mustContain(t, s, `"connected":true`)
	mustContain(t, s, `"monitoring":true`)
	mustContain(t, s, `"error":"oops"`)
	mustContain(t, s, `"reconnectAttempts":3`)
	mustContain(t, s, `"working_pids":5`)
	mustContain(t, s, `"total_pids":19`)
}

func TestPIDValue_JSON(t *testing.T) {
	pv := PIDValue{
		Name:      "Regime Motore",
		Value:     1726,
		Unit:      "RPM",
		Success:   true,
		Timestamp: "2025-01-01T00:00:00Z",
	}
	b, _ := json.Marshal(pv)
	s := string(b)

	mustContain(t, s, `"name":"Regime Motore"`)
	mustContain(t, s, `"value":1726`)
	mustContain(t, s, `"unit":"RPM"`)
	mustContain(t, s, `"success":true`)
	// Error is omitempty — must be absent when false.
	if strings.Contains(s, `"error"`) {
		t.Errorf(`"error" key should be omitted when false: %s`, s)
	}
}

func TestOBDLiveData_JSON_NestedParameters(t *testing.T) {
	d := OBDLiveData{
		Timestamp:  "2025-01-01T00:00:00Z",
		Monitoring: true,
		Parameters: map[string]PIDValue{
			"rpm":   {Value: 2000, Unit: "RPM", Success: true},
			"speed": {Value: 60, Unit: "km/h", Success: true},
		},
	}
	b, _ := json.Marshal(d)
	s := string(b)

	mustContain(t, s, `"monitoring":true`)
	mustContain(t, s, `"parameters"`)
	mustContain(t, s, `"rpm"`)
	mustContain(t, s, `"speed"`)
}

func TestWarningsData_JSON(t *testing.T) {
	d := WarningsData{
		Timestamp: "2025-01-01T00:00:00Z",
		Warnings: map[string]bool{
			"turnSignals": true,
			"battery":     false,
		},
		Changes: []WarningChange{
			{Warning: "turnSignals", State: true, Pin: 17, Name: "Frecce"},
		},
	}
	b, _ := json.Marshal(d)
	s := string(b)

	mustContain(t, s, `"warnings"`)
	mustContain(t, s, `"turnSignals":true`)
	mustContain(t, s, `"battery":false`)
	mustContain(t, s, `"changes"`)
	mustContain(t, s, `"pin":17`)
}

func TestTemperatureData_JSON(t *testing.T) {
	var d TemperatureData
	d.Timestamp = "2025-01-01T00:00:00Z"
	d.Temperature.External = 21.5
	d.Temperature.Unit = "°C"
	d.Temperature.Sensor = "DS18B20"

	b, _ := json.Marshal(d)
	s := string(b)

	mustContain(t, s, `"external":21.5`)
	mustContain(t, s, `"sensor":"DS18B20"`)
}

func TestFuelData_JSON(t *testing.T) {
	var d FuelData
	d.Timestamp = "2025-01-01T00:00:00Z"
	d.Fuel.Level = 75.3
	d.Fuel.Unit = "%"
	d.Fuel.Sensor = "ADS1115"
	d.Fuel.Raw.ADC = 16383
	d.Fuel.Raw.Voltage = 2.048

	b, _ := json.Marshal(d)
	s := string(b)

	mustContain(t, s, `"level":75.3`)
	mustContain(t, s, `"sensor":"ADS1115"`)
	mustContain(t, s, `"adc":16383`)
}

func TestIgnitionData_JSON(t *testing.T) {
	var d IgnitionData
	d.Timestamp = "2025-01-01T00:00:00Z"
	d.Ignition.On = true
	d.Ignition.State = "ON"

	b, _ := json.Marshal(d)
	s := string(b)

	mustContain(t, s, `"on":true`)
	mustContain(t, s, `"state":"ON"`)
}

// ---- benchmarks ----

func BenchmarkMarshalOBDLiveData(b *testing.B) {
	d := OBDLiveData{
		Timestamp:  "2025-01-01T00:00:00.000000000Z",
		Monitoring: true,
		Parameters: map[string]PIDValue{
			"rpm":          {Value: 2000, Unit: "RPM", Success: true},
			"speed":        {Value: 60, Unit: "km/h", Success: true},
			"coolant_temp": {Value: 90, Unit: "°C", Success: true},
		},
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		json.Marshal(d)
	}
}

// ---- helper ----

func mustContain(t *testing.T, s, sub string) {
	t.Helper()
	if !strings.Contains(s, sub) {
		t.Errorf("JSON missing %q\nFull output: %s", sub, s)
	}
}
