[🇬🇧 English](CONFIGURAZIONE.en.md) | [🇮🇹 Italiano](CONFIGURAZIONE.md) | [🇩🇪 Deutsch](CONFIGURAZIONE.de.md)

---

# ⚙️ Client-Konfiguration - PandaOS Cluster

Vollständiger Leitfaden zur Konfiguration des Cyberpandino-Clients.

---

## 📁 Konfigurationsdateien

Die Hauptkonfigurationsdatei befindet sich in:

```
client/src/config/environment.ts
```

---

## 🔧 Konfigurationsparameter

### 1. WebSocket

Konfiguration der Verbindung zum Backend-Server.

```typescript
websocket: {
  url: 'http://127.0.0.1:3001',
  mock: true,
  reconnectionAttempts: 3,
  reconnectionDelay: 1000,
  timeout: 5000,
}
```

#### `url` (string)
- **Beschreibung**: Adresse des WebSocket-Servers
- **Standard**: `'http://127.0.0.1:3001'`
- **Raspberry Pi-Produktion**: `'http://127.0.0.1:3001'` (localhost)
- **Remote-Entwicklung**: `'http://192.168.x.x:3001'` (Raspberry IP)

**Beispiele**:
```typescript
// Lokal (Entwicklung oder Produktion auf demselben Gerät)
url: 'http://127.0.0.1:3001'

// Raspberry Pi im lokalen Netzwerk
url: 'http://192.168.1.100:3001'

// Server auf benutzerdefiniertem Port
url: 'http://127.0.0.1:8080'
```

#### `mock` (boolean)
- **Beschreibung**: Client-Betriebsmodus
- **Werte**:
  - `true` = Demo-Modus mit simulierten Daten (kein Server erforderlich)
  - `false` = Echte Verbindung zum Backend-Server
- **Standard**: `true`

**Wann verwenden**:
```typescript
// Lokale Entwicklung auf Mac/Windows/Linux (ohne Raspberry Pi)
mock: true

// Produktion auf Raspberry Pi mit aktivem Server
mock: false

// Interface-Testing ohne Hardware
mock: true
```

#### `reconnectionAttempts` (number)
- **Beschreibung**: Maximale Anzahl von Wiederverbindungsversuchen
- **Standard**: `3`
- **Empfohlener Bereich**: 2-10
- **Verhalten**: Nach Erschöpfung der Versuche wechselt zu Mock-Modus

#### `reconnectionDelay` (number)
- **Beschreibung**: Verzögerung zwischen Wiederverbindungsversuchen (Millisekunden)
- **Standard**: `1000` (1 Sekunde)
- **Empfohlener Bereich**: 500-5000 ms

#### `timeout` (number)
- **Beschreibung**: WebSocket-Verbindungs-Timeout (Millisekunden)
- **Standard**: `5000` (5 Sekunden)
- **Empfohlener Bereich**: 3000-10000 ms

---

### 2. Splash Screen

Konfiguration des Startbildschirms.

```typescript
splashScreen: {
  path: '/splashscreen.mp4'
}
```

#### `path` (string)
- **Beschreibung**: Pfad des Splash-Screen-Videos
- **Standard**: `'/splashscreen.mp4'`
- **Unterstütztes Format**: MP4, WebM
- **Dateispeicherort**: `client/public/splashscreen.mp4`

**Anpassung**:
```typescript
// Benutzerdefiniertes Video
path: '/custom-splash.mp4'

// Splash deaktivieren (verwenden Sie statisches Bild oder nichts)
path: '' // Erfordert Änderung der SplashScreen-Komponente
```

---

### 3. Debug

Konfiguration des Debug- und Entwicklungsmodus.

```typescript
debug: {
  enabled: true,
  showConsoleViewer: true,
}
```

#### `enabled` (boolean)
- **Beschreibung**: Aktiviert Debug-Funktionen
- **Standard**: `true`
- **Auswirkungen**:
  - Zeigt detaillierte Logs in der Browser-Konsole
  - Aktiviert Debug-Schaltflächen
  - Zeigt WebSocket-Verbindungsinformationen
