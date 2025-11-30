[🇬🇧 English](DOCUMENTAZIONE.en.md) | [🇮🇹 Italiano](DOCUMENTAZIONE.md) | [🇩🇪 Deutsch](DOCUMENTAZIONE.de.md)

---

# 📚 Documentation Index - PandaOS Cluster

Complete guide to project documentation.

---

## 🎯 Where to Start?

### 👋 New to the Project?
**Start here**: [QUICK_START.en.md](QUICK_START.en.md)  
Quick guide to launch the project in 5 minutes.

### 📖 Want to Understand Everything?
**Read**: [README.en.md](README.en.md)  
Complete main documentation with setup, configuration, and troubleshooting.

### 🛒 Need to Buy Components?
**Check**: [HARDWARE.en.md](HARDWARE.en.md)  
Complete list of required hardware with specifications and costs.

### 🔧 Need to Configure Hardware?
**Go to**: [server/CONFIGURAZIONE_SERVER.en.md](server/CONFIGURAZIONE_SERVER.en.md)  
GPIO, OBD-II, sensors setup and Raspberry Pi configuration.

### 💻 Want to Customize the Client?
**Check**: [client/CONFIGURAZIONE.en.md](client/CONFIGURAZIONE.en.md)  
Complete interface configuration and operating modes.

### 🏗️ Want to Extend the System?
**Study**: [ARCHITETTURA.en.md](ARCHITETTURA.en.md)  
Technical architecture, data flows, and guides to add functionality.

### 💡 Want to Contribute but Have No Ideas?
**See**: [ROADMAP.en.md](ROADMAP.en.md)  
Future features, wishlist, and ideas to contribute to the project.

---

## 📋 Documentation Structure

### 1. [README.en.md](README.en.md) - Main Documentation
**Content**:
- 📋 Project description and features
- ⚠️ **Important disclaimer** (responsibility and safety)
- 🏗️ General architecture
- 🤔 **Tech stack choice** (why React/Electron)
- ⚙️ System requirements
- 🚀 Complete step-by-step setup
- 🎯 Project startup (local and Raspberry)
- 🔌 GPIO configuration (overview)
- 🔧 PM2 setup for production
- 🛠️ General troubleshooting
- 📦 Production build

**For Whom**: Everyone - essential starting point

---

### 2. [QUICK_START.en.md](QUICK_START.en.md) - Quick Start
**Content**:
- ⚡ Installation in 3 commands
- 🎛️ Minimal configuration
- 🚀 Quick startup
- 📋 Hardware checklist
- 🐛 Quick troubleshooting
- 🎯 Next steps

**For Whom**: Developers who want to start quickly

---

### 3. [client/CONFIGURAZIONE.en.md](client/CONFIGURAZIONE.en.md) - Client Configuration
**Content**:
- 📁 `environment.ts` file explained in detail
- 🔧 WebSocket parameters (URL, mock mode, timeout)
- 🎬 Splash screen configuration
- 🐛 Debug mode and console viewer
- 🌍 Locale, timezone, time format
- 📋 Configuration examples (development, production, testing)
- 🔄 Development workflow
- 🔍 Mock vs real mode
- 🚨 Client troubleshooting

**For Whom**: Frontend developers, those customizing the interface

---

### 4. [HARDWARE.en.md](HARDWARE.en.md) - Hardware List
**Content**:
- 🛒 Essential components (Raspberry Pi, ELM327, Display)
- 🔌 Optional sensors (DS18B20, ADS1115)
- 📺 LCD display specifications used
- 🔗 Required accessories and wiring
- 💰 Indicative cost estimate
- 📦 Recommended kits
- 🔍 Technical notes and compatibility
- ⚠️ Hardware disclaimer

**For Whom**: Those who need to buy components, hardware builders

---

### 5. [ROADMAP.en.md](ROADMAP.en.md) - Future Features and Wishlist
**Content**:
- 🚗 Hardware features (backup camera, parking sensors, 3D animations)
- 💻 Software features (trip computer, custom dashboards, mobile app)
- 📚 Documentation (photo/video tutorials, i18n, PCB design)
- 🧪 Testing & quality
- 🔧 Other Fiat vehicle compatibility
- 🎨 UI/UX improvements
- 🌐 Service integrations (Maps, Spotify, weather)
- 📊 Current development status
- 🤝 How to contribute to features

**For Whom**: Those who want to contribute but looking for ideas, those who want to know the project's future

