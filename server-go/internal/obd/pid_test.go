/*
 * PandaOS
 * Copyright (C) 2025  Cyberpandino
 */

package obd

import (
	"encoding/json"
	"strings"
	"testing"
)

// ---- parseBytes ----

func TestParseBytes_Spaced(t *testing.T) {
	got := parseBytes("41 0C 1A F8")
	want := []string{"41", "0C", "1A", "F8"}
	if !strSliceEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestParseBytes_NoSpaces(t *testing.T) {
	got := parseBytes("410C1AF8")
	want := []string{"41", "0C", "1A", "F8"}
	if !strSliceEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestParseBytes_FiltersShortTokens(t *testing.T) {
	// Tokens not exactly 2 chars should be dropped when space-delimited.
	got := parseBytes("41 0 0C")
	want := []string{"41", "0C"}
	if !strSliceEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestParseBytes_Empty(t *testing.T) {
	if got := parseBytes(""); len(got) != 0 {
		t.Errorf("expected empty, got %v", got)
	}
}

// ---- hexByte ----

func TestHexByte(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"00", 0},
		{"0C", 12},
		{"FF", 255},
		{"1A", 26},
		{"69", 105},
	}
	for _, c := range cases {
		if got := hexByte(c.in); got != c.want {
			t.Errorf("hexByte(%q) = %d, want %d", c.in, got, c.want)
		}
	}
}

// ---- round helpers ----

func TestRoundHelpers(t *testing.T) {
	if round1(50.196) != 50.2 {
		t.Errorf("round1 wrong")
	}
	if round2(3.14159) != 3.14 {
		t.Errorf("round2 wrong")
	}
	if round3(1.23456) != 1.235 {
		t.Errorf("round3 wrong")
	}
}

// ---- ParseResponse error paths ----

func TestParseResponse_NoData(t *testing.T) {
	def := PIDDef{PID: "010C", Name: "RPM", Key: "rpm"}
	r := ParseResponse(def, "NO DATA")
	assertErrorResult(t, r, "NON_SUPPORTATO")
}

func TestParseResponse_Empty(t *testing.T) {
	def := PIDDef{PID: "010C", Name: "RPM", Key: "rpm"}
	r := ParseResponse(def, "")
	assertErrorResult(t, r, "NON_SUPPORTATO")
}

func TestParseResponse_Error(t *testing.T) {
	def := PIDDef{PID: "010C", Name: "RPM", Key: "rpm"}
	r := ParseResponse(def, "UNABLE TO CONNECT ERROR")
	assertErrorResult(t, r, "NON_SUPPORTATO")
}

func TestParseResponse_Stopped(t *testing.T) {
	def := PIDDef{PID: "010C", Name: "RPM", Key: "rpm"}
	r := ParseResponse(def, "STOPPED")
	assertErrorResult(t, r, "ECU_DORMIENTE")
}

func TestParseResponse_7F(t *testing.T) {
	def := PIDDef{PID: "010C", Name: "RPM", Key: "rpm"}
	r := ParseResponse(def, "7F 0C 11")
	// parsed bytes: ["7F", "0C", "11"] — first byte is not "41"
	if r.Value != "ECU_DORMIENTE" && r.Value != "FORMATO_ERRATO" {
		t.Errorf("unexpected value %q", r.Value)
	}
	if r.Success {
		t.Errorf("expected !Success")
	}
}

func TestParseResponse_BadFormat(t *testing.T) {
	def := PIDDef{PID: "010C", Name: "RPM", Key: "rpm"}
	r := ParseResponse(def, "XX YY")
	assertErrorResult(t, r, "FORMATO_ERRATO")
}

// ---- ParseResponse happy paths (table-driven) ----

func TestParseResponse_Table(t *testing.T) {
	tests := []struct {
		name     string
		pid      string
		response string
		wantVal  interface{}
		wantUnit string
	}{
		// 010C — RPM: (0x1A*256 + 0xF8) / 4 = 6904/4 = 1726
		{"rpm_1726", "010C", "41 0C 1A F8", 1726, "RPM"},
		// Minimum RPM (800): 800*4 = 3200 = 0x0C80
		{"rpm_800", "010C", "41 0C 0C 80", 800, "RPM"},

		// 010D — Speed: 0x3C = 60 km/h
		{"speed_60", "010D", "41 0D 3C", 60, "km/h"},
		{"speed_0", "010D", "41 0D 00", 0, "km/h"},
		{"speed_255", "010D", "41 0D FF", 255, "km/h"},

		// 0105 — Coolant temp: 0x69 - 40 = 105 - 40 = 65°C
		{"coolant_65", "0105", "41 05 69", 65, "°C"},
		{"coolant_minus40", "0105", "41 05 00", -40, "°C"},
		{"coolant_215", "0105", "41 05 FF", 215, "°C"},

		// 010F — Air temp: same formula as coolant
		{"air_temp_25", "010F", "41 0F 41", 25, "°C"},

		// 0104 — Engine load: 0x80 * 100 / 255 = 50.2%
		{"engine_load_50pct", "0104", "41 04 80", 50.2, "%"},
		{"engine_load_0", "0104", "41 04 00", 0.0, "%"},
		{"engine_load_100", "0104", "41 04 FF", 100.0, "%"},

		// 0111 — Throttle: round(0x80 * 100 / 255) = 50
		{"throttle_50", "0111", "41 11 80", 50, "%"},
		{"throttle_0", "0111", "41 11 00", 0, "%"},
		{"throttle_100", "0111", "41 11 FF", 100, "%"},

		// 010E — Timing advance: (0x80 - 128) / 2 = 0.0°
		{"timing_zero", "010E", "41 0E 80", 0.0, "°"},
		// 0xA0 = 160: (160-128)/2 = 16.0°
		{"timing_16", "010E", "41 0E A0", 16.0, "°"},

		// 0106 — Short fuel trim: (0x80 - 128) * 100 / 128 = 0%
		{"fuel_trim_zero", "0106", "41 06 80", 0.0, "%"},
		// 0xFF=255: (255-128)*100/128 = 99.2%
		{"fuel_trim_max", "0106", "41 06 FF", 99.2, "%"},

		// 010B — MAP pressure: 0x60 = 96 kPa
		{"map_96kpa", "010B", "41 0B 60", 96, "kPa"},

		// 011C — OBD standards
		{"obd_std_eobd", "011C", "41 1C 06", "EOBD (Europe)", "standard"},

		// 0121 — Distance with MIL on: 0x01 0x F4 = 500 km
		{"distance_500", "0121", "41 21 01 F4", 500, "km"},

		// Unknown PID → hex dump
		{"unknown_pid", "0199", "41 99 AB CD", "41 99 AB CD", "hex"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			def := PIDDef{PID: tt.pid, Name: tt.name, Key: tt.name}
			r := ParseResponse(def, tt.response)
			if !r.Success {
				t.Fatalf("expected Success=true, got Error=%v Value=%v", r.Error, r.Value)
			}
			if r.Unit != tt.wantUnit {
				t.Errorf("unit: got %q, want %q", r.Unit, tt.wantUnit)
			}
			// Compare via JSON to handle int vs float64 from interface{}.
			gotJSON, _ := json.Marshal(r.Value)
			wantJSON, _ := json.Marshal(tt.wantVal)
			if string(gotJSON) != string(wantJSON) {
				t.Errorf("value: got %s, want %s", gotJSON, wantJSON)
			}
		})
	}
}

