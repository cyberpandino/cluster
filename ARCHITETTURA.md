# 🏗️ Architettura - PandaOS Cluster

Documentazione tecnica dell'architettura del sistema.

---

## 📊 Panoramica Generale

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
│  (Raspberry Pi 4B - Sensori - Attuatori)                    │
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
│  • Spie luminose 12V (optoaccoppiatori)                     │
│  • Sensore carburante (0-12V analogico)                     │
│  • Quadro accensione (12V on/off)                           │
└─────────────────────────────────────────────────────────────┘
```

---

## 🔄 Flusso Dati

### 1. Avvio Sistema

```
1. Electron (main.js)
   ↓
2. Avvia Server (port 3001)
   ↓
3. Avvia Client Vite (port 5173)
   ↓
4. Electron carica http://localhost:5173
   ↓
5. Client si connette via WebSocket a ws://localhost:3001
   ↓
6. Server inizializza servizi (GPIO, OBD, Sensori)
   ↓
7. Sistema operativo
```

### 2. Lettura Dati OBD

```
ECU (Magneti Marelli)
   ↓ (K-Line ISO 9141-2)
ELM327 Adapter
   ↓ (Serial USB 38400 baud)
Raspberry Pi /dev/ttyUSB0
   ↓
OBDCommunicationService
   ↓ (Comandi AT / PID)
PIDParserService
   ↓ (Parsing hex → valori)
MonitoringService
   ↓ (Polling continuo)
WebSocketService
   ↓ (Socket.IO emit)
Client WebSocketService
   ↓ (State update)
Valtio State
   ↓ (React re-render)
UI Components (Tachometer, Odometer, etc.)
```

### 3. Rilevamento Spie

```
Spia veicolo 12V
   ↓
Optoaccoppiatore (PC817)
   ↓
GPIO Pin (es. GPIO 17)
   ↓
GPIOService (polling 100ms)
   ↓ (debounce 50ms)
WebSocketService
   ↓ (Socket.IO emit)
Client
   ↓
State.warnings
   ↓
WarningLights Component
```

### 4. Lettura Sensori

#### Temperatura (DS18B20)
```
DS18B20 Sensor
   ↓ (1-Wire GPIO 4)
/sys/bus/w1/devices/28-xxx/w1_slave
   ↓ (lettura file ogni 5s)
TemperatureSensorService
   ↓
WebSocketService
   ↓
Client State
   ↓
Temperature Component
```

#### Carburante (ADS1115)
```
Sensore carburante veicolo (0-12V)
   ↓ (partitore resistivo)
ADS1115 Canale A0 (0-4V)
   ↓ (I2C 0x48)
FuelSensorService
   ↓ (lettura ogni 500ms)
Calibrazione voltage → percentage
   ↓
WebSocketService
   ↓
Client State
   ↓
