/*
 * PandaOS
 * Copyright (C) 2025  Cyberpandino
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3.
 */

package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cyberpandino/cluster/server/internal/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		s := <-sigs
		log.Printf("Signal received: %s — shutting down", s)
		cancel()
	}()

	// Mirror the JS 2-second startup delay.
	time.Sleep(2 * time.Second)

	srv := server.New()
	if err := srv.Start(ctx); err != nil {
		log.Printf("Server stopped: %v", err)
		os.Exit(1)
	}
}
