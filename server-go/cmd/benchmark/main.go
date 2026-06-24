/*
 * PandaOS — Benchmark Server
 * Copyright (C) 2025  Cyberpandino
 *
 * Mirrors server/BenchmarkServer.js: animates RPM 800→5500 at 20 Hz and
 * broadcasts `obd-live` events so the React client can be stress-tested
 * without real OBD hardware.
 */

package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"net/http"
	"sync"
	"time"

	"github.com/zishang520/socket.io/v2/socket"
)

const addr = ":3002"

// animState holds the shared RPM / speed animation shared across all clients,
// matching the JS global variables `rpm`, `speed`, `direction`.
type animState struct {
	mu        sync.Mutex
	rpm       float64
	speed     float64
	direction float64
}

func (a *animState) tick() (rpm float64, speed float64) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.rpm += 80 * a.direction
	a.speed += 0.8 * a.direction

	if a.rpm > 5500 {
		a.direction = -1
	}
	if a.rpm < 800 {
		a.direction = 1
	}
	return a.rpm, a.speed
}

var anim = &animState{rpm: 800, speed: 0, direction: 1}

// pidParam matches the JSON shape expected by the client's `obd-live` handler.
type pidParam struct {
	Value   interface{} `json:"value"`
	Success bool        `json:"success"`
}

type obdPacket struct {
	Parameters map[string]pidParam `json:"parameters"`
}

// clientRegistry maps socket IDs to per-client cancel funcs.
type clientRegistry struct {
	mu      sync.Mutex
	cancels map[socket.SocketId]context.CancelFunc
}

func (r *clientRegistry) add(id socket.SocketId, cancel context.CancelFunc) {
	r.mu.Lock()
	r.cancels[id] = cancel
	r.mu.Unlock()
}

func (r *clientRegistry) remove(id socket.SocketId) {
	r.mu.Lock()
	if cancel, ok := r.cancels[id]; ok {
		cancel()
		delete(r.cancels, id)
	}
	r.mu.Unlock()
}

var clients = &clientRegistry{cancels: make(map[socket.SocketId]context.CancelFunc)}

func main() {
	io := socket.NewServer(nil, nil)

	io.On("connection", func(args ...any) {
		s := args[0].(*socket.Socket)
		log.Printf("CLIENT RICONOSCIUTO: %s", s.Id())
		log.Println("Inizio invio dati ...")

		ctx, cancel := context.WithCancel(context.Background())
		clients.add(s.Id(), cancel)

		go runClientLoop(ctx, s)

		s.On("disconnect", func(reason ...any) {
			fmt.Printf("\nClient disconnesso: %s\n", s.Id())
			clients.remove(s.Id())
		})
	})

	mux := http.NewServeMux()
	mux.Handle("/socket.io/", corsMiddleware(io.ServeHandler(nil)))
	mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "Panda Benchmark Server Active")
	})

	fmt.Printf("BENCHMARK SERVER (GO) SU PORTA %s\n", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}

// runClientLoop emits `obd-live` events to s at 20 Hz until ctx is cancelled.
func runClientLoop(ctx context.Context, s *socket.Socket) {
	ticker := time.NewTicker(50 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			rpm, speed := anim.tick()

			if rpm > 5400 || rpm < 850 {
				fmt.Print(".")
			}

			pkt := obdPacket{
				Parameters: map[string]pidParam{
					"rpm":             {Value: math.Round(rpm), Success: true},
					"speed":           {Value: int(math.Abs(speed)), Success: true},
					"coolant_temp":    {Value: 90, Success: true},
					"engine_load":     {Value: 25, Success: true},
					"battery_voltage": {Value: 13.8, Success: true},
				},
			}
			if err := s.Emit("obd-live", pkt); err != nil {
				log.Printf("emit error: %v", err)
				return
			}
		}
	}
}

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