Fuel Component
```

---

## 📦 Moduli Server

### OBDServer (Orchestratore Principale)

**File**: `server/services/OBDServer.js`

**Responsabilità**:
- Inizializzazione e coordinamento tutti i servizi
- Gestione ciclo di vita (start/stop/restart)
- Retry logic per connessione OBD
- Scansione periodica PID
- Error handling e recovery

**Metodi Chiave**:
```javascript
start()                    // Avvia tutti i servizi
stop()                     // Ferma gracefully
startWithRetry()           // Connessione OBD con retry
scanAllPIDs()              // Scansione iniziale PID supportati
periodicPIDScan()          // Scansione periodica (ogni 30s)
forceReconnect()           // Riconnessione OBD
forceRestart()             // Riavvio completo processo
```

**Eventi Gestiti**:
- `SIGINT` → Shutdown graceful
- `unhandledRejection` → Force restart
- `uncaughtException` → Force restart
- WebSocket `force-restart` → Force restart

---

### OBDCommunicationService

**File**: `server/services/OBDCommunicationService.js`

**Responsabilità**:
- Comunicazione seriale con ELM327
- Invio comandi AT e PID
- Parsing risposte ELM327
- Gestione timeout e errori

**Configurazione**:
```javascript
portPath: '/dev/ttyUSB0'
baudRate: 38400
timeout: 4000ms (default)
```

**Metodi Principali**:
```javascript
connect()                  // Apre porta seriale
initialize()               // Setup ELM327 (ATZ, ATE0, etc.)
sendCommand(cmd)           // Invia comando raw
waitForResponse(timeout)   // Attende risposta con timeout
readPID(pid, name)         // Legge singolo PID
wakeupECU()                // Risveglia ECU se in sleep
disconnect()               // Chiude porta
```

**Comandi Inizializzazione**:
```javascript
'ATZ'     // Reset ELM327
'ATE0'    // Echo off
'ATL0'    // Linefeeds off
'ATS0'    // Spaces off
'ATSP0'   // Auto protocol detection
```

---

### PIDParserService

**File**: `server/services/PIDParserService.js`

**Responsabilità**:
- Definizione PID OBD-II
- Parsing risposte hex → valori fisici
- Formule di conversione specifiche per ECU Magneti Marelli

**PID Supportati**:
```javascript
'010C' // RPM (Giri motore)
'010D' // Speed (Velocità)
'0105' // Coolant Temperature (Temperatura liquido)
'010F' // Intake Air Temperature
'0111' // Throttle Position (Posizione acceleratore)
'0104' // Engine Load (Carico motore)
'010A' // Fuel Pressure
'010B' // Intake Manifold Pressure
'0106' // Short Term Fuel Trim
'0107' // Long Term Fuel Trim
'0142' // Control Module Voltage
```

**Esempio Parsing**:
```javascript
// PID 010C (RPM)
// Risposta: 41 0C 1A F8
// A=1A(hex)=26(dec), B=F8(hex)=248(dec)
// RPM = ((A*256)+B)/4 = (6656+248)/4 = 1726 RPM

