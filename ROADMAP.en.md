[🇬🇧 English](ROADMAP.en.md) | [🇮🇹 Italiano](ROADMAP.md) | [🇩🇪 Deutsch](ROADMAP.de.md)

---

# 🗺️ Roadmap & Wishlist - PandaOS Cluster

Features and improvements we'd like to implement in the project.

> 💡 **Want to contribute?** Pick a feature from the list and open a [Feature Request](https://github.com/cyberpandino/cluster/issues/new?template=feature_request.md)!

---

## 🚗 Hardware Features

### High Priority

#### 📹 Integrated Backup Camera
**Description**: Display rear camera feed in the cluster when reverse gear is engaged  
**Complexity**: Medium  
**Components**: USB/CSI camera, reverse gear detection via GPIO  
**Benefits**: Parking safety, replacement for damaged rearview mirror

#### 📡 Parking Sensors
**Description**: Graphical display of obstacle distance with ultrasonic radars  
**Complexity**: Medium  
**Components**: 4-8 ultrasonic sensors, Arduino/ESP32 for processing  
**Benefits**: Parking assistance like modern cars

### Medium Priority

- **Rain sensor** - Automatic windshield wiper adjustment
- **Light sensor** - Auto-brightness adjustment for display (day/night mode)
- **Tire pressure (TPMS)** - Integration of tire pressure sensors
- **CAN Bus support** - Beyond OBD-II, support for native CAN protocol
- **360° camera** - Multi-camera system for complete view

---

## 💻 Software Features

### High Priority

#### 🛣️ Trip Computer System
**Description**: Trip logs with consumption, distance, time, route  
**Complexity**: Medium  
**Technologies**: Local database (SQLite), geolocation  
**Benefits**: Detailed statistics, consumption analysis, trip history

#### 🎨 Customizable Dashboards
**Description**: Multiple layouts selectable by user (sport, eco, minimal, full)  
**Complexity**: High  
**Technologies**: React layout system, persistent storage  
**Benefits**: Personalized experience, adaptability to preferences

#### 🌓 Color Themes
**Description**: Dark mode, light mode, custom themes (retro, futuristic, etc.)  
**Complexity**: Low  
**Technologies**: CSS variables, theme provider  
**Benefits**: Ambient light adaptability, personalized aesthetics

#### 🎯 Assisted Calibration
**Description**: Step-by-step wizard to calibrate fuel/temperature sensors  
**Complexity**: Medium  
**Technologies**: UI wizard, calibration storage  
**Benefits**: Easier setup for non-technical users

#### 📱 Mobile Companion App
**Description**: Smartphone app for vehicle statistics, notifications, remote control  
**Complexity**: High  
**Technologies**: React Native / Flutter, REST/WebSocket API  
**Benefits**: Data access even away from vehicle

#### 🚪 3D Door Animations
**Description**: Represent open/closed doors on the Panda 3D model in the cluster  
**Complexity**: Low  
**Technologies**: Three.js, 3D model animations, GPIO integration  
**Benefits**: Immediate visual feedback, more immersive UI

#### 💡 Lights on 3D Model
**Description**: Show active lights (high beam, turn signals, fog lights) directly on the 3D model  
**Complexity**: Low  
**Technologies**: Three.js materials, emissive textures, existing GPIO data  
**Benefits**: Intuitive light status visualization, more engaging UI

### Medium Priority

- **Scheduled maintenance** - Alerts for service, oil changes, inspection based on km
- **Weather integration** - External temperature from API if sensor unavailable
- **Automatic night/day mode** - Based on time/GPS or light sensor
- **Gesture control** - Gestural controls (if touch display)
- **Voice control** - Voice commands for main functions (Annyang.js)
- **Multi-user profiles** - Separate statistics for different drivers
- **Data export** - CSV/JSON/Excel for external analysis
- **Push notification system** - Sound/visual alerts for anomalies/maintenance

---

## 📚 Documentation

### High Priority

#### 📸 Photographic Wiring Tutorial
**Description**: Step-by-step guide with real photos of optocoupler wiring  
**Complexity**: Low (but requires actual installation)  
**Content**: Detailed photos of each stage, wire identification, connections  
**Benefits**: Drastically reduces installation errors

#### 🎥 Video Installation Guide
**Description**: Complete video tutorial from wiring to software  
**Complexity**: Medium  
**Content**: Edited video with voice-over, subtitles, chapters  
**Benefits**: More accessible format for less technical users

#### 🌍 Internationalization (i18n)
**Description**: EN, ES, DE, FR interface translations  
**Complexity**: Medium  
**Technologies**: react-i18next (already present), JSON translation files  
**Benefits**: International accessibility, wider community

#### 📝 Centralized Translation File
**Description**: Move all hardcoded microcopy to JSON/i18n files  
**Complexity**: Low  
**Technologies**: i18next, JSON  
**Benefits**: Easy maintenance, community-driven translations

#### 🔌 Custom PCB Schematic
**Description**: Professional PCB design for optocouplers (KiCad/Eagle)  
**Complexity**: High  
**Technologies**: KiCad, Gerber export  
**Benefits**: Clean installation, no breadboard, cheap PCB ordering

### Medium Priority

- **Extended FAQ** - Frequently asked questions with detailed troubleshooting
- **Installation case studies** - Real examples with photos and logs
- **Compatibility guides** - List of compatible vehicles (Uno, Seicento, etc.)
- **Interactive wiring diagram** - Navigable electrical schematic online (SVG/HTML)

---

## 🧪 Testing & Quality

- **Unit tests** - Automated testing for backend services (Jest)
- **E2E tests** - Complete interface tests (Playwright/Cypress)
- **CI/CD Pipeline** - GitHub Actions for automatic build and deploy
- **Performance profiling** - Rendering and memory optimization
- **Hardware-in-the-loop testing** - Automated tests with simulated hardware (mock GPIO/Serial)
- **Stress testing** - Long-running stability tests
- **Code coverage** - >80% coverage for critical code

---

## 🔧 Compatibility & Extensions

### Vehicles

- **Fiat Uno** (1983-1995) - Similar OBD-I/II
- **Fiat Seicento** (1998-2010) - OBD-II compatible
- **Fiat Punto** (first series) - Same Magneti Marelli ECU
- **Lancia Y** (first series) - Similar ECUs
- **Fiat Tipo** - Related mechanics

### Protocols

- **J1850 PWM/VPW** - American protocols
- **CAN Bus (ISO 15765)** - Newer vehicles
- **LIN Bus** - Automotive accessories

### Hardware

- **Raspberry Pi Zero 2W** - Compact low-power version
- **Raspberry Pi CM4** - For custom integrations
- **Orange Pi / Banana Pi** - Cheap SBC alternatives
- **Android tablets** - Port app to native Android

### Integrations

- **Android Auto / CarPlay** - Integration with modern systems
- **Aftermarket Head Units** - Compatibility with Chinese Android stereos
- **Automatic backup** - Cloud/USB restore/backup settings system
- **OTA Updates** - Over-the-air software updates

---

## 🎨 UI/UX

- **Transition animations** - Smooth GSAP transitions between states/screens
- **Sport/eco modes** - Different visualizations for driving style (aggressive red / efficient green)
- **Customizable widgets** - Drag & drop components in the cluster
- **Community skins/themes** - Marketplace to share custom themes
- **Minimal mode** - Essential UI for less driving distraction
- **Screensaver** - Animations when vehicle parked
- **Startup animation** - Customizable boot animation
- **Easter eggs** - Special animations for events (Christmas, car birthday, etc.)

---

## 🌐 Service Integration

- **Google Maps / OpenStreetMap** - Navigation integrated in cluster
- **Spotify / Apple Music** - Music control from cluster
- **Telegram Bot** - Notifications and remote control
- **IFTTT / Home Assistant** - Smart home automations
- **Fuel prices API** - Real-time fuel prices
- **Traffic data** - Traffic information
- **Weather API** - Integrated weather forecast

---

## 🔐 Security & Privacy

- **User authentication** - Login for statistics/configuration access
- **Data encryption** - Protection of sensitive saved data
- **VPN client** - Secure connection for cloud data
- **Privacy mode** - Disable GPS tracking/statistics

---

## 🎯 Performance & Optimization

- **Improved lazy loading** - On-demand component loading
- **Service Workers** - PWA for cache and offline
- **WebGL optimization** - Optimized 3D rendering
- **Memory management** - Reduced memory footprint
- **Boot time** - Startup time reduction <10 seconds

---

## 📊 Current Status

### Version 0.9.0 (Current)

**Implemented**:
- ✅ OBD-II data reading via ELM327
- ✅ GPIO warning light detection with optocouplers
- ✅ DS18B20 temperature sensor
- ✅ ADS1115 fuel sensor
- ✅ 3D dashboard with interactive Panda model
- ✅ Mock mode for development
- ✅ Integrated debug console
- ✅ Ignition management (power-saving)
- ✅ Real-time WebSocket
- ✅ Complete documentation

**In Development**:
- 🔄 No features currently in active development

**Planned**:
- 📋 See roadmap above

### Version 1.0.0 (Goal)

**Target Features**:
- Working backup camera
- Parking sensors
- Complete 3D animations
- Customizable dashboards
- Internationalization
- Complete testing

---

## 🤝 How to Contribute to These Features

### 1. Choose a Feature
Browse the list and find something that:
- You're passionate about
- You have skills to implement
- Solves your problem/need

### 2. Open an Issue
Use [Feature Request](https://github.com/cyberpandino/cluster/issues/new?template=feature_request.md) to discuss:
- Implementation approach
- Technologies to use
- Compatibility with existing
- Estimated timeline

### 3. Develop
- Fork the repository
- Create dedicated branch
- Implement following [code style](.github/CONTRIBUTING.en.md#-code-style)
- Test thoroughly

### 4. Pull Request
- Open PR with [template](.github/PULL_REQUEST_TEMPLATE.md)
- Describe implementation
- Attach screenshots/videos
- Wait for review

---

## 💬 Discussion

Want to discuss features, propose alternatives, or share ideas?
- Open a [Discussion](https://github.com/cyberpandino/cluster/discussions) (if enabled)
- Or a [Question Issue](https://github.com/cyberpandino/cluster/issues/new?template=question.md)

---

## 🛠️ Native Stack Rewrite?

### The Great Refactoring

A "special" feature worth mentioning separately:

#### ⚙️ Port to C++/Qt Native

**Description**: Rewrite PandaOS with professional automotive tech stack  
**Complexity**: Very High (practically a complete rewrite)  
**Technologies**:
- **C++17/20** for backend and business logic
- **Qt 6 / QML** for graphical interface (automotive standard)
- **Qt 3D** for interactive Panda model
- **Yocto/Buildroot** for optimized embedded Linux
- **systemd** for service management
- **D-Bus** for IPC between processes

**Benefits**:
- 🚀 **Boot time** <3 seconds (vs ~15s current)
- 💾 **Memory** ~50MB total (vs ~500MB current)
- ⚡ **Performance** guaranteed 60fps rendering even on Pi Zero
- 🔋 **Power consumption** reduced by 60-70%
- 🏭 **Professional approach** worthy of production

**Why didn't we do it?**

Because after 2 hours of struggling with CMake we chose the simpler way: npm. And honestly, we don't regret it. 😅

The web stack allowed us to:
- Have something working in weeks, not months/years
- Avoid complexity of ARM cross-compilation
- Use powerful libraries like Three.js without writing OpenGL shaders by hand
- Focus on UX instead of debugging segfaults

**But if you want to accept the challenge...**

We'd be **thrilled** if someone wanted to do a native port! We could have:
- **PandaOS-Web** (current) - Fast and accessible approach
- **PandaOS-Native** (future?) - Professional and performant approach

If you're interested, open a discussion. We admire those with patience to master Qt and CMake. 🚀

---

## 📅 Timeline

There are no fixed timelines. The project is hobbyist and proceeds when there's time and passion.  
**Community contributions are essential to accelerate development!**

---

## 🤷 But Seriously, Why This Absurd Stack?

**Short answer**: Because it's fun and we don't want to go crazy.

**Long answer**: [README.en.md - But React + Electron on Automotive?!](../README.en.md#-but-react--electron-on-automotive-are-you-crazy)

**Honest answer**: If we had to do it "properly" we'd use C++/Qt. But it would require much more time and specific skills. For now, our "creative" approach works perfectly for the purpose. 😅

---

**Last revision**: November 2025  
**Next update**: When we have new brilliant ideas 💡