// ---- ParseResponse — oxygen sensor (multi-byte struct value) ----

func TestParseResponse_OxygenSensor(t *testing.T) {
	def := PIDDef{PID: "0114", Name: "O2 Bank1 S1", Key: "oxygen_sensor_1_1"}
	// 41 14 80 FF → voltage=0x80/200=0.640, fuelTrim=(0xFF-128)*100/128=99.2
	r := ParseResponse(def, "41 14 80 FF")
	if !r.Success {
		t.Fatalf("expected success")
	}
	m, ok := r.Value.(map[string]interface{})
	if !ok {
		t.Fatalf("expected map value, got %T", r.Value)
	}
	if m["voltage"] == nil || m["fuelTrim"] == nil {
		t.Errorf("missing keys in %v", m)
	}
}

// ---- ParseResponse — monitor status struct value ----

func TestParseResponse_MonitorStatus(t *testing.T) {
	def := PIDDef{PID: "0101", Name: "Monitor Status", Key: "monitor_status"}
	// MIL off (bit7=0), 3 DTCs: 0x03
	r := ParseResponse(def, "41 01 03 00 00 00")
	if !r.Success {
		t.Fatalf("expected success")
	}
	m, ok := r.Value.(map[string]interface{})
	if !ok {
		t.Fatalf("expected map value, got %T", r.Value)
	}
	if m["mil"] != false {
		t.Errorf("MIL should be off for 0x03")
	}
}

