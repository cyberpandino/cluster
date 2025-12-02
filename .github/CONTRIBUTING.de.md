[🇬🇧 English](CONTRIBUTING.en.md) | [🇮🇹 Italiano](CONTRIBUTING.md) | [🇩🇪 Deutsch](CONTRIBUTING.de.md)

---

# 🤝 Zu PandaOS beitragen

Vielen Dank für Ihr Interesse, zu PandaOS beizutragen! Jeder Beitrag ist willkommen, sei es Code, Dokumentation, Fehlerberichte oder Vorschläge.

## 📋 Inhaltsverzeichnis

- [Verhaltenskodex](#-verhaltenskodex)
- [Wie man beiträgt](#-wie-man-beiträgt)
- [Fehler melden](#-fehler-melden)
- [Neue Features vorschlagen](#-neue-features-vorschlagen)
- [Pull Requests](#-pull-requests)
- [Code-Stil](#-code-stil)
- [Dokumentation](#-dokumentation)
- [Lizenz](#-lizenz)

---

## 🤝 Verhaltenskodex

Dieses Projekt hält sich an einen impliziten Verhaltenskodex, der auf gegenseitigem Respekt basiert:

- Seien Sie respektvoll gegenüber allen Mitwirkenden
- Akzeptieren Sie konstruktive Kritik mit Offenheit
- Konzentrieren Sie sich darauf, was für die Community am besten ist
- Zeigen Sie Empathie gegenüber anderen Community-Mitgliedern

## 🎯 Wie man beiträgt

Es gibt viele Möglichkeiten, zu PandaOS beizutragen:

### 🐛 Fehler melden
Haben Sie einen Fehler gefunden? Öffnen Sie ein [Issue](https://github.com/cyberpandino/cluster/issues/new?template=bug_report.md) mit der Vorlage "Bug Report".

### ✨ Features vorschlagen
Haben Sie eine Idee zur Verbesserung von PandaOS? Öffnen Sie ein [Issue](https://github.com/cyberpandino/cluster/issues/new?template=feature_request.md) mit der Vorlage "Feature Request".

### 📚 Dokumentation verbessern
- Tippfehler oder Fehler korrigieren
- Bestehende Erklärungen verbessern
- Beispiele und Anleitungen hinzufügen
- Dokumentation übersetzen

### 💻 Mit Code beitragen
- Offene Fehler beheben
- Neue Features implementieren
- Performance optimieren
- Tests hinzufügen

### 🧪 Testing
- Auf unterschiedlicher Hardware testen
- Kompatibilität überprüfen
- Probleme spezifisch für Ihr Setup melden

---

## 🐛 Fehler melden

Bevor Sie einen Fehler melden:

1. **Suchen Sie nach bestehenden Issues**: Überprüfen Sie, ob der Fehler nicht bereits gemeldet wurde
2. **Verwenden Sie den Mock-Modus**: Testen Sie im Mock-Modus, um Hardware-Probleme auszuschließen
3. **Sammeln Sie Informationen**: Bereiten Sie Logs, Screenshots und Umgebungsdetails vor

**Verwenden Sie die Vorlage**: [Bug Report](https://github.com/cyberpandino/cluster/issues/new?template=bug_report.md)

### Zu enthaltende Informationen

- **Klare Beschreibung** des Problems
- **Schritte zur Reproduktion** des Verhaltens
- **Erwartetes Verhalten** vs. **tatsächliches Verhalten**
- **Umgebung**: OS, Hardware, Software-Versionen
- **Vollständige Logs** vom Server und/oder Client
- **Screenshots**, falls relevant

---

## ✨ Neue Features vorschlagen

Bevor Sie ein Feature vorschlagen:

1. **Überprüfen Sie, ob es bereits existiert**: Suchen Sie in offenen/geschlossenen Issues und PRs
2. **Berücksichtigen Sie den Umfang**: Passt das Feature zu den Projektzielen?
3. **Denken Sie über die Implementierung nach**: Haben Sie Ideen, wie es implementiert werden könnte?

**Verwenden Sie die Vorlage**: [Feature Request](https://github.com/cyberpandino/cluster/issues/new?template=feature_request.md)

### Diskussion

Für komplexe Features wird empfohlen:
1. Zuerst ein Issue öffnen, um zu diskutieren
2. Auf Feedback von den Maintainern warten
3. Nach Genehmigung mit der Implementierung fortfahren

---

## 🔀 Pull Requests

### Workflow

1. **Fork** des Repositories
2. **Branch erstellen** von `main`:
   ```bash
   git checkout -b feature/feature-name
   # oder
   git checkout -b fix/bug-name
   ```
3. **Änderungen vornehmen**
4. **Commit** mit klaren Nachrichten:
   ```bash
   git commit -m "feat: add support for XYZ sensor"
   git commit -m "fix: correct GPIO pin 17 reading"
   git commit -m "docs: update configuration guide"
   ```
5. **Push** zu Ihrem Fork:
   ```bash
   git push origin feature/feature-name
   ```
6. **Pull Request öffnen** zu `main`

### Commit-Konventionen

Verwenden Sie [Conventional Commits](https://www.conventionalcommits.org/):

- `feat:` neues Feature
- `fix:` Fehlerbehebung
- `docs:` Dokumentationsänderungen
- `style:` Formatierung, fehlende Semikolons, etc.
- `refactor:` Code-Refactoring
- `perf:` Performance-Verbesserungen
- `test:` Tests hinzufügen
- `chore:` Abhängigkeiten aktualisieren, Konfiguration, etc.

**Beispiele**:
```bash
feat(server): add CAN protocol support
fix(client): correct tachometer rendering on Safari
docs(readme): update Raspberry Pi installation instructions
refactor(gpio): simplify debouncing logic
```

### PR-Checkliste

Vor dem Öffnen der PR überprüfen Sie:

- [ ] Der Code kompiliert ohne Fehler
- [ ] Sie haben die Änderungen lokal getestet
- [ ] Sie haben die Dokumentation hinzugefügt/aktualisiert
- [ ] Sie haben den GPL-3.0-Header zu neuen Dateien hinzugefügt
- [ ] Die Commits sind atomar und gut beschrieben
- [ ] Sie haben auf den aktualisierten main rebased
- [ ] Es gibt keine Konflikte

### Review-Prozess

1. PR öffnen und die Vorlage ausfüllen
2. Die Maintainer überprüfen den Code
3. Angeforderte Änderungen anwenden (falls erforderlich)
4. Nach Genehmigung wird die PR gemergt

---

## 🎨 Code-Stil

### JavaScript/Node.js (Server)

- **Einrückung**: 2 Leerzeichen
- **Anführungszeichen**: Einfache Anführungszeichen `'`
- **Semikolons**: Ja
- **Benennung**:
  - `camelCase` für Variablen und Funktionen
  - `PascalCase` für Klassen
  - `UPPER_CASE` für Konstanten

**Beispiel**:
```javascript
const MAX_RETRIES = 3;

class OBDService {
  constructor() {
    this.retryCount = 0;
  }

  async readPID(pidCode) {
    // implementation
  }
}
```

### TypeScript/React (Client)

- **Einrückung**: 2 Leerzeichen
- **Anführungszeichen**: Doppelte Anführungszeichen `"`
- **Semikolons**: Ja
- **Components**: PascalCase
- **Hooks**: camelCase mit Präfix `use`

**Beispiel**:
```typescript
interface OdometerProps {
  speed: number;
  rpm: number;
}

export const Odometer: React.FC<OdometerProps> = ({ speed, rpm }) => {
  const [isActive, setIsActive] = useState(false);
  
  return <div>{speed} km/h</div>;
};
```

### Kommentare

- Kommentieren Sie komplexen oder nicht offensichtlichen Code
- Verwenden Sie JSDoc für öffentliche Funktionen
- Erklären Sie das "Warum", nicht das "Was"

**Beispiel**:
```javascript
/**
 * Reads a PID from the ECU with automatic retry
 * @param {string} pid - PID code in hex format (e.g. '010C')
 * @param {string} name - Descriptive parameter name
 * @returns {Promise<Object>} PID reading result
 */
async readPID(pid, name) {
  // Wake up ECU if in sleep mode to avoid timeout
  await this.wakeupECU();
  
  // Implementation...
}
```

---

## 📚 Dokumentation

Dokumentation ist entscheidend! Jede Änderung sollte Aktualisierungen der relevanten Dokumentation enthalten.

### Zu aktualisierende Dateien

| Änderung | Dokumentation |
|----------|---------------|
| Neues Client-Feature | `client/CONFIGURAZIONE.md` |
| Neues Server-Feature | `server/CONFIGURAZIONE_SERVER.md` |
| Architekturänderung | `ARCHITETTURA.md` |
| Neue Konfiguration | `README.md` + spezifische Datei |
| Setup/Installation | `README.md` + `QUICK_START.md` |

### Dokumentations-Stil

- **Sprache**: Deutsch (für .de.md-Dateien)
- **Format**: Markdown
- **Ton**: Informell aber technisch
- **Abschnitte**: Gut strukturiert mit Emojis
- **Beispiele**: Immer praktische Beispiele einschließen
- **Screenshots**: Wenn nützlich für UI

---

## 🔐 Sicherheit

### Schwachstellen melden

**KEINE** öffentlichen Issues für Sicherheitsschwachstellen öffnen. Kontaktieren Sie die Maintainer privat.

### Sicherheits-Checkliste

- [ ] Keine Zugangsdaten, Token, Passwörter committen
- [ ] Keine Endpoints ohne Authentifizierung exponieren (falls hinzugefügt)
- [ ] Benutzereingaben validieren
- [ ] Keine unsanitisierten Shell-Befehle ausführen
- [ ] Hardware-Implikationen berücksichtigen (GPIO, Seriell)

---

## 📄 Lizenz

Indem Sie zu PandaOS beitragen, akzeptieren Sie, dass Ihr Beitrag unter der Lizenz [GNU General Public License v3.0 or later](../LICENSE) veröffentlicht wird.

### Lizenz-Header

Fügen Sie diesen Header zu jeder neuen Quelldatei hinzu:

```javascript
/*
 * PandaOS
 * Copyright (C) 2025  Cyberpandino
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU General Public License for more details.
 */
```

---

## 🎓 Nützliche Ressourcen

### Projekt-Dokumentation
- [README.md](../README.md) - Hauptdokumentation
- [QUICK_START.md](../QUICK_START.md) - Schnellanleitung
- [ARCHITETTURA.md](../ARCHITETTURA.md) - Technische Architektur
- [DOCUMENTAZIONE.md](../DOCUMENTAZIONE.md) - Dokumentationsindex

### Entwicklungs-Setup
```bash
# Repository klonen
git clone https://github.com/cyberpandino/cluster.git
cd cluster

# Abhängigkeiten installieren
npm run install:all

# Im Entwicklungsmodus starten
npm start
```

### Testing
```bash
# Client mit Mock-Modus
cd client
npm run dev

# Server (benötigt Raspberry Pi)
cd server
node server.js
```

---

## 💡 Möchten Sie beitragen, haben aber keine Ideen?

Hier ist eine Liste von Features und Verbesserungen, die wir gerne implementieren würden, aber noch keine Zeit gefunden haben!

### 🚗 Hardware-Features

**Hohe Priorität**:
- [ ] **Integrierte Rückfahrkamera** - Anzeige der Rückfahrkamera im Cluster beim Einlegen des Rückwärtsgangs
- [ ] **Parksensoren** - Grafische Darstellung der Hindernisabstände mit Ultraschallsensoren
- [ ] **3D-Tür-Animation** - Darstellen von offenen/geschlossenen Türen am 3D-Modell des Panda im Cluster
- [ ] **Lichter am 3D-Modell** - Eingeschaltete Lichter (Fernlicht, Blinker, etc.) direkt am 3D-Modell anzeigen

**Mittlere Priorität**:
- [ ] **Regensensor** - Automatische Scheibenwischerregelung
- [ ] **Helligkeitssensor** - Auto-Anpassung der Display-Helligkeit
- [ ] **Reifendrucksensor (TPMS)** - Integration von Reifendrucksensoren
- [ ] **CAN-Bus-Unterstützung** - Neben OBD-II, Unterstützung für natives CAN-Protokoll
- [ ] **360°-Kamera** - Multi-Kamera-System für vollständige Sicht

### 💻 Software-Features

**Hohe Priorität**:
- [ ] **Bordcomputer-System** - Fahrten-Logging mit Verbrauch, Distanz, Zeit
- [ ] **Anpassbare Dashboards** - Mehrere vom Benutzer wählbare Layouts
- [ ] **Farbthemen** - Dark Mode, Light Mode, benutzerdefinierte Themen
- [ ] **Assistierte Kalibrierung** - Assistent zum Kalibrieren von Kraftstoff-/Temperatursensoren
- [ ] **Mobile Companion-App** - Fahrzeugstatistiken auf dem Smartphone

**Mittlere Priorität**:
- [ ] **Wartungsplanung** - Warnungen für Service, Ölwechsel, Inspektion
- [ ] **Wetterintegration** - Außentemperatur von Wetter-API, falls kein Sensor verfügbar
- [ ] **Automatischer Nacht-/Tagmodus** - Basierend auf Uhrzeit oder Helligkeitssensor

### 📚 Dokumentation

**Hohe Priorität**:
- [ ] **Fotografisches Verdrahtungs-Tutorial** - Schritt-für-Schritt-Anleitung mit echten Fotos der Optokoppler-Verdrahtung
- [ ] **Video-Installationsanleitung** - Vollständiges Video-Tutorial von der Verdrahtung zur Software
- [ ] **Internationalisierung (i18n)** - Übersetzungen EN, ES, DE, FR
- [ ] **Zentralisierte Übersetzungsdatei** - Alle Texte in JSON/i18n-Dateien verschieben
- [ ] **Benutzerdefiniertes PCB-Schema** - Professionelles PCB-Design für Optokoppler (KiCad/Eagle)

**Mittlere Priorität**:
- [ ] **Erweiterte FAQ** - Häufig gestellte Fragen mit detaillierter Fehlerbehebung
- [ ] **Installations-Fallstudien** - Echte Beispiele abgeschlossener Installationen
- [ ] **Kompatibilitätsanleitungen** - Liste kompatibler Fahrzeuge außer Panda 141
- [ ] **Interaktiver Schaltplan** - Online navigierbarer elektrischer Schaltplan

### 🧪 Testing & Qualität

- [ ] **Unit-Tests** - Automatisierte Tests für Backend-Services
- [ ] **E2E-Tests** - Vollständige Interface-Tests mit Playwright/Cypress
- [ ] **Performance-Profiling** - Rendering- und Speicheroptimierung

### 🔧 Kompatibilität & Erweiterungen

- [ ] **Unterstützung für andere Fahrzeuge** - Uno, Tipo, Punto, Seicento...

---

### 🚀 Wie man anfängt

1. **Wählen Sie eine Aufgabe** aus der obigen Liste, die Sie interessiert
2. **Öffnen Sie ein Issue** mit [Feature Request](https://github.com/cyberpandino/cluster/issues/new?template=feature_request.md)
3. **Diskutieren Sie die Implementierung** mit den Maintainern
4. **Fork und entwickeln** Sie gemäß diesem Leitfaden
5. **Öffnen Sie eine PR**, wenn bereit

Auch teilweise Implementierungen sind willkommen! Sie müssen nicht das gesamte Feature auf einmal abschließen.

---

## ❓ Fragen?

- **Issue**: Öffnen Sie ein [question issue](https://github.com/cyberpandino/cluster/issues/new?template=question.md)
- **Diskussionen**: Nehmen Sie an [GitHub Discussions](https://github.com/cyberpandino/cluster/discussions) teil (falls aktiviert)

---

## 🙏 Danksagungen

Vielen Dank an alle Mitwirkenden, die helfen, PandaOS besser zu machen! 🚗💨

Alle Mitwirkenden werden in der Datei [AUTHORS](../AUTHORS.md) gewürdigt.

---

**Für Informationen über Autoren und Maintainer siehe [AUTHORS](../AUTHORS.md)**

**Zuletzt aktualisiert**: November 2025
