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

/**
 * Servizio per la gestione della connessione WebSocket
 * con supporto per modalità mock
 */

import { io, Socket } from 'socket.io-client';
import { state } from '../store/state';
import { websocket as config } from '../config/environment';
import { CONNECTION_TIMEOUT } from '../config/constants';
import { WebSocketData, GPIOWarningsData, ExternalTemperatureData, FuelLevelData } from './types';
import { mockAnimationService } from './MockAnimationService';

class WebSocketService {
  private socket: Socket | null = null;
  private isConnected = false;
  private connectionTimeout: NodeJS.Timeout | null = null;
  public mock = config.mock;

  /**
   * Inizializza la connessione WebSocket o avvia modalità mock
   */
  connect(): void {
    if (this.mock) {
      this.startMockMode();
      return;
    }

    this.connectToServer();
  }

  /**
   * Disconnette il WebSocket e ferma le animazioni
   */
  disconnect(): void {
    this.clearConnectionTimeout();
    mockAnimationService.stop();

    if (this.socket) {
      this.socket.disconnect();
      this.socket = null;
      this.isConnected = false;
    }
  }

  /**
   * Verifica stato connessione
   */
  getConnectionStatus(): boolean {
    return this.isConnected;
  }

  /**
   * Invia evento al server
   */
  emit(event: string, data: any): void {
    if (this.socket && this.isConnected) {
      this.socket.emit(event, data);
    }
  }

  /**
   * Cambia modalità mock runtime
   */
  setMockMode(enabled: boolean): void {
    const wasInMockMode = this.mock;
    this.mock = enabled;

    if (enabled && !wasInMockMode) {
      this.disconnect();
      this.startMockMode();
    } else if (!enabled && wasInMockMode) {
      mockAnimationService.stop();
      this.connectToServer();
    }
  }

  /**
   * Verifica se è in modalità mock
   */
  isMockMode(): boolean {
    return this.mock;
  }

  /**
   * Avvia modalità mock con animazioni
   */
  private startMockMode(): void {
    mockAnimationService.start();
  }

  /**
   * Connessione al server WebSocket
   */
  private connectToServer(): void {
    try {
      this.socket = io(config.url, {
        transports: ['websocket', 'polling'],
        timeout: config.timeout,
        reconnection: true,
        reconnectionDelay: config.reconnectionDelay,
        reconnectionAttempts: config.reconnectionAttempts,
      });

      this.setupConnectionTimeout();
      this.setupEventHandlers();
    } catch (error) {
      console.error('Errore inizializzazione WebSocket:', error);
    }
  }

  /**
   * Configura timeout connessione
   */
  private setupConnectionTimeout(): void {
    this.connectionTimeout = setTimeout(() => {
      if (!this.isConnected) {
        // Timeout connessione scaduto
      }
    }, CONNECTION_TIMEOUT);
  }

  /**
   * Rimuove timeout connessione
   */
  private clearConnectionTimeout(): void {
    if (this.connectionTimeout) {
      clearTimeout(this.connectionTimeout);
      this.connectionTimeout = null;
    }
  }

  /**
   * Configura gestori eventi WebSocket
   */
  private setupEventHandlers(): void {
    if (!this.socket) return;

    // Connessione stabilita
    this.socket.on('connect', () => {
      this.isConnected = true;
      mockAnimationService.stop();
      this.clearConnectionTimeout();
    });

    // Connessione persa
    this.socket.on('disconnect', () => {
      this.isConnected = false;
    });

    // Errore connessione
    this.socket.on('connect_error', (error) => {
      console.error('Errore connessione WebSocket:', error);
      this.isConnected = false;
    });

    // Riavvio frontend
    this.socket.on('start_cockpit', () => {
      window.location.reload();
    });

    // Dati OBD live dal server Go (20 Hz)
    this.socket.on('obd-live', (data) => this.handleData(data));
    
    // Ricezione stato spie GPIO
    this.socket.on('gpio-warnings', (data) => this.handleGPIOWarnings(data));
    
    // Ricezione temperatura esterna da sensore DS18B20
    this.socket.on('external-temperature', (data) => this.handleExternalTemperature(data));
    
    // Ricezione livello carburante da sensore ADS1115
    this.socket.on('fuel-level', (data) => this.handleFuelLevel(data));
  }

  /**
   * Elabora dati ricevuti da WebSocket
   */
  private handleData(data: WebSocketData): void {
    try {
      if (!data || !data.parameters) return;

      const { parameters } = data;

      // Aggiorna temperatura
      if (parameters.coolant_temp?.success) {
        state.session.temparature.current = parameters.coolant_temp.value;
      }

      // Aggiorna RPM
      if (parameters.rpm?.success) {
        state.session.rpm.current = parameters.rpm.value;
      }

      // Aggiorna velocità
      if (parameters.speed?.success) {
        state.session.speed.current = parameters.speed.value;
      }
    } catch (error) {
      console.error('Errore elaborazione dati WebSocket:', error);
    }
  }

  /**
   * Gestisce aggiornamenti stato spie GPIO da optoaccoppiatori
   */
  private handleGPIOWarnings(data: GPIOWarningsData): void {
    try {
      if (!data || !data.warnings) return;

      // Aggiorna stato di tutte le spie ricevute
      Object.entries(data.warnings).forEach(([warningKey, isActive]) => {
        if (warningKey in state.warnings) {
          (state.warnings as any)[warningKey] = isActive;
        }
      });
    } catch (error) {
      console.error('Errore elaborazione dati GPIO:', error);
    }
  }

  /**
   * Gestisce temperatura esterna da sensore DS18B20
   */
  private handleExternalTemperature(data: ExternalTemperatureData): void {
    try {
      if (!data || !data.temperature) return;

      // Aggiorna temperatura ambiente nello stato
      (state.session.temparature as any).ambient = data.temperature.external;
    } catch (error) {
      console.error('Errore elaborazione temperatura esterna:', error);
    }
  }

  /**
   * Gestisce livello carburante da sensore ADS1115
   */
  private handleFuelLevel(data: FuelLevelData): void {
    try {
      if (!data || !data.fuel) return;

      // Aggiorna livello carburante nello stato
      state.session.fuel.current = data.fuel.level;
    } catch (error) {
      console.error('Errore elaborazione livello carburante:', error);
    }
  }
}

export const websocketService = new WebSocketService();
export default websocketService;