---

### 6. [server/CONFIGURAZIONE_SERVER.en.md](server/CONFIGURAZIONE_SERVER.en.md) - Hardware and Server Setup
**Content**:
- 🛠️ Complete hardware requirements
- 🔧 Raspberry Pi configuration step-by-step
- 🔌 **OBD-II Serial Port**:
  - ELM327 hardware setup
  - Port identification (`/dev/ttyUSB0`)
  - Connection testing
  - Baudrate configuration
  - OBD-II protocol
- 🔢 **GPIO Configuration**:
  - Complete pin mapping (BCM)
  - Raspberry Pi 4B pinout diagram
  - All warning lights table (17 GPIO)
  - Optocoupler wiring
  - Electrical schematic
  - GPIO testing
- 🌡️ **DS18B20 Temperature Sensor**:
  - Technical specifications
  - Connection diagram (pull-up resistor)
  - 1-Wire setup
  - Detection verification
  - Software configuration
  - Multiple sensors
- ⛽ **ADS1115 Fuel Sensor**:
  - ADC specifications
  - I2C connection diagram
  - Resistive divider (calculation)
  - I2C setup
  - Reading test
  - Gain and sample rate configuration
  - Calibration procedure
- ⚡ **Ignition Management**:
  - Hardware setup
  - Power-saving scripts
  - Customization (auto-shutdown, notifications)
- 🚨 Detailed hardware troubleshooting
- 📚 Datasheets and useful commands

**For Whom**: Hardware engineers, those configuring Raspberry Pi, sysadmins

---

### 7. [ARCHITETTURA.en.md](ARCHITETTURA.en.md) - Technical Documentation
**Content**:
- 📊 Complete architecture diagram
- 🔄 Data flows:
  - System startup
  - OBD data reading
  - GPIO warning light detection
  - Sensor reading
- 📦 **Server Modules** (detailed description):
  - OBDServer (orchestrator)
  - OBDCommunicationService (ELM327)
  - PIDParserService (hex parsing)
  - MonitoringService (polling)
  - GPIOService (warning lights)
  - IgnitionService (ignition)
  - TemperatureSensorService (DS18B20)
  - FuelSensorService (ADS1115)
  - WebSocketService (communication)
- 🎨 **Client Modules**:
  - State Management (Valtio)
  - WebSocketService client
  - MockAnimationService
- 🔐 Security and permissions
- 🧪 Testing (local, integration, Electron)
- 📈 Performance and optimizations
- 🔄 **Extensibility Guides**:
  - Add new OBD PID
  - Add new GPIO warning light
  - Add new sensor (BMP280 example)
- 📚 Main files table

**For Whom**: Advanced developers, those who want to contribute, those who want to extend the system

---

## 🗂️ File Organization

```
cockpit/
├── README.en.md                       ← 📖 Main documentation
├── QUICK_START.en.md                  ← ⚡ Quick guide
├── HARDWARE.en.md                     ← 🛒 Complete hardware list
├── ROADMAP.en.md                      ← 🗺️ Future features and wishlist
├── DOCUMENTAZIONE.en.md               ← 📚 This file (index)
├── ARCHITETTURA.en.md                 ← 🏗️ Technical architecture
│
├── client/
│   ├── CONFIGURAZIONE.en.md           ← 💻 Client configuration
│   └── src/config/
│       └── environment.ts             ← ⚙️ Configuration file
│
└── server/
    ├── CONFIGURAZIONE_SERVER.en.md    ← 🔧 Hardware setup
    ├── config/
    │   └── gpio-mapping.js            ← 🔢 GPIO mapping
    ├── services/                      ← 📦 Backend services
    └── scripts/                       ← ⚡ Power-saving scripts
```

---

## 🔍 Find Quickly

### Configuration
| What you're looking for | Where to find it |
|------------------------|------------------|
| **Disclaimer and responsibility** | **[README.en.md](README.en.md) § Disclaimer** |
| Initial project setup | [README.en.md](README.en.md) § Setup |
| Client configuration | [client/CONFIGURAZIONE.en.md](client/CONFIGURAZIONE.en.md) |
| Server configuration | [server/CONFIGURAZIONE_SERVER.en.md](server/CONFIGURAZIONE_SERVER.en.md) |
| Mock vs real mode | [client/CONFIGURAZIONE.en.md](client/CONFIGURAZIONE.en.md) § Mock vs Real |
| GPIO pin mapping | [server/CONFIGURAZIONE_SERVER.en.md](server/CONFIGURAZIONE_SERVER.en.md) § GPIO |
| OBD serial port | [server/CONFIGURAZIONE_SERVER.en.md](server/CONFIGURAZIONE_SERVER.en.md) § OBD |

