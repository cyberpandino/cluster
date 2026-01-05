[🇬🇧 English](QUICK_START.en.md) | [🇮🇹 Italiano](QUICK_START.md) | [🇩🇪 Deutsch](QUICK_START.de.md)

---

# ⚡ Schnellstart - PandaOS Cluster

Kurzanleitung zum Einstieg in 5 Minuten.

> ⚠️ **ACHTUNG**: PandaOS ist ein experimentelles Hobbyprojekt. Die Installation in echten Fahrzeugen erfolgt auf eigene Gefahr. Lesen Sie den [vollständigen Haftungsausschluss](README.de.md#️-haftungsausschluss) vor dem Fortfahren.

---

## 🚀 Schnellstart

### 0. Voraussetzungen

Node.js 18+, npm 9+, Git 2.0+

```bash
node --version && npm --version && git --version
```

> 📘 Raspberry Pi: siehe [CONFIGURAZIONE_SERVER.de.md](server/CONFIGURAZIONE_SERVER.de.md#2-installation-von-nodejs-und-npm)

---

### 1. Installation

```bash
# Repository klonen
git clone git@github.com:cyberpandino/cluster.git
cd cluster

# Alle Abhängigkeiten installieren
npm run install:all
```

### 2. Basis-Konfiguration

#### Für lokale Entwicklung (Mac/Windows/Linux)

`client/src/config/environment.ts` bearbeiten:

```typescript
websocket: {
  mock: true,  // ← Demo-Modus
  // ...
}
```

#### Für Raspberry Pi (Produktion)

```typescript
websocket: {
  mock: false,  // ← Echte Verbindung zum Server
  // ...
}
```

### 3. Start

#### Lokale Entwicklung (Nur Client)

```bash
# Nur Oberfläche mit simulierten Daten
cd client
npm run dev
```

Browser öffnen: `http://localhost:5173`

#### Raspberry Pi (Vollständiger Stack)

```bash
# Server + Client + Electron
npm start
```

---

## 📋 Hardware-Checkliste (Raspberry Pi)

> 💡 **Müssen Sie Komponenten kaufen?** Konsultieren Sie [HARDWARE.de.md](HARDWARE.de.md) für die vollständige Liste von allem Notwendigen.

Vor dem Produktionsstart überprüfen:

- [ ] ELM327 an USB-Port angeschlossen (`/dev/ttyUSB0`)
- [ ] Optokoppler an GPIO-Pins verkabelt
- [ ] Benutzerberechtigungen konfiguriert:
  ```bash
  sudo usermod -a -G dialout,gpio,i2c $USER
  ```
- [ ] Schnittstellen über `raspi-config` aktiviert:
  - [ ] I2C (wenn Sie ADS1115 verwenden)
  - [ ] 1-Wire (wenn Sie DS18B20 verwenden)
  - [ ] Serial Port
- [ ] Neustart nach Konfiguration: `sudo reboot`

---

## 🎛️ Minimalkonfiguration

### Client (`client/src/config/environment.ts`)

```typescript
export const environment = {
  websocket: {
    url: 'http://127.0.0.1:3001',
    mock: true,  // true=Demo | false=Echt
  },
  debug: {
    enabled: true,
  },
};
```

### Server (`server/config/gpio-mapping.js`)

```javascript
module.exports = {
  // OBD Serielle Schnittstelle (ändern Sie bei Bedarf)
  // In OBDCommunicationService.js:
  portPath: '/dev/ttyUSB0',
  
  // GPIO-Pins für Warnleuchten (siehe vollständige Tabelle)
  mapping: {
    turnSignals: { pin: 17 },
    battery: { pin: 27 },
    highBeam: { pin: 5 },
    // ... andere
  },
  
  // Optionale Sensoren
  temperature: { enabled: true },
  fuel: { enabled: true },
  ignition: { enabled: true },
};
```

---

## 🔑 Tastatursteuerung

Nach dem Start der Anwendung:

- **`d`** → Debug-Konsole öffnen
- **`ESC`** → Debug-Konsole schließen
- **`r`** → Anwendung neu laden

---

## 🐛 Schnelle Fehlerbehebung

### "Server startet nicht"

**Auf Mac/Windows**: Normal! Der Server benötigt Raspberry Pi.  
**Lösung**: Verwenden Sie `mock: true` im Client.

### "ELM327 nicht gefunden"

```bash
# Port überprüfen
ls -l /dev/ttyUSB*

# Berechtigungen erteilen
sudo usermod -a -G dialout $USER
# Abmelden und anmelden
```

### "Client verbindet sich nicht"

1. Server gestartet überprüfen: `npm run server`
2. URL prüfen: `websocket.url` in `environment.ts`
3. Port 3001 frei überprüfen: `lsof -i :3001`

### "GPIO funktioniert nicht"

```bash
# GPIO-Pin 17 testen
gpio -g read 17

# Bei Fehler:
sudo usermod -a -G gpio $USER
# Neustart
```

---

## 📚 Vollständige Dokumentation

- **[README.de.md](README.de.md)** → Vollständige Hauptdokumentation
- **[client/CONFIGURAZIONE.de.md](client/CONFIGURAZIONE.de.md)** → Detaillierte Client-Konfiguration
- **[server/CONFIGURAZIONE_SERVER.de.md](server/CONFIGURAZIONE_SERVER.de.md)** → Hardware- und Server-Setup
- **[ARCHITETTURA.de.md](ARCHITETTURA.de.md)** → Systemarchitektur

---

## 🎯 Nächste Schritte

1. **UI-Entwicklung**: Komponenten in `client/src/components/` ändern
2. **GPIO anpassen**: `server/config/gpio-mapping.js` an Ihre Verkabelung anpassen
3. **PID hinzufügen**: `server/services/PIDParserService.js` erweitern
4. **Styling**: `client/src/assets/scss/` ändern
5. **Produktion**: PM2-Setup (siehe README.de.md)

---

## 📞 Hilfe

Haben Sie Probleme?

1. **Dokumentation konsultieren**:
   - [README.de.md](README.de.md) - Allgemeine Fehlerbehebung
   - [server/CONFIGURAZIONE_SERVER.de.md](server/CONFIGURAZIONE_SERVER.de.md) - Hardware-Probleme
   - [client/CONFIGURAZIONE.de.md](client/CONFIGURAZIONE.de.md) - Client-Probleme

2. **Issue öffnen**:
   - [🐛 Bug Report](.github/ISSUE_TEMPLATE/bug_report.md) - Problem melden
   - [❓ Frage](.github/ISSUE_TEMPLATE/question.md) - Frage stellen

3. **Beitragen**:
   - [✨ Feature Request](.github/ISSUE_TEMPLATE/feature_request.md) - Verbesserungen vorschlagen
   - [CONTRIBUTING.de.md](.github/CONTRIBUTING.de.md) - Anleitung zum Beitragen

---

**Viel Spaß beim Coding und keinen Schaden anrichten! 🚗💨**
