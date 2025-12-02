[🇬🇧 English](ARCHITETTURA.en.md) | [🇮🇹 Italiano](ARCHITETTURA.md) | [🇩🇪 Deutsch](ARCHITETTURA.de.md)

---

# 🏗️ Architektur - PandaOS Cluster

Technische Dokumentation der Systemarchitektur.

> 💡 **Hinweis zum Stack**: Ja, wir verwenden JavaScript in einem Auto. Ja, wir wissen, dass es verrückt ist. Nein, wir bereuen es nicht. Siehe [README - Aber React + Electron in einem Automotive?!](README.de.md) für die vollständige Rechtfertigung.

---

## 📊 Allgemeiner Überblick

```
┌─────────────────────────────────────────────────────────────┐
│                    ELECTRON WRAPPER                          │
│  (main.js - Desktop Application - Port 5173)                │
└──────────────────┬──────────────────────────────────────────┘
                   │
                   ▼
┌─────────────────────────────────────────────────────────────┐
│                    CLIENT LAYER                              │
│  (React + TypeScript + Vite - Port 5173)                    │
│                                                               │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐      │
│  │  Components  │  │    Routes    │  │   Services   │      │
│  │  (UI/UX)     │  │  (Cockpit)   │  │  (WebSocket) │      │
│  └──────────────┘  └──────────────┘  └──────────────┘      │
│                                                               │
│  ┌──────────────────────────────────────────────────────┐   │
│  │            State Management (Valtio)                  │   │
│  │  - OBD Data    - GPIO Warnings    - Sensors          │   │
│  └──────────────────────────────────────────────────────┘   │
└──────────────────┬──────────────────────────────────────────┘
                   │ WebSocket (Socket.IO)
                   │ ws://localhost:3001
                   ▼
┌─────────────────────────────────────────────────────────────┐
│                    SERVER LAYER                              │
│  (Node.js + Express + Socket.IO - Port 3001)                │
│                                                               │
│  ┌──────────────────────────────────────────────────────┐   │
│  │                   OBDServer (Main)                    │   │
│  └──────────────────────────────────────────────────────┘   │
│                                                               │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐      │
│  │OBDComm       │  │  WebSocket   │  │  Monitoring  │      │
│  │Service       │  │  Service     │  │  Service     │      │
│  └──────┬───────┘  └──────────────┘  └──────────────┘      │
│         │                                                     │
│  ┌──────┴───────┐  ┌──────────────┐  ┌──────────────┐      │
│  │GPIO          │  │  Ignition    │  │  Temperature │      │
│  │Service       │  │  Service     │  │  Service     │      │
│  └──────────────┘  └──────────────┘  └──────────────┘      │
│                                                               │
│  ┌──────────────┐  ┌──────────────┐                         │
│  │Fuel          │  │  PIDParser   │                         │
│  │Service       │  │  Service     │                         │
│  └──────────────┘  └──────────────┘                         │
└──────────────────┬──────────────────────────────────────────┘
                   │
                   ▼
┌─────────────────────────────────────────────────────────────┐
│                    HARDWARE LAYER                            │
│  (Raspberry Pi 4B - Sensors - Actuators)                    │
│                                                               │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐      │
│  │ ELM327 OBD   │  │  GPIO Pins   │  │  DS18B20     │      │
│  │ /dev/ttyUSB0 │  │  (BCM mode)  │  │  (1-Wire)    │      │
│  └──────────────┘  └──────────────┘  └──────────────┘      │
│                                                               │
│  ┌──────────────┐  ┌──────────────┐                         │
│  │  ADS1115     │  │  Ignition    │                         │
│  │  (I2C 0x48)  │  │  GPIO 21     │                         │
│  └──────────────┘  └──────────────┘                         │
└─────────────────────────────────────────────────────────────┘
                   │
                   ▼
┌─────────────────────────────────────────────────────────────┐
│                  VEHICLE ECU LAYER                           │
│  (Fiat Panda 141 - Magneti Marelli IAW 4AF)                │
│                                                               │
│  • OBD-II K-Line (ISO 9141-2 / ISO 14230-4)                 │
│  • 12V-Warnleuchten (Optokoppler)                           │
│  • Kraftstoffsensor (0-12V analog)                          │
│  • Zündschloss (12V an/aus)                                 │
└─────────────────────────────────────────────────────────────┘
```

---

## 🔄 Datenfluss

### 1. Systemstart

