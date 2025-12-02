[🇬🇧 English](README.en.md) | [🇮🇹 Italiano](../README.md) | [🇩🇪 Deutsch](README.de.md)

---

# 🚗 Cyberpandino Cluster - PandaOS

[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)
[![Version](https://img.shields.io/badge/version-0.9.0-green.svg)](https://github.com/cyberpandino/cluster/releases)
[![Node](https://img.shields.io/badge/node-%3E%3D18.0.0-brightgreen.svg)](https://nodejs.org/)
[![Platform](https://img.shields.io/badge/platform-Raspberry%20Pi%204B%2F5-red.svg)](https://www.raspberrypi.com/)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](https://github.com/cyberpandino/cluster/blob/main/.github/CONTRIBUTING.en.md)

Digital instrument cluster for Fiat Panda 141 based on Raspberry Pi 4B.

## 📋 Description

Complete digital instrument cluster system that replaces the original analog instrumentation of the Fiat Panda 141. The system interfaces with the ECU via OBD-II protocol (ELM327) and reads warning lights through optocouplers connected to the Raspberry Pi GPIO pins.

### Main Features (v0.9.0)

- ✅ **OBD-II data reading**: Speed, engine RPM, temperature, oil pressure, etc.
- ✅ **Vehicle warning lights detection**: High beams, low beams, turn signals, oil level, etc.
- ✅ **External sensors**:
  - External temperature (DS18B20)
  - Fuel level (ADS1115)
- ✅ **Ignition management**: Automatic power-saving system
- ✅ **Modern interface**: 3D dashboard with interactive Panda model
- ✅ **Demo mode**: For development without hardware

---

## 📸 Preview

### Main Dashboard

The digital cluster completely replaces the original analog instrument cluster with a modern and customizable interface.

<div align="center">
  <img src="images/dashboard-main.png" alt="Main dashboard" width="800"/>
  <p><em>Main dashboard with interactive 3D model</em></p>
</div>

---

### 🗺️ Future Features

Discover what we're planning: [Roadmap & Wishlist](ROADMAP.en.md)

Some ideas on the list:
- 📹 Rear camera and parking sensors
- 🚪 Advanced 3D animations (doors, lights)
- 🎨 Customizable dashboards and themes
- 🌍 Internationalization
- 📱 Mobile companion app
- And much more!

Want to contribute? Every help is welcome! See the [contributing guide](.github/CONTRIBUTING.en.md).

---

## 📚 Documentation Index

### 🚀 Start Here
- **[Quick Start](QUICK_START.en.md)** - Quick guide to get started
- **[Hardware](HARDWARE.en.md)** - Complete component list and assembly diagram

### 📖 Technical Documentation
- **[Architecture](ARCHITETTURA.en.md)** - Detailed system architecture
- **[General Documentation](DOCUMENTAZIONE.en.md)** - Complete project overview
- **[Client Configuration](client/CONFIGURAZIONE.en.md)** - Frontend setup and configuration
- **[Server Configuration](server/CONFIGURAZIONE_SERVER.en.md)** - Backend setup and configuration
- **[Environment Configuration](client/config/README.en.md)** - Environment variables and parameters

### 🤝 Contribution
- **[How to Contribute](.github/CONTRIBUTING.en.md)** - Complete guide to contributing to the project

### 📋 Other
- **[Roadmap](ROADMAP.en.md)** - Development plan and wishlist
- **[Authors](AUTHORS.en.md)** - Who contributed to the project
- **[License](../LICENSE)** - GNU General Public License v3.0

---

## ⚠️ Disclaimer

PandaOS is a hobbyist and experimental project, born out of technical curiosity and a spirit of digital adventure. It is not a certified product, not designed for production, and makes no claim to meet industrial, automotive, or galactic standards.

All material in this repository, including code, guides, diagrams, and more or less sensible ideas, is provided "AS IS", without warranties of functionality, reliability, or compatibility with your heroic utility vehicle's electrical system.

The authors and contributors assume no responsibility in case of:

* electrical or electronic failures
* abnormal vehicle behavior
* unexpected short circuits
* damage to persons, property, animals, and the like
* any side effects resulting from the use of the software or following the instructions in this documentation

The use of PandaOS on vehicles in circulation or in any context where compliance, homologation, or common sense requirements might be necessary is strongly discouraged. Any installation or experimentation occurs at the user's sole risk, who assumes full responsibility for the technical and practical consequences of their choices.

---

## 🏗️ Architecture

The project consists of three main modules:

```
cluster/
├── client/          → Graphical interface (React + Vite + Electron)
├── server/          → Backend OBD-II and GPIO communication (Node.js)
└── main.js          → Electron wrapper for desktop app
```

### Technologies Used

- **Frontend**: React 18, TypeScript, Three.js, Socket.IO Client
- **Backend**: Node.js, Socket.IO Server, SerialPort, GPIO (onoff)
- **Desktop**: Electron 36
- **Hardware**: Raspberry Pi 4B, ELM327, DS18B20, ADS1115

### 🤔 But React + Electron on Automotive?! Are You Crazy?

Yes, we know. Any embedded engineer seeing this project is probably having a panic attack.

**How things should be done properly:**
- **C/C++** - Because JavaScript on a car is like putting square wheels
- **Qt/QML** - The industry standard (Tesla, Audi, BMW use it)
- **Yocto/Buildroot** - Serious embedded Linux, not Raspberry Pi OS with all the clutter
- **Direct framebuffer** - Not Electron running an entire browser to display 4 numbers

**So why React/Electron/Node.js?**

Because it's a **hobby project** and we want to **have fun**, not go crazy.

**Pros of our questionable approach**:
- ⚡ **Fast to develop** - Seen Three.js? Make a 3D model in 5 minutes. Try with native OpenGL.
- 🎨 **Libraries everywhere** - npm has everything. C++ has... um... boost?
- 🧑‍💻 **Accessible** - Know React? Welcome. Know CMake? Condolences.
- 🐛 **Debug** - F12 and see everything. GDB instead is... an experience.
- 🚀 **Fun** - More time tinkering, less fighting with toolchains
- 💡 **Proof of concept** - Does it work? Good! Then we'll see.

**Cons (that we consciously accept)**:
- 💾 **Eats RAM** like it's pasta (~500MB vs ~50MB)
- 🐌 **Slow boot** (~30s vs ~3s) - but with always-on standby mode it becomes instant
- 🔋 **Consumes** more than it should (but standby consumes only 0.4W, negligible)
- 📊 **JavaScript** - Yes, JavaScript. On a car. Deal with it.

**The point is**: We're talking about a **1990 Panda**. It's not an F-35. It doesn't have to go to the Moon.
It needs to show you the engine RPM in a cool way while you listen to Pink Floyd. And it does that very well. 🚗💨

> 💡 **Want to redo it in C++/Qt "properly"?** Fantastic! We'd be curious to see a native port and would happily help.

---

## ⚙️ System Requirements

### Software Prerequisites

| Software | Minimum Version | Recommended |
|----------|----------------|-------------|
| **Node.js** | 18.0.0 | 20.x LTS |
| **npm** | 9.0.0 | 10.x |
| **Git** | 2.0+ | Latest |

```bash
# Quick check
node --version  # >= v18.0.0
npm --version   # >= 9.0.0
git --version   # >= 2.0.0
```

⚠️ **Raspberry Pi**: Don't use `apt install nodejs` (obsolete version). See [CONFIGURAZIONE_SERVER.en.md](server/CONFIGURAZIONE_SERVER.en.md#2-installing-nodejs-and-npm) for NodeSource/nvm.

---

### For Raspberry Pi (Production)

- **Hardware**:
  - Raspberry Pi 4B (4GB or higher recommended) or Raspberry Pi 5
  - ELM327 USB adapter (serial port `/dev/ttyUSB0`)
  - Optocouplers for warning light detection (PC817 or similar)
  - Ultra-wide LCD display (1920×480 recommended)
  - DS18B20 temperature sensor (optional)
  - ADS1115 ADC converter (optional, for fuel sensor)

📋 **Complete hardware list**: See [HARDWARE.en.md](HARDWARE.en.md) for details on all required components

- **Operating System**:
  - Raspberry Pi OS Lite (64-bit) - Debian-based recommended
  - Boot time: ~30s (optimizable to ~20s, or instant with standby mode)
  - ARM/ARM64 architecture
  
  > 📘 **OS Choice and Boot Time**: See [CONFIGURAZIONE_SERVER.en.md](server/CONFIGURAZIONE_SERVER.en.md#1-operating-system-installation) for details on choosing the right distro, optimizing boot time, and configuring **always-on standby mode** (negligible consumption, instant startup)

- **Software**: See [CONFIGURAZIONE_SERVER.en.md](server/CONFIGURAZIONE_SERVER.en.md#2-installing-nodejs-and-npm) for installation instructions on Raspberry Pi

### For Local Development (Mac/Windows/Linux)

- Node.js 18+ (20 LTS recommended)
- npm 9+ (10.x recommended)
- Git 2.0+

> 💡 **Quick Setup**: See installation instructions in the [Software Prerequisites](#software-prerequisites) section above

⚠️ **NOTE**: Running the project on non-Raspberry Pi systems, the server will fail to start due to missing hardware-specific dependencies (GPIO, sensors, OBD serial port). You can use **mock mode** in the client for development without server.

---

## 🚀 Project Setup

### 1. Clone the Repository

```bash
git clone https://github.com/cyberpandino/cluster
cd cluster
```

### 2. Install Dependencies

The project provides an installation script that sets up all dependencies:

```bash
npm run install:all
```

This command installs dependencies for:
- Root (Electron + concurrently)
- Client (React + frontend dependencies)
- Server (Node.js + hardware dependencies)

### 3. Configuration

#### a) Client Configuration

Edit the client configuration file:

**File**: `client/src/config/environment.ts`

```typescript
export const environment: EnvironmentConfig = {
  websocket: {
    url: 'http://127.0.0.1:3001',  // WebSocket server URL
    mock: true,                      // true = demo mode | false = real connection
    reconnectionAttempts: 3,
    reconnectionDelay: 1000,
    timeout: 5000,
  },
  debug: {
    enabled: true,                   // Enable debug mode
    showConsoleViewer: true,         // Show console viewer (key 'd')
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

**Key Parameters**:
- `websocket.url`: WebSocket server address (default: `http://127.0.0.1:3001`)
- `websocket.mock`: 
  - `true` = Demo mode with simulated animations (for local development)
  - `false` = Real connection to server (for production on Raspberry Pi)
- `debug.enabled`: Enable debug features
- `debug.showConsoleViewer`: Show console viewer (activatable with key `d`)

#### b) Server Configuration

Edit the GPIO and sensor configuration file:

**File**: `server/config/gpio-mapping.js`

See the [GPIO and Sensors Configuration](#-gpio-and-sensors-configuration) section for complete details.

---

## 🎯 Starting the Project

### Full Mode (Raspberry Pi)

Start client, server, and Electron simultaneously:

```bash
npm start
```

This command runs:
1. OBD-II server on port 3001
2. React/Vite client on port 5173
3. Electron desktop app

### Development Mode (Local without Raspberry)

#### Option 1: Client Only (Mock Mode)

1. Ensure `websocket.mock = true` in `client/src/config/environment.ts`
2. Start only the client:

```bash
npm run client
```

The application will be available at `http://localhost:5173` with simulated data.

#### Option 2: Client + Electron

```bash
npm run client    # In one terminal
npm run electron  # In another terminal
```

### Individual Commands

```bash
# Server only (requires Raspberry Pi)
npm run server

# Client only
npm run client

# Electron only (waits for client on port 5173)
npm run electron
```

---

## 🔌 GPIO and Sensors Configuration

### GPIO Mapping for Vehicle Warning Lights

The `server/config/gpio-mapping.js` file contains the complete GPIO pin mapping.

> 📘 **Vehicle Electrical Diagram**: To identify the correct wiring for warning lights on the original Panda dashboard, consult the [Official Fiat Panda 141 Electrical Diagram](http://www.bunkeringegnere.altervista.org/esplosi/FIAT%20PANDA/panda%20141/1100%20mpi/55%20IMPIANTO%20ELETTRICO%20-%20SCHEMI%20-%20GAMMA%202000.pdf) with all color codes and connections.

#### Used Pins

| Warning Light/Function | GPIO Pin (BCM) | Description |
|------------------------|----------------|-------------|
| Turn Signals | 17 | Direction indicators |
| Alternator | 27 | Battery charging |
| Oil Pressure | 22 | Engine oil pressure |
| Brake System | 23 | Brakes |
| Injectors | 24 | Injection system |
| Key ON | 25 | Key inserted |
| High Beam | 5 | High beam headlights |
| Low Beam | 6 | Low beam headlights |
| Hazard Lights | 12 | Emergency lights |
| Fog Lights | 13 | Rear fog light |
| Coolant Temperature | 16 | Coolant fluid |
| Rear Window Heater | 19 | Rear window defroster |
| Fuel Reserve | 20 | Low fuel level |
| Ignition | 21 | Ignition on/off detection |

#### Optocoupler Configuration

```javascript
config: {
  mode: 'BCM',              // Broadcom GPIO numbering
  pullMode: 'PUD_DOWN',     // Internal pull-down resistor
  debounceTime: 50,         // Anti-bounce filter (ms)
  pollingInterval: 100,     // GPIO reading frequency (ms)
}
```

**Operating Logic**:
- `HIGH (1)` = Warning light on
- `LOW (0)` = Warning light off

### Ignition Management

```javascript
ignition: {
  enabled: true,
  pin: 21,                  // Dedicated GPIO pin
  activeOn: 0,              // 0 = active low | 1 = active high
  scripts: {
    lowPower: './scripts/low-power.sh',   // Executed when ignition turns off
    wake: './scripts/wake.sh',             // Executed when ignition turns on
  },
}
```

Power-saving scripts can be customized to:
- Turn off display
- Reduce brightness
- Disable non-essential services
- Start controlled shutdown

### External Temperature Sensor (DS18B20)

```javascript
temperature: {
  enabled: true,
  sensorId: null,           // null = auto-detect first sensor
  basePath: '/sys/bus/w1/devices',
  readInterval: 5000,       // Read interval (ms)
  pin: 4,                   // GPIO 4 (default for 1-Wire)
}
```

**Hardware Setup**:
1. Connect DS18B20 to GPIO 4
2. Enable 1-Wire: `sudo raspi-config` → Interface Options → 1-Wire
3. Verify sensor presence: `ls /sys/bus/w1/devices/`

### Fuel Sensor (ADS1115 - I2C ADC)

```javascript
fuel: {
  enabled: true,
  chip: 0,                  // 0 = ADS1115 | 1 = ADS1015
  channel: 0,               // Channel A0 (0-3 available)
  gain: 4096,               // ±4.096V full-scale
  sampleRate: 250,          // Sample rate (SPS)
  readInterval: 500,        // Read interval (ms)
  
  // Resistive divider configuration
  voltageDivider: {
    r1: 100000,             // 100kΩ
    r2: 33000,              // 33kΩ
  },
  
  // Voltage → percentage calibration
  calibration: {
    voltageEmpty: 0.5,      // Empty tank voltage (V)
    voltageFull: 4.0,       // Full tank voltage (V)
  },
  
  pins: {
    sda: 2,                 // GPIO 2 (I2C SDA)
    scl: 3,                 // GPIO 3 (I2C SCL)
  },
}
```

**Hardware Setup**:
1. Connect ADS1115:
   - VDD → 3.3V
   - GND → GND
   - SDA → GPIO 2
   - SCL → GPIO 3
   - A0 → Fuel sensor (via resistive divider)
2. Enable I2C: `sudo raspi-config` → Interface Options → I2C
3. Verify presence: `sudo i2cdetect -y 1`

### OBD-II Serial Port

**File**: `server/services/OBDCommunicationService.js`

```javascript
constructor() {
  this.portPath = '/dev/ttyUSB0';  // ELM327 port
  this.port = null;
  this.baudRate = 38400;            // Communication speed
}
```

**Hardware Setup**:
1. Connect ELM327 adapter via USB
2. Verify port: `ls -l /dev/ttyUSB*`
3. Grant permissions: `sudo usermod -a -G dialout $USER`
4. Reboot or re-login to apply permissions

**Alternative Serial Port Configuration**:

If the OBD adapter is on a different port (e.g., `/dev/ttyUSB1`, `/dev/ttyACM0`), modify:

```javascript
// In server/services/OBDCommunicationService.js (line 7)
this.portPath = '/dev/ttyUSB1';  // Change here
```

---

## 🔧 PM2 Configuration (Automatic Startup)

To run the server as a system service on Raspberry Pi:

### 1. Install PM2

```bash
sudo npm install -g pm2
```

### 2. Configure Ecosystem

Edit `server/ecosystem.config.js`:

```javascript
module.exports = {
  apps: [{
    name: 'obd-server',
    script: './server.js',
    cwd: '/home/pi/cockpit/server',  // ⚠️ CHANGE THIS PATH
    instances: 1,
    autorestart: true,
    watch: false,
    max_memory_restart: '200M',
    restart_delay: 2000,
    max_restarts: 15,
    min_uptime: '10s',
    exp_backoff_restart_delay: 100,
    env: {
      NODE_ENV: 'production',
      PORT: 3001
    },
    log_file: './logs/obd-combined.log',
    out_file: './logs/obd-out.log',
    error_file: './logs/obd-error.log',
    log_date_format: 'YYYY-MM-DD HH:mm:ss Z',
    merge_logs: true
  }]
};
```

### 3. Start with PM2

```bash
cd server
mkdir -p logs
pm2 start ecosystem.config.js
pm2 save
pm2 startup
```

### 4. Useful PM2 Commands

```bash
pm2 status              # Service status
pm2 logs obd-server     # View logs
pm2 restart obd-server  # Restart service
pm2 stop obd-server     # Stop service
pm2 monit               # Real-time monitor
```

---

## 🛠️ Troubleshooting

### Server doesn't start on non-Raspberry system

**Error**:
```
❌ ERROR: Essential Raspberry Pi dependencies not available
Unsupported platform: darwin arm64 - Linux ARM required
```

**Solution**: 
- Use mock mode in client (`websocket.mock = true`)
- Or run the server only on Raspberry Pi

### Installation error: Python 3.13 / node-gyp incompatible

**Error** (during `npm install` in server):
```
gyp ERR! stack TypeError: Cannot assign to read only property 'cflags'
gyp info using node-gyp@7.1.2
gyp info using Python version 3.13.5
```

**Cause**: The `epoll` dependency (used by `onoff` for GPIO) has an old version of `node-gyp` incompatible with Python 3.13+.

**Solutions**:

**Option 1: Install with --ignore-scripts (Recommended for dev)**
```bash
cd server
npm install --ignore-scripts
```

This skips compilation of native dependencies (GPIO, SerialPort). Perfect for:
- ✅ Development on laptop/desktop
- ✅ CI/CD pipelines
- ✅ Systems with Python 3.13+
- ❌ Does NOT work on Raspberry Pi (needs compilation)

**Option 2: Downgrade Python (only if needed for Raspberry)**
```bash
# Install Python 3.11 (compatible with node-gyp)
sudo apt install python3.11
sudo update-alternatives --install /usr/bin/python3 python3 /usr/bin/python3.11 1
```

**Option 3: DEV_MODE (server development)**
```bash
cd server
npm install --ignore-scripts
DEV_MODE=true node server.js
```

⚠️ In DEV_MODE the server starts but is not functional (no GPIO/OBD). Use only for testing.

**Note**: Hardware dependencies (`onoff`, `serialport`, `ads1x15`) are `optionalDependencies` - they fail without blocking installation of other dependencies.

### ELM327 not found

**Error**:
```
Port /dev/ttyUSB0 not found
```

**Solution**:
1. Verify port: `ls -l /dev/ttyUSB*`
2. Check permissions: `sudo usermod -a -G dialout $USER`
3. Modify port in `OBDCommunicationService.js` if different

### Temperature sensor not found

**Warning**:
```
⚠️ DS18B20 temperature sensor not available (1-Wire not found)
```

**Solution**:
1. Enable 1-Wire: `sudo raspi-config` → Interface Options → 1-Wire
2. Reboot: `sudo reboot`
3. Verify: `ls /sys/bus/w1/devices/`
4. If not needed, disable in `gpio-mapping.js`: `temperature.enabled = false`

### Fuel sensor not responding

**Warning**:
```
⚠️ ADS1115 fuel sensor not available
```

**Solution**:
1. Enable I2C: `sudo raspi-config` → Interface Options → I2C
2. Verify connection: `sudo i2cdetect -y 1`
3. Check ADS1115 wiring
4. If not needed, disable in `gpio-mapping.js`: `fuel.enabled = false`

### Electron won't start

**Error**:
```
Error: connect ECONNREFUSED 127.0.0.1:5173
```

**Solution**:
The Vite client must be started first. Use `npm start` which manages the order automatically.

### GPIO not responding

**Issue**: Warning lights not detected

**Solution**:
1. Verify optocoupler wiring
2. Test pin: `gpio readall` (install wiringpi if needed)
3. Check pin mapping in `gpio-mapping.js`
4. Verify active high/low logic of optocouplers

---

## 📱 Using the Application

### Keyboard Controls

- **`d`**: Open debug console
- **`ESC`**: Close debug console
- **`r`**: Reload application

### Debug Console

Press `d` to open the interactive console showing:
- WebSocket logs
- Connection errors
- Real-time OBD-II data
- GPIO and sensor status

---

## 📦 Production Build

### Client Build

```bash
cd client
npm run build
```

Output in `client/dist/`

### Electron Build

To create a distributable app:

1. Install electron-builder: `npm install --save-dev electron-builder`
2. Add script in `package.json`:

```json
"scripts": {
  "build:electron": "electron-builder"
}
```

3. Run: `npm run build:electron`

---

## 📝 Main File Structure

```
cockpit/
├── client/
│   ├── src/
│   │   ├── config/
│   │   │   └── environment.ts          ← Client configuration
│   │   ├── components/                 ← React components
│   │   ├── routes/
│   │   │   └── Cockpit/               ← Main dashboard
│   │   ├── services/
│   │   │   └── WebSocketService.ts    ← Client WebSocket management
│   │   └── store/                     ← State management (Valtio)
│   └── package.json
│
├── server/
│   ├── config/
│   │   └── gpio-mapping.js            ← ⚙️ GPIO and sensor configuration
│   ├── services/
│   │   ├── OBDServer.js               ← Main server
│   │   ├── OBDCommunicationService.js ← ELM327 communication
│   │   ├── GPIOService.js             ← GPIO management for warning lights
│   │   ├── IgnitionService.js         ← Ignition management
│   │   ├── TemperatureSensorService.js← DS18B20 temperature sensor
│   │   └── FuelSensorService.js       ← ADS1115 fuel sensor
│   ├── scripts/
│   │   ├── low-power.sh               ← Power-saving script
│   │   └── wake.sh                    ← Wake script
│   ├── ecosystem.config.js            ← PM2 configuration
│   └── package.json
│
├── main.js                            ← Electron wrapper
└── package.json                       ← Main scripts
```

---

## 🔒 Security and Notes

- ⚠️ **Don't run as root**: Use normal user permissions with `dialout` and `gpio` groups
- 🔋 **Power-saving**: Ignition scripts can protect the system from battery drain
- 🧪 **Testing**: Always use mock mode for testing without hardware
- 📊 **Monitoring**: Use PM2 to monitor the server in production

---

## 📄 License

This project is released under the **GNU General Public License v3.0 or later**.

```
PandaOS
Copyright (C) 2025  Cyberpandino

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License version 3.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU General Public License for more details.
```

The full license text is available in the [LICENSE](../LICENSE) file and at https://www.gnu.org/licenses/gpl-3.0.html

---

## 👥 Contributing

Every contribution is welcome! Whether it's code, documentation, bug reports, or suggestions.

### 🚀 How to Start

1. Read the [contributing guide](.github/CONTRIBUTING.en.md)
2. Choose how to contribute:
   - 🐛 [Report a bug](.github/ISSUE_TEMPLATE/bug_report.md)
   - ✨ [Propose a feature](.github/ISSUE_TEMPLATE/feature_request.md)
   - ❓ [Ask a question](.github/ISSUE_TEMPLATE/question.md)
   - 💻 Contribute with code
   - 💡 Look for inspiration in the [Roadmap & Wishlist](.github/CONTRIBUTING.en.md#-want-to-contribute-but-have-no-ideas)

### 📝 Contribution Workflow

1. **Fork** the repository
2. **Create a branch**: `git checkout -b feature/feature-name`
3. **Make your changes** following the [code style](.github/CONTRIBUTING.en.md#-code-style)
4. **Add GPL-3.0 header** to new source files
5. **Commit**: `git commit -m 'feat: added new feature'` ([Conventional Commits](https://www.conventionalcommits.org/))
6. **Push**: `git push origin feature/feature-name`
7. **Open a Pull Request** filling in the [template](.github/PULL_REQUEST_TEMPLATE.md)

### 📋 Available Templates

- [🐛 Bug Report](.github/ISSUE_TEMPLATE/bug_report.md) - Report issues
- [✨ Feature Request](.github/ISSUE_TEMPLATE/feature_request.md) - Propose improvements
- [❓ Question](.github/ISSUE_TEMPLATE/question.md) - Ask for help
- [🔀 Pull Request](.github/PULL_REQUEST_TEMPLATE.md) - Contribute with code

### 💡 Looking for Ideas?

Don't know where to start? We have a [Roadmap & Wishlist](.github/CONTRIBUTING.en.md#-want-to-contribute-but-have-no-ideas) of features we'd like to implement:
- Rear camera and parking sensors
- Advanced 3D animations (doors, lights on model)
- Customizable dashboards and themes
- Photographic and video tutorials
- Internationalization
- And much more!

Check the [complete contributing guide](.github/CONTRIBUTING.en.md) for all details.

---

## 📞 Support

For issues or questions, open an issue on GitHub.

---

## 👨‍💻 Authors

PandaOS is developed and maintained by:

- **[Matteo Errera](https://github.com/matteoerrera)**
- **[Roberto Zaccardi](https://github.com/rzaccardi)**
- **[Ludovico Verde](https://www.instagram.com/ludovico.verdee/)**
