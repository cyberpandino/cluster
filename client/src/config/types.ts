/**
 * Tipi e interfacce per la configurazione dell'ambiente
 */

export interface WebSocketConfig {
  url: string;
  mock: boolean;
  reconnectionAttempts: number;
  reconnectionDelay: number;
  timeout: number;
}

export interface SplashScreenConfig {
  path: string;
}

export interface DebugConfig {
  enabled: boolean;
  showConsoleViewer: boolean;
}

export interface AppConfig {
  name: string;
  version: string;
  locale: string;
  timezone: string;
  timeFormat: "24h" | "12h";
}

/**
 * Configurazione qualità grafica
 * - 3 (max): Tutto abilitato (modello 3D + effetti blur)
 * - 2 (medium): Modello 3D disabilitato, effetti blur attivi
 * - 1 (min): Modello 3D disabilitato + blur sostituiti con gradienti statici
 */
export interface GraphicsConfig {
  quality: 1 | 2 | 3;
}

export interface EnvironmentConfig {
  websocket: WebSocketConfig;
  splashScreen: SplashScreenConfig;
  debug: DebugConfig;
  app: AppConfig;
  graphics: GraphicsConfig;
}

