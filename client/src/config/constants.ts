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
  
  // RPM range covered in mock mode, from idle to the top of the sweep
  RPM: {
    IDLE: 800,
    MAX: 6500,
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
 * Gauge dials (odometer, tachometer)
 */
export const GAUGE = {
  // Time constant of the displayed value smoothing (ms)
  SMOOTHING: 120,
  // Below this difference the value snaps to the target (km/h)
  SETTLE_THRESHOLD: 0.5,
  // Portion of the circumference covered by the arc (%)
  ARC_SWEEP: 80,
  // Colour thresholds, as a fraction of max speed
  HIGH_LEVEL: 0.7,
  CRITICAL_LEVEL: 0.9,
} as const;

