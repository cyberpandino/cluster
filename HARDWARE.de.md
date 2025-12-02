[🇬🇧 English](HARDWARE.en.md) | [🇮🇹 Italiano](HARDWARE.md) | [🇩🇪 Deutsch](HARDWARE.de.md)

---

# 🛒 Erforderliche Hardware - PandaOS Cluster

Vollständige Liste der für das PandaOS-Projekt erforderlichen Hardwarekomponenten.

> ⚠️ **ACHTUNG**: Diese Liste wird nur zu Informationszwecken bereitgestellt. Kaufen und installieren Sie Hardware nur, wenn Sie über angemessene technische Fähigkeiten verfügen. Siehe [Haftungsausschluss](README.de.md#%EF%B8%8F-haftungsausschluss).

---

## 📋 Inhaltsverzeichnis

1. [Wesentliche Komponenten](#-wesentliche-komponenten)
2. [Optionale Sensoren](#-optionale-sensoren)
3. [Display](#-display)
4. [Zubehör und Verkabelung](#-zubehör-und-verkabelung)
5. [Technische Hinweise](#-technische-hinweise)

---

## 🔧 Wesentliche Komponenten

### 1. Single Board Computer

**Raspberry Pi 4 Model B** (empfohlen) oder **Raspberry Pi 5**

- **RAM**: Minimum 4GB (8GB empfohlen für optimale Leistung)
- **Speicher**: MicroSD 32GB minimum (Klasse 10 oder höher)
- **Erforderliche Schnittstellen**:
  - GPIO 40-polig
  - USB 2.0/3.0 (für ELM327)
  - HDMI (für Display)
  - Ethernet/WiFi (für Entwicklung)

**Warum Raspberry Pi 4B/5**:
- Vollständige GPIO für Optokoppler
- Angemessene Leistung für React + Electron
- Native I2C- und 1-Wire-Unterstützung
- Große Community und umfangreiche Dokumentation

**Raspberry Pi 5 vs 4B**:
- ✅ **Deutlich bessere Leistung**, besonders für 3D-Animationen des Panda-Modells
- ✅ Flüssigeres und reaktionsschnelleres Interface-Rendering
- ⚠️ **Komplizierte Stromversorgung**: Benötigt **5 Ampere** statt 3A des 4B
- ⚠️ Leistungsfähigeres Netzteil und robustere Verkabelung erforderlich
- 💰 Höhere Kosten

**Empfehlung**: Der **Raspberry Pi 4B mit 4GB** ist der beste Kompromiss zwischen Leistung und einfacher Installation. Der Pi 5 wird nur empfohlen, wenn Sie ultra-flüssige 3D-Animationen wünschen und über ein geeignetes Netzteil verfügen.

---

### 2. OBD-II-Adapter

**ELM327 USB**

- **Typ**: USB (nicht Bluetooth/WiFi)
- **Chipsatz**: ELM327 v1.5 oder höher
- **Unterstützte Protokolle**: 
  - ISO 9141-2 (K-Line)
  - ISO 14230-4 (KWP2000)
- **Anschluss**: USB Type-A
- **Kompatibilität**: Magneti Marelli IAW 4AF (Fiat Panda 141)

**Technische Spezifikationen**:
- Baudrate: 38400 bps
- Spannung: 12V vom Fahrzeug
- Schnittstelle: Serial USB (/dev/ttyUSB0)

---

### 3. Stromversorgung

**Display-Netzteil**

- **Eingangsspannung**: 12V DC (vom Fahrzeug)
- **Ausgangsspannung**: 12V DC
- **Strom**: Minimum 2A (abhängig vom Display)
- **Anschluss**: DC-Buchse oder direkte Kabel
- **Schutz**: Überspannung, Kurzschluss

**Raspberry Pi Netzteil** (optional bei USB-Stromversorgung)

- **Spannung**: 5V DC
- **Strom**: Minimum 3A (Raspberry Pi 4B/5)
- **Anschluss**: USB-C
- **Typ**: Step-down DC-DC-Wandler 12V→5V für Fahrzeug

---

## 📺 Display

### LCD-Panel

**Im Projekt verwendete Display-Spezifikationen**:

- **Größe**: 12,6 Zoll
- **Auflösung**: 1920×480 Pixel
- **Seitenverhältnis**: 4:1 (ultra-breit)
- **Panel-Typ**: IPS
- **Bildwiederholrate**: 60Hz
- **Schnittstelle**: HDMI
- **Krümmung**: Flach (nicht gekrümmt)
- **Helligkeit**: Geeignet für Automotive-Einsatz
- **Spannung**: 12V DC

**Wichtige Eigenschaften**:
- Ultra-breites Format ideal für Automotive-Dashboard
- IPS für weite Betrachtungswinkel
- Auflösung optimiert für Instrumentierung (1920×480)

**Alternativen**:
- Displays 10-14 Zoll mit Auflösung 1920×480 oder ähnlich
- Automotive-zertifizierte Panels für erweiterten Temperaturbereich
- Touchscreen optional (nicht im Basisprojekt verwendet)

---

## 🔌 Optionale Sensoren

### 1. Außentemperatursensor

**DS18B20 - Digitaler 1-Wire-Temperatursensor**

- **Typ**: Versiegelter digitaler Sensor (wasserdicht)
- **Temperaturbereich**: -55°C bis +125°C
- **Genauigkeit**: ±0,5°C
- **Protokoll**: 1-Wire (Dallas)
- **Stromversorgung**: 3,0V - 5,5V
- **Gehäuse**: TO-92 oder wasserdichte Version mit Kabel
- **Kabel**: 1-2 Meter (so lang wie möglich für Außeninstallation)

**Zusätzliche Komponenten**:
- Pull-up-Widerstand 4,7kΩ (essentiell)

---

### 2. Kraftstoffstandssensor

**ADS1115 - 16-Bit-I2C-ADC**

- **Typ**: Analog-Digital-Wandler
- **Auflösung**: 16 Bit
- **Kanäle**: 4 Single-Ended oder 2 Differenziell
- **Spannungsbereich**: ±0,256V bis ±6,144V (programmierbar)
- **Schnittstelle**: I2C
- **Adresse**: 0x48 (Standard)
- **Sample-Rate**: 8-860 SPS
- **Stromversorgung**: 2,0V - 5,5V

**Zusätzliche Komponenten**:
- 2× Widerstände für Spannungsteiler (z.B. 100kΩ + 33kΩ)
- Kabel für Verbindung zum Original-Kraftstoffsensor

---

### 3. Optokoppler für Warnleuchten

**PC817 oder gleichwertig**

- **Menge**: 13 Einheiten (eine für jede zu erfassende Warnleuchte)
- **Typ**: Transistor-Fotokoppler
- **Isolationsspannung**: 5000V
- **LED-Strom**: 20mA (typisch)
- **Ausgangsspannung**: 5V kompatibel mit Raspberry GPIO
- **Gehäuse**: DIP-4

**Zusätzliche Komponenten**:
- 13× LED-Vorwiderstände (1kΩ - 2,2kΩ)
- Breadboard oder individuelles PCB für Montage

> 📘 **Offizielles Schaltbild**: Zur Identifizierung der korrekten Kabel für die Warnleuchten am Fahrzeug, konsultieren Sie das [Schaltbild Fiat Panda 141 - Werkstatthandbuch](http://www.bunkeringegnere.altervista.org/esplosi/FIAT%20PANDA/panda%20141/1100%20mpi/55%20IMPIANTO%20ELETTRICO%20-%20SCHEMI%20-%20GAMMA%202000.pdf). Enthält alle Schaltpläne der elektrischen Anlage Gamma 2000 mit Kabel-Farbcodes.

---

## 🔗 Zubehör und Verkabelung

### Kabel und Stecker

- **HDMI-Kabel**: Für Verbindung Raspberry Pi → Display
- **USB Type-A-Kabel**: Für ELM327 → Raspberry Pi
- **Dupont-Kabel**: Male-Female für GPIO (40-teiliges Set)
- **Ethernet-Kabel**: Für Erstkonfiguration (optional)

### Elektrisches Material

- **Automotive-Kabel**: 0,5-1,0 mm² für 12V-Verbindungen
- **Faston-Stecker**: Für Verbindungen zu Fahrzeug-Warnleuchten
- **Schrumpfschlauch**: Verbindungsschutz
- **Kabelbinder**: Kabelorganisation

### Schutzvorrichtungen

**⚡ GRUNDLEGEND: ALLES mit separaten Sicherungen schützen**

- **Raspberry Pi Sicherung**:
  - Wert: **2A** für Raspberry Pi 4B (3A für Pi 5)
  - Typ: Automotive Blade-Sicherung oder Inline
  - Position: Am 12V-Kabel vor dem DC-DC-Wandler
  - Grund: Schützt vor Kurzschlüssen im Wandler oder Raspberry

- **Display-Sicherung**:
  - Wert: **3A** (Display-spezifischen Verbrauch prüfen)
  - Typ: Automotive Blade-Sicherung oder Inline
  - Position: Am 12V-Display-Stromkabel
  - Grund: Schützt vor Kurzschlüssen im Display oder der Verkabelung

- **GPIO-Schaltkreise-Sicherung** (optional aber empfohlen):
  - Wert: **1A**
  - Typ: Inline-Sicherung oder auf Breadboard
  - Position: An der gemeinsamen 12V-Leitung zu den Optokopplern
  - Grund: Schützt Optokoppler vor Verkabelungsfehlern

**Warum separate Sicherungen**:
- ✅ Fehlerisolation: Ein Display-Problem schaltet nicht den Raspberry aus
- ✅ Einfachere Diagnose: Sie erkennen sofort, welcher Schaltkreis Probleme hat
- ✅ Gezielter Schutz: Jede Sicherung ist auf ihre Last dimensioniert
- ✅ Sicherheit: Verhindert das Durchbrennen von Komponenten oder Schlimmeres, verursacht Brände

**Erforderliche Komponenten**:
- **Sicherungshalter**: Inline für 12V-Schaltkreise (mindestens 3 Einheiten)
- **Ersatzsicherungen**: Immer ein paar extra haben
- **Dioden**: 1N4001 oder ähnlich für Verpolungsschutz (eine pro 12V-Leitung)

---

## 🧰 Erforderliche Werkzeuge

### Für Hardware-Installation

- Digitales Multimeter
- Lötkolben (für PCB/dauerhafte Verbindungen)
- Abisolierzange / Zange
- Schraubendreher-Set
- Schaltkreistester

### Für Software-Entwicklung

- Entwicklungscomputer (Mac/Windows/Linux)
- USB-MicroSD-Kartenleser
- Ethernet-Kabel (Raspberry-Ersteinrichtung)

---

## 💰 Geschätzte Kosten

| Komponente | Richtwert |
|-----------|-----------|
| Raspberry Pi 4B (4GB) | €50-70 |
| ELM327 USB | €15-30 |
| LCD-Display 12,6" | €80-150 |
| DC-DC-Wandler 12V→5V 3A | €8-15 |
| DS18B20 Wasserdicht | €3-8 |
| ADS1115 | €5-10 |
| PC817 (10er-Set) | €2-5 |
| Sicherungen + Halter (×3) | €5-10 |
| Kabel und Zubehör | €20-40 |
| **GESAMT GESCHÄTZT** | **€188-338** |

> 💡 Die Preise sind Richtwerte und variieren je nach Lieferanten, Verfügbarkeit und Komponentenqualität.

---

## 📦 Empfohlene Kits

Um den Kauf zu vereinfachen, ziehen Sie diese Kits in Betracht:

### Basis-Kit (Nur Software-Entwicklung)
- Raspberry Pi 4B/5
- USB-C-Netzteil
- MicroSD 32GB
- Raspberry Pi Gehäuse
- ELM327 USB

**Für**: Entwicklung und Tests im Mock-Modus

### Komplettes Kit (Produktion)
- Alles aus dem Basis-Kit
- LCD-Display 12,6" 1920×480
- DC-DC-Wandler 12V→5V (3A für Pi 4B, 5A für Pi 5)
- **3× Automotive-Sicherungen** (2A, 3A, 1A)
- **3× Inline-Sicherungshalter**
- PC817-Optokoppler-Set
- Sortierte Widerstände
- Kabel und Stecker
- Schutzdioden (1N4001 oder ähnlich)

**Für**: Vollständige Fahrzeuginstallation

### Sensor-Kit (Optional)
- DS18B20 wasserdicht
- ADS1115
- Widerstände (4,7kΩ, 100kΩ, 33kΩ)

**Für**: Erweiterte Funktionen (Temperatur, Kraftstoff)

---

## 🔍 Technische Hinweise

### Display-Kompatibilität

Das Projekt ist für **Ultra-Wide 1920×480**-Displays optimiert, kann aber angepasst werden an:
- 1280×480 (niedrigere Auflösung)
- 1920×720 (Standard 16:9, erfordert UI-Anpassung)
- 1024×600 (7" Tablet, erfordert Interface-Resizing)

### Raspberry Pi: 4B vs 5

| Merkmal | Raspberry Pi 4B | Raspberry Pi 5 |
|---------|-----------------|----------------|
| CPU | Quad-Core ARM Cortex-A72 1,5GHz | Quad-Core ARM Cortex-A76 2,4GHz |
| RAM | 2/4/8 GB | 4/8 GB |
| GPIO | 40-polig | 40-polig |
| Leistung | Angemessen | Besser |
| Verbrauch | Niedriger | Höher |
| Kosten | Niedriger | Höher |
| **Empfehlung** | ✅ Ausgezeichnetes Preis-Leistungs-Verhältnis | ✅ Für maximale Leistung |

**Empfehlung**: Raspberry Pi 4B mit 4GB ist mehr als ausreichend für das Projekt.

### Stromversorgung vom Fahrzeug

**Empfohlenes Stromversorgungsschema**:

```
12V Fahrzeugbatterie
    │
    ├─→ [SICHERUNG 2A] → DC-DC-Wandler 12V→5V → Raspberry Pi 4B
    │
    ├─→ [SICHERUNG 3A] → LCD-Display 12V
    │
    └─→ [SICHERUNG 1A] → Optokoppler-Schaltkreis 12V
```

**Spezifikationen**:
- Eingang: Auto 12V (Toleranz 9-16V zur Kompensation von Motorschwankungen)
- Raspberry-Ausgang: 5V 3A USB-C (5A für Pi 5)
- Display-Ausgang: 12V 2-3A (Display-Spezifikationen prüfen)

**Empfohlener DC-DC-Wandler**:
- Eingang: 9-30V DC
- Ausgang: 5V 3A (oder 5A für Pi 5)
- Typ: Automotive Step-Down Buck Converter
- Schutz: Überspannung, Überstrom, Kurzschluss, Verpolung
- Effizienz: >85%

### Praktische Stromversorgungsinstallation

#### Wo 12V im Panda 141 entnehmen

Sie haben mehrere Optionen für die 12V-Stromversorgung:

**Option 1: Von der Sicherungsdose (EMPFOHLEN)**
- **Position**: Unter der Motorhaube, linke Seite nahe der Batterie
- **Empfohlene Sicherungen zur Ableitung**:
  - **F15 (10A)** - Instrumententafel: Aktiviert sich mit Zündung ON, perfekt zur Systemsynchronisation
  - **F16 (7,5A)** - Zubehör: Immer gespeist, nützlich wenn System immer aktiv sein soll
  - **F17 (15A)** - Zigarettenanzünder: Praktisch für Tests, immer gespeist
- **Vorteile**: Sicherungen bereits vorhanden, einfache Ableitung mit Faston-Steckern, vorhandener Schutz
- **Vorgehensweise**: Verwenden Sie Y-Faston-Stecker zur Ableitung ohne Schneiden von Originalkabeln

**Option 2: Batterie direkt (für permanente Installationen)**
- **Position**: Unter der Motorhaube, Pluspol der Batterie
- **Vorteile**: Immer verfügbare Stromversorgung, keine Interferenz mit vorhandenen Sicherungen
- **Nachteile**: Erfordert dedizierte Inline-Sicherung unmittelbar nach Batteriepol
- **Vorgehensweise**: 
  1. Rotes Kabel direkt an Batterie-Pluspol anschließen
  2. Sofort 5A-Sicherung inline einfügen (max 10cm vom Pol)
  3. Geschütztes Kabel bis zur Armaturentafel führen

**Option 3: Hinter der Instrumententafel (sauberer)**
- **Position**: Hinter der Tafel, Verteilerzentrale
- **Vorteile**: Kürzere Kabel, sauberere Installation, bereits in der Armaturentafel
- **Vorgehensweise**: Konsultieren Sie das [Offizielle Schaltbild](http://www.bunkeringegnere.altervista.org/esplosi/FIAT%20PANDA/panda%20141/1100%20mpi/55%20IMPIANTO%20ELETTRICO%20-%20SCHEMI%20-%20GAMMA%202000.pdf) zur Identifizierung der korrekten Kabel

⚠️ **WICHTIG**: Unabhängig von der gewählten Option, verwenden Sie IMMER separate Sicherungen für jede Komponente!

#### Physische Anordnung der Sicherungen

**Empfohlenes Layout im Fahrzeug**:

```
12V Entnahmepunkt (Batterie oder Sicherungen)
    │
    ├─── Rotes Kabel 1,0mm² (30cm) ──→ [INLINE-HALTER 2A] ──→ DC-DC-Wandler
    │                                                                │
    ├─── Rotes Kabel 1,0mm² (50cm) ──→ [INLINE-HALTER 3A] ─────────┼──→ Display 12V
    │                                                                │
    └─── Rotes Kabel 0,5mm² (20cm) ──→ [INLINE-HALTER 1A] ─────────┼──→ Optokoppler
                                                                     │
                                                                     ▼
                                                               DC-DC-Wandler
                                                                     │
                                                                     ├──→ USB-C Raspberry Pi 5V 3A
```

**Physische Positionierung**:
- **Raspberry-Sicherung**: Nahe dem DC-DC-Wandler (unter Armaturentafel)
- **Display-Sicherung**: Hinter dem Display selbst (erleichtert Austausch)
- **Optokoppler-Sicherung**: Nahe Breadboard/PCB der Optokoppler
- **Masse (GND)**: Mit Öse an Fahrzeugmasse anschließen (Armaturentafel-Rahmen-Schraube)

#### DC-DC-Wandler-Verkabelung

**Installationsschritte**:

1. **Positionierung**: DC-DC-Wandler unter Armaturentafel mit Kabelbindern oder Klettverschluss befestigen
2. **12V-Eingang**:
   - Rotes Kabel von 2A-Sicherung → INPUT+-Klemme Wandler
   - Schwarzes Kabel von Fahrzeugmasse → INPUT--Klemme Wandler
3. **5V-Ausgang**:
   - USB-C-Kabel von OUTPUT+ / OUTPUT- Wandler → Raspberry Pi
   - Oder direkt an USB-C-Pins löten (stabiler)
4. **Spannungsregelung**:
   - Vor Raspberry-Anschluss Ausgang mit Multimeter prüfen
   - Trimmer am Wandler auf exakt **5,0V - 5,2V** einstellen
   - ⚠️ 5,3V NICHT überschreiten (würde Raspberry beschädigen!)
5. **Lasttest**:
   - Ausgeschalteten Raspberry anschließen
   - Spannung unter Last messen (muss 5,0-5,2V bleiben)
   - Fällt unter 4,8V, ist Wandler unterdimensioniert

#### Empfohlener Kabeldurchschnitt

| Komponente | Max. Strom | Länge | Kabeldurchschnitt |
|------------|------------|-------|-------------------|
| Raspberry Pi (12V→5V) | 2A @ 12V | <1m | 1,0 mm² |
| Display 12V | 3A @ 12V | <1m | 1,0 mm² |
| Optokoppler | 0,5A @ 12V | <0,5m | 0,5 mm² |
| Masse (gemeinsamer GND) | 5A gesamt | <1m | 1,5 mm² |

**Hinweis**: Querschnitte größer als Minimum zur Kompensation von Verlusten und Automotive-Vibrationen.

#### Verbindung zur Zündung ON/OFF

Um Systemstart mit Zündung zu synchronisieren:

**Methode 1: Entnahme von F15 (Instrumententafel)**
- Pro: Schaltet automatisch mit Schlüssel ein/aus
- Pro: Entlädt Batterie nicht bei ausgeschaltetem Fahrzeug
- Contra: Keine Verzögerung, sofortiges Ausschalten bei Schlüsselentnahme

**Methode 2: Batterie Direkt + GPIO Ignition (EMPFOHLEN)**
- Pro: Kontinuierliche Stromversorgung, Software-gesteuerter Shutdown
- Pro: Verzögerter Shutdown-Timer implementierbar
- Pro: Vermeidet SD-Karten-Korruption mit kontrolliertem Shutdown
- Erfordert: GPIO 21 an "Schlüssel eingesteckt"-Signal angeschlossen (siehe CONFIGURAZIONE_SERVER.de.md)

**Vorgehensweise Methode 2**:
1. Raspberry von Batterie direkt speisen (immer ON)
2. GPIO 21 via Optokoppler an "Schlüssel eingesteckt"-Warnleuchte anschließen
3. `low-power.sh`-Script wird bei Schlüsselentnahme ausgeführt
4. Verzögerten Shutdown von 5-10 Minuten implementieren (siehe CONFIGURAZIONE_SERVER.de.md § Ignition)

#### Checkliste vor dem Einschalten

Vor dem ersten Einschalten des Systems:

- [ ] Polarität mit Multimeter prüfen (rot=12V, schwarz=0V)
- [ ] Batteriespannung prüfen: muss 12,0-14,5V sein
- [ ] Massekontinuität (GND) zur Karosserie prüfen
- [ ] Sicherungen korrekt eingesetzt (2A, 3A, 1A)
- [ ] DC-DC-Wandler auf 5,0-5,2V eingestellt (ohne Last)
- [ ] Kabel gut isoliert (Schrumpfschlauch)
- [ ] Verbindungen fest (nicht fliegend)
- [ ] Raspberry Pi NOCH NICHT angeschlossen (Wandler zuerst im Leerlauf testen)

#### Fehlerbehebung Stromversorgung

**Raspberry startet nicht**:
1. Wandler-OUTPUT-Spannung mit Multimeter prüfen (muss 5,0-5,2V sein)
2. Raspberry-Power-LED prüfen (rot durchgehend = gespeist)
3. 2A-Sicherung nicht durchgebrannt prüfen
4. USB-C-Verbindung gut eingesteckt prüfen

**Raspberry startet zufällig neu**:
- Ursache: Instabile Spannung oder Lastspitzen
- Lösung: Elektrolytkondensator 1000µF 16V am Wandlerausgang verwenden
- Oder: Wandler unterdimensioniert, auf 5A-Modell wechseln

**Display schaltet nicht ein**:
1. 12V-Stromversorgung an Display-Pins prüfen
2. 3A-Sicherung prüfen
3. HDMI-Kabel an Raspberry angeschlossen prüfen

**System entlädt Batterie**:
- Gesamtstrom bei ausgeschalteter Zündung prüfen: <50mA ist akzeptabel
- Wenn >200mA: Leckageproblem, Verkabelung prüfen
- Manuellen Schalter verwenden oder von F15 entnehmen (schaltet mit Schlüssel aus)

⚠️ **WICHTIG**: Sicherungen immer VOR jeder Komponente verwenden, nicht danach!

### GPIO-Schutz

⚠️ **WICHTIG**: Die GPIO des Raspberry Pi sind 3,3V und **NICHT** 5V-tolerant. Die Optokoppler dienen genau dazu, die GPIO vor den 12V des Fahrzeugs zu schützen.

**Schutzschema**:
```
Fahrzeug 12V → Widerstand → LED Optokoppler
                              ↓ (optische Isolierung)
Optokoppler-Transistor → GPIO (3,3V)
```

---

## 📚 Referenzen

### Komponenten-Datenblätter

- **Raspberry Pi 4B**: [raspberrypi.com/products/raspberry-pi-4-model-b](https://www.raspberrypi.com/products/raspberry-pi-4-model-b/)
- **DS18B20**: Maxim Integrated DS18B20 Datenblatt
- **ADS1115**: Texas Instruments ADS1115 Datenblatt
- **PC817**: Sharp PC817 Datenblatt
- **ELM327**: ELM Electronics ELM327 Datenblatt

### Fahrzeug-Schaltbild

- **Fiat Panda 141 - Elektrisches System Gamma 2000**: [Werkstatthandbuch - Schaltpläne](http://www.bunkeringegnere.altervista.org/esplosi/FIAT%20PANDA/panda%20141/1100%20mpi/55%20IMPIANTO%20ELETTRICO%20-%20SCHEMI%20-%20GAMMA%202000.pdf)
  - Vollständiger Schaltplan elektrisches System 1100 MPI
  - Kabel-Farbcodes (z.B. R=Rot, N=Schwarz, BN=Weiß-Schwarz, etc.)
  - Position Warnleuchten und Stecker an Instrumententafel
  - Identifikation Verteilerzentrale und Sicherungen
  - **Essentiell zur Identifizierung der korrekten Kabel für Optokoppler-Anschluss**

### Konfigurationsleitfäden

- [Raspberry Pi Konfiguration](server/CONFIGURAZIONE_SERVER.de.md#-raspberry-pi-konfiguration)
- [GPIO-Setup](server/CONFIGURAZIONE_SERVER.de.md#-gpio-konfiguration)
- [Sensoren](server/CONFIGURAZIONE_SERVER.de.md#-ds18b20-temperatursensor)

---

## ⚠️ Hardware-Haftungsausschluss

Kauf und Installation von Hardware erfolgen in Ihrer vollständigen Verantwortung. Stellen Sie sicher, dass Sie:

- ✅ Über angemessene elektrische/elektronische Fähigkeiten verfügen
- ✅ **IMMER separate Sicherungen** für jede gespeiste Komponente verwenden
- ✅ Kabel mit angemessenem Querschnitt für den zu transportierenden Strom verwenden
- ✅ Lokale Vorschriften zur Automotive-Sicherheit einhalten
- ✅ Fahrzeugsicherheit nicht beeinträchtigen
- ✅ Alles auf der Werkbank testen vor Installation
- ✅ Polarität vor Stromversorgungsanschluss prüfen
- ✅ Schutzvorrichtungen (Sicherungen, Dioden) NIEMALS umgehen

Siehe vollständigen [Haftungsausschluss](README.de.md#%EF%B8%8F-haftungsausschluss) für weitere Informationen.

---

**Letzte Aktualisierung**: November 2025  
**Getestete Hardware**: Fiat Panda 141 1100 mpi
