[🇬🇧 English](QUICK_START.en.md) | [🇮🇹 Italiano](QUICK_START.md) | [🇩🇪 Deutsch](QUICK_START.de.md)

---

# ⚡ Quick Start - PandaOS Cluster

Quick guide to get started in 5 minutes.

> ⚠️ **WARNING**: PandaOS is an experimental hobby project. Installation on real vehicles is at your own risk. Read the [full Disclaimer](README.en.md#️-disclaimer) before proceeding.

---

## 🚀 Quick Start

### 0. Prerequisites

Node.js 18+, npm 9+, Git 2.0+

```bash
node --version && npm --version && git --version
```

> 📘 Raspberry Pi: see [CONFIGURAZIONE_SERVER.en.md](server/CONFIGURAZIONE_SERVER.en.md#2-installing-nodejs-and-npm)

---

### 1. Installation

```bash
# Clone repository
git clone git@github.com:cyberpandino/cluster.git
cd cluster

# Install all dependencies
npm run install:all
```

### 2. Basic Configuration

#### For Local Development (Mac/Windows/Linux)

Edit `client/src/config/environment.ts`:

```typescript
websocket: {
  mock: true,  // ← Demo mode
  // ...
}
```

#### For Raspberry Pi (Production)

```typescript
websocket: {
  mock: false,  // ← Real connection to server
  // ...
}
```

### 3. Launch

#### Local Development (Client Only)

```bash
# Interface only with simulated data
cd client
npm run dev
```

Open browser: `http://localhost:5173`

#### Raspberry Pi (Complete Stack)

```bash
# Server + Client + Electron
npm start
```

---

## 📋 Hardware Checklist (Raspberry Pi)

> 💡 **Need to purchase components?** Check [HARDWARE.en.md](HARDWARE.en.md) for the complete list of everything needed.

Before starting in production, verify:

- [ ] ELM327 connected to USB port (`/dev/ttyUSB0`)
- [ ] Optocouplers wired to GPIO pins
- [ ] User permissions configured:
  ```bash
  sudo usermod -a -G dialout,gpio,i2c $USER
  ```
- [ ] Interfaces enabled via `raspi-config`:
  - [ ] I2C (if using ADS1115)
  - [ ] 1-Wire (if using DS18B20)
  - [ ] Serial Port
- [ ] Reboot after configuration: `sudo reboot`

---

## 🎛️ Minimal Configuration

### Client (`client/src/config/environment.ts`)

```typescript
export const environment = {
  websocket: {
    url: 'http://127.0.0.1:3001',
    mock: true,  // true=demo | false=real
  },
  debug: {
    enabled: true,
  },
};
```

### Server (`server/config/gpio-mapping.js`)

```javascript
module.exports = {
  // OBD serial port (modify if necessary)
  // In OBDCommunicationService.js:
  portPath: '/dev/ttyUSB0',
  
  // GPIO pins for warning lights (see full table)
  mapping: {
    turnSignals: { pin: 17 },
    battery: { pin: 27 },
    highBeam: { pin: 5 },
    // ... others
  },
  
  // Optional sensors
  temperature: { enabled: true },
  fuel: { enabled: true },
  ignition: { enabled: true },
};
```

---

## 🔑 Keyboard Controls

Once the application is running:

- **`d`** → Open debug console
- **`ESC`** → Close debug console
- **`r`** → Reload application

---

## 🐛 Quick Troubleshooting

### "Server won't start"

**On Mac/Windows**: Normal! Server requires Raspberry Pi.  
**Solution**: Use `mock: true` in client.

### "ELM327 not found"

```bash
# Check port
ls -l /dev/ttyUSB*

# Grant permissions
sudo usermod -a -G dialout $USER
# Logout and login
```

### "Client won't connect"

1. Verify server started: `npm run server`
2. Check URL: `websocket.url` in `environment.ts`
3. Verify port 3001 is free: `lsof -i :3001`

### "GPIO not working"

```bash
# Test GPIO pin 17
gpio -g read 17

# If error:
sudo usermod -a -G gpio $USER
# Reboot
```

---

## 📚 Complete Documentation

- **[README.en.md](README.en.md)** → Complete main documentation
- **[client/CONFIGURAZIONE.en.md](client/CONFIGURAZIONE.en.md)** → Detailed client configuration
- **[server/CONFIGURAZIONE_SERVER.en.md](server/CONFIGURAZIONE_SERVER.en.md)** → Hardware and server setup
- **[ARCHITETTURA.en.md](ARCHITETTURA.en.md)** → System architecture

---

## 🎯 Next Steps

1. **UI Development**: Modify components in `client/src/components/`
2. **Customize GPIO**: Adapt `server/config/gpio-mapping.js` to your wiring
3. **Add PID**: Extend `server/services/PIDParserService.js`
4. **Styling**: Modify `client/src/assets/scss/`
5. **Production**: PM2 setup (see README.en.md)

---

## 📞 Help

Having problems? 

1. **Check the documentation**:
   - [README.en.md](README.en.md) - General troubleshooting
   - [server/CONFIGURAZIONE_SERVER.en.md](server/CONFIGURAZIONE_SERVER.en.md) - Hardware issues
   - [client/CONFIGURAZIONE.en.md](client/CONFIGURAZIONE.en.md) - Client issues

2. **Open an issue**:
   - [🐛 Bug Report](.github/ISSUE_TEMPLATE/bug_report.md) - Report a problem
   - [❓ Question](.github/ISSUE_TEMPLATE/question.md) - Ask a question

3. **Contribute**:
   - [✨ Feature Request](.github/ISSUE_TEMPLATE/feature_request.md) - Propose improvements
   - [CONTRIBUTING.en.md](.github/CONTRIBUTING.en.md) - Contribution guide

---

**Happy coding and don't break anything! 🚗💨**
