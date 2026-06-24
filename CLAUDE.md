# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What This Project Is

PandaOS is a digital dashboard (quadro strumenti) for a Fiat Panda 141, running on a Raspberry Pi 4B. It replaces the original analog instruments with a React-based UI that reads OBD-II data via ELM327 and GPIO signals via optocouplers. The codebase is intentionally built with web tech (React, Electron, Node.js) instead of embedded C/C++ for development speed.

## Project Structure

```
cluster/
├── client/          # React + Vite + TypeScript frontend (ESM)
├── server/          # Node.js backend for OBD-II and GPIO (CommonJS)
├── main.js          # Electron wrapper — loads http://localhost:5173
└── scripts/pm.js    # Cross-platform install/run helper
```

## Commands

### Install all dependencies
```bash
npm run install:all
```
On non-Raspberry systems, if native modules fail to compile (Python 3.13+ / node-gyp incompatibility):
```bash
cd server && npm install --ignore-scripts
```

### Development (local, no hardware)
```bash
# Client only with mock data (no server needed)
npm run client          # starts Vite dev server at http://localhost:5173
```
Ensure `websocket.mock = true` in `client/src/config/environment.ts`.

### Full stack (Raspberry Pi only)
```bash
npm start              # concurrently: server (port 3001) + client (port 5173) + Electron
```

### Individual processes
```bash
npm run server         # Node.js OBD server — exits immediately on non-ARM Linux
npm run electron       # Electron window — waits for client on port 5173
```

### Client build
```bash
cd client && npm run build    # output to client/dist/
```

### Lint (client only)
```bash
cd client && npm run lint
```

There are no tests (server `package.json` has a placeholder `test` script that exits 1).

## Architecture

### Data flow
```
ELM327 (USB serial /dev/ttyUSB0)
  → OBDCommunicationService (SerialPort, 38400 baud)
  → MonitoringService (PID polling loop)
  → WebSocketService (Socket.IO server, port 3001)
  → Client WebSocketService (socket.io-client)
  → Valtio state store (proxy)
  → React components

GPIO pins (optocouplers)
  → GPIOService (onoff, BCM numbering, 100ms polling)
  → WebSocketService → client `gpio-warnings` event → state.warnings

DS18B20 (1-Wire GPIO 4) → TemperatureSensorService → `external-temperature` event
ADS1115 (I2C GPIO 2/3) → FuelSensorService → `fuel-level` event
IgnitionService (GPIO 21) → runs low-power.sh / wake.sh scripts
```

### Server (`server/` — CommonJS)
- **`server.js`**: Entry point; instantiates `OBDServer` with a 2-second start delay.
- **`services/OBDServer.js`**: Orchestrator. On non-ARM-Linux, the process exits immediately unless `DEV_MODE=true` is set. Retries OBD connection indefinitely (force-restarts process after 20 failures). Runs a periodic PID re-scan every 30s.
- **`services/OBDCommunicationService.js`**: SerialPort + ELM327 AT commands. Port hardcoded to `/dev/ttyUSB0` (line 7 — change here if on a different port).
- **`services/GPIOService.js`**: Polls GPIO pins for vehicle warning lights. Configured via `server/config/gpio-mapping.js`.
- **`config/gpio-mapping.js`**: Single source of truth for all GPIO pin assignments, sensor enable flags, and calibration values.
- **`ecosystem.config.js`**: PM2 config for production — update `cwd` path before deploying.

### Client (`client/` — ESM, React 18 + TypeScript + Vite)
- **`src/config/environment.ts`**: Single config file — change `websocket.mock`, `graphics.quality`, locale, timezone here.
- **`src/store/state.tsx`**: Valtio proxy holding all runtime state (`session.*` for OBD data, `warnings.*` for GPIO lights). Components read via `useSnapshot(state)`.
- **`src/services/WebSocketService.ts`**: Singleton (`websocketService`). In mock mode, delegates to `MockAnimationService`; in real mode, connects to Socket.IO and updates state directly.
- **`src/routes/Cockpit/`**: Main dashboard view.
- **`src/components/`**: Individual gauge/indicator components (Tachometer, Temperature, Fuel, WarningLights, ModelViewer for Three.js 3D model, etc.).

### Electron (`main.js`)
Opens a 1920×580 window targeting display index 0 (change `targetDisplayIndex` for multi-monitor setups). Loads `http://localhost:5173`. GPU acceleration flags are enabled for performance on Raspberry Pi.

## Key Configuration Points

| What to change | Where |
|---|---|
| Mock vs real hardware | `client/src/config/environment.ts` → `websocket.mock` |
| Graphics quality (1/2/3) | `client/src/config/environment.ts` → `graphics.quality` |
| OBD serial port | `server/services/OBDCommunicationService.js` line 7 |
| GPIO pin assignments | `server/config/gpio-mapping.js` |
| Sensor enable/disable | `server/config/gpio-mapping.js` → `temperature.enabled`, `fuel.enabled` |
| Target display index | `main.js` → `targetDisplayIndex` |
| PM2 install path | `server/ecosystem.config.js` → `cwd` |

## Hardware Notes

- Server **requires** Linux ARM (Raspberry Pi) and the `onoff` module — it will `process.exit(1)` on any other platform unless `DEV_MODE=true`.
- Optional dependencies (`onoff`, `serialport`, `ads1x15`) are in `optionalDependencies` — install failures don't block `npm install`.
- For dev on macOS/Windows: run client only in mock mode. The server cannot run on non-Raspberry hardware without `DEV_MODE=true` (which disables all hardware I/O).

## Commit Style

Uses [Conventional Commits](https://www.conventionalcommits.org/): `feat:`, `fix:`, `docs:`, etc. New source files must include the GPL-3.0 header (see any existing `.ts`/`.js` file for the template).