### Hardware
| Component | Documentation |
|-----------|---------------|
| **Complete component list** | **[HARDWARE.en.md](HARDWARE.en.md)** |
| Raspberry Pi setup | [server/CONFIGURAZIONE_SERVER.en.md](server/CONFIGURAZIONE_SERVER.en.md) § Raspberry |
| ELM327 OBD | [server/CONFIGURAZIONE_SERVER.en.md](server/CONFIGURAZIONE_SERVER.en.md) § Serial Port |
| GPIO optocouplers | [server/CONFIGURAZIONE_SERVER.en.md](server/CONFIGURAZIONE_SERVER.en.md) § GPIO |
| DS18B20 temperature | [server/CONFIGURAZIONE_SERVER.en.md](server/CONFIGURAZIONE_SERVER.en.md) § Temperature |
| ADS1115 fuel | [server/CONFIGURAZIONE_SERVER.en.md](server/CONFIGURAZIONE_SERVER.en.md) § Fuel |
| Ignition/Power saving | [server/CONFIGURAZIONE_SERVER.en.md](server/CONFIGURAZIONE_SERVER.en.md) § Ignition |

### Troubleshooting
| Problem | Solution |
|---------|----------|
| Server won't start | [README.en.md](README.en.md) § Troubleshooting |
| ELM327 not found | [server/CONFIGURAZIONE_SERVER.en.md](server/CONFIGURAZIONE_SERVER.en.md) § Troubleshooting |
| GPIO not working | [server/CONFIGURAZIONE_SERVER.en.md](server/CONFIGURAZIONE_SERVER.en.md) § Troubleshooting |
| Sensors not detecting | [server/CONFIGURAZIONE_SERVER.en.md](server/CONFIGURAZIONE_SERVER.en.md) § Troubleshooting |
| Client won't connect | [client/CONFIGURAZIONE.en.md](client/CONFIGURAZIONE.en.md) § Troubleshooting |

### Development
| Task | Guide |
|------|-------|
| Add OBD PID | [ARCHITETTURA.en.md](ARCHITETTURA.en.md) § Extensibility |
| Add GPIO warning light | [ARCHITETTURA.en.md](ARCHITETTURA.en.md) § Extensibility |
| Add sensor | [ARCHITETTURA.en.md](ARCHITETTURA.en.md) § Extensibility |
| Modify UI | [client/CONFIGURAZIONE.en.md](client/CONFIGURAZIONE.en.md) + [ARCHITETTURA.en.md](ARCHITETTURA.en.md) |
| Testing | [ARCHITETTURA.en.md](ARCHITETTURA.en.md) § Testing |

---

## 📞 Support

### Documentation Not Clear?
Open an [issue](https://github.com/cyberpandino/cluster/issues/new?template=question.md) to improve documentation.

### Bug or Problem?
1. Check [README.en.md](README.en.md) § Troubleshooting
2. Check module-specific documentation
3. Open a [Bug Report](https://github.com/cyberpandino/cluster/issues/new?template=bug_report.md) with:
   - Operating system
   - Hardware used
   - Complete logs
   - Steps to reproduce

### Want to Contribute?
1. Read [CONTRIBUTING.en.md](.github/CONTRIBUTING.en.md) - Complete guide
2. Look for ideas in the [Roadmap & Wishlist](.github/CONTRIBUTING.en.md#-want-to-contribute-but-have-no-ideas)
3. Choose an appropriate [issue template](.github/ISSUE_TEMPLATE/)
4. Study [ARCHITETTURA.en.md](ARCHITETTURA.en.md) for code modifications
5. Use the [PR template](.github/PULL_REQUEST_TEMPLATE.md) for contributions

### Available Templates
- [🐛 Bug Report](.github/ISSUE_TEMPLATE/bug_report.md)
- [✨ Feature Request](.github/ISSUE_TEMPLATE/feature_request.md)
- [❓ Question](.github/ISSUE_TEMPLATE/question.md)
- [🔀 Pull Request](.github/PULL_REQUEST_TEMPLATE.md)

---

## 🔄 Documentation Updates

**Version**: 0.9.0  
**Last update**: November 2025
