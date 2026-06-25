/*
 * PandaOS — Stress Server
 * Copyright (C) 2025  Cyberpandino
 *
 * Configurable via environment variables:
 *   TICK_MS  – interval between emits in ms; 0 = uncapped (default 50)
 *   PAYLOAD  – "small" (5 params) or "full" (17 params, default "small")
 *   PORT     – listen port (default 3002)
 */

package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/zishang520/socket.io/v2/socket"
)

// ── config ────────────────────────────────────────────────────────────────────

var (
	tickMs      = envInt("TICK_MS", 50)
	payloadMode = envStr("PAYLOAD", "small")
	port        = envStr("PORT", "3002")
)

// ── animated state (same sine-wave across all clients) ────────────────────────

type state struct {
	mu        sync.Mutex
	startTime time.Time
}

var sim = &state{startTime: time.Now()}

// values returns a time-based snapshot so values animate smoothly.
func (s *state) values() map[string]any {
	s.mu.Lock()
	t := time.Since(s.startTime).Seconds()
	s.mu.Unlock()

	// slow 10-second RPM sweep 800→4500
	rpm := 800 + 3700*(0.5+0.5*math.Sin(t*2*math.Pi/10))
	speed := rpm / 60
	load := 20 + 60*(0.5+0.5*math.Sin(t*2*math.Pi/7))
	coolant := 85 + 10*math.Sin(t*2*math.Pi/60)
	throttle := 5 + 40*(0.5+0.5*math.Sin(t*2*math.Pi/8))
	manifold := 30 + 50*(0.5+0.5*math.Sin(t*2*math.Pi/9))
	timing := 8 + 12*(0.5+0.5*math.Sin(t*2*math.Pi/12))
	airTemp := 22 + 3*math.Sin(t*2*math.Pi/120)
	stFuel := -3 + 4*math.Sin(t*2*math.Pi/15)
	ltFuel := 2 + 2*math.Sin(t*2*math.Pi/90)

	small := map[string]any{
		"rpm":          round1(rpm),
		"speed":        int(speed),
		"coolant_temp": round1(coolant),
		"engine_load":  round1(load),
		"battery_voltage": 13.8,
	}
	if payloadMode != "full" {
		return small
	}

	// full: all 17 non-discovery PIDs
	full := map[string]any{
		"rpm":                    round1(rpm),
		"speed":                  int(speed),
		"coolant_temp":           round1(coolant),
		"engine_load":            round1(load),
		"battery_voltage":        13.8,
		"throttle_pos":           round1(throttle),
		"manifold_pressure":      int(manifold),
		"timing_advance":         round1(timing),
		"air_temp":               round1(airTemp),
		"short_fuel_trim_1":      round2(stFuel),
		"long_fuel_trim_1":       round2(ltFuel),
		"monitor_status":         0,
		"fuel_system_status":     "CL",
		"oxygen_sensors_present": 3,
		"oxygen_sensor_1_1":      round2(0.72 + 0.1*math.Sin(t*2*math.Pi/3)),
		"oxygen_sensor_2_1":      round2(0.68 + 0.1*math.Sin(t*2*math.Pi/4)),
		"obd_standards":          "OBD-II",
		"distance_with_mil_on":   int(t * 0.5),
	}
	return full
}

// ── counters ──────────────────────────────────────────────────────────────────

var totalEmitted atomic.Int64

// ── main ─────────────────────────────────────────────────────────────────────

func main() {
	io := socket.NewServer(nil, nil)

	var regMu sync.Mutex
	cancels := map[socket.SocketId]context.CancelFunc{}

	io.On("connection", func(args ...any) {
		s := args[0].(*socket.Socket)
		log.Printf("+ %s  goroutines=%d", s.Id(), runtime.NumGoroutine())

		ctx, cancel := context.WithCancel(context.Background())
		regMu.Lock()
		cancels[s.Id()] = cancel
		regMu.Unlock()

		go clientLoop(ctx, s)

		s.On("disconnect", func(...any) {
			regMu.Lock()
			if c, ok := cancels[s.Id()]; ok {
				c()
				delete(cancels, s.Id())
			}
			regMu.Unlock()
		})
	})

	mux := http.NewServeMux()
	mux.Handle("/socket.io/", corsMiddleware(io.ServeHandler(nil)))

	// /stats returns live server metrics for the benchmark harness.
	mux.HandleFunc("/stats", func(w http.ResponseWriter, _ *http.Request) {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"goroutines":%d,"heap_alloc_kb":%d,"rss_kb":%d,"total_emitted":%d}`,
			runtime.NumGoroutine(),
			m.HeapAlloc/1024,
			m.Sys/1024,
			totalEmitted.Load(),
		)
	})

	fmt.Printf("Go stress server :%s  tick=%dms  payload=%s  GOMAXPROCS=%d\n",
		port, tickMs, payloadMode, runtime.GOMAXPROCS(0))
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}

// ── client loop ───────────────────────────────────────────────────────────────

func clientLoop(ctx context.Context, s *socket.Socket) {
	emit := func() bool {
		pkt := map[string]any{"parameters": buildParams()}
		if err := s.Emit("obd-live", pkt); err != nil {
			return false
		}
		totalEmitted.Add(1)
		return true
	}

	if tickMs <= 0 {
		// Uncapped: send as fast as socket allows; yield between sends
		// so other goroutines (including other clients) get CPU time.
		for {
			select {
			case <-ctx.Done():
				return
			default:
				if !emit() {
					return
				}
				runtime.Gosched()
			}
		}
	}

	ticker := time.NewTicker(time.Duration(tickMs) * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if !emit() {
				return
			}
		}
	}
}

// buildParams wraps each value in {value, success} as the client expects.
func buildParams() map[string]any {
	raw := sim.values()
	out := make(map[string]any, len(raw))
	for k, v := range raw {
		out[k] = map[string]any{"value": v, "success": true}
	}
	return out
}

// ── helpers ───────────────────────────────────────────────────────────────────

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func envInt(key string, def int) int {
	if v := os.Getenv(key); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			return n
		}
	}
	return def
}

func envStr(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func round1(v float64) float64 { return math.Round(v*10) / 10 }
func round2(v float64) float64 { return math.Round(v*100) / 100 }