parseResponse(pid, response, name) {
  // ... logica parsing specifica per PID
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

**File**: `server/services/MonitoringService.js`

**Responsabilità**:
- Polling continuo PID funzionanti
- Invio dati real-time via WebSocket
- Watchdog per rilevare freeze/timeout
- Gestione lista PID dinamica

**Configurazione**:
```javascript
pollingDelay: 200ms        // Delay tra letture PID
watchdogInterval: 30s      // Controlla attività ogni 30s
watchdogTimeout: 60s       // Timeout inattività
```

**Metodi**:
```javascript
startMonitoring(workingPIDs)  // Avvia polling
stopMonitoring()              // Ferma polling
updateWorkingPIDs(newList)    // Aggiorna PID da monitorare
isPIDCurrentlyMonitored(key)  // Check se PID attivo
startWatchdog()               // Avvia watchdog
```

**Flusso Monitoring**:
```javascript
Loop infinito:
  Per ogni PID in workingPIDs:
    1. Leggi PID da ECU
    2. Emetti dato via WebSocket
    3. Attendi 200ms
  Ripeti
```

---

### GPIOService

**File**: `server/services/GPIOService.js`

**Responsabilità**:
- Inizializzazione pin GPIO
- Polling stato spie veicolo
- Debouncing segnali
- Emissione eventi cambio stato

**Configurazione**:
```javascript
pollingInterval: 100ms     // Frequenza lettura GPIO
debounceTime: 50ms         // Anti-rimbalzo
```

**Metodi**:
```javascript
initializeGPIO()           // Setup tutti i pin
startPolling()             // Avvia polling GPIO
stopPolling()              // Ferma polling
readGPIOState(pin)         // Legge singolo pin
cleanup()                  // Libera risorse GPIO
```

**Algoritmo Debouncing**:
```javascript
lastStableState[pin] = null
lastReadTime[pin] = 0

onPoll():
  currentState = gpio.read(pin)
  now = Date.now()
  
  if (currentState != lastStableState[pin]):
    if (now - lastReadTime[pin] > debounceTime):
      // Stato cambiato e stabile per >50ms
      lastStableState[pin] = currentState
      emitStateChange(pin, currentState)
  
  lastReadTime[pin] = now
```

---

### IgnitionService

**File**: `server/services/IgnitionService.js`

**Responsabilità**:
- Monitoraggio stato quadro accensione
- Esecuzione script power-saving
- Gestione transizioni ON/OFF

**Configurazione** (da `gpio-mapping.js`):
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

**Stati**:
```javascript
'ON'   // Quadro acceso
'OFF'  // Quadro spento
null   // Stato iniziale/sconosciuto
```

**Flusso**:
```javascript
GPIO 21 cambia:
  Leggi nuovo stato
  
  Se transizione OFF→ON:
    Esegui wake.sh
    Emetti 'ignition:on' via WebSocket
  
  Se transizione ON→OFF:
    Esegui low-power.sh
    Emetti 'ignition:off' via WebSocket
```

---

### TemperatureSensorService

**File**: `server/services/TemperatureSensorService.js`

**Responsabilità**:
- Lettura sensore DS18B20 via 1-Wire
- Parsing file sysfs
- Invio dati temperatura via WebSocket

**Path Lettura**:
```
/sys/bus/w1/devices/28-xxxxxxxxxxxx/w1_slave
```

**Formato Lettura**:
```
7d 01 4b 46 7f ff 0c 10 57 : crc=57 YES
7d 01 4b 46 7f ff 0c 10 57 t=23812
                             ^^^^^^
                             23.812°C (valore raw)
```

**Parsing**:
```javascript
readTemperature():
  1. Leggi file w1_slave
  2. Cerca pattern "t=XXXXX"
  3. Estrai valore (es. 23812)
  4. Converti: 23812 / 1000 = 23.8°C
  5. Emetti via WebSocket
```

---

### FuelSensorService

**File**: `server/services/FuelSensorService.js`

**Responsabilità**:
- Lettura ADC ADS1115 via I2C
- Conversione tensione → percentuale carburante
- Applicazione calibrazione
- Invio dati via WebSocket

**Algoritmo**:
```javascript
readFuelLevel():
  1. Leggi tensione da ADS1115 canale A0
  2. Applica correzione partitore:
     V_reale = V_misurata × ((R1+R2)/R2)
  3. Applica calibrazione:
     percentage = ((V_reale - V_empty) / (V_full - V_empty)) × 100
  4. Clamp tra 0-100%
  5. Emetti via WebSocket
```

**Esempio**:
```javascript
V_misurata = 2.5V
Partitore: R1=100kΩ, R2=33kΩ
Calibrazione: V_empty=0.5V, V_full=4.0V

V_reale = 2.5 × (133/33) = 10.08V
percentage = ((10.08 - 0.5) / (4.0 - 0.5)) × 100
           = (9.58 / 3.5) × 100
           = 273.7% → clamp → 100%
```

---

### WebSocketService (Server)

**File**: `server/services/WebSocketService.js`

**Responsabilità**:
- Gestione connessioni Socket.IO
- Broadcasting eventi a tutti i client
- Gestione namespace e rooms (future)

**Eventi Emessi**:
```javascript
'status'           // Stato server/connessione OBD
'obd:data'         // Dati singolo PID
'obd:scan'         // Risultati scansione PID
'gpio:warning'     // Cambio stato spia
'sensor:temp'      // Temperatura esterna
'sensor:fuel'      // Livello carburante
'ignition:on'      // Quadro acceso
'ignition:off'     // Quadro spento
'error'            // Errore generico
```

**Eventi Ricevuti**:
```javascript
'force-restart'    // Client richiede restart server
```

**Metodi**:
```javascript
emitStatus(data)           // Invia stato
emitOBDData(data)          // Invia dato PID
emitWarning(key, state)    // Invia stato spia
emitTemperature(temp)      // Invia temperatura
emitFuelLevel(level)       // Invia carburante
emitIgnitionState(state)   // Invia stato quadro
```

---

## 🎨 Moduli Client

### State Management (Valtio)

**File**: `client/src/store/state.tsx`

**Store Globale**:
```typescript
export const state = proxy({
  // Dati OBD
  obd: {
    rpm: 0,
    speed: 0,
    coolantTemp: 0,
    intakeTemp: 0,
    throttle: 0,
    engineLoad: 0,
    // ... altri PID
  },
  
  // Spie veicolo
  warnings: {
    highBeam: false,
    lowBeam: false,
    turnSignals: false,
    battery: false,
    engineOil: false,
    // ... altre spie
  },
  
  // Sensori
  sensors: {
    temperature: null,
    fuel: null,
  },
  
  // Sistema
  system: {
    connected: false,
    ignition: null,
  }
})
```

**Utilizzo nei Componenti**:
```typescript
function Tachometer() {
  const snap = useSnapshot(state);
  const rpm = snap.obd.rpm;
  
  return <div>RPM: {rpm}</div>;
}
```

---

### WebSocketService (Client)

**File**: `client/src/services/WebSocketService.ts`

**Modalità Operazione**:
```typescript
1. Mock Mode (websocket.mock = true)
   → MockAnimationService genera dati simulati
   → Nessuna connessione Socket.IO

2. Real Mode (websocket.mock = false)
   → Connessione Socket.IO a server
   → Dati reali da hardware
```

**Eventi Ascoltati**:
```typescript
socket.on('status', handleStatus)
socket.on('obd:data', handleOBDData)
socket.on('gpio:warning', handleWarning)
socket.on('sensor:temp', handleTemperature)
socket.on('sensor:fuel', handleFuel)
socket.on('ignition:on', handleIgnitionOn)
socket.on('ignition:off', handleIgnitionOff)
```

**Handlers**:
```typescript
handleOBDData(data) {
  // Aggiorna state.obd con nuovi valori PID
  state.obd.rpm = data.parameters.rpm?.value || 0;
  state.obd.speed = data.parameters.speed?.value || 0;
  // ...
}

handleWarning(data) {
  // Aggiorna state.warnings
  state.warnings[data.key] = data.state;
}
```

---

### MockAnimationService

**File**: `client/src/services/MockAnimationService.ts`

**Responsabilità**:
- Simulazione dati realistici per sviluppo
- Animazioni fluide RPM/velocità
- Cicli accensione/spegnimento spie

**Animazioni**:
```typescript
// RPM: 800 (idle) ↔ 5500 (redline)
rpm: Math.sin(t * 0.5) * 2000 + 2500

// Velocità: 0 ↔ 120 km/h
speed: Math.abs(Math.sin(t * 0.3)) * 120

// Spie: Toggle casuale ogni 3-5 secondi
warnings.turnSignals = Math.random() > 0.8
```

---

## 🔐 Sicurezza e Permessi

### Permessi Utente Richiesti

```bash
# Porta seriale
sudo usermod -a -G dialout $USER

# GPIO
sudo usermod -a -G gpio $USER

# I2C
sudo usermod -a -G i2c $USER
```

### Script Ignition (Low-Power / Wake)

Gli script vengono eseguiti con permessi utente corrente.  
Per operazioni privilegiate (es. shutdown):

```bash
# Configura sudo NOPASSWD per comandi specifici
sudo visudo

# Aggiungi:
pi ALL=(ALL) NOPASSWD: /sbin/shutdown
```

### WebSocket Security

Attualmente Socket.IO è **non autenticato**.

**Per produzione**, considera:
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

### Test Locale (Mock Mode)

```bash
cd client
npm run dev
```

Imposta `websocket.mock = true` in `environment.ts`

### Test Integrazione (con Server)

```bash
# Terminale 1 (Raspberry Pi o locale)
cd server
node server.js

# Terminale 2
cd client
npm run dev
```

Imposta `websocket.mock = false` in `environment.ts`

### Test Electron

```bash
npm start
```

---

## 📈 Performance

### Ottimizzazioni Implementate

1. **Debouncing GPIO**: Riduce chiamate spurie (50ms)
2. **Polling OBD ottimizzato**: 200ms tra PID (bilanciato)
3. **Lazy loading componenti**: React.lazy() per code-splitting
4. **Memoizzazione**: useMemo() per calcoli pesanti
5. **Virtualizzazione liste**: Per log e dati estesi

### Metriche Target

- **Latency OBD→UI**: <500ms
- **GPIO Response**: <100ms
- **Frame rate UI**: 60 FPS
- **Memoria Raspberry**: <200MB server + <500MB Electron

---

## 🔄 Estendibilità

### Aggiungere Nuovo PID OBD

1. Aggiungi definizione in `PIDParserService.js`:

```javascript
getPIDDefinitions() {
  return [
    // ... existing
    {
      pid: '0143',
      name: 'Absolute Load Value',
      key: 'absoluteLoad'
    }
  ]
}
```

2. Aggiungi parsing in `parseResponse()`:

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

3. Aggiorna `state.tsx` client:

```typescript
obd: {
  // ... existing
  absoluteLoad: 0
}
```

### Aggiungere Nuova Spia GPIO

1. Cabla optoaccoppiatore al pin desiderato (es. GPIO 26)

2. Aggiungi mapping in `gpio-mapping.js`:

```javascript
mapping: {
  // ... existing
  customWarning: {
    pin: 26,
    name: 'Avviso Custom',
    description: 'Descrizione spia custom'
  }
}
```

3. Aggiorna `state.tsx` client:

```typescript
warnings: {
  // ... existing
  customWarning: false
}
```

4. Aggiungi componente visuale in `WarningLights.tsx`

### Aggiungere Nuovo Sensore

Esempio: Pressione atmosferica BMP280

1. Installa libreria: `npm install i2c-bus bmp280-sensor`

2. Crea servizio: `server/services/PressureSensorService.js`

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

3. Integra in `OBDServer.js`:

```javascript
this.pressureService = new PressureSensorService(this.webSocketService);
await this.pressureService.initialize();
this.pressureService.startReading();
```

4. Aggiungi handling nel client `WebSocketService.ts`

---

## 📚 Riferimenti Codice

### File Principali

| Componente | Path | Righe | Responsabilità |
|------------|------|-------|----------------|
| **Server** |
| OBDServer | `server/services/OBDServer.js` | 418 | Orchestrazione |
| OBDComm | `server/services/OBDCommunicationService.js` | 220 | Comunicazione ELM327 |
| PIDParser | `server/services/PIDParserService.js` | ~300 | Parsing PID |
| Monitoring | `server/services/MonitoringService.js` | ~200 | Polling OBD |
| GPIO | `server/services/GPIOService.js` | ~150 | Gestione GPIO |
| Ignition | `server/services/IgnitionService.js` | ~100 | Power management |
| WebSocket | `server/services/WebSocketService.js` | ~100 | Comunicazione |
| **Client** |
| App | `client/src/App.tsx` | 83 | Entry point |
| State | `client/src/store/state.tsx` | ~150 | State management |
| WebSocket | `client/src/services/WebSocketService.ts` | ~200 | Connessione server |
| Cockpit | `client/src/routes/Cockpit/Cockpit.tsx` | ~300 | Dashboard principale |

---

## 📄 Licenza

PandaOS è software libero rilasciato sotto **GNU General Public License v3.0 or later**.  
Vedi [LICENSE](LICENSE) per i dettagli completi.

---

## 👨‍💻 Autori

- **[Matteo Errera](https://github.com/matteoerrera)** - Lead Developer
- **[Roberto Zaccardi](https://github.com/rzaccardi)** - Co-Developer
- **[Ludovico Verde](https://www.instagram.com/ludovico.verdee/)** - Design & Creative

---

**Ultimo aggiornamento**: v0.9.0  
**Team**: PandaOS - Cyberpandino