```
1. Electron (main.js)
   ↓
2. Server starten (Port 3001)
   ↓
3. Vite Client starten (Port 5173)
   ↓
4. Electron lädt http://localhost:5173
   ↓
5. Client verbindet via WebSocket zu ws://localhost:3001
   ↓
6. Server initialisiert Services (GPIO, OBD, Sensors)
   ↓
7. System betriebsbereit
```

### 2. OBD-Datenlesung

```
ECU (Magneti Marelli)
   ↓ (K-Line ISO 9141-2)
ELM327 Adapter
   ↓ (Serial USB 38400 Baud)
Raspberry Pi /dev/ttyUSB0
   ↓
OBDCommunicationService
   ↓ (AT-Befehle / PID)
PIDParserService
   ↓ (Parsing hex → Werte)
MonitoringService
   ↓ (Kontinuierliches Polling)
WebSocketService
   ↓ (Socket.IO emit)
Client WebSocketService
   ↓ (State-Update)
Valtio State
   ↓ (React Re-Render)
UI Components (Tachometer, Odometer, etc.)
```

### 3. Warnleuchten-Erkennung

```
Fahrzeugwarnleuchte 12V
   ↓
Optokoppler (PC817)
   ↓
GPIO-Pin (z.B. GPIO 17)
   ↓
GPIOService (Polling 100ms)
   ↓ (Debounce 50ms)
WebSocketService
   ↓ (Socket.IO emit)
Client
   ↓
State.warnings
   ↓
WarningLights Component
```

### 4. Sensor-Auslesung

#### Temperatur (DS18B20)
```
DS18B20 Sensor
   ↓ (1-Wire GPIO 4)
/sys/bus/w1/devices/28-xxx/w1_slave
   ↓ (Datei-Lesung alle 5s)
TemperatureSensorService
   ↓
WebSocketService
   ↓
Client State
   ↓
Temperature Component
```

#### Kraftstoff (ADS1115)
```
Fahrzeug-Kraftstoffsensor (0-12V)
   ↓ (Spannungsteiler)
ADS1115 Kanal A0 (0-4V)
   ↓ (I2C 0x48)
FuelSensorService
   ↓ (Lesung alle 500ms)
Kalibrierung Spannung → Prozent
   ↓
WebSocketService
   ↓
Client State
   ↓
Fuel Component
```

---

## 📦 Server-Module

### OBDServer (Hauptorchestrator)

**Datei**: `server/services/OBDServer.js`

**Verantwortlichkeiten**:
- Initialisierung und Koordination aller Services
- Verwaltung des Lebenszyklus (Start/Stop/Neustart)
- Retry-Logik für OBD-Verbindung
- Periodisches PID-Scanning
- Fehlerbehandlung und Recovery

**Hauptmethoden**:
```javascript
start()                    // Alle Services starten
stop()                     // Graceful Stop
startWithRetry()           // OBD-Verbindung mit Retry
scanAllPIDs()              // Initiales Scannen unterstützter PIDs
periodicPIDScan()          // Periodisches Scannen (alle 30s)
forceReconnect()           // OBD-Wiederverbindung
forceRestart()             // Vollständiger Prozess-Neustart
```

**Behandelte Events**:
- `SIGINT` → Graceful Shutdown
- `unhandledRejection` → Force Restart
- `uncaughtException` → Force Restart
- WebSocket `force-restart` → Force Restart

---

### OBDCommunicationService

**Datei**: `server/services/OBDCommunicationService.js`

**Verantwortlichkeiten**:
- Serielle Kommunikation mit ELM327
- Senden von AT- und PID-Befehlen
- Parsen von ELM327-Antworten
- Timeout- und Fehlerverwaltung

**Konfiguration**:
```javascript
portPath: '/dev/ttyUSB0'
baudRate: 38400
timeout: 4000ms (Standard)
```

**Hauptmethoden**:
```javascript
connect()                  // Serielle Port öffnen
initialize()               // ELM327-Setup (ATZ, ATE0, etc.)
sendCommand(cmd)           // Raw-Befehl senden
waitForResponse(timeout)   // Auf Antwort mit Timeout warten
readPID(pid, name)         // Einzelnen PID lesen
wakeupECU()                // ECU aufwecken falls im Sleep
disconnect()               // Port schließen
```

**Initialisierungs-Befehle**:
```javascript
'ATZ'     // ELM327 zurücksetzen
'ATE0'    // Echo aus
'ATL0'    // Linefeeds aus
'ATS0'    // Leerzeichen aus
'ATSP0'   // Auto-Protokoll-Erkennung
```

