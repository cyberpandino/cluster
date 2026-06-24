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
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"go.bug.st/serial"
)

const (
	portPath = "/dev/ttyUSB0"
	baudRate = 38400
	// lineBuffer is the number of lines we keep queued from the ELM327.
	lineBuffer = 64
)

// Comm manages the serial connection to the ELM327 adapter.
type Comm struct {
	mu           sync.Mutex
	port         serial.Port
	lines        chan string // lines received from ELM327, split on '\r'
	connected    bool
	lastActivity time.Time
}

// NewComm creates an unconnected Comm. Call Connect before anything else.
func NewComm() *Comm {
	return &Comm{lines: make(chan string, lineBuffer)}
}

// Connect opens the serial port and starts the read loop.
func (c *Comm) Connect(ctx context.Context) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Check port exists before opening.
	ports, err := serial.GetPortsList()
	if err != nil {
		return fmt.Errorf("listing serial ports: %w", err)
	}
	found := false
	for _, p := range ports {
		if p == portPath {
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("porta %s non trovata", portPath)
	}

	mode := &serial.Mode{BaudRate: baudRate}
	port, err := serial.Open(portPath, mode)
	if err != nil {
		return fmt.Errorf("aprendo %s: %w", portPath, err)
	}
	if err := port.SetReadTimeout(100 * time.Millisecond); err != nil {
		port.Close()
		return fmt.Errorf("set read timeout: %w", err)
	}

	c.port = port
	c.connected = true
	c.lastActivity = time.Now()

	log.Println("ELM327 connesso alla porta seriale")

	go c.readLoop(ctx)
	return nil
}

// readLoop reads bytes from the port, splits on '\r', and feeds non-empty,
// non-prompt lines into c.lines. Runs until ctx is cancelled or port closes.
func (c *Comm) readLoop(ctx context.Context) {
	buf := make([]byte, 1)
	var current []byte

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		n, err := c.port.Read(buf)
		if err != nil {
			c.mu.Lock()
			c.connected = false
			c.mu.Unlock()
			log.Printf("Errore lettura porta seriale: %v", err)
			return
		}
		if n == 0 {
			continue
		}

		ch := buf[0]
		if ch == '\r' || ch == '\n' {
			line := strings.TrimSpace(string(current))
			current = current[:0]
			if line == "" || line == ">" {
				continue
			}
			select {
			case c.lines <- line:
			default:
				// Drop oldest line to prevent stale data.
				select {
				case <-c.lines:
				default:
				}
				c.lines <- line
			}
		} else {
			current = append(current, ch)
		}
	}
}

// drain clears any pending lines before sending a new command.
func (c *Comm) drain() {
	for {
		select {
		case <-c.lines:
		default:
			return
		}
	}
}

// SendCommand writes cmd + CR to the ELM327 and clears the response buffer.
func (c *Comm) SendCommand(cmd string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if !c.connected {
		return errors.New("porta non aperta")
	}
	c.drain()
	_, err := c.port.Write([]byte(cmd + "\r"))
	if err != nil {
		return fmt.Errorf("scrittura comando %q: %w", cmd, err)
	}
	c.lastActivity = time.Now()
	return nil
}

// WaitForResponse blocks until a valid (non-SEARCHING, non-prompt) line arrives
// or until timeout elapses. ctx cancellation also terminates the wait.
func (c *Comm) WaitForResponse(ctx context.Context, timeout time.Duration) (string, error) {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		case line := <-c.lines:
			if strings.Contains(line, "SEARCHING") {
				// ECU is still looking for the protocol; give it a second.
				select {
				case <-ctx.Done():
					return "", ctx.Err()
				case <-time.After(time.Second):
				}
				continue
			}
			c.mu.Lock()
			c.lastActivity = time.Now()
			c.mu.Unlock()
			return line, nil
		case <-time.After(100 * time.Millisecond):
			// No data yet; check connection and loop.
			c.mu.Lock()
			conn := c.connected
			c.mu.Unlock()
			if !conn {
				return "", errors.New("connessione persa durante attesa risposta")
			}
		}
	}
	return "", errors.New("timeout: nessuna risposta ricevuta")
}

// Initialize sends the standard ELM327 AT initialisation sequence.
func (c *Comm) Initialize(ctx context.Context) error {
	log.Println("Inizializzazione ELM327...")

	if err := c.SendCommand("ATZ"); err != nil {
		return err
	}
	if _, err := c.WaitForResponse(ctx, 4*time.Second); err != nil {
		log.Printf("Reset ELM327 parziale: %v", err)
	} else {
		log.Println("Reset ELM327 completato")
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(2 * time.Second):
	}

	cmds := []struct{ cmd, desc string }{
		{"ATE0", "Echo off"},
		{"ATL0", "Linefeeds off"},
		{"ATS0", "Spaces off"},
		{"ATSP0", "Auto protocol"},
	}
	for _, c_ := range cmds {
		if err := c.SendCommand(c_.cmd); err != nil {
			return err
		}
		if _, err := c.WaitForResponse(ctx, 2*time.Second); err != nil {
			log.Printf("Timeout durante inizializzazione (%s)", c_.desc)
		} else {
			log.Printf("Comando inizializzazione eseguito: %s", c_.desc)
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(300 * time.Millisecond):
		}
	}
	return nil
}

// ReadPID sends one PID and parses the response.
func (c *Comm) ReadPID(ctx context.Context, def PIDDef) PIDResult {
	c.mu.Lock()
	conn := c.connected
	c.mu.Unlock()
	if !conn {
		return PIDResult{PID: def.PID, Name: def.Name, Key: def.Key,
			Error: true, Value: "NON_CONNESSO",
			Timestamp: time.Now().UTC().Format(time.RFC3339Nano)}
	}

	if err := c.SendCommand(def.PID); err != nil {
		return errorResult(def, err.Error())
	}
	resp, err := c.WaitForResponse(ctx, 5*time.Second)
	if err != nil {
		return errorResult(def, err.Error())
	}

	result := ParseResponse(def, resp)

	// If ECU is sleeping, try a wake-up and retry once.
	if result.Value == "ECU_DORMIENTE" {
		c.wakeupECU(ctx)
		if err := c.SendCommand(def.PID); err != nil {
			return errorResult(def, err.Error())
		}
		if resp2, err := c.WaitForResponse(ctx, 5*time.Second); err == nil {
			result = ParseResponse(def, resp2)
		}
	}
	return result
}

func (c *Comm) wakeupECU(ctx context.Context) {
	_ = c.SendCommand("0100")
	_, _ = c.WaitForResponse(ctx, 2*time.Second)
	select {
	case <-ctx.Done():
	case <-time.After(300 * time.Millisecond):
	}
}

// IsConnected reports whether the serial port is open.
func (c *Comm) IsConnected() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.connected
}

// LastActivity returns the time of the last serial port activity.
func (c *Comm) LastActivity() time.Time {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.lastActivity
}

// Disconnect closes the serial port.
func (c *Comm) Disconnect() {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.port != nil {
		c.port.Close()
		c.port = nil
	}
	c.connected = false
}

func errorResult(def PIDDef, msg string) PIDResult {
	return PIDResult{
		PID:       def.PID,
		Name:      def.Name,
		Key:       def.Key,
		Value:     nil,
		Raw:       msg,
		Error:     true,
		Timestamp: time.Now().UTC().Format(time.RFC3339Nano),
	}
}
