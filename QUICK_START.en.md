# ⚡ Quick Start - PandaOS Cluster

Quick guide to get started in 5 minutes.

> ⚠️ **WARNING**: PandaOS is an experimental hobby project. Installing it on real vehicles is at your own risk. Read the [Full Disclaimer](README.en.md#️-disclaimer) before proceeding.

---

## 🚀 Quick Setup

### 0. Prerequisites

Node.js 18+, npm 9+, Git 2.0+

```bash
node --version && npm --version && git --version
```

> 📘 Raspberry Pi: See [CONFIGURAZIONE_SERVER.md](server/CONFIGURAZIONE_SERVER.md#2-installazione-nodejs-e-npm)

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
  mock: true,  // ← Demo mode with simulated data
  // ...
}
```

#### For Raspberry Pi (Production)

```typescript
websocket: {
  mock: false,  // ← Real connection to the server
  // ...
}
```

### 3. Start

#### Local Development (Client Only)

```bash
# UI only, with simulated data
cd client
npm run dev
```

Open your browser: `http://localhost:5173`

#### Raspberry Pi (Full Stack)

```bash
# Server + Client + Electron
npm start
```

---

## 📋 Hardware Checklist (Raspberry Pi)

> 💡 **Need to buy components?** Check [HARDWARE.md](HARDWARE.md) for the complete list.

Before running in production, verify:

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

## 🎛️ Basic Configuration

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
  // OBD serial port (modify if needed)
  // In OBDCommunicationService.js:
  portPath: '/dev/ttyUSB0',
  
  // GPIO pins for LEDs (see full mapping table)
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

### "Server doesn't start"

**On Mac/Windows**: Normal! The server requires a Raspberry Pi.  
**Fix**: Use `mock: true` in the client.

### "ELM327 not found"

```bash
# Check serial port
ls -l /dev/ttyUSB*

# Give permissions
sudo usermod -a -G dialout $USER
# Logout and login
```

### "Client cannot connect to server"

1. Ensure the server is running: `npm run server`
2. Check the URL: `websocket.url` in `environment.ts`
3. Check if port 3001 is free: `lsof -i :3001`

### "GPIO not working"

```bash
# Test GPIO pin 17
gpio -g read 17

# If error:
sudo usermod -a -G gpio $USER
# Reboot
```

---

## 📚 Full Documentation

- **[README.md](README.en.md)** → Complete main documentation
- **[client/CONFIGURAZIONE.md](client/CONFIGURAZIONE.md)** → Detailed client configuration
- **[server/CONFIGURAZIONE_SERVER.md](server/CONFIGURAZIONE_SERVER.md)** → Hardware + server setup
- **[ARCHITETTURA.md](ARCHITETTURA.md)** → System architecture

---

## 🎯 Next Steps

1. **UI Development**: Edit components in `client/src/components/`
2. **Customize GPIO**: Adapt `server/config/gpio-mapping.js` to your wiring
3. **Add PIDs**: Extend `server/services/PIDParserService.js`
4. **Styling**: Modify `client/src/assets/scss/`
5. **Production**: Configure PM2 (see README.md)

---

## 📞 Help

Need Help? 

1. **Check the documentation**:
   - [README.md](README.md) - general troubleshooting
   - [server/CONFIGURAZIONE_SERVER.md](server/CONFIGURAZIONE_SERVER.md) - hardware issues
   - [client/CONFIGURAZIONE.md](client/CONFIGURAZIONE.md) - client issues

2. **Open an Issue**:
   - [🐛 Bug Report](.github/ISSUE_TEMPLATE/bug_report.md) - report a problem
   - [❓ Question](.github/ISSUE_TEMPLATE/question.md) - ask a question

3. **Contribute**:
   - [✨ Feature Request](.github/ISSUE_TEMPLATE/feature_request.md) - propose improvements
   - [CONTRIBUTING.md](.github/CONTRIBUTING.md) - contribution guide

---

**Happy coding, and try not to break anything! 🚗💨**

