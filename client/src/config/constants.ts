/**
 * Costanti globali dell'applicazione
 */

/**
 * Costanti per animazioni mock
 */
export const MOCK_ANIMATION = {
  // Durata animazioni velocità (ms)
  SPEED: {
    ACCELERATION_DURATION: 15000,
    DECELERATION_DURATION: 15000,
    PAUSE_DURATION: 5000,
  },
  
  // RPM fisso in modalità mock
  RPM: {
    FIXED_VALUE: 1650,
  },
  
  // Ciclo spie di warning (ms)
  WARNING: {
    CYCLE_DURATION: 25000,
    ACTIVE_DURATION: 10000,
  },
  
  // Range temperatura (°C)
  TEMPERATURE: {
    MIN: 75,
    MAX: 105,
    CYCLE_DURATION: 60000,
  },
  
  // Incremento chilometraggio (ms)
  KILOMETRES: {
    INCREMENT_INTERVAL: 5000,
    INCREMENT_VALUE: 1,
  },
} as const;

/**
 * Costanti per la UI del cockpit
 */
export const COCKPIT_UI = {
  CLOCK_REFRESH_INTERVAL: 1000, // ms
} as const;

/**
 * Timeout connessione WebSocket (ms)
 */
export const CONNECTION_TIMEOUT = 3000;

/**
 * Lista spie di warning disponibili
 */
export const WARNING_LIGHTS = [
  'doors',
  'light',
  'lowBeam',
  'highBeam',
  'fogLight',
  'engineColant',
  'warning',
  'hazard',
  'turnSignals',
  'battery',
  'brakeSystem',
  'fuel',
] as const;

/**
 * Velocità animazione tachimetro/odometro
 */
export const ANIMATION_SPEED = {
  STEP: 80,
  THRESHOLD: 1,
} as const;

/**
 * Odometer (speed)
 */
export const ODOMETER = {
  // Portion of the circumference covered by the arc (%)
  ARC_SWEEP: 80,
} as const;

