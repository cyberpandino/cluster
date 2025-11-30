[🇬🇧 English](CONFIGURAZIONE.en.md) | [🇮🇹 Italiano](CONFIGURAZIONE.md) | [🇩🇪 Deutsch](CONFIGURAZIONE.de.md)

---

# ⚙️ Client Configuration - PandaOS Cluster

Complete guide to Cyberpandino client configuration.

---

## 📁 Configuration File

The main configuration file is located at:

```
client/src/config/environment.ts
```

---

## 🔧 Configuration Parameters

### 1. WebSocket

Backend server connection configuration.

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
- **Description**: WebSocket server address
- **Default**: `'http://127.0.0.1:3001'`
- **Production Raspberry Pi**: `'http://127.0.0.1:3001'` (localhost)
- **Remote development**: `'http://192.168.x.x:3001'` (Raspberry IP)

**Examples**:
```typescript
// Local (development or production on same device)
url: 'http://127.0.0.1:3001'

// Raspberry Pi on local network
url: 'http://192.168.1.100:3001'

// Server on custom port
url: 'http://127.0.0.1:8080'
```

#### `mock` (boolean)
- **Description**: Client operating mode
- **Values**:
  - `true` = Demo mode with simulated data (doesn't require server)
  - `false` = Real connection to backend server
- **Default**: `true`

**When to use**:
```typescript
// Local development on Mac/Windows/Linux (without Raspberry Pi)
mock: true

// Production on Raspberry Pi with active server
mock: false

// Interface testing without hardware
mock: true
```

#### `reconnectionAttempts` (number)
- **Description**: Maximum number of reconnection attempts
- **Default**: `3`
- **Recommended range**: 2-10
- **Behavior**: After exhausting attempts, switches to mock mode

#### `reconnectionDelay` (number)
- **Description**: Delay between reconnection attempts (milliseconds)
- **Default**: `1000` (1 second)
- **Recommended range**: 500-5000 ms

#### `timeout` (number)
- **Description**: WebSocket connection timeout (milliseconds)
- **Default**: `5000` (5 seconds)
- **Recommended range**: 3000-10000 ms

---

### 2. Splash Screen

Startup screen configuration.

```typescript
splashScreen: {
  path: '/splashscreen.mp4'
}
```

#### `path` (string)
- **Description**: Splash screen video path
- **Default**: `'/splashscreen.mp4'`
- **Supported format**: MP4, WebM
- **File location**: `client/public/splashscreen.mp4`

**Customization**:
```typescript
// Custom video
path: '/custom-splash.mp4'

// Disable splash (use static image or none)
path: '' // Requires SplashScreen component modification
```

---

### 3. Debug

Debug and development mode configuration.

```typescript
debug: {
  enabled: true,
  showConsoleViewer: true,
}
```

#### `enabled` (boolean)
- **Description**: Enable debug functionality
- **Default**: `true`
- **Effects**:
  - Shows detailed logs in browser console
  - Enables debug buttons
  - Shows WebSocket connection information
- **Production**: Set to `false` for optimal performance

#### `showConsoleViewer` (boolean)
- **Description**: Enable integrated console viewer (press `d`)
- **Default**: `true`
- **Features**:
  - Real-time log viewing
  - Error monitoring
  - WebSocket debugging
  - Sensor and GPIO status

**Console Viewer Controls**:
- Press **`d`** to open
- Press **`ESC`** to close
- "Clear" button to clear logs

---

### 4. App

General application configuration.

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
- **Description**: Displayed application name
- **Default**: `"PandaOS Cluster"`
- **Usage**: Window title, splash screen, about

#### `version` (string)
- **Description**: Application version
- **Default**: `"0.9.0"`
- **Format**: Semantic versioning (MAJOR.MINOR.PATCH)

#### `locale` (string)
- **Description**: Application language
- **Default**: `"it"` (Italian)
- **Supported**: 
  - `"it"` - Italian
  - `"en"` - English

#### `timezone` (string)
- **Description**: Timezone for date/time display
- **Default**: `"Europe/Rome"`
- **Format**: IANA timezone database
- **Other examples**:
  - `"Europe/London"`
  - `"America/New_York"`
  - `"Asia/Tokyo"`

#### `timeFormat` (string)
- **Description**: Time format
- **Values**:
  - `"24h"` - 24-hour format (e.g. 15:30)
  - `"12h"` - 12-hour format AM/PM (e.g. 3:30 PM)
- **Default**: `"24h"`

---

## 📋 Configuration Examples

### Local Development Configuration (Mac/Windows)

```typescript
export const environment: EnvironmentConfig = {
  websocket: {
    url: 'http://127.0.0.1:3001',
    mock: true,                    // ← Demo mode
    reconnectionAttempts: 3,
    reconnectionDelay: 1000,
    timeout: 5000,
  },
  splashScreen: {
    path: '/splashscreen.mp4'
  },
  debug: {
    enabled: true,                 // ← Active debug
    showConsoleViewer: true,       // ← Console viewer available
  },
  app: {
    name: "PandaOS Cluster [DEV]",
    version: "0.9.0",
    locale: "en",
    timezone: "Europe/Rome",
    timeFormat: "24h",
  },
};
```

### Production Configuration (Raspberry Pi)

```typescript
export const environment: EnvironmentConfig = {
  websocket: {
    url: 'http://127.0.0.1:3001',
    mock: false,                   // ← Real connection
    reconnectionAttempts: 10,      // ← More attempts for stability
    reconnectionDelay: 2000,
    timeout: 10000,
  },
  splashScreen: {
    path: '/splashscreen.mp4'
  },
  debug: {
    enabled: false,                // ← Debug disabled for performance
    showConsoleViewer: false,
  },
  app: {
    name: "PandaOS Cluster",
    version: "0.9.0",
    locale: "en",
    timezone: "Europe/Rome",
    timeFormat: "24h",
  },
};
```

### Remote Testing Configuration

```typescript
export const environment: EnvironmentConfig = {
  websocket: {
    url: 'http://192.168.1.100:3001', // ← Raspberry Pi IP
    mock: false,                       // ← Real connection
    reconnectionAttempts: 5,
    reconnectionDelay: 1500,
    timeout: 7000,
  },
  splashScreen: {
    path: '/splashscreen.mp4'
  },
  debug: {
    enabled: true,                     // ← Debug for troubleshooting
    showConsoleViewer: true,
  },
  app: {
    name: "PandaOS Cluster [TEST]",
    version: "0.9.0",
    locale: "en",
    timezone: "Europe/Rome",
    timeFormat: "24h",
  },
};
```

---

## 🔄 Development Workflow

### 1. Interface Development (without hardware)

```typescript
websocket.mock = true
debug.enabled = true
```

Start only the client:
```bash
cd client
npm run dev
```

### 2. Integration Testing (with Raspberry Pi)

```typescript
websocket.mock = false
websocket.url = 'http://192.168.1.100:3001'
debug.enabled = true
```

Start server on Raspberry Pi and client locally

### 3. Production (on Raspberry Pi)

```typescript
websocket.mock = false
websocket.url = 'http://127.0.0.1:3001'
debug.enabled = false
```

Start full stack:
```bash
npm start
```

---

## 🔍 Mock vs Real Mode

### Mock Mode (`mock: true`)

**Characteristics**:
- ✅ No server required
- ✅ Realistic simulated data
- ✅ Smooth animations
- ✅ Works on any OS
- ❌ Doesn't read real OBD data
- ❌ Doesn't read GPIO/sensors

**Use**: UI development, demos, visual testing

### Real Mode (`mock: false`)

**Characteristics**:
- ✅ Real OBD data from ECU
- ✅ GPIO reading for vehicle warning lights
- ✅ Temperature and fuel sensors
- ❌ Requires active server
- ❌ Requires hardware (Raspberry Pi + ELM327)

**Use**: Production, hardware testing, complete integration

---

## 🚨 Troubleshooting

### Client won't connect to server

**Symptom**: Console shows connection errors

**Solutions**:
1. Verify `mock: false`
2. Check server URL: `websocket.url`
3. Verify server is running: `npm run server`
4. Check firewall/port 3001 open
5. Test connection: `curl http://127.0.0.1:3001`

### Splash screen won't load

**Symptom**: White screen at startup

**Solutions**:
1. Verify file exists: `client/public/splashscreen.mp4`
2. Check path in configuration
3. Verify video format (MP4 preferred)
4. Check browser console for errors

### Debug console won't open

**Symptom**: Pressing `d` does nothing

**Solutions**:
1. Verify `debug.showConsoleViewer: true`
2. Check `debug.enabled: true`
3. Check browser console for errors
4. Try `CTRL+D` or `CMD+D`

### Wrong timezone/time

**Symptom**: Incorrect time in cluster

**Solutions**:
1. Verify `app.timezone` correct for your zone
2. Timezone list: https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
3. Verify Raspberry Pi system time: `timedatectl`
4. Sync NTP: `sudo timedatectl set-ntp true`

---

## 📚 References

- **Source code**: `client/src/config/environment.ts`
- **TypeScript types**: `client/src/config/types.ts`
- **Main documentation**: `README.en.md`
- **Server configuration**: `server/CONFIGURAZIONE_SERVER.en.md`

---

**Last update**: v0.9.0