- **Produktion**: Auf `false` setzen für optimale Performance

#### `showConsoleViewer` (boolean)
- **Beschreibung**: Aktiviert integrierten Console Viewer (Taste `d`)
- **Standard**: `true`
- **Funktionen**:
  - Echtzeit-Log-Anzeige
  - Fehlerüberwachung
  - WebSocket-Debugging
  - Sensor- und GPIO-Status

**Console Viewer-Steuerung**:
- Drücken Sie **`d`** zum Öffnen
- Drücken Sie **`ESC`** zum Schließen
- "Clear"-Schaltfläche zum Löschen der Logs

---

### 4. App

Allgemeine Anwendungskonfiguration.

```typescript
app: {
  name: "PandaOS Cluster",
  version: "0.9.0",
  locale: "it",
  timezone: "Europe/Rome",
  timeFormat: "24h",
}
```

#### `name` (string)
- **Beschreibung**: Angezeigter Anwendungsname
- **Standard**: `"PandaOS Cluster"`
- **Verwendung**: Fenstertitel, Splash-Screen, About

#### `version` (string)
- **Beschreibung**: Anwendungsversion
- **Standard**: `"0.9.0"`
- **Format**: Semantic Versioning (MAJOR.MINOR.PATCH)

#### `locale` (string)
- **Beschreibung**: Anwendungssprache
- **Standard**: `"it"` (Italienisch)
- **Unterstützt**: 
  - `"it"` - Italienisch

#### `timezone` (string)
- **Beschreibung**: Zeitzone für Datums-/Zeitanzeige
- **Standard**: `"Europe/Rome"`
- **Format**: IANA-Zeitzonendatenbank
- **Andere Beispiele**:
  - `"Europe/Berlin"`
  - `"America/New_York"`
  - `"Asia/Tokyo"`

#### `timeFormat` (string)
- **Beschreibung**: Zeitformat
- **Werte**:
  - `"24h"` - 24-Stunden-Format (z.B. 15:30)
  - `"12h"` - 12-Stunden-Format AM/PM (z.B. 3:30 PM)
- **Standard**: `"24h"`

---

## 📋 Konfigurationsbeispiele

### Lokale Entwicklungskonfiguration (Mac/Windows)

```typescript
export const environment: EnvironmentConfig = {
  websocket: {
    url: 'http://127.0.0.1:3001',
    mock: true,                    // ← Demo-Modus
    reconnectionAttempts: 3,
    reconnectionDelay: 1000,
    timeout: 5000,
  },
  splashScreen: {
    path: '/splashscreen.mp4'
  },
  debug: {
    enabled: true,                 // ← Debug aktiv
    showConsoleViewer: true,       // ← Console Viewer verfügbar
  },
  app: {
    name: "PandaOS Cluster [DEV]",
    version: "0.9.0",
    locale: "it",
    timezone: "Europe/Rome",
    timeFormat: "24h",
  },
};
```

### Produktionskonfiguration (Raspberry Pi)

```typescript
export const environment: EnvironmentConfig = {
  websocket: {
    url: 'http://127.0.0.1:3001',
    mock: false,                   // ← Echte Verbindung
    reconnectionAttempts: 10,      // ← Mehr Versuche für Stabilität
    reconnectionDelay: 2000,
    timeout: 10000,
  },
  splashScreen: {
    path: '/splashscreen.mp4'
  },
  debug: {
    enabled: false,                // ← Debug deaktiviert für Performance
    showConsoleViewer: false,
  },
  app: {
    name: "PandaOS Cluster",
    version: "0.9.0",
    locale: "it",
    timezone: "Europe/Rome",
    timeFormat: "24h",
  },
};
```

### Remote-Test-Konfiguration

