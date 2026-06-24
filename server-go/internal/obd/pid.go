/*
 * PandaOS
 * Copyright (C) 2025  Cyberpandino
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3.
 */

// Package obd implements ELM327 serial communication and OBD-II PID parsing.
package obd

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

// PIDDef is a single OBD-II PID definition.
type PIDDef struct {
	PID  string // 4-char hex, e.g. "010C"
	Name string
	Key  string // matches frontend parameter key, e.g. "rpm"
}

// PIDResult is the parsed output of a single PID read.
type PIDResult struct {
	PID       string
	Name      string
	Key       string
	Value     any
	Unit      string
	Raw       string
	Success   bool
	Error     bool
	Timestamp string
}

// pidDefinitions is the list of PIDs we attempt to read from the ECU.
var pidDefinitions = []PIDDef{
	{"0100", "PID Supportati (01-20)", "supported_pids_01_20"},
	{"0101", "Monitor Status", "monitor_status"},
	{"0103", "Fuel System Status", "fuel_system_status"},
	{"0104", "Calculated Engine Load", "engine_load"},
	{"0105", "Temperatura Refrigerante", "coolant_temp"},
	{"0106", "Short Term Fuel Trim Bank 1", "short_fuel_trim_1"},
	{"0107", "Long Term Fuel Trim Bank 1", "long_fuel_trim_1"},
	{"010B", "Pressione Collettore", "manifold_pressure"},
	{"010C", "Regime Motore", "rpm"},
	{"010D", "Velocità Veicolo", "speed"},
	{"010E", "Timing Advance", "timing_advance"},
	{"010F", "Temperatura Aria", "air_temp"},
	{"0111", "Posizione Farfalla", "throttle_pos"},
	{"0113", "Oxygen Sensors Present", "oxygen_sensors_present"},
	{"0114", "Oxygen Sensor 1 Bank 1", "oxygen_sensor_1_1"},
	{"0115", "Oxygen Sensor 2 Bank 1", "oxygen_sensor_2_1"},
	{"011C", "OBD Standards", "obd_standards"},
	{"0120", "PID Supportati (21-40)", "supported_pids_21_40"},
	{"0121", "Distance with MIL On", "distance_with_mil_on"},
}

// monitorKeys is the subset of PID keys used for continuous monitoring.
var monitorKeys = map[string]bool{
	"coolant_temp":       true,
	"rpm":               true,
	"speed":             true,
	"throttle_pos":      true,
	"manifold_pressure": true,
	"engine_load":       true,
	"timing_advance":    true,
	"air_temp":          true,
	"short_fuel_trim_1": true,
	"long_fuel_trim_1":  true,
}

// Definitions returns the full ordered PID list.
func Definitions() []PIDDef { return pidDefinitions }

// IsMonitorKey reports whether a PID key belongs to the continuous-monitoring set.
func IsMonitorKey(key string) bool { return monitorKeys[key] }

// ParseResponse converts the raw ELM327 response string into a PIDResult.
func ParseResponse(def PIDDef, response string) PIDResult {
	r := PIDResult{
		PID:       def.PID,
		Name:      def.Name,
		Key:       def.Key,
		Raw:       response,
		Timestamp: time.Now().UTC().Format(time.RFC3339Nano),
	}

	if response == "" ||
		strings.Contains(response, "NO DATA") ||
		strings.Contains(response, "ERROR") {
		r.Error = true
		r.Value = "NON_SUPPORTATO"
		return r
	}
	if strings.Contains(response, "STOPPED") || strings.Contains(response, "7F") {
		r.Error = true
		r.Value = "ECU_DORMIENTE"
		return r
	}

	clean := strings.ReplaceAll(response, ">", "")
	clean = strings.TrimSpace(clean)
	bytes_ := parseBytes(clean)

	if len(bytes_) < 2 || strings.ToUpper(bytes_[0]) != "41" {
		r.Error = true
		r.Value = "FORMATO_ERRATO"
		return r
	}

	r.Success = true
	parsePIDValue(def.PID, bytes_, &r)
	return r
}

func parseBytes(s string) []string {
	if strings.Contains(s, " ") {
		var out []string
		for _, b := range strings.Fields(s) {
			if len(b) == 2 {
				out = append(out, b)
			}
		}
		return out
	}
	var out []string
	for i := 0; i+1 < len(s); i += 2 {
		out = append(out, s[i:i+2])
	}
	return out
}

func hexByte(b string) int {
	v, _ := strconv.ParseInt(b, 16, 32)
	return int(v)
}

