[🇬🇧 English](ROADMAP.en.md) | [🇮🇹 Italiano](ROADMAP.md) | [🇩🇪 Deutsch](ROADMAP.de.md)

---

# 🗺️ Roadmap & Wishlist - PandaOS Cluster

Features und Verbesserungen, die wir im Projekt implementieren möchten.

> 💡 **Möchten Sie beitragen?** Wählen Sie ein Feature aus der Liste und öffnen Sie eine [Feature Request](https://github.com/cyberpandino/cluster/issues/new?template=feature_request.md)!

---

## 🚗 Hardware-Features

### Hohe Priorität

#### 📹 Integrierte Rückfahrkamera
**Beschreibung**: Anzeige der Rückfahrkamera im Cluster beim Einlegen des Rückwärtsgangs  
**Komplexität**: Mittel  
**Komponenten**: USB/CSI-Kamera, Rückwärtsgang-Erkennung via GPIO  
**Vorteile**: Parksicherheit, Ersatz für beschädigten Rückspiegel

#### 📡 Parksensoren
**Beschreibung**: Grafische Darstellung der Hindernisentfernung mit Ultraschallsensoren  
**Komplexität**: Mittel  
**Komponenten**: 4-8 Ultraschallsensoren, Arduino/ESP32 für Verarbeitung  
**Vorteile**: Parkassistent wie bei modernen Autos

### Mittlere Priorität

- **Regensensor** - Automatische Scheibenwischerregelung
- **Helligkeitssensor** - Auto-Regelung der Display-Helligkeit (Tag/Nacht-Modus)
- **Reifendruckkontrolle (TPMS)** - Integration von Reifendrucksensoren
- **CAN-Bus-Unterstützung** - Zusätzlich zu OBD-II, Unterstützung für natives CAN-Protokoll
- **360°-Kamera** - Multi-Kamera-System für Rundumansicht

---

## 💻 Software-Features

### Hohe Priorität

#### 🛣️ Bordcomputer-System
**Beschreibung**: Fahrtaufzeichnung mit Verbrauch, Strecke, Zeit, Route  
**Komplexität**: Mittel  
**Technologien**: Lokale Datenbank (SQLite), Geolokalisierung  
**Vorteile**: Detaillierte Statistiken, Verbrauchsanalyse, Fahrthistorie

#### 🎨 Anpassbare Dashboards
**Beschreibung**: Mehrere vom Benutzer wählbare Layouts (Sport, Eco, Minimal, Full)  
**Komplexität**: Hoch  
**Technologien**: React-Layout-System, persistente Speicherung  
**Vorteile**: Personalisierte Erfahrung, Anpassung an Präferenzen

#### 🌓 Farbthemen
**Beschreibung**: Dunkelmodus, Hellmodus, benutzerdefinierte Themen (Retro, Futuristisch, etc.)  
**Komplexität**: Niedrig  
**Technologien**: CSS-Variablen, Theme-Provider  
**Vorteile**: Anpassung an Umgebungshelligkeit, personalisierte Ästhetik

#### 🎯 Unterstützte Kalibrierung
**Beschreibung**: Schritt-für-Schritt-Assistent zur Kalibrierung von Kraftstoff-/Temperatursensoren  
**Komplexität**: Mittel  
**Technologien**: UI-Assistent, Kalibrierungsspeicherung  
**Vorteile**: Einfacheres Setup für nicht-technische Benutzer

#### 📱 Mobile Companion-App
**Beschreibung**: Smartphone-App für Fahrzeugstatistiken, Benachrichtigungen, Fernsteuerung  
**Komplexität**: Hoch  
**Technologien**: React Native / Flutter, REST/WebSocket-API  
**Vorteile**: Datenzugriff auch außerhalb des Fahrzeugs

#### 🚪 3D-Türanimation
**Beschreibung**: Darstellung geöffneter/geschlossener Türen am 3D-Panda-Modell im Cluster  
**Komplexität**: Niedrig  
**Technologien**: Three.js, 3D-Modellanimationen, GPIO-Integration  
**Vorteile**: Sofortiges visuelles Feedback, immersivere Benutzeroberfläche

#### 💡 Lichter am 3D-Modell
**Beschreibung**: Anzeige eingeschalteter Lichter (Fernlicht, Blinker, Nebelscheinwerfer) direkt am 3D-Modell  
**Komplexität**: Niedrig  
**Technologien**: Three.js-Materialien, emissive Texturen, bestehende GPIO-Daten  
**Vorteile**: Intuitive Visualisierung des Lichtstatus, ansprechendere Benutzeroberfläche

### Mittlere Priorität

- **Geplante Wartung** - Warnungen für Service, Ölwechsel, TÜV basierend auf km
- **Wetterintegration** - Außentemperatur von API, falls kein Sensor verfügbar
- **Automatischer Nacht-/Tagmodus** - Basierend auf Zeit/GPS oder Helligkeitssensor
- **Gestensteuerung** - Gestenbedienung (bei Touch-Display)
- **Sprachsteuerung** - Sprachbefehle für Hauptfunktionen (Annyang.js)
- **Multi-Benutzerprofil** - Separate Statistiken für verschiedene Fahrer
- **Datenexport** - CSV/JSON/Excel für externe Analyse
- **Push-Benachrichtigungssystem** - Audio-/visuelle Warnungen für Anomalien/Wartung

---

## 📚 Dokumentation

### Hohe Priorität

#### 📸 Foto-Verkabelungs-Tutorial
**Beschreibung**: Schritt-für-Schritt-Anleitung mit echten Fotos der Optokoppler-Verkabelung  
**Komplexität**: Niedrig (erfordert aber echte Installation)  
**Inhalt**: Detaillierte Fotos jeder Phase, Kabelidentifikation, Verbindungen  
**Vorteile**: Reduziert Installationsfehler drastisch

#### 🎥 Video-Installationsanleitung
**Beschreibung**: Vollständiges Video-Tutorial von Verkabelung bis Software  
**Komplexität**: Mittel  
**Inhalt**: Montiertes Video mit Voice-over, Untertiteln, Kapiteln  
**Vorteile**: Zugänglicheres Format für weniger technische Benutzer

#### 🌍 Internationalisierung (i18n)
**Beschreibung**: EN, ES, DE, FR Übersetzungen der Benutzeroberfläche  
**Komplexität**: Mittel  
**Technologien**: react-i18next (bereits vorhanden), JSON-Übersetzungsdateien  
**Vorteile**: Internationale Zugänglichkeit, breitere Community

#### 📝 Zentralisierte Übersetzungsdatei
**Beschreibung**: Alle hartkodierten Microcopy in JSON/i18n-Dateien verschieben  
**Komplexität**: Niedrig  
**Technologien**: i18next, JSON  
**Vorteile**: Einfache Wartung, community-gesteuerte Übersetzungen

#### 🔌 Benutzerdefiniertes PCB-Schema
**Beschreibung**: Professionelles PCB-Design für Optokoppler (KiCad/Eagle)  
**Komplexität**: Hoch  
**Technologien**: KiCad, Gerber-Export  
**Vorteile**: Saubere Installation, kein Breadboard, günstige PCB-Bestellung

### Mittlere Priorität

- **Erweiterte FAQ** - Häufig gestellte Fragen mit detailliertem Troubleshooting
- **Installation-Fallstudien** - Echte Beispiele mit Fotos und Protokollen
- **Kompatibilitätsleitfäden** - Liste kompatibler Fahrzeuge (Uno, Seicento, etc.)
- **Interaktives Verkabelungsdiagramm** - Navigierbarer Online-Schaltplan (SVG/HTML)

---

## 🧪 Testing & Qualität

- **Unit-Tests** - Automatisiertes Testen von Backend-Services (Jest)
- **E2E-Tests** - Vollständige Oberflächentests (Playwright/Cypress)
- **CI/CD-Pipeline** - GitHub Actions für automatische Builds und Deployments
- **Performance-Profiling** - Optimierung von Rendering und Speicher
- **Hardware-in-the-Loop-Testing** - Automatisierte Tests mit simulierter Hardware (Mock GPIO/Serial)
- **Stresstest** - Langzeit-Stabilitätstest
- **Code-Abdeckung** - Abdeckung >80% für kritischen Code

---

## 🔧 Kompatibilität & Erweiterungen

### Fahrzeuge

- **Fiat Uno** (1983-1995) - Ähnliches OBD-I/II
- **Fiat Seicento** (1998-2010) - OBD-II kompatibel
- **Fiat Punto** (erste Serie) - Gleiche Magneti Marelli Steuergerät
- **Lancia Y** (erste Serie) - Ähnliche ECUs
- **Fiat Tipo** - Verwandte Mechanik

### Protokolle

- **J1850 PWM/VPW** - Amerikanische Protokolle
- **CAN-Bus (ISO 15765)** - Neuere Fahrzeuge
- **LIN-Bus** - Automotive-Zubehör

### Hardware

- **Raspberry Pi Zero 2W** - Kompakte Version mit niedrigem Stromverbrauch
- **Raspberry Pi CM4** - Für individuelle Integrationen
- **Orange Pi / Banana Pi** - Günstige SBC-Alternativen
- **Android-Tablets** - Portierung der App auf natives Android

### Integrationen

- **Android Auto / CarPlay** - Integration mit modernen Systemen
- **Aftermarket-Head-Unit** - Kompatibilität mit chinesischen Android-Radios
- **Automatisches Backup** - Cloud/USB-Wiederherstellung/Backup von Einstellungen
- **OTA-Updates** - Over-the-Air-Software-Updates

---

## 🎨 UI/UX

- **Übergangsanimationen** - Flüssige GSAP-Übergänge zwischen Zuständen/Bildschirmen
- **Sport-/Eco-Modus** - Verschiedene Visualisierungen für Fahrstil (aggressiv rot / effizient grün)
- **Anpassbare Widgets** - Drag & Drop von Komponenten im Cluster
- **Skins/Community-Themen** - Marktplatz zum Teilen individueller Themen
- **Minimaler Modus** - Essentielle Benutzeroberfläche für weniger Ablenkung beim Fahren
- **Bildschirmschoner** - Animationen bei stehendem Fahrzeug
- **Startanimation** - Anpassbare Boot-Animation
- **Easter Eggs** - Spezialanimationen für Ereignisse (Weihnachten, Auto-Geburtstag, etc.)

---

## 🌐 Service-Integrationen

- **Google Maps / OpenStreetMap** - Integrierte Navigation im Cluster
- **Spotify / Apple Music** - Musiksteuerung vom Cluster
- **Telegram-Bot** - Benachrichtigungen und Fernsteuerung
- **IFTTT / Home Assistant** - Smart-Home-Automatisierungen
- **Kraftstoffpreis-API** - Kraftstoffpreise in Echtzeit
- **Verkehrsdaten** - Verkehrsinformationen
- **Wetter-API** - Integrierte Wettervorhersage

---

## 🔐 Sicherheit & Datenschutz

- **Benutzerauthentifizierung** - Login für Zugriff auf Statistiken/Konfiguration
- **Datenverschlüsselung** - Schutz sensibler gespeicherter Daten
- **VPN-Client** - Sichere Verbindung für Cloud-Daten
- **Datenschutzmodus** - Deaktivierung von GPS-Tracking/Statistiken

---

## 🎯 Performance & Optimierung

- **Verbessertes Lazy Loading** - On-Demand-Laden von Komponenten
- **Service Workers** - PWA für Cache und Offline-Betrieb
- **WebGL-Optimierung** - Optimiertes 3D-Rendering
- **Speicherverwaltung** - Reduzierung des Speicher-Footprints
- **Startzeit** - Reduzierung der Startzeit auf <10 Sekunden

---

## 📊 Aktueller Status

### Version 0.9.0 (Aktuell)

**Implementiert**:
- ✅ OBD-II-Datenauslese via ELM327
- ✅ GPIO-Warnlicht-Erkennung mit Optokopplern
- ✅ DS18B20-Temperatursensor
- ✅ ADS1115-Kraftstoffsensor
- ✅ 3D-Dashboard mit interaktivem Panda-Modell
- ✅ Mock-Modus für Entwicklung
- ✅ Integrierte Debug-Konsole
- ✅ Zündungsverwaltung (Energiesparen)
- ✅ WebSocket-Echtzeit
- ✅ Vollständige Dokumentation

**In Entwicklung**:
- 🔄 Derzeit keine Feature in aktiver Entwicklung

**Geplant**:
- 📋 Siehe Roadmap oben

### Version 1.0.0 (Ziel)

**Ziel-Features**:
- Funktionierende Rückfahrkamera
- Parksensoren
- Vollständige 3D-Animationen
- Anpassbare Dashboards
- Internationalisierung
- Vollständige Tests

---

## 🤝 Wie Sie zu diesen Features beitragen können

### 1. Wählen Sie ein Feature
Durchsuchen Sie die Liste und finden Sie etwas, das:
- Sie begeistert
- Sie haben die Fähigkeiten zur Implementierung
- Ein Problem/Bedürfnis von Ihnen löst

### 2. Öffnen Sie eine Issue
Verwenden Sie [Feature Request](https://github.com/cyberpandino/cluster/issues/new?template=feature_request.md) zur Diskussion:
- Implementierungsansatz
- Zu verwendende Technologien
- Kompatibilität mit Bestehendem
- Geschätzte Timeline

### 3. Entwickeln
- Repository forken
- Dediziertes Branch erstellen
- Implementierung gemäß [Code-Stil](.github/CONTRIBUTING.de.md#-code-stil)
- Gründlich testen

### 4. Pull Request
- PR mit [Vorlage](.github/PULL_REQUEST_TEMPLATE.md) öffnen
- Implementierung beschreiben
- Screenshots/Videos anhängen
- Auf Review warten

---

## 💬 Diskussion

Möchten Sie Features diskutieren, Alternativen vorschlagen oder Ideen teilen?
- Öffnen Sie eine [Diskussion](https://github.com/cyberpandino/cluster/discussions) (falls aktiviert)
- Oder eine [Question Issue](https://github.com/cyberpandino/cluster/issues/new?template=question.md)

---

## 🛠️ Rewrite auf Native Stack?

### Das große Refactoring

Ein "spezielles" Feature, das eine separate Erwähnung verdient:

#### ⚙️ Port auf C++/Qt Native

**Beschreibung**: PandaOS mit professionellem Automotive-Technologie-Stack neu schreiben  
**Komplexität**: Sehr hoch (praktisch ein komplettes Rewrite)  
**Technologien**:
- **C++17/20** für Backend und Business-Logik
- **Qt 6 / QML** für grafische Benutzeroberfläche (Automotive-Standard)
- **Qt 3D** für interaktives Panda-Modell
- **Yocto/Buildroot** für optimiertes Embedded-Linux
- **systemd** für Service-Verwaltung
- **D-Bus** für IPC zwischen Prozessen

**Vorteile**:
- 🚀 **Startzeit** <3 Sekunden (vs ~15s aktuell)
- 💾 **Speicher** ~50MB insgesamt (vs ~500MB aktuell)
- ⚡ **Performance** garantierte 60fps Rendering auch auf Pi Zero
- 🔋 **Energieverbrauch** um 60-70% reduziert
- 🏭 **Professioneller Ansatz** produktionsreif

**Warum haben wir es nicht gemacht?**

Weil wir nach 2 Stunden Kampf mit CMake den einfacheren Weg gewählt haben: npm. Und ehrlich gesagt bereuen wir es nicht. 😅

Der Web-Stack ermöglichte uns:
- Etwas Funktionierendes in Wochen, nicht Monaten/Jahren
- Vermeidung der Komplexität der ARM-Cross-Compilation
- Nutzung leistungsstarker Bibliotheken wie Three.js ohne manuelles Schreiben von OpenGL-Shadern
- Fokus auf UX statt auf Segfault-Debugging

**Aber wenn Sie die Herausforderung annehmen möchten...**

Wir wären **begeistert**, wenn jemand einen nativen Port machen würde! Wir könnten haben:
- **PandaOS-Web** (aktuell) - Schneller und zugänglicher Ansatz
- **PandaOS-Native** (Zukunft?) - Professioneller und leistungsstarker Ansatz

Wenn Sie interessiert sind, öffnen Sie eine Diskussion. Wir bewundern jeden, der die Geduld hat, Qt und CMake zu meistern. 🚀

---

## 📅 Timeline

Es gibt keine festen Timelines. Das Projekt ist ein Hobby und schreitet voran, wenn Zeit und Leidenschaft vorhanden sind.  
**Community-Beiträge sind entscheidend, um die Entwicklung zu beschleunigen!**

---

## 🤷 Aber ernsthaft, warum dieser absurde Stack?

**Kurze Antwort**: Weil es Spaß macht und wir nicht verrückt werden wollen.

**Lange Antwort**: [README - Aber React + Electron auf Automotive?!](README.de.md#-aber-react--electron-auf-einem-automotive-seid-ihr-verrückt)

**Ehrliche Antwort**: Wenn wir es "richtig" machen würden, würden wir C++/Qt verwenden. Aber das würde viel mehr Zeit und spezifisches Know-how erfordern. Vorerst funktioniert unser "kreativer" Ansatz perfekt für den Zweck. 😅

---

**Letzte Revision**: November 2025  
**Nächstes Update**: Wenn wir neue brillante Ideen haben 💡