```typescript
export const environment: EnvironmentConfig = {
  websocket: {
    url: 'http://192.168.1.100:3001', // ← Raspberry Pi IP
    mock: false,                       // ← Echte Verbindung
    reconnectionAttempts: 5,
    reconnectionDelay: 1500,
    timeout: 7000,
  },
  splashScreen: {
    path: '/splashscreen.mp4'
  },
  debug: {
    enabled: true,                     // ← Debug für Fehlerbehebung
    showConsoleViewer: true,
  },
  app: {
    name: "PandaOS Cluster [TEST]",
    version: "0.9.0",
    locale: "it",
    timezone: "Europe/Rome",
    timeFormat: "24h",
  },
};
```

---

## 🔄 Entwicklungs-Workflow

### 1. Interface-Entwicklung (ohne Hardware)

```typescript
websocket.mock = true
debug.enabled = true
```

Nur Client starten:
```bash
cd client
npm run dev
```

### 2. Integrationstest (mit Raspberry Pi)

```typescript
websocket.mock = false
websocket.url = 'http://192.168.1.100:3001'
debug.enabled = true
```

Server auf Raspberry Pi und lokalen Client starten

### 3. Produktion (auf Raspberry Pi)

```typescript
websocket.mock = false
websocket.url = 'http://127.0.0.1:3001'
debug.enabled = false
```

Vollständigen Stack starten:
```bash
npm start
```

---

## 🔍 Mock- vs. Real-Modus

### Mock-Modus (`mock: true`)

**Eigenschaften**:
- ✅ Kein Server erforderlich
- ✅ Realistische simulierte Daten
- ✅ Flüssige Animationen
- ✅ Funktioniert auf jedem OS
- ❌ Liest keine echten OBD-Daten
- ❌ Liest keine GPIO/Sensoren

**Verwendung**: UI-Entwicklung, Demo, visuelles Testing

### Real-Modus (`mock: false`)

**Eigenschaften**:
- ✅ Echte OBD-Daten vom Steuergerät
- ✅ GPIO-Lesung für Fahrzeugwarnleuchten
- ✅ Temperatur- und Kraftstoffsensoren
- ❌ Erfordert aktiven Server
- ❌ Erfordert Hardware (Raspberry Pi + ELM327)

**Verwendung**: Produktion, Hardware-Testing, vollständige Integration

---

## 🚨 Fehlerbehebung

### Client verbindet sich nicht mit Server

**Symptom**: Konsole zeigt Verbindungsfehler

**Lösungen**:
1. Überprüfen Sie, dass `mock: false`
2. Server-URL prüfen: `websocket.url`
3. Überprüfen Sie, dass der Server gestartet ist: `npm run server`
4. Firewall/Port 3001 offen prüfen
5. Verbindung testen: `curl http://127.0.0.1:3001`

### Splash-Screen lädt nicht

**Symptom**: Weißer Bildschirm beim Start

**Lösungen**:
1. Überprüfen Sie, dass Datei existiert: `client/public/splashscreen.mp4`
2. Pfad in Konfiguration prüfen
3. Videoformat überprüfen (MP4 bevorzugt)
4. Browser-Konsole auf Fehler prüfen

### Debug-Konsole öffnet sich nicht

**Symptom**: Beim Drücken von `d` passiert nichts

**Lösungen**:
1. Überprüfen Sie `debug.showConsoleViewer: true`
2. Prüfen Sie, dass `debug.enabled: true`
3. Browser-Konsole auf Fehler prüfen
4. Versuchen Sie `CTRL+D` oder `CMD+D`

### Zeitzone/Uhrzeit falsch

**Symptom**: Uhrzeit im Cluster nicht korrekt

**Lösungen**:
1. Überprüfen Sie `app.timezone` korrekt für Ihre Zone
2. Zeitzonenliste: https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
3. Raspberry Pi-Systemzeit überprüfen: `timedatectl`
4. NTP synchronisieren: `sudo timedatectl set-ntp true`

---

## 📚 Referenzen

- **Quellcode**: `client/src/config/environment.ts`
- **TypeScript-Typen**: `client/src/config/types.ts`
- **Hauptdokumentation**: `README.md`
- **Server-Konfiguration**: `server/CONFIGURAZIONE_SERVER.md`

---

**Zuletzt aktualisiert**: v0.9.0