---

### PIDParserService

**Datei**: `server/services/PIDParserService.js`

**Verantwortlichkeiten**:
- Definition der OBD-II-PIDs
- Parsen von Hex-Antworten → physikalische Werte
- Spezifische Umrechnungsformeln für Magneti Marelli ECU

**Unterstützte PIDs**:
```javascript
'010C' // RPM (Motordrehzahl)
'010D' // Speed (Geschwindigkeit)
'0105' // Coolant Temperature (Kühlmitteltemperatur)
'010F' // Intake Air Temperature (Ansauglufttemperatur)
'0111' // Throttle Position (Gaspedalstellung)
'0104' // Engine Load (Motorlast)
'010A' // Fuel Pressure (Kraftstoffdruck)
'010B' // Intake Manifold Pressure (Ansaugkrümmerdruck)
'0106' // Short Term Fuel Trim (Kurzzeit-Kraftstoffanpassung)
'0107' // Long Term Fuel Trim (Langzeit-Kraftstoffanpassung)
'0142' // Control Module Voltage (Steuergeräte-Spannung)
```

**Parsing-Beispiel**:
```javascript
// PID 010C (RPM)
// Antwort: 41 0C 1A F8
// A=1A(hex)=26(dez), B=F8(hex)=248(dez)
// RPM = ((A*256)+B)/4 = (6656+248)/4 = 1726 RPM

parseResponse(pid, response, name) {
  // ... spezifische PID-Parsing-Logik
  return {
    pid: '010C',
    name: 'Engine RPM',
    value: 1726,
    unit: 'RPM',
    raw: '41 0C 1A F8',
    success: true,
    timestamp: '2025-01-01T12:00:00.000Z'
  }
}
```

---

### MonitoringService

**Datei**: `server/services/MonitoringService.js`

**Verantwortlichkeiten**:
- Kontinuierliches Polling funktionierender PIDs
- Echtzeit-Datenübertragung via WebSocket
- Watchdog zur Erkennung von Freeze/Timeout
- Verwaltung dynamischer PID-Liste

**Konfiguration**:
```javascript
pollingDelay: 200ms        // Verzögerung zwischen PID-Lesungen
watchdogInterval: 30s      // Aktivitätsprüfung alle 30s
watchdogTimeout: 60s       // Inaktivitäts-Timeout
```

**Methoden**:
```javascript
startMonitoring(workingPIDs)  // Polling starten
stopMonitoring()              // Polling stoppen
updateWorkingPIDs(newList)    // Zu überwachende PIDs aktualisieren
isPIDCurrentlyMonitored(key)  // Prüfen ob PID aktiv
startWatchdog()               // Watchdog starten
```

**Monitoring-Ablauf**:
```javascript
Endlosschleife:
  Für jeden PID in workingPIDs:
    1. PID vom ECU lesen
    2. Daten via WebSocket senden
    3. 200ms warten
  Wiederholen
```

---

### GPIOService

**Datei**: `server/services/GPIOService.js`

**Verantwortlichkeiten**:
- GPIO-Pin-Initialisierung
- Polling des Warnleuchten-Status
- Signal-Debouncing
- Events bei Statusänderung senden

**Konfiguration**:
```javascript
pollingInterval: 100ms     // GPIO-Lesefrequenz
debounceTime: 50ms         // Entprellung
```

**Methoden**:
```javascript
initializeGPIO()           // Alle Pins einrichten
startPolling()             // GPIO-Polling starten
stopPolling()              // Polling stoppen
readGPIOState(pin)         // Einzelnen Pin lesen
cleanup()                  // GPIO-Ressourcen freigeben
```

**Debouncing-Algorithmus**:
```javascript
lastStableState[pin] = null
lastReadTime[pin] = 0

onPoll():
  currentState = gpio.read(pin)
  now = Date.now()
  
  if (currentState != lastStableState[pin]):
    if (now - lastReadTime[pin] > debounceTime):
      // Status geändert und stabil für >50ms
      lastStableState[pin] = currentState
      emitStateChange(pin, currentState)
  
  lastReadTime[pin] = now
```

---

### IgnitionService

**Datei**: `server/services/IgnitionService.js`

**Verantwortlichkeiten**:
- Überwachung des Zündschloss-Status
- Ausführung von Power-Saving-Skripten
- Verwaltung von ON/OFF-Übergängen

