[🇬🇧 English](README.en.md) | [🇮🇹 Italiano](README.md) | [🇩🇪 Deutsch](README.de.md)

---

# 🚗 Cyberpandino Cluster - PandaOS

[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)
[![Version](https://img.shields.io/badge/version-0.9.0-green.svg)](https://github.com/cyberpandino/cluster/releases)
[![Node](https://img.shields.io/badge/node-%3E%3D18.0.0-brightgreen.svg)](https://nodejs.org/)
[![Platform](https://img.shields.io/badge/platform-Raspberry%20Pi%204B%2F5-red.svg)](https://www.raspberrypi.com/)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](https://github.com/cyberpandino/cluster/blob/main/.github/CONTRIBUTING.md)

Digitales Armaturenbrett für Fiat Panda 141 basierend auf Raspberry Pi 4B.

## 📋 Beschreibung

Vollständiges digitales Armaturenbrett-System, das die originale analoge Instrumentierung des Fiat Panda 141 ersetzt. Das System kommuniziert mit dem Steuergerät über das OBD-II-Protokoll (ELM327) und liest die Kontrollleuchten über Optokoppler, die an die GPIO-Pins des Raspberry Pi angeschlossen sind.

### Hauptmerkmale (v0.9.0)

- ✅ **OBD-II-Datenauslesen**: Geschwindigkeit, Motordrehzahl, Temperatur, Öldruck, etc.
- ✅ **Fahrzeugwarnleuchten-Erkennung**: Fernlicht, Abblendlicht, Blinker, Ölstand, etc.
- ✅ **Externe Sensoren**:
  - Außentemperatur (DS18B20)
  - Kraftstoffstand (ADS1115)
- ✅ **Zündungsverwaltung**: Automatisches Energiesparsystem
- ✅ **Moderne Oberfläche**: 3D-Dashboard mit interaktivem Panda-Modell
- ✅ **Demo-Modus**: Für Entwicklung ohne Hardware

---

## 📸 Vorschau

### Haupt-Dashboard

Das digitale Cluster ersetzt das ursprüngliche analoge Armaturenbrett vollständig durch eine moderne und anpassbare Benutzeroberfläche.

<div align="center">
  <img src="docs/images/dashboard-main.png" alt="Haupt-Dashboard" width="800"/>
  <p><em>Haupt-Dashboard mit interaktivem 3D-Modell</em></p>
</div>

---

### 🗺️ Zukünftige Features

Entdecken Sie unsere Planung: [Roadmap & Wunschliste](ROADMAP.de.md)

Einige Ideen auf der Liste:
- 📹 Rückfahrkamera und Parksensoren
- 🚪 Erweiterte 3D-Animationen (Türen, Lichter)
- 🎨 Anpassbare Dashboards und Themes
- 🌍 Internationalisierung
- 📱 Begleit-App für Mobilgeräte
- Und vieles mehr!

Möchten Sie beitragen? Jede Hilfe ist willkommen! Siehe die [Anleitung zum Beitragen](.github/CONTRIBUTING.de.md).

---

## 📚 Dokumentationsverzeichnis

### 🚀 Hier starten
- **[Schnellstart](QUICK_START.de.md)** - Kurzanleitung für den sofortigen Einstieg
- **[Hardware](HARDWARE.de.md)** - Vollständige Komponentenliste und Montageplan

### 📖 Technische Dokumentation
- **[Architektur](ARCHITETTURA.de.md)** - Detaillierte Systemarchitektur
- **[Allgemeine Dokumentation](DOCUMENTAZIONE.de.md)** - Vollständiger Projektüberblick
- **[Client-Konfiguration](client/CONFIGURAZIONE.de.md)** - Frontend-Setup und Konfiguration
- **[Server-Konfiguration](server/CONFIGURAZIONE_SERVER.de.md)** - Backend-Setup und Konfiguration
- **[Environment-Konfiguration](client/src/config/README.de.md)** - Umgebungsvariablen und Parameter

### 🤝 Beitragen
- **[Wie man beiträgt](.github/CONTRIBUTING.de.md)** - Vollständiger Leitfaden zum Beitragen zum Projekt

### 📋 Sonstiges
- **[Roadmap](ROADMAP.de.md)** - Entwicklungsplan und Wunschliste
- **[Autoren](AUTHORS.de.md)** - Wer zum Projekt beigetragen hat
- **[Lizenz](LICENSE)** - GNU General Public License v3.0

---

## ⚠️ Haftungsausschluss

PandaOS ist ein Hobby- und Versuchsprojekt, entstanden aus technischer Neugier und dem Geist des digitalen Abenteuers. Es ist kein zertifiziertes Produkt, nicht für die Produktion gedacht und erhebt keinerlei Anspruch darauf, industrielle, automotive oder galaktische Standards zu erfüllen.

Alle in diesem Repository enthaltenen Materialien, einschließlich Code, Anleitungen, Schemata und mehr oder weniger vernünftiger Ideen, werden "WIE BESEHEN" bereitgestellt, ohne Garantien für Funktionalität, Zuverlässigkeit oder Kompatibilität mit der elektrischen Anlage Ihres heroischen Kleinwagens.

Die Autoren und Mitwirkenden übernehmen keinerlei Verantwortung im Falle von:

* elektrischen oder elektronischen Ausfällen
* anomalem Fahrzeugverhalten
* unvorhergesehenen Kurzschlüssen
* Schäden an Personen, Sachen, Tieren und dergleichen
* jeglichen Nebenwirkungen, die sich aus der Verwendung der Software oder dem Befolgen der in dieser Dokumentation enthaltenen Anweisungen ergeben

Die Verwendung von PandaOS in Fahrzeugen im Straßenverkehr oder in jedem Kontext, in dem Konformitäts-, Homologations- oder gesunder Menschenverstand erforderlich sein könnte, wird dringend abgeraten. Jede Installation oder Experimentation erfolgt auf ausschließliches Risiko des Benutzers, der alle Verantwortung für die technischen und praktischen Folgen seiner Entscheidungen übernimmt.

---

## 🏗️ Architektur

Das Projekt besteht aus drei Hauptmodulen:

```
cluster/
├── client/          → Grafische Oberfläche (React + Vite + Electron)
├── server/          → Backend OBD-II und GPIO-Kommunikation (Node.js)
└── main.js          → Electron-Wrapper für Desktop-App
```

### Verwendete Technologien

- **Frontend**: React 18, TypeScript, Three.js, Socket.IO Client
- **Backend**: Node.js, Socket.IO Server, SerialPort, GPIO (onoff)
- **Desktop**: Electron 36
- **Hardware**: Raspberry Pi 4B, ELM327, DS18B20, ADS1115

### 🤔 Aber React + Electron in einem Automobil?! Seid ihr verrückt?

Ja, wir wissen es. Jeder Embedded-Ingenieur, der dieses Projekt sieht, bekommt wahrscheinlich eine Panikattacke.

**Wie man es richtig machen würde:**
- **C/C++** - Weil JavaScript in einem Auto ist wie quadratische Räder anzubringen
- **Qt/QML** - Der Industriestandard (Tesla, Audi, BMW verwenden es)
- **Yocto/Buildroot** - Seriöses eingebettetes Linux, nicht Raspberry Pi OS mit all dem Ballast
- **Direct framebuffer** - Nicht Electron, das einen ganzen Browser ausführt, um 4 Zahlen anzuzeigen

**Und warum dann React/Electron/Node.js?**

Weil es ein **Hobbyprojekt** ist und wir **Spaß haben** wollen, nicht verrückt werden.

**Vorteile unseres fragwürdigen Ansatzes**:
- ⚡ **Schnell zu entwickeln** - Haben Sie Three.js gesehen? Erstellen Sie ein 3D-Modell in 5 Minuten. Versuchen Sie das mit nativem OpenGL.
- 🎨 **Bibliotheken überall** - npm hat alles. C++ hat... ähm... boost?
- 🧑‍💻 **Zugänglich** - Können Sie React? Willkommen. Können Sie CMake? Unser Beileid.
- 🐛 **Debug** - F12 und Sie sehen alles. GDB hingegen ist... eine Erfahrung.
- 🚀 **Spaß** - Mehr Zeit zum Basteln, weniger zum Kämpfen mit Toolchains
- 💡 **Konzeptbeweis** - Funktioniert es? Gut! Dann sehen wir weiter.

**Nachteile (die wir bewusst akzeptieren)**:
- 💾 **Frisst RAM** wie Pasta (~500MB vs ~50MB)
- 🐌 **Langsamer Start** (~30s vs ~3s) - aber mit Standby-Modus, der immer an ist, wird es sofort
- 🔋 **Verbraucht** mehr als es sollte (aber Standby verbraucht nur 0,4W, vernachlässigbar)
- 📊 **JavaScript** - Ja, JavaScript. In einem Auto. Deal with it.

**Der Punkt ist**: Wir sprechen über einen **Panda von 1990**. Es ist kein F-35. Es muss nicht zum Mond fliegen.  
Es muss Ihnen die Motordrehzahl auf coole Weise zeigen, während Sie Pink Floyd hören. Und das macht es perfekt. 🚗💨

> 💡 **Möchten Sie es in C++/Qt "richtig" neu machen?** Fantastisch! Wir wären neugierig, einen nativen Port zu sehen und würden gerne helfen.

---

## ⚙️ Systemanforderungen

### Software-Voraussetzungen

| Software | Mindestversion | Empfohlen |
|----------|----------------|-----------|
| **Node.js** | 18.0.0 | 20.x LTS |
| **npm** | 9.0.0 | 10.x |
| **Git** | 2.0+ | Neueste |

```bash
# Schnellprüfung
node --version  # >= v18.0.0
npm --version   # >= 9.0.0
git --version   # >= 2.0.0
```

⚠️ **Raspberry Pi**: Verwenden Sie nicht `apt install nodejs` (veraltete Version). Siehe [CONFIGURAZIONE_SERVER.de.md](server/CONFIGURAZIONE_SERVER.de.md#2-installation-von-nodejs-und-npm) für NodeSource/nvm.

---

### Für Raspberry Pi (Produktion)

- **Hardware**:
  - Raspberry Pi 4B (4GB oder höher empfohlen) oder Raspberry Pi 5
  - ELM327 USB-Adapter (Serielle Schnittstelle `/dev/ttyUSB0`)
  - Optokoppler zur Warnsignalerfassung (PC817 oder ähnlich)
  - LCD-Display ultra-breit (1920×480 empfohlen)
  - DS18B20 Temperatursensor (optional)
  - ADS1115 ADC-Wandler (optional, für Kraftstoffsensor)

📋 **Vollständige Hardware-Liste**: Siehe [HARDWARE.de.md](HARDWARE.de.md) für Details zu allen erforderlichen Komponenten

- **Betriebssystem**:
  - Raspberry Pi OS Lite (64-bit) - Debian-basiert empfohlen
  - Bootzeit: ~30s (optimierbar auf ~20s, oder sofort mit Standby-Modus)
  - ARM/ARM64-Architektur
  
  > 📘 **OS-Wahl und Bootzeit**: Siehe [CONFIGURAZIONE_SERVER.de.md](server/CONFIGURAZIONE_SERVER.de.md#1-betriebssystem-installation) für Details zur Auswahl der richtigen Distribution, Optimierung der Bootzeit und Konfiguration des **Standby-Modus, der immer an ist** (vernachlässigbarer Verbrauch, sofortiger Start)

- **Software**: Siehe [CONFIGURAZIONE_SERVER.de.md](server/CONFIGURAZIONE_SERVER.de.md#2-installation-von-nodejs-und-npm) für Installationsanleitungen auf Raspberry Pi

### Für lokale Entwicklung (Mac/Windows/Linux)

- Node.js 18+ (20 LTS empfohlen)
- npm 9+ (10.x empfohlen)
- Git 2.0+

> 💡 **Schnelleinrichtung**: Siehe Installationsanweisungen im Abschnitt [Software-Voraussetzungen](#software-voraussetzungen) oben

⚠️ **HINWEIS**: Wenn Sie das Projekt auf Nicht-Raspberry-Pi-Systemen ausführen, schlägt der Server beim Start aufgrund fehlender hardwarespezifischer Abhängigkeiten (GPIO, Sensoren, OBD-Serielle Schnittstelle) fehl. Es ist möglich, den **Mock-Modus** im Client für Entwicklung ohne Server zu verwenden.

---

## 🚀 Projekt-Setup

### 1. Repository klonen

```bash
git clone https://github.com/cyberpandino/cluster
cd cluster
```

### 2. Abhängigkeiten installieren

Das Projekt bietet ein Installationsskript, das alle Abhängigkeiten konfiguriert:

```bash
npm run install:all
```

Dieser Befehl installiert die Abhängigkeiten für:
- Root (Electron + concurrently)
- Client (React + Frontend-Abhängigkeiten)
- Server (Node.js + Hardware-Abhängigkeiten)

### 3. Konfiguration

#### a) Client-Konfiguration

Bearbeiten Sie die Client-Konfigurationsdatei:

**Datei**: `client/src/config/environment.ts`

```typescript
export const environment: EnvironmentConfig = {
  websocket: {
    url: 'http://127.0.0.1:3001',  // WebSocket-Server-URL
    mock: true,                      // true = Demo-Modus | false = Echte Verbindung
    reconnectionAttempts: 3,
    reconnectionDelay: 1000,
    timeout: 5000,
  },
  debug: {
    enabled: true,                   // Debug-Modus aktivieren
    showConsoleViewer: true,         // Console-Viewer anzeigen (Taste 'd')
  },
  app: {
    name: "PandaOS Cluster",
    version: "0.9.0",
    locale: "de",
    timezone: "Europe/Berlin",
    timeFormat: "24h",
  },
};
```

**Schlüssel-Parameter**:
- `websocket.url`: WebSocket-Server-Adresse (Standard: `http://127.0.0.1:3001`)
- `websocket.mock`: 
  - `true` = Demo-Modus mit simulierten Animationen (für lokale Entwicklung)
  - `false` = Echte Verbindung zum Server (für Produktion auf Raspberry Pi)
- `debug.enabled`: Debug-Funktionalität aktivieren
- `debug.showConsoleViewer`: Console-Viewer anzeigen (mit Taste `d` aktivierbar)

#### b) Server-Konfiguration

Bearbeiten Sie die GPIO- und Sensor-Konfigurationsdatei:

**Datei**: `server/config/gpio-mapping.js`

Siehe Abschnitt [GPIO-Konfiguration](#-gpio-und-sensoren-konfiguration) für vollständige Details.

---

## 🎯 Projekt starten

### Vollständiger Modus (Raspberry Pi)

Client, Server und Electron gleichzeitig starten:

```bash
npm start
```

Dieser Befehl führt aus:
1. OBD-II-Server auf Port 3001
2. React/Vite-Client auf Port 5173
3. Electron Desktop-App

### Entwicklungsmodus (Lokal ohne Raspberry)

#### Option 1: Nur Client (Mock-Modus)

1. Stellen Sie sicher, dass `websocket.mock = true` in `client/src/config/environment.ts`
2. Nur den Client starten:

```bash
npm run client
```

Die Anwendung ist auf `http://localhost:5173` mit simulierten Daten verfügbar.

#### Option 2: Client + Electron

```bash
npm run client    # In einem Terminal
npm run electron  # In einem anderen Terminal
```

### Einzelne Befehle

```bash
# Nur Server (erfordert Raspberry Pi)
npm run server

# Nur Client
npm run client

# Nur Electron (wartet auf Client auf Port 5173)
npm run electron
```

---

## 🔌 GPIO und Sensoren-Konfiguration

### GPIO-Mapping für Fahrzeugwarnleuchten

Die Datei `server/config/gpio-mapping.js` enthält das vollständige GPIO-Pin-Mapping.

> 📘 **Fahrzeug-Schaltplan**: Um die richtigen Kabel für die Warnleuchten im Original-Dashboard des Panda zu identifizieren, konsultieren Sie den [Offiziellen Elektrischen Schaltplan Fiat Panda 141](http://www.bunkeringegnere.altervista.org/esplosi/FIAT%20PANDA/panda%20141/1100%20mpi/55%20IMPIANTO%20ELETTRICO%20-%20SCHEMI%20-%20GAMMA%202000.pdf) mit allen Farbcodes und Verbindungen.

#### Verwendete Pins

| Warnleuchte/Funktion | GPIO-Pin (BCM) | Beschreibung |
|----------------------|----------------|--------------|
| Blinker | 17 | Fahrtrichtungsanzeiger |
| Lichtmaschine | 27 | Batterieladung |
| Öldruck | 22 | Motoröldruck |
| Bremssystem | 23 | Bremsen |
| Einspritzdüsen | 24 | Einspritzsystem |
| Zündung an (KEY) | 25 | Schlüssel eingesteckt |
| Fernlicht | 5 | Fernlicht-Scheinwerfer |
| Abblendlicht | 6 | Abblendlicht-Scheinwerfer |
| Warnblinker | 12 | Warnlicht |
| Nebelschlussleuchte | 13 | Hintere Nebelschlussleuchte |
| Kühlmitteltemperatur | 16 | Kühlflüssigkeit |
| Heckscheibenheizung | 19 | Heckscheibenentfroster |
| Kraftstoffreserve | 20 | Niedriger Kraftstoffstand |
| Zündung (Quadro) | 21 | Zündung ein/aus-Erkennung |

#### Optokoppler-Konfiguration

```javascript
config: {
  mode: 'BCM',              // Broadcom GPIO-Nummerierung
  pullMode: 'PUD_DOWN',     // Interner Pull-Down-Widerstand
  debounceTime: 50,         // Anti-Prellfilter (ms)
  pollingInterval: 100,     // GPIO-Lesefrequenz (ms)
}
```

**Betriebslogik**:
- `HIGH (1)` = Warnleuchte an
- `LOW (0)` = Warnleuchte aus

### Zündungsverwaltung (Ignition)

```javascript
ignition: {
  enabled: true,
  pin: 21,                  // Dedizierter GPIO-Pin
  activeOn: 0,              // 0 = active low | 1 = active high
  scripts: {
    lowPower: './scripts/low-power.sh',   // Ausgeführt, wenn Zündung ausgeschaltet wird
    wake: './scripts/wake.sh',             // Ausgeführt, wenn Zündung eingeschaltet wird
  },
}
```

Die Energiespar-Skripte können angepasst werden für:
- Display ausschalten
- Helligkeit reduzieren
- Nicht-essentielle Dienste deaktivieren
- Kontrolliertes Herunterfahren starten

### Außentemperatursensor (DS18B20)

```javascript
temperature: {
  enabled: true,
  sensorId: null,           // null = automatische Erkennung des ersten Sensors
  basePath: '/sys/bus/w1/devices',
  readInterval: 5000,       // Leseintervall (ms)
  pin: 4,                   // GPIO 4 (Standard für 1-Wire)
}
```

**Hardware-Setup**:
1. DS18B20 an GPIO 4 anschließen
2. 1-Wire aktivieren: `sudo raspi-config` → Interface Options → 1-Wire
3. Sensor-Vorhandensein überprüfen: `ls /sys/bus/w1/devices/`

### Kraftstoffsensor (ADS1115 - ADC I2C)

```javascript
fuel: {
  enabled: true,
  chip: 0,                  // 0 = ADS1115 | 1 = ADS1015
  channel: 0,               // Kanal A0 (0-3 verfügbar)
  gain: 4096,               // ±4.096V Vollausschlag
  sampleRate: 250,          // Abtastrate (SPS)
  readInterval: 500,        // Leseintervall (ms)
  
  // Spannungsteiler-Konfiguration
  voltageDivider: {
    r1: 100000,             // 100kΩ
    r2: 33000,              // 33kΩ
  },
  
  // Kalibrierung Spannung → Prozent
  calibration: {
    voltageEmpty: 0.5,      // Spannung bei leerem Tank (V)
    voltageFull: 4.0,       // Spannung bei vollem Tank (V)
  },
  
  pins: {
    sda: 2,                 // GPIO 2 (SDA I2C)
    scl: 3,                 // GPIO 3 (SCL I2C)
  },
}
```

**Hardware-Setup**:
1. ADS1115 anschließen:
   - VDD → 3.3V
   - GND → GND
   - SDA → GPIO 2
   - SCL → GPIO 3
   - A0 → Kraftstoffsensor (über Spannungsteiler)
2. I2C aktivieren: `sudo raspi-config` → Interface Options → I2C
3. Vorhandensein überprüfen: `sudo i2cdetect -y 1`

### OBD-II Serielle Schnittstelle

**Datei**: `server/services/OBDCommunicationService.js`

```javascript
constructor() {
  this.portPath = '/dev/ttyUSB0';  // ELM327-Port
  this.port = null;
  this.baudRate = 38400;            // Kommunikationsgeschwindigkeit
}
```

**Hardware-Setup**:
1. ELM327-Adapter über USB anschließen
2. Port überprüfen: `ls -l /dev/ttyUSB*`
3. Berechtigungen erteilen: `sudo usermod -a -G dialout $USER`
4. Neustart oder erneute Anmeldung zum Anwenden der Berechtigungen

**Alternative Serielle Schnittstellen-Konfiguration**:

Wenn der OBD-Adapter an einem anderen Port ist (z.B. `/dev/ttyUSB1`, `/dev/ttyACM0`), ändern Sie:

```javascript
// In server/services/OBDCommunicationService.js (Zeile 7)
this.portPath = '/dev/ttyUSB1';  // Hier ändern
```

---

## 🔧 PM2-Konfiguration (Automatischer Start)

Um den Server als Systemdienst auf Raspberry Pi auszuführen:

### 1. PM2 installieren

```bash
sudo npm install -g pm2
```

### 2. Ecosystem konfigurieren

`server/ecosystem.config.js` bearbeiten:

```javascript
module.exports = {
  apps: [{
    name: 'obd-server',
    script: './server.js',
    cwd: '/home/pi/cockpit/server',  // ⚠️ DIESEN PFAD ÄNDERN
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

### 3. Mit PM2 starten

```bash
cd server
mkdir -p logs
pm2 start ecosystem.config.js
pm2 save
pm2 startup
```

### 4. Nützliche PM2-Befehle

```bash
pm2 status              # Dienststatus
pm2 logs obd-server     # Protokolle anzeigen
pm2 restart obd-server  # Dienst neu starten
pm2 stop obd-server     # Dienst stoppen
pm2 monit               # Echtzeitüberwachung
```

---

## 🛠️ Fehlerbehebung

### Server startet nicht auf Nicht-Raspberry-System

**Fehler**:
```
❌ FEHLER: Wesentliche Raspberry Pi-Abhängigkeiten nicht verfügbar
Plattform nicht unterstützt: darwin arm64 - Linux ARM erforderlich
```

**Lösung**: 
- Verwenden Sie den Mock-Modus im Client (`websocket.mock = true`)
- Oder führen Sie den Server nur auf Raspberry Pi aus

### Installationsfehler: Python 3.13 / node-gyp inkompatibel

**Fehler** (während `npm install` im Server):
```
gyp ERR! stack TypeError: Cannot assign to read only property 'cflags'
gyp info using node-gyp@7.1.2
gyp info using Python version 3.13.5
```

**Ursache**: Die Abhängigkeit `epoll` (von `onoff` für GPIO verwendet) hat eine alte `node-gyp`-Version, die mit Python 3.13+ inkompatibel ist.

**Lösungen**:

**Option 1: Installation mit --ignore-scripts (Empfohlen für Entwicklung)**
```bash
cd server
npm install --ignore-scripts
```

Dies überspringt die Kompilierung nativer Abhängigkeiten (GPIO, SerialPort). Perfekt für:
- ✅ Entwicklung auf Laptop/Desktop
- ✅ CI/CD-Pipelines
- ✅ Systeme mit Python 3.13+
- ❌ Funktioniert NICHT auf Raspberry Pi (Kompilierung erforderlich)

**Option 2: Python-Downgrade (nur wenn nötig für Raspberry)**
```bash
# Python 3.11 installieren (kompatibel mit node-gyp)
sudo apt install python3.11
sudo update-alternatives --install /usr/bin/python3 python3 /usr/bin/python3.11 1
```

**Option 3: DEV_MODE (Server-Entwicklung)**
```bash
cd server
npm install --ignore-scripts
DEV_MODE=true node server.js
```

⚠️ Im DEV_MODE startet der Server, ist aber nicht funktional (kein GPIO/OBD). Nur für Tests verwenden.

**Hinweis**: Hardware-Abhängigkeiten (`onoff`, `serialport`, `ads1x15`) sind `optionalDependencies` - sie schlagen fehl, ohne die Installation anderer Abhängigkeiten zu blockieren.

### ELM327 nicht gefunden

**Fehler**:
```
Port /dev/ttyUSB0 nicht gefunden
```

**Lösung**:
1. Port überprüfen: `ls -l /dev/ttyUSB*`
2. Berechtigungen prüfen: `sudo usermod -a -G dialout $USER`
3. Port in `OBDCommunicationService.js` ändern, falls anders

### Temperatursensor nicht gefunden

**Warnung**:
```
⚠️ DS18B20 Temperatursensor nicht verfügbar (1-Wire nicht gefunden)
```

**Lösung**:
1. 1-Wire aktivieren: `sudo raspi-config` → Interface Options → 1-Wire
2. Neustart: `sudo reboot`
3. Überprüfen: `ls /sys/bus/w1/devices/`
4. Wenn nicht erforderlich, in `gpio-mapping.js` deaktivieren: `temperature.enabled = false`

### Kraftstoffsensor antwortet nicht

**Warnung**:
```
⚠️ ADS1115 Kraftstoffsensor nicht verfügbar
```

**Lösung**:
1. I2C aktivieren: `sudo raspi-config` → Interface Options → I2C
2. Verbindung überprüfen: `sudo i2cdetect -y 1`
3. ADS1115-Verkabelung überprüfen
4. Wenn nicht erforderlich, in `gpio-mapping.js` deaktivieren: `fuel.enabled = false`

### Electron startet nicht

**Fehler**:
```
Error: connect ECONNREFUSED 127.0.0.1:5173
```

**Lösung**:
Der Vite-Client muss zuerst gestartet werden. Verwenden Sie `npm start`, das die Reihenfolge automatisch verwaltet.

### GPIO antwortet nicht

**Problem**: Warnleuchten werden nicht erkannt

**Lösung**:
1. Optokoppler-Verkabelung überprüfen
2. Pins testen: `gpio readall` (wiringpi installieren, falls erforderlich)
3. Pin-Mapping in `gpio-mapping.js` überprüfen
4. Active high/low-Logik der Optokoppler überprüfen

---

## 📱 Anwendungsnutzung

### Tastatursteuerung

- **`d`**: Debug-Konsole öffnen
- **`ESC`**: Debug-Konsole schließen
- **`r`**: Anwendung neu laden

### Debug-Konsole

Drücken Sie `d`, um die interaktive Konsole zu öffnen, die anzeigt:
- WebSocket-Protokolle
- Verbindungsfehler
- OBD-II-Daten in Echtzeit
- GPIO- und Sensorstatus

---

## 📦 Build für Produktion

### Client-Build

```bash
cd client
npm run build
```

Ausgabe in `client/dist/`

### Electron-Build

Um eine verteilbare App zu erstellen:

1. electron-builder installieren: `npm install --save-dev electron-builder`
2. Skript in `package.json` hinzufügen:

```json
"scripts": {
  "build:electron": "electron-builder"
}
```

3. Ausführen: `npm run build:electron`

---

## 📝 Hauptdateistruktur

```
cockpit/
├── client/
│   ├── src/
│   │   ├── config/
│   │   │   └── environment.ts          ← Client-Konfiguration
│   │   ├── components/                 ← React-Komponenten
│   │   ├── routes/
│   │   │   └── Cockpit/               ← Haupt-Dashboard
│   │   ├── services/
│   │   │   └── WebSocketService.ts    ← Client-WebSocket-Verwaltung
│   │   └── store/                     ← State Management (Valtio)
│   └── package.json
│
├── server/
│   ├── config/
│   │   └── gpio-mapping.js            ← ⚙️ GPIO- und Sensorkonfiguration
│   ├── services/
│   │   ├── OBDServer.js               ← Hauptserver
│   │   ├── OBDCommunicationService.js ← ELM327-Kommunikation
│   │   ├── GPIOService.js             ← GPIO-Verwaltung für Warnleuchten
│   │   ├── IgnitionService.js         ← Zündungsverwaltung
│   │   ├── TemperatureSensorService.js← DS18B20 Temperatursensor
│   │   └── FuelSensorService.js       ← ADS1115 Kraftstoffsensor
│   ├── scripts/
│   │   ├── low-power.sh               ← Energiespar-Skript
│   │   └── wake.sh                    ← Aufwach-Skript
│   ├── ecosystem.config.js            ← PM2-Konfiguration
│   └── package.json
│
├── main.js                            ← Electron-Wrapper
└── package.json                       ← Hauptskripte
```

---

## 🔒 Sicherheit und Hinweise

- ⚠️ **Nicht als Root ausführen**: Verwenden Sie normale Benutzerberechtigungen mit den Gruppen `dialout` und `gpio`
- 🔋 **Energiesparen**: Die Zündungsskripte können das System vor Batterieentladung schützen
- 🧪 **Testen**: Verwenden Sie immer den Mock-Modus zum Testen ohne Hardware
- 📊 **Überwachung**: Verwenden Sie PM2 zur Serverüberwachung in der Produktion

---

## 📄 Lizenz

Dieses Projekt wird unter der **GNU General Public License v3.0 oder später** veröffentlicht.

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

Der vollständige Lizenztext ist in der Datei [LICENSE](LICENSE) und unter https://www.gnu.org/licenses/gpl-3.0.html verfügbar

---

## 👥 Beitragen

Jeder Beitrag ist willkommen! Ob Code, Dokumentation, Fehlerberichte oder Vorschläge.

### 🚀 Wie man anfängt

1. Lesen Sie die [Anleitung zum Beitragen](.github/CONTRIBUTING.de.md)
2. Wählen Sie, wie Sie beitragen möchten:
   - 🐛 [Fehler melden](.github/ISSUE_TEMPLATE/bug_report.md)
   - ✨ [Feature vorschlagen](.github/ISSUE_TEMPLATE/feature_request.md)
   - ❓ [Frage stellen](.github/ISSUE_TEMPLATE/question.md)
   - 💻 Mit Code beitragen
   - 💡 Suchen Sie Inspiration in der [Roadmap & Wunschliste](.github/CONTRIBUTING.de.md#-möchten-sie-beitragen-aber-haben-keine-ideen)

### 📝 Beitrags-Workflow

1. **Fork** des Repositorys
2. **Branch erstellen**: `git checkout -b feature/feature-name`
3. **Änderungen vornehmen** gemäß [Code-Stil](.github/CONTRIBUTING.de.md#-code-stil)
4. **GPL-3.0-Header hinzufügen** zu neuen Quelldateien
5. **Commit**: `git commit -m 'feat: neue Funktion hinzugefügt'` ([Conventional Commits](https://www.conventionalcommits.org/))
6. **Push**: `git push origin feature/feature-name`
7. **Pull Request öffnen** mit ausgefüllter [Vorlage](.github/PULL_REQUEST_TEMPLATE.md)

### 📋 Verfügbare Vorlagen

- [🐛 Fehlerbericht](.github/ISSUE_TEMPLATE/bug_report.md) - Probleme melden
- [✨ Feature-Anfrage](.github/ISSUE_TEMPLATE/feature_request.md) - Verbesserungen vorschlagen
- [❓ Frage](.github/ISSUE_TEMPLATE/question.md) - Um Hilfe bitten
- [🔀 Pull Request](.github/PULL_REQUEST_TEMPLATE.md) - Mit Code beitragen

### 💡 Suchen Sie Ideen?

Sie wissen nicht, wo Sie anfangen sollen? Wir haben eine [Roadmap & Wunschliste](.github/CONTRIBUTING.de.md#-möchten-sie-beitragen-aber-haben-keine-ideen) von Features, die wir implementieren möchten:
- Rückfahrkamera und Parksensoren
- Erweiterte 3D-Animationen (Türen, Lichter am Modell)
- Anpassbare Dashboards und Themes
- Foto- und Video-Tutorials
- Internationalisierung
- Und vieles mehr!

Konsultieren Sie die [vollständige Anleitung zum Beitragen](.github/CONTRIBUTING.de.md) für alle Details.

---

## 📞 Support

Bei Problemen oder Fragen öffnen Sie ein Issue auf GitHub.

---

## 👨‍💻 Autoren

PandaOS wird entwickelt und gepflegt von:

- **[Matteo Errera](https://github.com/matteoerrera)**
- **[Roberto Zaccardi](https://github.com/rzaccardi)**
- **[Ludovico Verde](https://www.instagram.com/ludovico.verdee/)**

Siehe [AUTHORS.de.md](AUTHORS.de.md) für die vollständige Liste und Details zu den Mitwirkenden.
