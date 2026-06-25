/*
 * PandaOS
 * Copyright (C) 2025  Cyberpandino
 */

package obd

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestNewComm_InitialState(t *testing.T) {
	c := NewComm()
	if c.IsConnected() {
		t.Error("new Comm should not be connected")
	}
	if c.lines == nil {
		t.Error("lines channel should be initialised")
	}
	if cap(c.lines) != lineBuffer {
		t.Errorf("channel capacity = %d, want %d", cap(c.lines), lineBuffer)
	}
}

func TestErrorResult(t *testing.T) {
	def := PIDDef{PID: "010C", Name: "RPM", Key: "rpm"}
	r := errorResult(def, "test error")
	if r.PID != def.PID {
		t.Errorf("PID: got %q, want %q", r.PID, def.PID)
	}
	if r.Key != def.Key {
		t.Errorf("Key: got %q, want %q", r.Key, def.Key)
	}
	if r.Raw != "test error" {
		t.Errorf("Raw: got %q", r.Raw)
	}
	if r.Success || !r.Error {
		t.Errorf("expected Error=true, Success=false")
	}
	if r.Timestamp == "" {
		t.Error("Timestamp must be set")
	}
}

func TestDrain_ClearsChannel(t *testing.T) {
	c := NewComm()
	c.lines <- "line1"
	c.lines <- "line2"
	c.drain()
	select {
	case leftover := <-c.lines:
		t.Errorf("channel should be empty after drain, got %q", leftover)
	default:
	}
}

func TestDrain_EmptyChannel(t *testing.T) {
	c := NewComm()
	// Must not block on an already-empty channel.
	done := make(chan struct{})
	go func() {
		c.drain()
		close(done)
	}()
	select {
	case <-done:
	case <-time.After(100 * time.Millisecond):
		t.Fatal("drain blocked on empty channel")
	}
}

func TestWaitForResponse_ContextAlreadyCancelled(t *testing.T) {
	c := NewComm()
	c.connected = true

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // already done

	_, err := c.WaitForResponse(ctx, time.Second)
	if !errors.Is(err, context.Canceled) {
		t.Errorf("expected context.Canceled, got %v", err)
	}
}

func TestWaitForResponse_ReceivesLine(t *testing.T) {
	c := NewComm()
	c.connected = true

	// Pre-load a response before calling WaitForResponse.
	c.lines <- "41 0C 1A F8"

	ctx := context.Background()
	got, err := c.WaitForResponse(ctx, time.Second)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "41 0C 1A F8" {
		t.Errorf("got %q, want %q", got, "41 0C 1A F8")
	}
}

func TestWaitForResponse_SkipsSearching(t *testing.T) {
	c := NewComm()
	c.connected = true

	// First line is SEARCHING, then the real response.
	c.lines <- "SEARCHING..."
	go func() {
		time.Sleep(50 * time.Millisecond)
		c.lines <- "41 0D 3C"
	}()

	ctx := context.Background()
	got, err := c.WaitForResponse(ctx, 3*time.Second)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "41 0D 3C" {
		t.Errorf("got %q, want %q", got, "41 0D 3C")
	}
}

func TestWaitForResponse_Timeout(t *testing.T) {
	c := NewComm()
	c.connected = true

	ctx := context.Background()
	start := time.Now()
	_, err := c.WaitForResponse(ctx, 150*time.Millisecond)
	elapsed := time.Since(start)

	if err == nil {
		t.Fatal("expected timeout error")
	}
	if elapsed < 100*time.Millisecond {
		t.Errorf("returned too quickly: %v", elapsed)
	}
}

func TestSendCommand_NotConnected(t *testing.T) {
	c := NewComm()
	err := c.SendCommand("ATZ")
	if err == nil {
		t.Fatal("expected error when not connected")
	}
}

func TestDisconnect_Idempotent(t *testing.T) {
	c := NewComm()
	// Calling Disconnect on an unconnected Comm must not panic.
	c.Disconnect()
	c.Disconnect()
	if c.IsConnected() {
		t.Error("should not be connected after Disconnect")
	}
}

func TestLastActivity_InitiallyZero(t *testing.T) {
	c := NewComm()
	if !c.LastActivity().IsZero() {
		t.Error("LastActivity should be zero on new Comm")
	}
}

// ---- benchmarks ----

func BenchmarkWaitForResponse_ImmediateHit(b *testing.B) {
	b.ReportAllocs()
	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		c := NewComm()
		c.connected = true
		c.lines <- "41 0C 1A F8"
		_, _ = c.WaitForResponse(ctx, time.Second)
	}
}
