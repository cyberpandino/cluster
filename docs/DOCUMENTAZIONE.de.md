[🇬🇧 English](DOCUMENTAZIONE.en.md) | [🇮🇹 Italiano](DOCUMENTAZIONE.md) | [🇩🇪 Deutsch](DOCUMENTAZIONE.de.md)

---

# 📚 PandaOS Cluster - Dokumentationsindex

Willkommen bei der deutschen technischen Dokumentation von **PandaOS Cluster**.

Dieser Index führt Sie durch alle verfügbaren Ressourcen, um das Projekt zu verstehen, zu konfigurieren und zu verwenden.

---

## 📖 Hauptdokumentation

### 🚀 [Quick Start](QUICK_START.de.md)
Schnelle Anleitung zum ersten Start des Projekts in weniger als 15 Minuten.

**Enthält**:
- Grundlegende Hardware-/Softwarevoraussetzungen
- Schnellinstallation von Client und Server
- Erste Start- und Verbindungstests
- Typische Probleme und Lösungen

**Ideal für**: Benutzer, die das System sofort testen möchten.

---

### 🏗️ [Architektur](ARCHITETTURA.de.md)
Detaillierte technische Dokumentation der Software- und Hardwarearchitektur des Systems.

**Enthält**:
- Architekturdiagramme (Client-Server, Datenfluss)
- Beschreibung der Module und Services
- Kommunikationsprotokolle (WebSocket, OBD-II)
- Hardware-Integration (GPIO, I2C, 1-Wire, Serial)
- Performance- und Sicherheitsüberlegungen

**Ideal für**: Entwickler, die den Quellcode verstehen oder erweitern möchten.

---

### 🔧 [Hardware](HARDWARE.de.md)
Vollständige Liste der Hardware-Komponenten und Montageanleitungen.

**Enthält**:
- Raspberry Pi 4B und Peripheriegeräte
- OBD-II-Adapter (ELM327)
- Digitales Instrumententafel (Waveshare)
- Sensoren (DS18B20, ADS1115)
- Optokoppler für Warnleuchten
- Detaillierte Verkabelungsschaltpläne
- Einkaufsliste mit Links

**Ideal für**: Diejenigen, die das physische System bauen.

---

## ⚙️ Konfiguration

### 🖥️ [Client-Konfiguration](client/CONFIGURAZIONE.de.md)
Konfiguration der Electron/React-Anwendung.

**Enthält**:
- Umgebungsvariablen (`environment.ts`)
- Themenkonfiguration (Hell/Dunkel)
- WebSocket-Verbindungseinstellungen
- Mock-Modus vs. Produktionsmodus
- Konstanten und Limits (RPM, Geschwindigkeit)

**Ideal für**: Anpassung der Benutzeroberfläche und Entwicklungseinstellungen.

---

### 🖧 [Server-Konfiguration](server/CONFIGURAZIONE_SERVER.de.md)
Konfiguration des Node.js-Servers und der Hardware-Integration.

**Enthält**:
- OBD-II-Verbindung (serielle Port-, Baudrate)
- GPIO-Mapping für Warnleuchten
- DS18B20-Temperatursensor (1-Wire)
- ADS1115-Kraftstoffsensor (I2C)
- Zündungsverwaltung (Ignition)
- Skripte für Energiesparmodus (`low-power.sh`, `wake.sh`)
- Vollständige Fehlerbehebungsanleitungen

**Ideal für**: Diejenigen, die Hardware an den Raspberry Pi anschließen.

---

### 🎛️ [Konfigurationsübersicht](client/src/config/README.de.md)
Vereinfachter Leitfaden zu Konfigurationsdateien im Client-Modul.

**Enthält**:
- Struktur von `client/src/config/`
- Beschreibung von `constants.ts`, `environment.ts`, `time.ts`
- Verwendungsbeispiele in Komponenten

**Ideal für**: Schneller Überblick über Client-Konfigurationsoptionen.

---

## 🤝 Mitwirken

### 👥 [Mitwirkende](AUTHORS.de.md)
Liste der Autoren, Mitwirkenden und Anerkennungen.

---

### 🛠️ [Contributing Guide](../.github/CONTRIBUTING.de.md)
Leitfaden zur Teilnahme an der Projektentwicklung.

**Enthält**:
- Code-Konventionen (TypeScript, JavaScript)
- Git-Workflow (Branching, Commits)
- Wie man Issues und Pull Requests öffnet
- Code-Review-Richtlinien

**Ideal für**: Externe Entwickler, die einen Beitrag leisten möchten.

---

## 🗺️ Projektplanung

### 🛣️ [Roadmap](ROADMAP.de.md)
Entwicklungsplan und zukünftige Features.

**Enthält**:
- Bereits abgeschlossene Features (v0.1.0 → v0.9.0)
- In Arbeit befindliche Features
- Zukünftige Prioritäten (v1.0.0 und darüber hinaus)
- Experimentelle Ideen

**Ideal für**: Verständnis der Projektrichtung.

---

## 📂 Dokumentationsstruktur

```
docs/
├── ARCHITETTURA.de.md           # Architektur (dieses Dokument)
├── HARDWARE.de.md               # Hardware-Komponenten
├── QUICK_START.de.md            # Schnellstart-Anleitung
├── ROADMAP.de.md                # Entwicklungs-Roadmap
├── AUTHORS.de.md                # Autoren und Mitwirkende
├── DOCUMENTAZIONE.de.md         # Index (dieses Dokument)
├── client/
│   ├── CONFIGURAZIONE.de.md     # Client-Konfiguration
│   └── src/config/
│       └── README.de.md         # Konfig-Dateien-Übersicht
└── server/
    └── CONFIGURAZIONE_SERVER.de.md  # Server-/Hardware-Konfiguration
```

---

## 🔗 Externe Ressourcen

### Offizielle Dokumentation
- **Raspberry Pi**: https://www.raspberrypi.org/documentation/
- **ELM327 Protokoll**: https://www.elmelectronics.com/wp-content/uploads/2017/01/ELM327DS.pdf
- **OBD-II PIDs**: https://en.wikipedia.org/wiki/OBD-II_PIDs
- **Socket.IO**: https://socket.io/docs/v4/
- **React**: https://react.dev/
- **Electron**: https://www.electronjs.org/docs/latest

### Schaltpläne und Datenblätter
- **Fiat Panda 141 - Offizieller elektrischer Schaltplan**: [PDF](http://www.bunkeringegnere.altervista.org/esplosi/FIAT%20PANDA/panda%20141/1100%20mpi/55%20IMPIANTO%20ELETTRICO%20-%20SCHEMI%20-%20GAMMA%202000.pdf)
- **DS18B20**: https://datasheets.maximintegrated.com/en/ds/DS18B20.pdf
- **ADS1115**: https://www.ti.com/lit/ds/symlink/ads1115.pdf
- **PC817 Optokoppler**: https://www.farnell.com/datasheets/73758.pdf

---

## 📞 Support

Bei Fragen, Problemen oder Anregungen:

1. **GitHub Issues**: [github.com/cyberpandino/cluster/issues](https://github.com/cyberpandino/cluster/issues)
2. **Diskussionen**: [github.com/cyberpandino/cluster/discussions](https://github.com/cyberpandino/cluster/discussions)

---

**Zuletzt aktualisiert**: v0.9.0  
**Sprachen**: [🇬🇧 English](DOCUMENTAZIONE.en.md) | [🇮🇹 Italiano](DOCUMENTAZIONE.md) | [🇩🇪 Deutsch](DOCUMENTAZIONE.de.md)
