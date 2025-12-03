[🇬🇧 English](CONFIGURAZIONE_SERVER.en.md) | [🇮🇹 Italiano](CONFIGURAZIONE_SERVER.md) | [🇩🇪 Deutsch](CONFIGURAZIONE_SERVER.de.md)

---

# 🔧 Server-Konfiguration - PandaOS Cluster

Vollständiger Leitfaden zur Hardware- und Software-Konfiguration des Backend-Servers.

> ⚠️ **WARNUNG**: Dieser Leitfaden beschreibt elektrische Verbindungen und Hardware-Modifikationen. Jeder Eingriff in Fahrzeugelektrik birgt Risiken. Lesen Sie den [vollständigen Haftungsausschluss](../README.md#️-disclaimer) und fahren Sie nur fort, wenn Sie wissen, was Sie tun. Die Autoren übernehmen keine Verantwortung für Schäden, die aus der Nutzung dieser Informationen entstehen.

---

## 📋 Inhaltsverzeichnis

> 💡 **Müssen Sie Komponenten kaufen?** Konsultieren Sie zuerst [HARDWARE.md](../HARDWARE.md) für die vollständige Liste aller benötigten Teile.

1. [Hardware-Anforderungen](#-hardware-anforderungen)
2. [Raspberry Pi-Konfiguration](#-raspberry-pi-konfiguration)
3. [OBD-II-Serielle-Schnittstelle](#-obd-ii-serielle-schnittstelle)
4. [GPIO-Konfiguration](#-gpio-konfiguration)
5. [DS18B20-Temperatursensor](#-ds18b20-temperatursensor)
6. [ADS1115-Kraftstoffsensor](#-ads1115-kraftstoffsensor)
7. [Zündungsverwaltung](#-zündungsverwaltung)
8. [Fehlerbehebung](#-fehlerbehebung)

---

## 🛠️ Hardware-Anforderungen

### Wesentliche Komponenten

| Komponente | Modell | Zweck |
|------------|--------|-------|
| **SBC** | Raspberry Pi 4B (4GB+) | Hauptverarbeitung |
| **OBD-Adapter** | ELM327 USB | ECU-Kommunikation |
| **Optokoppler** | PC817 oder ähnlich | Elektrische Isolation für Warnleuchten |
| **Netzteil** | 5V 3A USB-C | Raspberry-Stromversorgung |
| **Display** | HDMI 1920x480+ | Cluster-Anzeige |

### Optionale Komponenten

| Komponente | Modell | Zweck |
|------------|--------|-------|
| **Temperatursensor** | DS18B20 | Außentemperatur |
| **ADC** | ADS1115 | Kraftstoffsensor-Auslesung |
| **Widerstände** | 4.7kΩ, 100kΩ, 33kΩ | Pull-up und Teiler |

---

## 🔧 Raspberry Pi-Konfiguration

### 1. Betriebssystem-Installation

#### OS-Auswahl

**TL;DR**: Raspberry Pi OS Lite (64-bit) Debian-basiert, von unnötigen Diensten befreit.

**Empfohlene Distribution**: Raspberry Pi OS Lite (64-bit)
- **Download**: [raspberrypi.com/software](https://www.raspberrypi.com/software/)
- **Version**: Bookworm (Debian 12) oder neuer
- **Architektur**: 64-bit (bessere Performance für Node.js/Electron)

**Warum Lite statt Desktop?**
- ✅ Bootzeit ~30 Sekunden (vs ~60s mit Desktop)
- ✅ Freier RAM: ~200MB (vs ~500MB mit Desktop-Umgebung)
- ✅ Keine unnötigen Hintergrunddienste
- ✅ Electron bietet bereits UI, kein Desktop-Manager nötig
- ❌ Komplexer zu konfigurieren (keine GUI, alles via SSH)

**Getestete Alternativen**:
- **Raspberry Pi OS Desktop**: Funktioniert, aber langsamer Boot (~60s) und RAM-Verschwendung
- **DietPi**: Ausgezeichnet für ultra-schnellen Boot (~15-20s), aber erfordert mehr manuelle Konfiguration

#### Bootzeit: Die Realität

Nach verschiedenen Tests, mit **abgespecktem Raspberry Pi OS Lite** erreichten wir:

- **~30 Sekunden** vollständiger Boot (POST → Login → PandaOS betriebsbereit)
- **~20 Sekunden** wenn Sie nicht-essentielle Dienste deaktivieren (siehe Optimierungen unten)

**Ist das zu viel?** Kommt darauf an:
- ❌ Wenn Sie die Zündung häufig ein-/ausschalten: ja, Warten ist nervig
- ✅ Wenn Sie **immer-an-Standby-Modus** verwenden (was wir benutzt haben): kein Problem.

#### Immer-An-Standby-Ansatz:

**Wie es in unserem Setup funktioniert**:

1. Raspberry Pi **immer mit Strom versorgt** (direkt Batterie)
2. GPIO 21 erkennt "Schlüssel eingesteckt" (siehe § Zündung)
3. Wenn **Zündung aus**:
   - Skript `low-power.sh` schaltet HDMI-Display aus
   - System im Standby: ~0.4W Verbrauch (vernachlässigbar für Autobatterie)
4. Wenn **Zündung an**:
   - Skript `wake.sh` schaltet Display wieder ein
   - System **sofort betriebsbereit** (0 Sekunden Boot!)

**Vorteile**:
- ⚡ Cluster sofort bei Zündung verfügbar
- 🔋 Sehr niedriger Standby-Verbrauch (~30mA @ 12V)
- 🛡️ SD-Karte geschützt (keine abrupten Shutdowns)
- 🕐 Bootzeit wird irrelevant

**Tatsächlich gemessener Verbrauch**:
- **Standby** (Display aus, CPU Leerlauf): ~0.3-0.5W
- **Betrieb** (Display an, OBD-Daten): ~6-8W
- **Batterie-Auswirkung**: Vernachlässigbar (<0.01% Ladung/Tag)

⚠️ **Hinweis**: Wenn Sie das Auto >2 Wochen stehen lassen, erwägen Sie einen manuellen Schalter oder automatisches Herunterfahren nach 7 Tagen Inaktivität.

#### Basis-Installation

```bash
# Raspberry Pi Imager herunterladen
# https://www.raspberrypi.com/software/

# 1. OS wählen: "Raspberry Pi OS Lite (64-bit)"
# 2. Konfigurieren (Zahnrad-Symbol):
#    - Hostname: pandaos
#    - SSH aktivieren
#    - Benutzername/Passwort: pi/ihr-passwort
#    - WiFi (SSID und Passwort)
#    - Locale: de_DE, Zeitzone Europe/Berlin (oder Ihre Zeitzone)
# 3. Auf microSD schreiben
# 4. In Raspberry einlegen und einschalten
```

#### Boot-Zeit-Optimierungen (Fortgeschritten)

> 💡 **Hinweis**: Dieser Abschnitt ist für diejenigen, die den schnellstmöglichen Boot wollen. Wenn Sie **immer-an-Standby** verwenden, können Sie ihn sicher überspringen.

Mit diesen Anpassungen können Sie von 30s auf ~15-20s reduzieren:

**1. Unnötige Dienste deaktivieren**

```bash
# Bluetooth (falls nicht benötigt)
sudo systemctl disable bluetooth.service
sudo systemctl disable hciuart.service

# ModemManager (falls Sie kein USB-Modem haben)
sudo systemctl disable ModemManager.service

# Druckerdienste (nicht im Auto benötigt)
sudo systemctl disable cups.service
sudo systemctl disable cups-browsed.service

# Triggerhappy (nicht benötigt)
sudo systemctl disable triggerhappy.service

# Neustarten und Bootzeit überprüfen
sudo reboot
systemd-analyze  # Zeigt Gesamtzeit
systemd-analyze blame  # Zeigt langsamste Dienste
```

**2. Kernel-Boot optimieren**

`/boot/firmware/cmdline.txt` bearbeiten:

```bash
sudo nano /boot/firmware/cmdline.txt

# Am Ende der Zeile hinzufügen (alles in EINER Zeile):
quiet splash fastboot noatime nodiratime
```

**3. Netzwerk-Warten deaktivieren**

```bash
# Falls statische IP verwendet oder Netzwerk beim Boot nicht benötigt
sudo systemctl disable NetworkManager-wait-online.service
sudo systemctl disable systemd-networkd-wait-online.service
```

**4. Boot-Timeouts reduzieren**

In `/etc/systemd/system.conf`:

```bash
sudo nano /etc/systemd/system.conf

# Auskommentieren und ändern:
DefaultTimeoutStartSec=10s
DefaultTimeoutStopSec=5s
```

**5. OBD-Server vor Desktop starten**

PM2 + systemd für parallelen Start (siehe § PM2-Konfiguration in README.md).

**Erwartete Ergebnisse**:
- **OS-Boot**: ~8-12 Sekunden
- **Dienste-Start**: ~5-8 Sekunden
- **Gesamt**: ~15-20 Sekunden (vs 30s original)

#### Zukünftige Optimierungen (TODO)

> 📝 WIP-Abschnitt - Beiträge willkommen!

Für diejenigen, die mit <10 Sekunden Boot experimentieren möchten:

**Zu testende Ansätze**:
- **Custom Init**: Systemd durch leichteres Init ersetzen (runit, OpenRC)
- **Minimaler Kernel**: Custom Linux-Kernel mit nur notwendigen Treibern kompilieren
- **Read-only Root**: Root-Dateisystem im Read-only-Modus (schneller, stabiler)
- **Optimiertes Initramfs**: Beim Boot geladene Dienste reduzieren

**Roadmap**:
1. Schritt-für-Schritt "Debian minimal"-Prozedur dokumentieren
2. Automatisches Skript für Boot-Optimierungen
3. Vorkonfiguriertes herunterladbares SD-Image

Falls Sie Erfahrung mit eingebettetem Linux haben und beitragen möchten, öffnen Sie ein [Issue](https://github.com/cyberpandino/cluster/issues)!

---

### 2. Node.js und npm-Installation

⚠️ **Wichtig**: `apt install nodejs` installiert eine veraltete Version (v12-14). PandaOS benötigt **Node.js 18+**.

**Empfohlene Methode: NodeSource**

```bash
# Node.js 20 LTS installieren
curl -fsSL https://deb.nodesource.com/setup_20.x | sudo -E bash -
sudo apt install -y nodejs

# Überprüfen
node --version   # v20.x.x
npm --version    # 10.x.x
```

**Alternative: nvm** (falls mehrere Versionen benötigt)

```bash
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.0/install.sh | bash
source ~/.bashrc
nvm install 20 && nvm use 20
```

> 💡 NodeSource ist stabiler mit PM2/systemd (empfohlen für Produktion)

**Git und Build-Tools**

```bash
sudo apt update && sudo apt upgrade -y
sudo apt install -y git build-essential python3
```

`build-essential` ist notwendig zum Kompilieren nativer Module (SerialPort, onoff, i2c-bus).

**Überprüfung**

```bash
node --version    # >= v18.0.0
npm --version     # >= 9.0.0
gcc --version     # Compiler überprüfen
```

---

### 3. Hardware-Schnittstellen-Konfiguration

```bash
sudo raspi-config
```

Aktivieren:
- **Interface Options** → **I2C** → Yes (für ADS1115)
- **Interface Options** → **Serial Port** → 
  - "Would you like a login shell...?" → **No**
  - "Would you like the serial port hardware...?" → **Yes**
- **Interface Options** → **1-Wire** → Yes (für DS18B20)

Neustart:
```bash
sudo reboot
```

### 4. Benutzerberechtigungen

```bash
# Benutzer zu notwendigen Gruppen hinzufügen
sudo usermod -a -G dialout $USER    # Serielle Schnittstelle
sudo usermod -a -G gpio $USER       # GPIO
sudo usermod -a -G i2c $USER        # I2C

# Abmelden und anmelden zum Anwenden
```

### 5. Konfigurationsüberprüfung

```bash
# I2C überprüfen
ls -l /dev/i2c*
# Erwartete Ausgabe: /dev/i2c-1

# Serial überprüfen
ls -l /dev/ttyUSB* /dev/ttyAMA*
# Erwartete Ausgabe: /dev/ttyUSB0 (mit verbundenem ELM327)

# 1-Wire überprüfen
ls -l /dev/1-wire*
# Oder prüfen: ls /sys/bus/w1/devices/
```

---

## 🔌 OBD-II-Serielle-Schnittstelle

### Hardware-Setup

#### 1. ELM327-Adapter

**Erforderliche Spezifikationen**:
- Protokoll: OBD-II (ISO 9141-2, ISO 14230-4)
- Verbindung: USB (FTDI oder CH340)
- Kompatibilität: ELM327 v1.5 oder höher

**Verbindung**:
1. ELM327 an Magneti Marelli IAW 4AF-Diagnosestecker anschließen
2. ELM327 via USB an Raspberry Pi anschließen
3. LED am ELM327 leuchtet überprüfen

#### 2. Port-Identifikation

```bash
# Verfügbare serielle Ports auflisten
ls -l /dev/ttyUSB*
ls -l /dev/ttyAMA*

# Typische Ausgabe:
# /dev/ttyUSB0 → ELM327 USB
# /dev/ttyAMA0 → GPIO UART (alternativ)

# Detaillierte Info
dmesg | grep tty
```

#### 3. Verbindungstest

```bash
# minicom installieren
sudo apt install minicom

# Zu Port verbinden (38400 Baud)
minicom -D /dev/ttyUSB0 -b 38400

# ELM327-Befehle testen:
ATZ          # Reset (Antwort: ELM327 v1.5)
ATE0         # Echo aus (Antwort: OK)
0100         # PIDs unterstützt (Antwort: Hex-Daten)

# Beenden: CTRL+A dann X
```

### Software-Konfiguration

#### Datei: `services/OBDCommunicationService.js`

```javascript
constructor() {
  this.portPath = '/dev/ttyUSB0';  // ← HIER ÄNDERN falls anders
  this.baudRate = 38400;            // ELM327-Standard
}
```

#### Alternative Ports

Falls ELM327 an anderem Port erkannt wird:

```javascript
// Alternativer USB-Port
this.portPath = '/dev/ttyUSB1';

// GPIO UART (falls direkt verkabelt)
this.portPath = '/dev/ttyAMA0';

// CH340-Adapter (einige Klone)
this.portPath = '/dev/ttyACM0';
```

#### Alternative Baudraten

```javascript
// Standard ELM327
this.baudRate = 38400;

// Einige Adapter anders konfiguriert
this.baudRate = 115200;
this.baudRate = 9600;
```

### OBD-Protokoll

Der Server unterstützt automatisch Fiat Panda 141 / Magneti Marelli IAW 4AF-Protokolle:
- **ISO 9141-2** (K-Line)
- **ISO 14230-4** (KWP2000)

Der Befehl `ATSP0` setzt Auto-Erkennung des Protokolls.

---

## 🔢 GPIO-Konfiguration

### Konfigurationsdatei

**Pfad**: `server/config/gpio-mapping.js`

Diese Datei enthält die gesamte GPIO-Zuordnung für Warnleuchten, Sensoren und Zündung.

### Optokoppler-Schaltplan

```
                  RASPBERRY PI
                  +------------+
    12V Leuchte ->|PC817   GPIO|----> Software-Auslesung
                  |            |
            GND-->|GND         |
                  +------------+
```

**Logik**:
- Leuchte **AN** (12V) → Optokoppler AN → GPIO **HIGH** (1)
- Leuchte **AUS** (0V) → Optokoppler AUS → GPIO **LOW** (0)

### Vollständige Pin-Zuordnung

#### Globale Konfiguration

```javascript
config: {
  mode: 'BCM',              // Broadcom-Nummerierung (nicht physisch)
  pullMode: 'PUD_DOWN',     // Interner Pull-down-Widerstand
  debounceTime: 50,         // Entprellung (ms)
  pollingInterval: 100,     // Lesefrequenz (ms)
}
```

**Parameter-Erklärung**:
- **mode**: `'BCM'` verwendet GPIO-Nummerierung, nicht physische Pin-Nummern
- **pullMode**: `'PUD_DOWN'` stellt 0V sicher wenn Optokoppler offen
- **debounceTime**: Filtert Störsignale (z.B. LED-Flackern)
- **pollingInterval**: Alle 100ms GPIO-Status prüfen

#### GPIO-Pin-Tabelle

| Leuchte/Funktion | GPIO (BCM) | Physischer Pin | Beschreibung |
|------------------|------------|----------------|--------------|
| **Beleuchtung** |
| Fernlicht | 5 | 29 | Fernlichtscheinwerfer |
| Abblendlicht | 6 | 31 | Abblendlichtscheinwerfer |
| Nebelschlussleuchte | 13 | 33 | Nebelschlussleuchte hinten |
| **Fahrtrichtungsanzeiger** |
| Blinker | 17 | 11 | Fahrtrichtungsanzeiger |
| Warnblinker | 12 | 32 | Warnblinkanlage |
| **Motorsystem** |
| Kühlmitteltemp | 16 | 36 | Kühlmitteltemperatur |
| Öldruck | 22 | 15 | Motoröldruck |
| Einspritzung | 24 | 18 | Einspritzsystem |
| **Elektrisches System** |
| Lichtmaschine/Batterie | 27 | 13 | Batterieladung |
| Schlüssel eingesteckt (KEY) | 25 | 22 | Zündung an |
| **Andere Systeme** |
| Bremssystem | 23 | 16 | Bremsen |
| Heckscheibenheizung | 19 | 35 | Defroster |
| Kraftstoffreserve | 20 | 38 | Niedriger Stand |
| **Power-Management** |
| Zündung (Armaturenbrett) | 21 | 40 | Zündung AN/AUS erkennen |

### Raspberry Pi 4B Pinout-Diagramm

```
        3V3  (1) (2)  5V
GPIO  2 SDA  (3) (4)  5V
GPIO  3 SCL  (5) (6)  GND
GPIO  4      (7) (8)  GPIO 14 TXD
        GND  (9) (10) GPIO 15 RXD
GPIO 17     (11) (12) GPIO 18
GPIO 27     (13) (14) GND
GPIO 22     (15) (16) GPIO 23
       3V3  (17) (18) GPIO 24
GPIO 10     (19) (20) GND
GPIO  9     (21) (22) GPIO 25
GPIO 11     (23) (24) GPIO 8
       GND  (25) (26) GPIO 7
GPIO  0     (27) (28) GPIO 1
GPIO  5     (29) (30) GND
GPIO  6     (31) (32) GPIO 12
GPIO 13     (33) (34) GND
GPIO 19     (35) (36) GPIO 16
GPIO 26     (37) (38) GPIO 20
       GND  (39) (40) GPIO 21
```

### Zuordnungs-Beispiel

#### Code in Datei

```javascript
mapping: {
  highBeam: {
    pin: 5,                    // GPIO 5 (Physischer Pin 29)
    name: 'Fernlicht',
    description: 'Fernlichtscheinwerfer',
  },
  
  turnSignals: {
    pin: 17,                   // GPIO 17 (Physischer Pin 11)
    name: 'Blinker',
    description: 'Fahrtrichtungsanzeiger',
  },
  
  // ... andere Leuchten
}
```

#### Zuordnung ändern

Falls Sie Optokoppler anders verkabelt haben:

```javascript
// BEISPIEL: Fernlicht von GPIO 5 zu GPIO 26 verschieben
highBeam: {
  pin: 26,  // ← Nur diese Nummer ändern
  name: 'Fernlicht',
  description: 'Fernlichtscheinwerfer',
}
```

### Optokoppler-Verkabelung

> 📘 **WICHTIG**: Bevor Sie Optokoppler anschließen, konsultieren Sie den [Offiziellen Fiat Panda 141 Schaltplan](http://www.bunkeringegnere.altervista.org/esplosi/FIAT%20PANDA/panda%20141/1100%20mpi/55%20IMPIANTO%20ELETTRICO%20-%20SCHEMI%20-%20GAMMA%202000.pdf), um die korrekten Kabel mit Farbcodes zu identifizieren (z.B. R=Rot, BN=Weiß-Schwarz, GV=Gelb-Grün).

#### Einzel-Optokoppler-Schaltplan (PC817)

```
        Fahrzeug (Eingangsseite)      |  Raspberry Pi (Ausgangsseite)
                                      |
    +12V (von Leuchte) --[R 1kΩ]--+  |
                                   |  |
                            LED+ (1)  |
                                      |
                            LED- (2)--|--- Gemeinsames GND
                                      |
                                 (3)--|--- GPIO (z.B. GPIO 5)
                                      |
                                 (4)--|--- GND
```

**Komponenten**:
- **R**: LED-Strombegrenzungswiderstand (1kΩ - 2.2kΩ)
- **PC817**: Standard-Optokoppler
- **Pin 1-2**: Interne LED (Fahrzeugseite)
- **Pin 3-4**: Ausgangstransistor (Raspberry-Seite)

**Wie Kabel identifizieren**:
1. Offiziellen Schaltplan konsultieren (Link oben)
2. Instrumententafel finden (Seite "Instrumententafel-Anschluss")
3. Gewünschte Leuchte identifizieren (z.B. Fernlicht, Blinker, etc.)
4. Kabel-Farbcode notieren (z.B. "BN" = Weiß-Schwarz)
5. Mit Multimeter 12V-Präsenz prüfen wenn Leuchte an ist

#### Vollständiger Multi-Leuchten-Schaltkreis

```
Leuchte 1 (12V) --[1kΩ]--+
                          |
                      PC817 #1 -----> GPIO 17
                          |
                         GND

Leuchte 2 (12V) --[1kΩ]--+
                          |
                      PC817 #2 -----> GPIO 27
                          |
                         GND

... (für jede Leuchte wiederholen)
```

**Hinweise**:
- Gemeinsames GND für alle Optokoppler verwenden
- Jede Leuchte hat ihren eigenen dedizierten Optokoppler
- Widerstände in Reihe zum Schutz der internen LED

### GPIO-Tests

#### Manueller CLI-Test

```bash
# wiringpi installieren (optional)
sudo apt install wiringpi

# GPIO 17-Status lesen
gpio -g read 17

# Ausgabe:
# 0 = LOW (Leuchte aus)
# 1 = HIGH (Leuchte an)

# Watch-Modus (aktualisiert jede Sekunde)
watch -n 1 'gpio -g read 17'
```

#### Python-Test

```python
#!/usr/bin/env python3
import RPi.GPIO as GPIO
import time

# Setup
GPIO.setmode(GPIO.BCM)
pin = 17  # GPIO 17 (Blinker)
GPIO.setup(pin, GPIO.IN, pull_up_down=GPIO.PUD_DOWN)

try:
    print(f"GPIO {pin} überwachen (CTRL+C zum Beenden)")
    while True:
        state = GPIO.input(pin)
        print(f"GPIO {pin}: {'HIGH (Leuchte an)' if state else 'LOW (Leuchte aus)'}")
        time.sleep(0.5)
except KeyboardInterrupt:
    print("\nUnterbrochen")
finally:
    GPIO.cleanup()
```

Als `test_gpio.py` speichern und ausführen:
```bash
python3 test_gpio.py
```

---

## 🌡️ DS18B20-Temperatursensor

### Hardware-Setup

#### Spezifikationen

- **Typ**: Digitaler Temperatursensor 1-Wire
- **Bereich**: -55°C bis +125°C
- **Genauigkeit**: ±0.5°C
- **Stromversorgung**: 3.0V - 5.5V
- **Protokoll**: Dallas 1-Wire

#### Anschlussdiagramm

```
DS18B20 (TO-92)                Raspberry Pi
                               
Pin 1 (GND)    ---------------- GND (Pin 6, 9, 14, 20, 25, 30, 34, 39)
Pin 2 (DATA)   ----[4.7kΩ]----- 3.3V (Pin 1, 17)
    |                           
    +-------------------------- GPIO 4 (Pin 7)
                               
Pin 3 (VDD)    ---------------- 3.3V (Pin 1, 17)
```

**Komponenten**:
- **Pull-up-Widerstand**: 4.7kΩ zwischen DATA und 3.3V (essentiell!)
- **Kabel**: Empfohlene Maximallänge 3 Meter

#### Verkabelung

1. **GND** (Pin 1) → Raspberry GND verbinden
2. **VDD** (Pin 3) → Raspberry 3.3V verbinden
3. **DATA** (Pin 2) → GPIO 4 (Pin 7) verbinden
4. 4.7kΩ-Widerstand zwischen DATA und 3.3V einfügen

### Software-Setup

#### 1. 1-Wire aktivieren

```bash
# Via raspi-config
sudo raspi-config
# Interface Options → 1-Wire → Yes

# Oder manuell in /boot/config.txt
sudo nano /boot/config.txt

# Hinzufügen (falls nicht vorhanden):
dtoverlay=w1-gpio,gpiopin=4

# Speichern und neu starten
sudo reboot
```

#### 2. Erkennung überprüfen

```bash
# Kernel-Module laden (automatisch nach Neustart)
sudo modprobe w1-gpio
sudo modprobe w1-therm

# Erkannte Sensoren auflisten
ls /sys/bus/w1/devices/

# Erwartete Ausgabe:
# 28-xxxxxxxxxxxx  w1_bus_master1
#
# "28-xxxxxxxxxxxx" ist die DS18B20-Sensor-ID
```

#### 3. Lesetest

```bash
# Temperatur lesen (Sensor-ID ersetzen)
cat /sys/bus/w1/devices/28-xxxxxxxxxxxx/w1_slave

# Ausgabe:
# 7d 01 4b 46 7f ff 0c 10 57 : crc=57 YES
# 7d 01 4b 46 7f ff 0c 10 57 t=23812
#                              ^^^^^^
#                              23.812°C
```

### Software-Konfiguration

#### Datei: `config/gpio-mapping.js`

```javascript
temperature: {
  enabled: true,                    // Sensor aktivieren/deaktivieren
  sensorId: null,                   // null = ersten Sensor auto-erkennen
  basePath: '/sys/bus/w1/devices',
  readInterval: 5000,               // Alle 5 Sekunden lesen
  pin: 4,                           // GPIO 4 (Standard 1-Wire)
}
```

#### Parameter

- **enabled**: `false` deaktiviert Sensor vollständig
- **sensorId**: 
  - `null` → ersten DS18B20 automatisch erkennen
  - `'28-xxxxxxxxxxxx'` → spezifische ID erzwingen (Multi-Sensoren)
- **readInterval**: Lesefrequenz in Millisekunden (min 1000)
- **pin**: GPIO für 1-Wire (Standard 4, änderbar)

#### Mehrere Sensoren

Um mehrere DS18B20 gleichzeitig zu verwenden:

```javascript
temperature: {
  enabled: true,
  sensorId: '28-0123456789ab',  // Zu verwendenden Sensor angeben
  // ... andere Optionen
}
```

Und duplizierten Service für zweiten Sensor erstellen durch Änderung von `TemperatureSensorService.js`.

### Fehlerbehebung

#### Sensor nicht erkannt

```bash
# Geladene Module überprüfen
lsmod | grep w1

# Erwartete Ausgabe:
# w1_therm
# w1_gpio

# Falls fehlend, manuell laden
sudo modprobe w1-gpio
sudo modprobe w1-therm
```

#### Falsche CRC-Lesung

```bash
# Ausgabe mit Fehler:
# xx xx xx xx xx xx xx xx xx : crc=xx NO

# Ursachen:
# 1. Pull-up-Widerstand fehlt oder falscher Wert
# 2. Kabel zu lang (>3m)
# 3. Elektrische Störungen
# 4. Defekter Sensor
```

**Lösungen**:
1. 4.7kΩ-Widerstand vorhanden überprüfen
2. Kabel verkürzen
3. Abgeschirmtes Kabel verwenden
4. Mit anderem Sensor testen

---

## ⛽ ADS1115-Kraftstoffsensor

### Hardware-Setup

#### ADS1115-Spezifikationen

- **Typ**: 16-Bit I2C ADC (Analog-Digital-Wandler)
- **Kanäle**: 4 single-ended oder 2 differenziell
- **Auflösung**: 16 Bit (65536 Stufen)
- **Bereich**: ±0.256V bis ±6.144V (programmierbar)
- **Schnittstelle**: I2C (Standardadresse 0x48)
- **Sample-Rate**: 8-860 SPS

#### Anschlussdiagramm

```
ADS1115-Modul                Raspberry Pi
                             
VDD    --------------------- 3.3V (Pin 1 oder 17)
GND    --------------------- GND (Pin 6, 9, 14, ...)
SCL    --------------------- GPIO 3 / SCL (Pin 5)
SDA    --------------------- GPIO 2 / SDA (Pin 3)
ADDR   --------------------- GND (Adresse 0x48)
A0     --------------------- Kraftstoffsensor (Teiler)
A1     --------------------- Nicht verwendet
A2     --------------------- Nicht verwendet
A3     --------------------- Nicht verwendet
```

#### Spannungsteiler für Kraftstoffsensor

```
         +12V Fahrzeug
              |
              R1 (100kΩ)
              |
              +-------> A0 (ADS1115)
              |
              R2 (33kΩ)
              |
             GND

Ausgangsspannung = Vin × (R2 / (R1 + R2))
                 = 12V × (33kΩ / 133kΩ)
                 = ~3.0V (max)
```

**Warum benötigt**:
- Original-Kraftstoffsensor variiert 0-12V
- ADS1115 akzeptiert max ±4.096V (Verstärkung 4096)
- Teiler reduziert 12V → ~3V

**Individuelle Berechnung**:
```
R2 / (R1 + R2) = Vout_max / Vin_max

Beispiel für Vin_max=12V, Vout_max=3V:
R2 / (R1 + R2) = 3 / 12 = 0.25

Falls R2 = 33kΩ:
33kΩ / (R1 + 33kΩ) = 0.25
R1 = 99kΩ ≈ 100kΩ
```

### Software-Setup

#### 1. I2C aktivieren

```bash
# Via raspi-config
sudo raspi-config
# Interface Options → I2C → Yes

# Oder manuell
sudo nano /boot/config.txt

# Hinzufügen:
dtparam=i2c_arm=on

# Speichern und neu starten
sudo reboot
```

#### 2. Erkennung überprüfen

```bash
# i2c-tools installieren
sudo apt install i2c-tools

# I2C-Bus scannen
sudo i2cdetect -y 1

# Erwartete Ausgabe:
#      0  1  2  3  4  5  6  7  8  9  a  b  c  d  e  f
# 00:          -- -- -- -- -- -- -- -- -- -- -- -- --
# 10: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
# 20: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
# 30: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
# 40: -- -- -- -- -- -- -- -- 48 -- -- -- -- -- -- --
#                              ^^
#                         ADS1115 gefunden!
# 50: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
# 60: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
# 70: -- -- -- -- -- -- -- --
```

#### 3. Lesetest

```bash
# Python-Bibliothek installieren (für Schnelltest)
sudo apt install python3-pip
pip3 install adafruit-circuitpython-ads1x15

# Test-Skript
python3 << EOF
import board
import busio
import adafruit_ads1x15.ads1115 as ADS
from adafruit_ads1x15.analog_in import AnalogIn

i2c = busio.I2C(board.SCL, board.SDA)
ads = ADS.ADS1115(i2c, gain=1)  # gain=1 → ±4.096V
chan = AnalogIn(ads, ADS.P0)    # Kanal A0

print(f"Spannung: {chan.voltage:.3f}V")
print(f"Rohwert: {chan.value}")
EOF
```

### Software-Konfiguration

#### Datei: `config/gpio-mapping.js`

```javascript
fuel: {
  enabled: true,                // Sensor aktivieren/deaktivieren
  chip: 0,                      // 0=ADS1115 (16-bit) | 1=ADS1015 (12-bit)
  channel: 0,                   // ADC-Kanal: 0=A0, 1=A1, 2=A2, 3=A3
  gain: 4096,                   // ±4.096V Vollausschlag
  sampleRate: 250,              // Sample-Rate (SPS)
  readInterval: 500,            // Alle 500ms lesen
  
  // Spannungsteiler
  voltageDivider: {
    r1: 100000,                 // 100kΩ
    r2: 33000,                  // 33kΩ
  },
  
  // Kalibrierung Spannung → Kraftstoffprozent
  calibration: {
    voltageEmpty: 0.5,          // Spannung leerer Tank
    voltageFull: 4.0,           // Spannung voller Tank
  },
  
  pins: {
    sda: 2,                     // GPIO 2 (SDA)
    scl: 3,                     // GPIO 3 (SCL)
  },
}
```

#### Verstärkungsparameter

| gain | Bereich | Auflösung (16-bit) |
|------|---------|-------------------|
| 256 | ±0.256V | 7.8 µV |
| 512 | ±0.512V | 15.6 µV |
| 1024 | ±1.024V | 31.2 µV |
| 2048 | ±2.048V | 62.5 µV |
| 4096 | ±4.096V | 125 µV |
| 6144 | ±6.144V | 187.5 µV |

**Verstärkungswahl**:
- Niedrigste Verstärkung verwenden, die Ihren Bereich enthält
- Für Sensor 0-3V (mit Teiler): `gain: 4096`
- Falls max. Spannung <2V: `gain: 2048` (mehr Auflösung)

#### Kalibrierung

Kalibrierung bildet ADC-Spannung → Kraftstoffprozent (0-100%) ab.

**Prozedur**:

1. **Leerer Tank**: 
   - Zündung bei leerem Tank einschalten
   - Spannung in Debug-Konsole ablesen
   - `voltageEmpty` setzen

2. **Voller Tank**:
   - Tank füllen
   - Zündung einschalten
   - Spannung ablesen
   - `voltageFull` setzen

**Beispiel**:
```javascript
calibration: {
  voltageEmpty: 0.8,   // Gemessen: 0.8V bei leerem Tank
  voltageFull: 3.2,    // Gemessen: 3.2V bei vollem Tank
}
```

System berechnet linear:
```
prozent = ((V_gemessen - V_leer) / (V_voll - V_leer)) × 100
```

#### Sample-Rate

| SPS | Verwendung |
|-----|------------|
| 8 | Maximale Präzision, langsam |
| 16-64 | Ausgewogen |
| 128-250 | Standard (empfohlen) |
| 475-860 | Hohe Geschwindigkeit, weniger präzise |

**Empfehlung**: 250 SPS ist ideal für Kraftstoffsensor (ändert sich langsam).

### Fehlerbehebung

#### ADS1115 nicht erkannt

```bash
# I2C-Modul geladen überprüfen
lsmod | grep i2c

# Gerät vorhanden überprüfen
ls /dev/i2c*
# Erwartete Ausgabe: /dev/i2c-1

# Bus scannen
sudo i2cdetect -y 1
# Falls "48" nicht erscheint, prüfen:
# - VDD- und GND-Stromversorgung
# - SDA/SCL-Verkabelung
# - Lötstellen (falls Custom-Modul)
```

#### Liest immer 0V oder festen Wert

**Ursachen**:
1. A0 nicht verbunden → liest 0V
2. Falscher Kanal in Konfiguration
3. Verstärkung zu niedrig (Signal sättigt)

**Lösungen**:
```javascript
// Korrekten Kanal überprüfen
channel: 0,  // 0=A0, 1=A1, 2=A2, 3=A3

// Höhere Verstärkung probieren
gain: 6144,  // Falls Signal >4.096V
```

#### Falscher Kraftstoffprozentsatz

**Ursache**: Falsche Kalibrierung

**Lösung**: Mit obiger Prozedur neu kalibrieren

---

## ⚡ Zündungsverwaltung

### Zweck

System erkennt, wann Fahrzeugzündung ein-/ausgeschaltet wird und führt automatische Aktionen aus:
- **Zündung aus** → `low-power.sh`-Skript ausführen (Stromsparen)
- **Zündung an** → `wake.sh`-Skript ausführen (Reaktivierung)

### Hardware-Setup

#### Verbindung

Dedizierten Optokoppler an "Schlüssel eingesteckt"-Signal anschließen (12V wenn Zündung an):

```
Zündung 12V (KEY) --[1kΩ]--+
                            |
                         PC817 -----> GPIO 21
                            |
                           GND
```

**Logik**:
- Zündung **an** (12V) → GPIO 21 **HIGH**
- Zündung **aus** (0V) → GPIO 21 **LOW**

Oder umgekehrt bei Verwendung von active-low-Optokoppler (`activeOn: 0` setzen)

### Konfiguration

#### Datei: `config/gpio-mapping.js`

```javascript
ignition: {
  enabled: true,                    // Zündungsverwaltung aktivieren
  pin: 21,                          // Dediziertes GPIO 21
  activeOn: 1,                      // 1=active high | 0=active low
  scripts: {
    lowPower: './scripts/low-power.sh',   // Wenn Zündung AUSSCHALTET
    wake: './scripts/wake.sh',             // Wenn Zündung EINSCHALTET
  },
}
```

#### Parameter

- **enabled**: `false` deaktiviert Funktion vollständig
- **pin**: Dediziertes GPIO (verschieden von denen für Leuchten verwendet)
- **activeOn**:
  - `1` → HIGH = Zündung an (active high Optokoppler)
  - `0` → LOW = Zündung an (active low Optokoppler)
- **scripts**: Pfade relativ zu `server/`-Verzeichnis

### Stromspar-Skripte

#### Datei: `scripts/low-power.sh`

Wird ausgeführt wenn Zündung **ausgeschaltet** wird.

```bash
#!/bin/bash
# Skript ausgeführt wenn Zündung ausschaltet

logger "PandaOS: Zündung aus - Stromsparmodus starten"

# HDMI-Display ausschalten
vcgencmd display_power 0

# Hintergrundbeleuchtung dimmen (falls GPIO-gesteuert)
# echo 0 > /sys/class/backlight/rpi_backlight/brightness

# Nicht-essentielle Dienste stoppen
# systemctl stop bluetooth

# Optional: Nach X Minuten herunterfahren
# sleep 600 && sudo shutdown -h now &

logger "PandaOS: Stromsparmodus aktiviert"
```

Ausführbar machen:
```bash
chmod +x server/scripts/low-power.sh
```

#### Datei: `scripts/wake.sh`

Wird ausgeführt wenn Zündung **eingeschaltet** wird.

```bash
#!/bin/bash
# Skript ausgeführt wenn Zündung einschaltet

logger "PandaOS: Zündung an - System reaktivieren"

# HDMI-Display einschalten
vcgencmd display_power 1

# Hintergrundbeleuchtung wiederherstellen
# echo 255 > /sys/class/backlight/rpi_backlight/brightness

# Dienste neu starten
# systemctl start bluetooth

logger "PandaOS: System reaktiviert"
```

Ausführbar machen:
```bash
chmod +x server/scripts/wake.sh
```

### Skript-Anpassung

Skripte können beliebige Bash-Befehle ausführen, zum Beispiel:

#### Auto-Shutdown nach 10 Minuten

In `low-power.sh`:
```bash
# Herunterfahren falls Zündung 10 Minuten aus bleibt
(sleep 600 && sudo shutdown -h now) &
echo $! > /tmp/pandaos-shutdown.pid
```

In `wake.sh`:
```bash
# Shutdown abbrechen falls Zündung wieder eingeschaltet wird
if [ -f /tmp/pandaos-shutdown.pid ]; then
  kill $(cat /tmp/pandaos-shutdown.pid) 2>/dev/null
  rm /tmp/pandaos-shutdown.pid
fi
```

#### Telegram-Benachrichtigung

```bash
# In low-power.sh
curl -s -X POST "https://api.telegram.org/bot<TOKEN>/sendMessage" \
  -d chat_id=<CHAT_ID> \
  -d text="🚗 Zündung aus - PandaOS im Standby"
```

---

## 🚨 Fehlerbehebung

### Server startet nicht

#### Fehler: "Nicht unterstützte Plattform"

```
❌ FEHLER: Essentielle Raspberry Pi-Abhängigkeiten nicht verfügbar
Nicht unterstützte Plattform: linux x64 - benötigt Linux ARM
```

**Ursache**: Läuft auf Nicht-Raspberry-Pi-System

**Lösung**: Server nur auf Raspberry Pi ausführen, oder Prüfung in `OBDServer.js` deaktivieren (nicht empfohlen)

#### Fehler: "GPIO-Modul nicht verfügbar"

```
GPIO-Modul (onoff) nicht verfügbar
```

**Ursache**: `onoff`-Bibliothek nicht installiert oder inkompatibles System

**Lösung**:
```bash
cd server
npm install onoff
```

### ELM327 antwortet nicht

#### Symptom: "Port /dev/ttyUSB0 nicht gefunden"

**Lösung**:
1. ELM327 verbunden überprüfen: `lsusb`
2. Port überprüfen: `ls -l /dev/ttyUSB*`
3. Berechtigungen prüfen: `groups` (muss `dialout` enthalten)

#### Symptom: "Timeout - Keine Antwort erhalten"

**Ursache**: ELM327 kommuniziert nicht

**Lösung**:
1. Baudrate in `OBDCommunicationService.js` überprüfen
2. Mit minicom testen: `minicom -D /dev/ttyUSB0 -b 38400`
3. Reset versuchen: ELM327 trennen, 10 Sekunden warten, wieder verbinden

### GPIO funktioniert nicht

#### Symptom: Leuchten nicht erkannt

**Diagnose**:
```bash
# Manueller GPIO-Test
gpio -g mode 17 in
gpio -g read 17

# Fahrzeugleuchte ein-/ausschalten und prüfen ob Wert sich ändert
```

**Lösungen**:
1. Optokoppler-Verkabelung überprüfen
2. Mit LED und Widerstand statt Optokoppler testen
3. Pin-Nummer in `gpio-mapping.js` prüfen (BCM vs. physisch)

### Sensoren erkennen nicht

#### DS18B20: "1-Wire nicht gefunden"

```bash
# Modul geladen überprüfen
lsmod | grep w1

# Falls fehlend
sudo modprobe w1-gpio
sudo modprobe w1-therm

# Permanente Konfiguration überprüfen
grep w1-gpio /boot/config.txt
```

#### ADS1115: "Adresse 0x48 antwortet nicht"

```bash
# I2C-Bus scannen
sudo i2cdetect -y 1

# Falls "48" nicht erscheint:
# - ADS1115-Stromversorgung prüfen
# - SDA/SCL nicht vertauscht überprüfen
# - Mit kürzerem Kabel testen
```

---

## 📚 Technische Referenzen

### Datenblätter und Dokumentation

- **Raspberry Pi GPIO**: https://pinout.xyz/
- **DS18B20**: https://datasheets.maximintegrated.com/en/ds/DS18B20.pdf
- **ADS1115**: https://www.ti.com/lit/ds/symlink/ads1115.pdf
- **PC817 Optokoppler**: https://www.farnell.com/datasheets/73758.pdf
- **ELM327**: https://www.elmelectronics.com/wp-content/uploads/2017/01/ELM327DS.pdf
- **Fiat Panda 141 - Offizieller Schaltplan**: http://www.bunkeringegnere.altervista.org/esplosi/FIAT%20PANDA/panda%20141/1100%20mpi/55%20IMPIANTO%20ELETTRICO%20-%20SCHEMI%20-%20GAMMA%202000.pdf

### Nützliche Befehle

```bash
# GPIO
gpio readall                    # Status aller Pins
gpio -g read <pin>              # Spezifischen Pin lesen

# I2C
sudo i2cdetect -y 1             # Bus-Scan
sudo i2cget -y 1 0x48 0x00 w    # ADS1115-Register lesen

# 1-Wire
ls /sys/bus/w1/devices/         # Sensoren auflisten
cat /sys/bus/w1/devices/28-*/w1_slave  # Temperatur lesen

# Serial
ls -l /dev/tty*                 # Ports auflisten
sudo minicom -D /dev/ttyUSB0    # Serieller Monitor

# Logs
journalctl -u obd-server -f     # Server-Log (falls PM2/systemd)
pm2 logs obd-server             # PM2-Log
```

---

**Letzte Aktualisierung**: v0.9.0  
**Ziel-Hardware**: Raspberry Pi 4B + Fiat Panda 141
