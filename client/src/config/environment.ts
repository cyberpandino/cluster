/*
 * PandaOS
 * Copyright (C) 2025  Cyberpandino
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU General Public License for more details.
 */

import { EnvironmentConfig } from './types';

/**
 * Configurazione unica dell'applicazione
 * 
 * Per cambiare modalità mock: modifica il valore di websocket.mock
 * - true  = Modalità demo con animazioni simulate
 * - false = Connessione WebSocket reale al server
 */
export const environment: EnvironmentConfig = {
  websocket: {
    url: 'http://127.0.0.1:3001',
    mock: true,
    reconnectionAttempts: 3,
    reconnectionDelay: 1000,
    timeout: 5000,
  },
  splashScreen: {
    path: '/splashscreen.mp4'
  },
  debug: {
    enabled: true,
    showConsoleViewer: true,
  },
  app: {
    name: "PandaOS Cluster",
    version: "0.9.0",
    locale: "it",
    timezone: "Europe/Rome",
    timeFormat: "24h",
  },
  /**
   * Qualità grafica:
   * - 3 (max): Modello 3D + effetti blur (per PC/Mac)
   * - 2 (medium): Solo effetti blur, no modello 3D
   * - 1 (min): No blur, no modello 3D (per hardware con scarse risorse computazionali)
   */
  graphics: {
    quality: 3,
  },
};

// Export dei singoli moduli per comodità
export const { websocket, splashScreen, debug, app, graphics } = environment;

export default environment;
