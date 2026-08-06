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
 * Servizio per gestire le animazioni mock in modalità demo
 */

import { state } from '../store/state';
import { MOCK_ANIMATION, WARNING_LIGHTS } from '../config/constants';

export class MockAnimationService {
  private isActive = false;
  private speedAnimationFrame: number | null = null;
  private rpmAnimationFrame: number | null = null;
  private warningTimeout: NodeJS.Timeout | null = null;
  private temperatureAnimationFrame: number | null = null;
  private kilometresTimeout: NodeJS.Timeout | null = null;

  /**
   * Avvia tutte le animazioni mock
   */
  start(): void {
    if (this.isActive) return;

    this.isActive = true;
    this.resetState();
    this.startSpeedAnimation();
    this.startRpmAnimation();
    this.startWarningAnimation();
    this.startTemperatureAnimation();
    this.startKilometresAnimation();
  }

  /**
   * Ferma tutte le animazioni mock
   */
  stop(): void {
    if (!this.isActive) return;

    this.isActive = false;

    if (this.speedAnimationFrame) {
      cancelAnimationFrame(this.speedAnimationFrame);
      this.speedAnimationFrame = null;
    }

    if (this.rpmAnimationFrame) {
      cancelAnimationFrame(this.rpmAnimationFrame);
      this.rpmAnimationFrame = null;
    }

    if (this.warningTimeout) {
      clearTimeout(this.warningTimeout);
      this.warningTimeout = null;
    }

    if (this.temperatureAnimationFrame) {
      cancelAnimationFrame(this.temperatureAnimationFrame);
      this.temperatureAnimationFrame = null;
    }

    if (this.kilometresTimeout) {
      clearTimeout(this.kilometresTimeout);
      this.kilometresTimeout = null;
    }

    this.resetWarnings();
  }

  /**
   * Verifica se le animazioni sono attive
   */
  isRunning(): boolean {
    return this.isActive;
  }

  /**
   * Reset stato iniziale
   */
  private resetState(): void {
    state.session.rpm.current = 0;
    state.session.speed.current = 0;
    state.session.temparature.current = MOCK_ANIMATION.TEMPERATURE.MIN;
    this.resetWarnings();
  }

  /**
   * Spegne tutte le spie di warning
   */
  private resetWarnings(): void {
    Object.keys(state.warnings).forEach(key => {
      (state.warnings as any)[key] = false;
    });
  }

  /**
   * Animazione velocità: ciclo accelerazione/decelerazione
   */
  private startSpeedAnimation(): void {
    const maxSpeed = state.session.speed.max;
    const { ACCELERATION_DURATION, DECELERATION_DURATION, PAUSE_DURATION } = MOCK_ANIMATION.SPEED;
    const cycleDuration = ACCELERATION_DURATION + DECELERATION_DURATION + PAUSE_DURATION;
    const startTime = Date.now();

    const animate = (): void => {
      if (!this.isActive) return;

      const elapsed = (Date.now() - startTime) % cycleDuration;
      let currentSpeed: number;

      if (elapsed < ACCELERATION_DURATION) {
        // Fase accelerazione
        const progress = elapsed / ACCELERATION_DURATION;
        currentSpeed = Math.round(progress * maxSpeed);
      } else if (elapsed < ACCELERATION_DURATION + DECELERATION_DURATION) {
        // Fase decelerazione
        const decelerationElapsed = elapsed - ACCELERATION_DURATION;
        const progress = decelerationElapsed / DECELERATION_DURATION;
        currentSpeed = Math.round(maxSpeed * (1 - progress));
      } else {
        // Fase pausa
        currentSpeed = 0;
      }

      state.session.speed.current = currentSpeed;
      this.speedAnimationFrame = requestAnimationFrame(animate);
    };

    animate();
  }

  /**
   * Animazione RPM: derivata dalla velocità corrente.
   * Keeping the two gauges on the same source stops them telling
   * contradictory stories, which a fixed value did
   */
  private startRpmAnimation(): void {
    const { IDLE, MAX } = MOCK_ANIMATION.RPM;
    const maxSpeed = state.session.speed.max;

    const animate = (): void => {
      if (!this.isActive) return;

      const progress = Math.min(state.session.speed.current / maxSpeed, 1);

      state.session.rpm.current = Math.round(IDLE + (MAX - IDLE) * progress);
      this.rpmAnimationFrame = requestAnimationFrame(animate);
    };

    animate();
  }

  /**
   * Animazione spie di warning: accensione/spegnimento casuale
   */
  private startWarningAnimation(): void {
    const { CYCLE_DURATION, ACTIVE_DURATION } = MOCK_ANIMATION.WARNING;

    const animate = (): void => {
      if (!this.isActive) {
        if (this.warningTimeout) {
          clearTimeout(this.warningTimeout);
          this.warningTimeout = null;
        }
        return;
      }

      // Seleziona spia casuale
      const randomIndex = Math.floor(Math.random() * WARNING_LIGHTS.length);
      const selectedWarning = WARNING_LIGHTS[randomIndex];

      // Accende la spia
      (state.warnings as any)[selectedWarning] = true;

      // Programma spegnimento
      setTimeout(() => {
        (state.warnings as any)[selectedWarning] = false;
      }, ACTIVE_DURATION);

      // Programma prossimo ciclo
      this.warningTimeout = setTimeout(() => {
        animate();
      }, CYCLE_DURATION);
    };

    animate();
  }

  /**
   * Animazione temperatura: oscillazione ciclica
   */
  private startTemperatureAnimation(): void {
    const { MIN, MAX, CYCLE_DURATION } = MOCK_ANIMATION.TEMPERATURE;
    const startTime = Date.now();

    const animate = (): void => {
      if (!this.isActive) return;

      const elapsed = (Date.now() - startTime) % CYCLE_DURATION;
      const halfCycle = CYCLE_DURATION / 2;

      let currentTemp: number;
      if (elapsed < halfCycle) {
        // Prima metà: MIN → MAX
        const progress = elapsed / halfCycle;
        currentTemp = MIN + Math.round(progress * (MAX - MIN));
      } else {
        // Seconda metà: MAX → MIN
        const progress = (elapsed - halfCycle) / halfCycle;
        currentTemp = MAX - Math.round(progress * (MAX - MIN));
      }

      state.session.temparature.current = currentTemp;
      this.temperatureAnimationFrame = requestAnimationFrame(animate);
    };

    animate();
  }

  /**
   * Animazione chilometri: incremento periodico
   */
  private startKilometresAnimation(): void {
    const { INCREMENT_INTERVAL, INCREMENT_VALUE } = MOCK_ANIMATION.KILOMETRES;

    const animate = (): void => {
      if (!this.isActive) return;

      state.session.kilometres.current += INCREMENT_VALUE;

      this.kilometresTimeout = setTimeout(() => {
        animate();
      }, INCREMENT_INTERVAL);
    };

    animate();
  }
}

export const mockAnimationService = new MockAnimationService();