// ---- Definitions ----

func TestDefinitions_NonEmpty(t *testing.T) {
	defs := Definitions()
	if len(defs) == 0 {
		t.Fatal("expected at least one PID definition")
	}
	for _, d := range defs {
		if d.PID == "" || d.Name == "" || d.Key == "" {
			t.Errorf("incomplete PID definition: %+v", d)
		}
		if len(d.PID) != 4 {
			t.Errorf("PID %q should be 4 chars", d.PID)
		}
	}
}

func TestDefinitions_NoDuplicateKeys(t *testing.T) {
	seen := make(map[string]bool)
	for _, d := range Definitions() {
		if seen[d.Key] {
			t.Errorf("duplicate key %q", d.Key)
		}
		seen[d.Key] = true
	}
}

func TestDefinitions_NoDuplicatePIDs(t *testing.T) {
	seen := make(map[string]bool)
	for _, d := range Definitions() {
		if seen[d.PID] {
			t.Errorf("duplicate PID %q", d.PID)
		}
		seen[d.PID] = true
	}
}

// ---- IsMonitorKey ----

func TestIsMonitorKey(t *testing.T) {
	must := []string{"coolant_temp", "rpm", "speed", "throttle_pos",
		"manifold_pressure", "engine_load", "timing_advance",
		"air_temp", "short_fuel_trim_1", "long_fuel_trim_1"}
	mustNot := []string{"supported_pids_01_20", "monitor_status",
		"oxygen_sensor_1_1", "obd_standards"}

	for _, k := range must {
		if !IsMonitorKey(k) {
			t.Errorf("expected %q to be a monitor key", k)
		}
	}
	for _, k := range mustNot {
		if IsMonitorKey(k) {
			t.Errorf("expected %q to NOT be a monitor key", k)
		}
	}
}

// ---- decode helpers ----

func TestDecodeFuelSystemStatus(t *testing.T) {
	s := decodeFuelSystemStatus(0x02)
	if !strings.Contains(s, "Closed loop") {
		t.Errorf("unexpected: %s", s)
	}
	unknown := decodeFuelSystemStatus(0xFF)
	if !strings.Contains(unknown, "Unknown") {
		t.Errorf("expected Unknown for 0xFF, got %s", unknown)
	}
}

func TestDecodeOBDStandards(t *testing.T) {
	s := decodeOBDStandards(0x06)
	if s != "EOBD (Europe)" {
		t.Errorf("got %q", s)
	}
	unknown := decodeOBDStandards(0x7F)
	if !strings.Contains(unknown, "Unknown") {
		t.Errorf("expected Unknown for 0x7F, got %s", unknown)
	}
}

// ---- benchmarks ----

func BenchmarkParseResponse_RPM(b *testing.B) {
	def := PIDDef{PID: "010C", Name: "RPM", Key: "rpm"}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ParseResponse(def, "41 0C 1A F8")
	}
}

func BenchmarkParseResponse_Speed(b *testing.B) {
	def := PIDDef{PID: "010D", Name: "Speed", Key: "speed"}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ParseResponse(def, "41 0D 3C")
	}
}

func BenchmarkParseBytes_Spaced(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		parseBytes("41 0C 1A F8 AB CD EF 00")
	}
}

func BenchmarkParseBytes_NoSpaces(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		parseBytes("410C1AF8ABCDEF00")
	}
}

// ---- helpers ----

func assertErrorResult(t *testing.T, r PIDResult, wantVal interface{}) {
	t.Helper()
	if r.Success {
		t.Errorf("expected Success=false")
	}
	if !r.Error {
		t.Errorf("expected Error=true")
	}
	if r.Value != wantVal {
		t.Errorf("value: got %v, want %v", r.Value, wantVal)
	}
}

func strSliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
