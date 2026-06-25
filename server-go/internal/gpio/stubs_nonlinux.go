//go:build !linux

/*
 * PandaOS
 * Copyright (C) 2025  Cyberpandino
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3.
 */

// Stub implementations for non-Linux builds (e.g. local development on macOS).
// The real implementations live in service.go and ignition.go (linux-only).
package gpio

import (
	"context"
	"log"

	"github.com/cyberpandino/cluster/server/internal/hub"
)

type Service struct{}

func NewService(h *hub.Hub) *Service {
	log.Println("GPIO: piattaforma non-Linux, polling disabilitato")
	return &Service{}
}
func (s *Service) StartPolling(ctx context.Context) {}
func (s *Service) AllWarningStates() map[string]bool { return nil }
func (s *Service) Cleanup()                          {}

type IgnitionService struct{}

func NewIgnitionService(h *hub.Hub) *IgnitionService {
	log.Println("IgnitionService: piattaforma non-Linux, disabilitato")
	return &IgnitionService{}
}
func (s *IgnitionService) Cleanup() {}
