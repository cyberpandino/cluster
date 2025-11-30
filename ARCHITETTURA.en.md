[🇬🇧 English](ARCHITETTURA.en.md) | [🇮🇹 Italiano](ARCHITETTURA.md) | [🇩🇪 Deutsch](ARCHITETTURA.de.md)

---

# 🏗️ Architecture - PandaOS Cluster

Technical documentation of the system architecture.

> 💡 **Note on Stack**: Yes, we use JavaScript on a car. Yes, we know it's crazy. No, we don't regret it. See [README - But React + Electron on Automotive?!](README.en.md#-but-react--electron-on-automotive-are-you-crazy) for the complete justification.

---

## 📊 General Overview

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
│  • 12V warning lights (optocouplers)                         │
│  • Fuel sensor (0-12V analog)                                │
│  • Ignition switch (12V on/off)                              │
└─────────────────────────────────────────────────────────────┘
```

---

## 🔄 Data Flow

### 1. System Startup

```
1. Electron (main.js)
   ↓
2. Start Server (port 3001)
   ↓
3. Start Vite Client (port 5173)
   ↓
4. Electron loads http://localhost:5173
   ↓
5. Client connects via WebSocket to ws://localhost:3001
   ↓
6. Server initializes services (GPIO, OBD, Sensors)
   ↓
7. System operational
```

### 2. OBD Data Reading

```
ECU (Magneti Marelli)
   ↓ (K-Line ISO 9141-2)
ELM327 Adapter
   ↓ (Serial USB 38400 baud)
Raspberry Pi /dev/ttyUSB0
   ↓
OBDCommunicationService
   ↓ (AT / PID Commands)
PIDParserService
   ↓ (Parsing hex → values)
MonitoringService
   ↓ (Continuous polling)
WebSocketService
   ↓ (Socket.IO emit)
Client WebSocketService
   ↓ (State update)
Valtio State
   ↓ (React re-render)
UI Components (Tachometer, Odometer, etc.)
```

### 3. Warning Lights Detection

```
Vehicle warning light 12V
   ↓
Optocoupler (PC817)
   ↓
GPIO Pin (e.g. GPIO 17)
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

### 4. Sensor Reading

#### Temperature (DS18B20)
```
DS18B20 Sensor
   ↓ (1-Wire GPIO 4)
/sys/bus/w1/devices/28-xxx/w1_slave
   ↓ (file read every 5s)
TemperatureSensorService
   ↓
WebSocketService
   ↓
Client State
   ↓
Temperature Component
```

#### Fuel (ADS1115)
```
Vehicle fuel sensor (0-12V)
   ↓ (voltage divider)
ADS1115 Channel A0 (0-4V)
   ↓ (I2C 0x48)
FuelSensorService
   ↓ (read every 500ms)
Voltage → percentage calibration
   ↓
WebSocketService
   ↓
Client State
   ↓
Fuel Component
```

---

## 📦 Server Modules

### OBDServer (Main Orchestrator)

**File**: `server/services/OBDServer.js`

**Responsibilities**:
- Initialization and coordination of all services
- Lifecycle management (start/stop/restart)
- Retry logic for OBD connection
- Periodic PID scanning
- Error handling and recovery

**Key Methods**:
```javascript
start()                    // Start all services
stop()                     // Stop gracefully
startWithRetry()           // OBD connection with retry
scanAllPIDs()              // Initial PID scan
periodicPIDScan()          // Periodic scan (every 30s)
forceReconnect()           // OBD reconnection
forceRestart()             // Full process restart
```

**Handled Events**:
- `SIGINT` → Graceful shutdown
- `unhandledRejection` → Force restart
- `uncaughtException` → Force restart
- WebSocket `force-restart` → Force restart

---

### OBDCommunicationService

**File**: `server/services/OBDCommunicationService.js`

**Responsibilities**:
- Serial communication with ELM327
- Sending AT and PID commands
- Parsing ELM327 responses
- Timeout and error handling

**Configuration**:
```javascript
portPath: '/dev/ttyUSB0'
baudRate: 38400
timeout: 4000ms (default)
```

**Main Methods**:
```javascript
connect()                  // Open serial port
initialize()               // Setup ELM327 (ATZ, ATE0, etc.)
sendCommand(cmd)           // Send raw command
waitForResponse(timeout)   // Wait for response with timeout
readPID(pid, name)         // Read single PID
wakeupECU()                // Wake ECU if sleeping
disconnect()               // Close port
```

**Initialization Commands**:
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

**Responsibilities**:
- OBD-II PID definitions
- Parsing hex responses → physical values
- Conversion formulas specific to Magneti Marelli ECU

**Supported PIDs**:
```javascript
'010C' // RPM (Engine speed)
'010D' // Speed (Vehicle speed)
'0105' // Coolant Temperature
'010F' // Intake Air Temperature
'0111' // Throttle Position
'0104' // Engine Load
'010A' // Fuel Pressure
'010B' // Intake Manifold Pressure
'0106' // Short Term Fuel Trim
'0107' // Long Term Fuel Trim
'0142' // Control Module Voltage
```

**Parsing Example**:
```javascript
// PID 010C (RPM)
// Response: 41 0C 1A F8
// A=1A(hex)=26(dec), B=F8(hex)=248(dec)
// RPM = ((A*256)+B)/4 = (6656+248)/4 = 1726 RPM

parseResponse(pid, response, name) {
  // ... PID-specific parsing logic
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

**Responsibilities**:
- Continuous polling of working PIDs
- Real-time data transmission via WebSocket
- Watchdog for freeze/timeout detection
- Dynamic PID list management

**Configuration**:
```javascript
pollingDelay: 200ms        // Delay between PID reads
watchdogInterval: 30s      // Check activity every 30s
watchdogTimeout: 60s       // Inactivity timeout
```

**Methods**:
```javascript
startMonitoring(workingPIDs)  // Start polling
stopMonitoring()              // Stop polling
updateWorkingPIDs(newList)    // Update PIDs to monitor
isPIDCurrentlyMonitored(key)  // Check if PID active
startWatchdog()               // Start watchdog
```

**Monitoring Flow**:
```javascript
Infinite loop:
  For each PID in workingPIDs:
    1. Read PID from ECU
    2. Emit data via WebSocket
    3. Wait 200ms
  Repeat
```

---

### GPIOService

**File**: `server/services/GPIOService.js`

**Responsibilities**:
- GPIO pin initialization
- Vehicle warning lights polling
- Signal debouncing
- State change event emission

**Configuration**:
```javascript
pollingInterval: 100ms     // GPIO read frequency
debounceTime: 50ms         // Anti-bounce
```

**Methods**:
```javascript
initializeGPIO()           // Setup all pins
startPolling()             // Start GPIO polling
stopPolling()              // Stop polling
readGPIOState(pin)         // Read single pin
cleanup()                  // Release GPIO resources
```

**Debouncing Algorithm**:
```javascript
lastStableState[pin] = null
lastReadTime[pin] = 0

onPoll():
  currentState = gpio.read(pin)
  now = Date.now()
  
  if (currentState != lastStableState[pin]):
    if (now - lastReadTime[pin] > debounceTime):
      // State changed and stable for >50ms
      lastStableState[pin] = currentState
      emitStateChange(pin, currentState)
  
  lastReadTime[pin] = now
```

---

### IgnitionService

**File**: `server/services/IgnitionService.js`

**Responsibilities**:
- Ignition switch status monitoring
- Power-saving script execution
- ON/OFF transition handling

**Configuration** (from `gpio-mapping.js`):
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

**States**:
```javascript
'ON'   // Ignition on
'OFF'  // Ignition off
null   // Initial/unknown state
```

**Flow**:
```javascript
GPIO 21 changes:
  Read new state
  
  If transition OFF→ON:
    Execute wake.sh
    Emit 'ignition:on' via WebSocket
  
  If transition ON→OFF:
    Execute low-power.sh
    Emit 'ignition:off' via WebSocket
```

---

### TemperatureSensorService

**File**: `server/services/TemperatureSensorService.js`

**Responsibilities**:
- DS18B20 sensor reading via 1-Wire
- Sysfs file parsing
- Temperature data transmission via WebSocket

**Read Path**:
```
/sys/bus/w1/devices/28-xxxxxxxxxxxx/w1_slave
```

**Read Format**:
```
7d 01 4b 46 7f ff 0c 10 57 : crc=57 YES
7d 01 4b 46 7f ff 0c 10 57 t=23812
                             ^^^^^^
                             23.812°C (raw value)
```

**Parsing**:
```javascript
readTemperature():
  1. Read w1_slave file
  2. Search for pattern "t=XXXXX"
  3. Extract value (e.g. 23812)
  4. Convert: 23812 / 1000 = 23.8°C
  5. Emit via WebSocket
```

---

### FuelSensorService

**File**: `server/services/FuelSensorService.js`

**Responsibilities**:
- ADS1115 ADC reading via I2C
- Voltage → fuel percentage conversion
- Calibration application
- Data transmission via WebSocket

**Algorithm**:
```javascript
readFuelLevel():
  1. Read voltage from ADS1115 channel A0
  2. Apply divider correction:
     V_real = V_measured × ((R1+R2)/R2)
  3. Apply calibration:
     percentage = ((V_real - V_empty) / (V_full - V_empty)) × 100
  4. Clamp between 0-100%
  5. Emit via WebSocket
```

**Example**:
```javascript
V_measured = 2.5V
Divider: R1=100kΩ, R2=33kΩ
Calibration: V_empty=0.5V, V_full=4.0V

V_real = 2.5 × (133/33) = 10.08V
percentage = ((10.08 - 0.5) / (4.0 - 0.5)) × 100
           = (9.58 / 3.5) × 100
           = 273.7% → clamp → 100%
```

---

### WebSocketService (Server)

**File**: `server/services/WebSocketService.js`

**Responsibilities**:
- Socket.IO connection management
- Broadcasting events to all clients
- Namespace and rooms handling (future)

**Emitted Events**:
```javascript
'status'           // Server/OBD connection status
'obd:data'         // Single PID data
'obd:scan'         // PID scan results
'gpio:warning'     // Warning light state change
'sensor:temp'      // External temperature
'sensor:fuel'      // Fuel level
'ignition:on'      // Ignition on
'ignition:off'     // Ignition off
'error'            // Generic error
```

**Received Events**:
```javascript
'force-restart'    // Client requests server restart
```

**Methods**:
```javascript
emitStatus(data)           // Send status
emitOBDData(data)          // Send PID data
emitWarning(key, state)    // Send warning state
emitTemperature(temp)      // Send temperature
emitFuelLevel(level)       // Send fuel level
emitIgnitionState(state)   // Send ignition state
```

---

## 🎨 Client Modules

### State Management (Valtio)

**File**: `client/src/store/state.tsx`

**Global Store**:
```typescript
export const state = proxy({
  // OBD data
  obd: {
    rpm: 0,
    speed: 0,
    coolantTemp: 0,
    intakeTemp: 0,
    throttle: 0,
    engineLoad: 0,
    // ... other PIDs
  },
  
  // Vehicle warnings
  warnings: {
    highBeam: false,
    lowBeam: false,
    turnSignals: false,
    battery: false,
    engineOil: false,
    // ... other warnings
  },
  
  // Sensors
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

**Usage in Components**:
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

**Operation Modes**:
```typescript
1. Mock Mode (websocket.mock = true)
   → MockAnimationService generates simulated data
   → No Socket.IO connection

2. Real Mode (websocket.mock = false)
   → Socket.IO connection to server
   → Real data from hardware
```

**Listened Events**:
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
  // Update state.obd with new PID values
  state.obd.rpm = data.parameters.rpm?.value || 0;
  state.obd.speed = data.parameters.speed?.value || 0;
  // ...
}

handleWarning(data) {
  // Update state.warnings
  state.warnings[data.key] = data.state;
}
```

---

### MockAnimationService

**File**: `client/src/services/MockAnimationService.ts`

**Responsibilities**:
- Realistic data simulation for development
- Smooth RPM/speed animations
- Warning light on/off cycles

**Animations**:
```typescript
// RPM: 800 (idle) ↔ 5500 (redline)
rpm: Math.sin(t * 0.5) * 2000 + 2500

// Speed: 0 ↔ 120 km/h
speed: Math.abs(Math.sin(t * 0.3)) * 120

// Warnings: Random toggle every 3-5 seconds
warnings.turnSignals = Math.random() > 0.8
```

---

## 🔐 Security and Permissions

### Required User Permissions

```bash
# Serial port
sudo usermod -a -G dialout $USER

# GPIO
sudo usermod -a -G gpio $USER

# I2C
sudo usermod -a -G i2c $USER
```

### Ignition Scripts (Low-Power / Wake)

Scripts are executed with current user permissions.  
For privileged operations (e.g. shutdown):

```bash
# Configure sudo NOPASSWD for specific commands
sudo visudo

# Add:
pi ALL=(ALL) NOPASSWD: /sbin/shutdown
```

### WebSocket Security

Currently Socket.IO is **unauthenticated**.

**For production**, consider:
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

### Local Testing (Mock Mode)

```bash
cd client
npm run dev
```

Set `websocket.mock = true` in `environment.ts`

### Integration Testing (with Server)

```bash
# Terminal 1 (Raspberry Pi or local)
cd server
node server.js

# Terminal 2
cd client
npm run dev
```

Set `websocket.mock = false` in `environment.ts`

### Electron Testing

```bash
npm start
```

---

## 📈 Performance

### Implemented Optimizations

1. **GPIO Debouncing**: Reduces spurious calls (50ms)
2. **Optimized OBD Polling**: 200ms between PIDs (balanced)
3. **Lazy loading components**: React.lazy() for code-splitting
4. **Memoization**: useMemo() for heavy calculations
5. **List virtualization**: For logs and extended data

### Target Metrics

- **OBD→UI Latency**: <500ms
- **GPIO Response**: <100ms
- **UI Frame rate**: 60 FPS
- **Raspberry Memory**: <200MB server + <500MB Electron

---

## 🔄 Extensibility

### Adding New OBD PID

1. Add definition in `PIDParserService.js`:

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

2. Add parsing in `parseResponse()`:

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

3. Update client `state.tsx`:

```typescript
obd: {
  // ... existing
  absoluteLoad: 0
}
```

### Adding New GPIO Warning Light

1. Wire optocoupler to desired pin (e.g. GPIO 26)

2. Add mapping in `gpio-mapping.js`:

```javascript
mapping: {
  // ... existing
  customWarning: {
    pin: 26,
    name: 'Custom Warning',
    description: 'Custom warning light description'
  }
}
```

3. Update client `state.tsx`:

```typescript
warnings: {
  // ... existing
  customWarning: false
}
```

4. Add visual component in `WarningLights.tsx`

### Adding New Sensor

Example: BMP280 atmospheric pressure

1. Install library: `npm install i2c-bus bmp280-sensor`

2. Create service: `server/services/PressureSensorService.js`

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

3. Integrate in `OBDServer.js`:

```javascript
this.pressureService = new PressureSensorService(this.webSocketService);
await this.pressureService.initialize();
this.pressureService.startReading();
```

4. Add handling in client `WebSocketService.ts`

---

## 📚 Code References

### Main Files

| Component | Path | Lines | Responsibility |
|-----------|------|-------|----------------|
| **Server** |
| OBDServer | `server/services/OBDServer.js` | 418 | Orchestration |
| OBDComm | `server/services/OBDCommunicationService.js` | 220 | ELM327 communication |
| PIDParser | `server/services/PIDParserService.js` | ~300 | PID parsing |
| Monitoring | `server/services/MonitoringService.js` | ~200 | OBD polling |
| GPIO | `server/services/GPIOService.js` | ~150 | GPIO management |
| Ignition | `server/services/IgnitionService.js` | ~100 | Power management |
| WebSocket | `server/services/WebSocketService.js` | ~100 | Communication |
| **Client** |
| App | `client/src/App.tsx` | 83 | Entry point |
| State | `client/src/store/state.tsx` | ~150 | State management |
| WebSocket | `client/src/services/WebSocketService.ts` | ~200 | Server connection |
| Cockpit | `client/src/routes/Cockpit/Cockpit.tsx` | ~300 | Main dashboard |

---

**Last update**: v0.9.0