**Konfiguration** (aus `gpio-mapping.js`):
```javascript
ignition: {
  enabled: true,
  pin: 21,
  activeOn: 0,
  scripts: {
    lowPower: './scripts/low-power.sh',
    wake: './scripts/wake.sh'
  }
}
```

**Zustände**:
```javascript
'ON'   // Zündung an
'OFF'  // Zündung aus
null   // Initial/Unbekannt
```

**Ablauf**:
```javascript
GPIO 21 ändert sich:
  Neuen Status lesen
  
  Falls Übergang OFF→ON:
    wake.sh ausführen
    'ignition:on' via WebSocket senden
  
  Falls Übergang ON→OFF:
    low-power.sh ausführen
    'ignition:off' via WebSocket senden
```

---

### TemperatureSensorService

**Datei**: `server/services/TemperatureSensorService.js`

**Verantwortlichkeiten**:
- DS18B20-Sensor via 1-Wire lesen
- Sysfs-Datei parsen
- Temperaturdaten via WebSocket senden

**Lesepfad**:
```
/sys/bus/w1/devices/28-xxxxxxxxxxxx/w1_slave
```

**Leseformat**:
```
7d 01 4b 46 7f ff 0c 10 57 : crc=57 YES
7d 01 4b 46 7f ff 0c 10 57 t=23812
                             ^^^^^^
                             23.812°C (Rohwert)
```

**Parsing**:
```javascript
readTemperature():
  1. w1_slave-Datei lesen
  2. Nach Muster "t=XXXXX" suchen
  3. Wert extrahieren (z.B. 23812)
  4. Umrechnen: 23812 / 1000 = 23.8°C
  5. Via WebSocket senden
```

---

### FuelSensorService

**Datei**: `server/services/FuelSensorService.js`

**Verantwortlichkeiten**:
- ADC ADS1115 via I2C lesen
- Umrechnung Spannung → Kraftstoffprozent
- Kalibrierung anwenden
- Daten via WebSocket senden

**Algorithmus**:
```javascript
readFuelLevel():
  1. Spannung von ADS1115 Kanal A0 lesen
  2. Spannungsteiler-Korrektur anwenden:
     V_real = V_measured × ((R1+R2)/R2)
  3. Kalibrierung anwenden:
     percentage = ((V_real - V_empty) / (V_full - V_empty)) × 100
  4. Zwischen 0-100% begrenzen
  5. Via WebSocket senden
```

**Beispiel**:
```javascript
V_measured = 2.5V
Spannungsteiler: R1=100kΩ, R2=33kΩ
Kalibrierung: V_empty=0.5V, V_full=4.0V

V_real = 2.5 × (133/33) = 10.08V
percentage = ((10.08 - 0.5) / (4.0 - 0.5)) × 100
           = (9.58 / 3.5) × 100
           = 273.7% → begrenzt → 100%
```

---

### WebSocketService (Server)

**Datei**: `server/services/WebSocketService.js`

**Verantwortlichkeiten**:
- Verwaltung von Socket.IO-Verbindungen
- Broadcasting von Events an alle Clients
- Verwaltung von Namespaces und Rooms (zukünftig)

**Gesendete Events**:
```javascript
'status'           // Server-/OBD-Verbindungsstatus
'obd:data'         // Einzelne PID-Daten
'obd:scan'         // PID-Scan-Ergebnisse
'gpio:warning'     // Warnleuchten-Statusänderung
'sensor:temp'      // Außentemperatur
'sensor:fuel'      // Kraftstoffstand
'ignition:on'      // Zündung an
'ignition:off'     // Zündung aus
'error'            // Allgemeiner Fehler
```

**Empfangene Events**:
```javascript
'force-restart'    // Client fordert Server-Neustart an
```

**Methoden**:
```javascript
emitStatus(data)           // Status senden
emitOBDData(data)          // PID-Daten senden
emitWarning(key, state)    // Warnleuchten-Status senden
emitTemperature(temp)      // Temperatur senden
emitFuelLevel(level)       // Kraftstoffstand senden
emitIgnitionState(state)   // Zündungs-Status senden
```

---

## 🎨 Client-Module

### State Management (Valtio)

**Datei**: `client/src/store/state.tsx`