func parsePIDValue(pid string, b []string, r *PIDResult) {
	pid = strings.ToUpper(pid)
	switch pid {
	case "0100", "0120", "0140":
		r.Value = fmt.Sprintf("%d PID disponibili", len(b)-2)
		r.Unit = "bitmap"

	case "0101":
		if len(b) >= 3 {
			v := hexByte(b[2])
			r.Value = map[string]any{
				"mil":      (v & 0x80) != 0,
				"dtcCount": v & 0x7F,
			}
			r.Unit = "status"
		}

	case "0103":
		if len(b) >= 3 {
			r.Value = decodeFuelSystemStatus(hexByte(b[2]))
			r.Unit = "status"
		}

	case "0104":
		if len(b) >= 3 {
			r.Value = round1(float64(hexByte(b[2])) * 100 / 255)
			r.Unit = "%"
		}

	case "0105":
		if len(b) >= 3 {
			r.Value = hexByte(b[2]) - 40
			r.Unit = "°C"
		}

	case "0106", "0107", "0108", "0109":
		if len(b) >= 3 {
			r.Value = round1((float64(hexByte(b[2])-128) * 100) / 128)
			r.Unit = "%"
		}

	case "010A":
		if len(b) >= 3 {
			r.Value = hexByte(b[2]) * 3
			r.Unit = "kPa"
		}

	case "010B":
		if len(b) >= 3 {
			r.Value = hexByte(b[2])
			r.Unit = "kPa"
		}

	case "010C":
		if len(b) >= 4 {
			r.Value = int(math.Round(float64(hexByte(b[2])*256+hexByte(b[3])) / 4))
			r.Unit = "RPM"
		}

	case "010D":
		if len(b) >= 3 {
			r.Value = hexByte(b[2])
			r.Unit = "km/h"
		}

	case "010E":
		if len(b) >= 3 {
			r.Value = round1(float64(hexByte(b[2])-128) / 2)
			r.Unit = "°"
		}

	case "010F":
		if len(b) >= 3 {
			r.Value = hexByte(b[2]) - 40
			r.Unit = "°C"
		}

	case "0110":
		if len(b) >= 4 {
			r.Value = round2(float64(hexByte(b[2])*256+hexByte(b[3])) / 100)
			r.Unit = "g/s"
		}

	case "0111":
		if len(b) >= 3 {
			r.Value = int(math.Round(float64(hexByte(b[2])) * 100 / 255))
			r.Unit = "%"
		}

	case "0113":
		if len(b) >= 3 {
			r.Value = fmt.Sprintf("%08b", hexByte(b[2]))
			r.Unit = "bitmap"
		}

	case "0114", "0115", "0116", "0117", "0118", "0119", "011A", "011B":
		if len(b) >= 4 {
			r.Value = map[string]any{
				"voltage":  round3(float64(hexByte(b[2])) / 200),
				"fuelTrim": round1(float64(hexByte(b[3])-128) * 100 / 128),
			}
			r.Unit = "V, %"
		}

	case "011C":
		if len(b) >= 3 {
			r.Value = decodeOBDStandards(hexByte(b[2]))
			r.Unit = "standard"
		}

	case "011F":
		if len(b) >= 4 {
			r.Value = hexByte(b[2])*256 + hexByte(b[3])
			r.Unit = "seconds"
		}

	case "0121":
		if len(b) >= 4 {
			r.Value = hexByte(b[2])*256 + hexByte(b[3])
			r.Unit = "km"
		}

	case "0142":
		if len(b) >= 4 {
			r.Value = round3(float64(hexByte(b[2])*256+hexByte(b[3])) / 1000)
			r.Unit = "V"
		}

	case "0143":
		if len(b) >= 4 {
			r.Value = round1(float64(hexByte(b[2])*256+hexByte(b[3])) * 100 / 255)
			r.Unit = "%"
		}

	case "0145":
		if len(b) >= 3 {
			r.Value = round1(float64(hexByte(b[2])) * 100 / 255)
			r.Unit = "%"
		}

	case "0146":
		if len(b) >= 3 {
			r.Value = hexByte(b[2]) - 40
			r.Unit = "°C"
		}

	default:
		parts := make([]string, len(b))
		copy(parts, b)
		r.Value = strings.Join(parts, " ")
		r.Unit = "hex"
	}
}

func round1(f float64) float64 { return math.Round(f*10) / 10 }
func round2(f float64) float64 { return math.Round(f*100) / 100 }
func round3(f float64) float64 { return math.Round(f*1000) / 1000 }

func decodeFuelSystemStatus(s int) string {
	statuses := map[int]string{
		0x01: "Open loop due to insufficient engine temperature",
		0x02: "Closed loop, using oxygen sensor feedback",
		0x04: "Open loop due to engine load OR fuel cut due to deceleration",
		0x08: "Open loop due to system failure",
		0x10: "Closed loop, using at least one oxygen sensor but there is a fault",
	}
	if v, ok := statuses[s]; ok {
		return v
	}
	return fmt.Sprintf("Unknown (0x%02X)", s)
}

func decodeOBDStandards(s int) string {
	standards := map[int]string{
		0x01: "OBD-II as defined by the CARB",
		0x02: "OBD as defined by the EPA",
		0x03: "OBD and OBD-II",
		0x04: "OBD-I",
		0x05: "Not OBD compliant",
		0x06: "EOBD (Europe)",
		0x07: "EOBD and OBD-II",
		0x08: "EOBD and OBD",
		0x09: "EOBD, OBD and OBD II",
		0x0A: "JOBD (Japan)",
		0x0B: "JOBD and OBD II",
		0x0C: "JOBD and EOBD",
		0x0D: "JOBD, EOBD, and OBD II",
	}
	if v, ok := standards[s]; ok {
		return v
	}
	return fmt.Sprintf("Unknown (0x%02X)", s)
}