**Globaler Store**:
```typescript
export const state = proxy({
  // OBD-Daten
  obd: {
    rpm: 0,
    speed: 0,
    coolantTemp: 0,
    intakeTemp: 0,
    throttle: 0,
    engineLoad: 0,
    // ... weitere PIDs
  },
  
  // Fahrzeug-Warnleuchten
  warnings: {
    highBeam: false,
    lowBeam: false,
    turnSignals: false,
    battery: false,
    engineOil: false,
    // ... weitere Warnleuchten
  },
  
  // Sensoren
  sensors: {
    temperature: null,
    fuel: null,
  },
  
  // System
  system: {
    connected: false,
    ignition: null,
  }
})
```

**Verwendung in Komponenten**:
```typescript
function Tachometer() {
  const snap = useSnapshot(state);
  const rpm = snap.obd.rpm;
  
  return <div>RPM: {rpm}</div>;
}
```

---

### WebSocketService (Client)

**Datei**: `client/src/services/WebSocketService.ts`

**Betriebsmodi**:
```typescript
1. Mock Mode (websocket.mock = true)
   → MockAnimationService generiert simulierte Daten
   → Keine Socket.IO-Verbindung

2. Real Mode (websocket.mock = false)
   → Socket.IO-Verbindung zum Server
   → Echte Daten von Hardware
```

**Empfangene Events**:
```typescript
socket.on('status', handleStatus)
socket.on('obd:data', handleOBDData)
socket.on('gpio:warning', handleWarning)
socket.on('sensor:temp', handleTemperature)
socket.on('sensor:fuel', handleFuel)
socket.on('ignition:on', handleIgnitionOn)
socket.on('ignition:off', handleIgnitionOff)
```

**Handler**:
```typescript
handleOBDData(data) {
  // state.obd mit neuen PID-Werten aktualisieren
  state.obd.rpm = data.parameters.rpm?.value || 0;
  state.obd.speed = data.parameters.speed?.value || 0;
  // ...
}

handleWarning(data) {
  // state.warnings aktualisieren
  state.warnings[data.key] = data.state;
}
```

---

### MockAnimationService

**Datei**: `client/src/services/MockAnimationService.ts`

**Verantwortlichkeiten**:
- Simulation realistischer Daten für Entwicklung
- Flüssige RPM/Geschwindigkeitsanimationen
- Warnleuchten-Ein-/Ausschalt-Zyklen

**Animationen**:
```typescript
// RPM: 800 (Leerlauf) ↔ 5500 (Drehzahlbegrenzung)
rpm: Math.sin(t * 0.5) * 2000 + 2500

// Geschwindigkeit: 0 ↔ 120 km/h
speed: Math.abs(Math.sin(t * 0.3)) * 120

// Warnleuchten: Zufälliger Toggle alle 3-5 Sekunden
warnings.turnSignals = Math.random() > 0.8
```

---

## 🔐 Sicherheit und Berechtigungen

### Erforderliche Benutzerberechtigungen

```bash
# Serielle Schnittstelle
sudo usermod -a -G dialout $USER

# GPIO
sudo usermod -a -G gpio $USER

# I2C
sudo usermod -a -G i2c $USER
```

### Ignition-Skripte (Low-Power / Wake)

Die Skripte werden mit den Berechtigungen des aktuellen Benutzers ausgeführt.  
Für privilegierte Operationen (z.B. Shutdown):

```bash
# Sudo NOPASSWD für spezifische Befehle konfigurieren
sudo visudo

# Hinzufügen:
pi ALL=(ALL) NOPASSWD: /sbin/shutdown
```

### WebSocket-Sicherheit

Derzeit ist Socket.IO **nicht authentifiziert**.

**Für Produktion** sollten Sie in Betracht ziehen:
```javascript
// Server
io.use((socket, next) => {
  const token = socket.handshake.auth.token;
  if (isValidToken(token)) {
    next();
  } else {
    next(new Error('Unauthorized'));
  }
});

// Client
const socket = io(url, {
  auth: { token: 'secret-token' }
});
```

---

## 🧪 Testing

### Lokaler Test (Mock-Modus)

```bash
cd client
npm run dev
```

In `environment.ts` einstellen: `websocket.mock = true`

### Integrationstest (mit Server)

```bash
# Terminal 1 (Raspberry Pi oder lokal)
cd server
node server.js

# Terminal 2
cd client
npm run dev
```

In `environment.ts` einstellen: `websocket.mock = false`

### Electron-Test

```bash
npm start
```

---

## 📈 Performance

### Implementierte Optimierungen

1. **GPIO-Debouncing**: Reduziert Störimpulse (50ms)
2. **Optimiertes OBD-Polling**: 200ms zwischen PIDs (ausgewogen)
3. **Lazy Loading Components**: React.lazy() für Code-Splitting
4. **Memoization**: useMemo() für rechenintensive Operationen
5. **Listen-Virtualisierung**: Für Logs und erweiterte Daten

### Zielmetriken

- **Latenz OBD→UI**: <500ms
- **GPIO-Antwort**: <100ms
- **UI-Framerate**: 60 FPS
- **Raspberry-Speicher**: <200MB Server + <500MB Electron

---

## 🔄 Erweiterbarkeit

### Neuen OBD-PID hinzufügen

1. Definition in `PIDParserService.js` hinzufügen:

```javascript
getPIDDefinitions() {
  return [
    // ... bestehende
    {
      pid: '0143',
      name: 'Absolute Load Value',
      key: 'absoluteLoad'
    }
  ]
}
```

2. Parsing in `parseResponse()` hinzufügen:

```javascript
if (pid === '0143') {
  const A = parseInt(bytes[2] + bytes[3], 16);
  const B = parseInt(bytes[4] + bytes[5], 16);
  return {
    value: ((A * 256) + B) * 100 / 255,
    unit: '%'
  };
}
```

3. `state.tsx` Client aktualisieren:

```typescript
obd: {
  // ... bestehende
  absoluteLoad: 0
}
```

### Neue GPIO-Warnleuchte hinzufügen

1. Optokoppler an gewünschten Pin verkabeln (z.B. GPIO 26)

2. Mapping in `gpio-mapping.js` hinzufügen:

```javascript
mapping: {
  // ... bestehende
  customWarning: {
    pin: 26,
    name: 'Custom-Warnung',
    description: 'Beschreibung der Custom-Warnleuchte'
  }
}
```

3. `state.tsx` Client aktualisieren:

```typescript
warnings: {
  // ... bestehende
  customWarning: false
}
```

4. Visuelle Komponente in `WarningLights.tsx` hinzufügen

### Neuen Sensor hinzufügen

Beispiel: Luftdruck BMP280

1. Bibliothek installieren: `npm install i2c-bus bmp280-sensor`

2. Service erstellen: `server/services/PressureSensorService.js`

```javascript
const BMP280 = require('bmp280-sensor');

class PressureSensorService {
  constructor(webSocketService) {
    this.ws = webSocketService;
    this.sensor = null;
    this.interval = null;
  }
  
  async initialize() {
    this.sensor = await BMP280({ address: 0x76 });
  }
  
  startReading() {
    this.interval = setInterval(async () => {
      const data = await this.sensor.read();
      this.ws.io.emit('sensor:pressure', {
        pressure: data.pressure,
        altitude: data.altitude
      });
    }, 5000);
  }
  
  stopReading() {
    clearInterval(this.interval);
  }
}
```

3. In `OBDServer.js` integrieren:

```javascript
this.pressureService = new PressureSensorService(this.webSocketService);
await this.pressureService.initialize();
this.pressureService.startReading();
```

4. Handling im Client `WebSocketService.ts` hinzufügen

---

## 📚 Code-Referenzen

### Hauptdateien

| Komponente | Pfad | Zeilen | Verantwortlichkeit |
|------------|------|--------|-------------------|
| **Server** |
| OBDServer | `server/services/OBDServer.js` | 418 | Orchestrierung |
| OBDComm | `server/services/OBDCommunicationService.js` | 220 | ELM327-Kommunikation |
| PIDParser | `server/services/PIDParserService.js` | ~300 | PID-Parsing |
| Monitoring | `server/services/MonitoringService.js` | ~200 | OBD-Polling |
| GPIO | `server/services/GPIOService.js` | ~150 | GPIO-Verwaltung |
| Ignition | `server/services/IgnitionService.js` | ~100 | Power-Management |
| WebSocket | `server/services/WebSocketService.js` | ~100 | Kommunikation |
| **Client** |
| App | `client/src/App.tsx` | 83 | Einstiegspunkt |
| State | `client/src/store/state.tsx` | ~150 | State-Management |
| WebSocket | `client/src/services/WebSocketService.ts` | ~200 | Server-Verbindung |
| Cockpit | `client/src/routes/Cockpit/Cockpit.tsx` | ~300 | Haupt-Dashboard |

---

**Zuletzt aktualisiert**: v0.9.0
