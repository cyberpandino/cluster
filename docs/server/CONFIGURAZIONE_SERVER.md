[🇬🇧 English](CONFIGURAZIONE_SERVER.en.md) | [🇮🇹 Italiano](CONFIGURAZIONE_SERVER.md) | [🇩🇪 Deutsch](CONFIGURAZIONE_SERVER.de.md)

---

# 🔧 Configurazione Server - PandaOS Cluster

Guida completa alla configurazione hardware e software del server backend.

> ⚠️ **ATTENZIONE**: Questa guida descrive collegamenti elettrici e modifiche hardware. Ogni intervento su impianti elettrici di veicoli comporta rischi. Leggi il [Disclaimer completo](../README.md#️-disclaimer) e procedi solo se sai cosa stai facendo. Gli autori non si assumono responsabilità per danni derivanti dall'uso di queste informazioni.

---

## 📋 Indice

> 💡 **Devi acquistare componenti?** Consulta prima [HARDWARE.md](../HARDWARE.md) per la lista completa di tutto il necessario.

1. [Requisiti Hardware](#-requisiti-hardware)
2. [Configurazione Raspberry Pi](#-configurazione-raspberry-pi)
3. [Porta Seriale OBD-II](#-porta-seriale-obd-ii)
4. [Configurazione GPIO](#-configurazione-gpio)
5. [Sensore Temperatura DS18B20](#-sensore-temperatura-ds18b20)
6. [Sensore Carburante ADS1115](#-sensore-carburante-ads1115)
7. [Gestione Quadro Accensione](#-gestione-quadro-accensione)
8. [Troubleshooting](#-troubleshooting)

---

## 🛠️ Requisiti Hardware

### Componenti Essenziali

| Componente | Modello | Scopo |
|------------|---------|-------|
| **SBC** | Raspberry Pi 4B (4GB+) | Elaborazione principale |
| **Adattatore OBD** | ELM327 USB | Comunicazione con centralina |
| **Optoaccoppiatori** | PC817 o simili | Isolamento elettrico spie |
| **Alimentatore** | 5V 3A USB-C | Alimentazione Raspberry |
| **Display** | HDMI 1920x480+ | Visualizzazione cluster |

### Componenti Opzionali

| Componente | Modello | Scopo |
|------------|---------|-------|
| **Sensore temperatura** | DS18B20 | Temperatura esterna |
| **ADC** | ADS1115 | Lettura sensore carburante |
| **Resistenze** | 4.7kΩ, 100kΩ, 33kΩ | Pull-up e partitore |

---

## 🔧 Configurazione Raspberry Pi

### 1. Installazione Sistema Operativo

#### Scelta del Sistema Operativo

**TL;DR**: Raspberry Pi OS Lite (64-bit) Debian-based, alleggerito dai servizi superflui.

**Distribuzione Consigliata**: Raspberry Pi OS Lite (64-bit)
- **Download**: [raspberrypi.com/software](https://www.raspberrypi.com/software/)
- **Versione**: Bookworm (Debian 12) o successiva
- **Architettura**: 64-bit (migliori performance per Node.js/Electron)

**Perché Lite e non Desktop?**
- ✅ Boot time ~30 secondi (vs ~60s con Desktop)
- ✅ RAM libera: ~200MB (vs ~500MB con Desktop Environment)
- ✅ Niente servizi inutili in background
- ✅ Electron fornisce già la UI, non serve Desktop Manager
- ❌ Più complesso da configurare (niente GUI, tutto via SSH)

**Alternative Testate**:
- **Raspberry Pi OS Desktop**: Funziona ma boot lento (~60s) e spreco RAM
- **DietPi**: Ottima per boot ultra-rapidi (~15-20s) ma richiede più configurazione manuale

#### Boot Time: La Realtà

Dopo vari test, con **Raspberry Pi OS Lite alleggerito** siamo arrivati a:

- **~30 secondi** di boot completo (POST → Login → PandaOS operativo)
- **~20 secondi** se disabiliti servizi non essenziali (vedi ottimizzazioni sotto)

**È tanto?** Dipende:
- ❌ Se accendi/spegni il quadro spesso: sì, è noioso aspettare
- ✅ Se usi modalità **standby sempre acceso** (quello che abbiamo usato noi): non è un problema.

#### Approccio Standby Sempre Acceso:

**Come funziona nel nostro setup**:

1. Raspberry Pi **sempre alimentato** (batteria diretta)
2. GPIO 21 rileva "chiave inserita" (vedi § Ignition)
3. Quando **spegni quadro**:
   - Script `low-power.sh` spegne display HDMI
   - Sistema in standby: ~0.4W di consumo (trascurabile per batteria auto)
4. Quando **accendi quadro**:
   - Script `wake.sh` riaccende display
   - Sistema **immediatamente operativo** (0 secondi boot!)

**Vantaggi**:
- ⚡ Cluster disponibile istantaneamente all'accensione
- 🔋 Consumo standby bassissimo (~30mA @ 12V)
- 🛡️ SD card protetta (nessun shutdown brusco)
- 🕐 Boot time diventa irrilevante

**Consumo Reale Misurato**:
- **Standby** (display off, CPU idle): ~0.3-0.5W
- **Operativo** (display on, dati OBD): ~6-8W
- **Impatto batteria**: Trascurabile (<0.01% carica/giorno)

⚠️ **Nota**: Se lasci l'auto ferma per >2 settimane, considera interruttore manuale o shutdown automatico dopo 7 giorni di inattività.

#### Installazione Base

```bash
# Scarica Raspberry Pi Imager
# https://www.raspberrypi.com/software/

# 1. Seleziona OS: "Raspberry Pi OS Lite (64-bit)"
# 2. Configura (icona ingranaggio):
#    - Hostname: pandaos
#    - Abilita SSH
#    - Username/Password: pi/tua-password
#    - WiFi (SSID e password)
#    - Locale: it_IT, timezone Europe/Rome
# 3. Scrivi su microSD
# 4. Inserisci nel Raspberry e accendi
```

#### Ottimizzazioni Boot Time (Avanzato)

> 💡 **Nota**: Questa sezione è per chi vuole il boot più veloce possibile. Se usi **standby sempre acceso**, puoi saltarla tranquillamente.

Con questi tweak puoi scendere da 30s a ~15-20s:

**1. Disabilita Servizi Inutili**

```bash
# Bluetooth (se non serve)
sudo systemctl disable bluetooth.service
sudo systemctl disable hciuart.service

# ModemManager (se non hai modem USB)
sudo systemctl disable ModemManager.service

# Servizi stampante (non servono su auto)
sudo systemctl disable cups.service
sudo systemctl disable cups-browsed.service

# Triggerhappy (non serve)
sudo systemctl disable triggerhappy.service

# Riavvia e verifica tempo boot
sudo reboot
systemd-analyze  # Mostra tempo totale
systemd-analyze blame  # Mostra servizi più lenti
```

**2. Ottimizza Boot Kernel**

Modifica `/boot/firmware/cmdline.txt`:

```bash
sudo nano /boot/firmware/cmdline.txt

# Aggiungi alla fine della riga (tutto su UNA riga):
quiet splash fastboot noatime nodiratime
```

**3. Disabilita Attesa Rete**

```bash
# Se usi IP statico o non serve rete all'avvio
sudo systemctl disable NetworkManager-wait-online.service
sudo systemctl disable systemd-networkd-wait-online.service
```

**4. Riduci Timeout Boot**

In `/etc/systemd/system.conf`:

```bash
sudo nano /etc/systemd/system.conf

# Decommentare e modificare:
DefaultTimeoutStartSec=10s
DefaultTimeoutStopSec=5s
```

**5. Avvia Server OBD Prima del Desktop**

PM2 + systemd per avvio in parallelo (vedi § Configurazione PM2 nel README.md).

**Risultati Attesi**:
- **Boot OS**: ~8-12 secondi
- **Avvio servizi**: ~5-8 secondi
- **Totale**: ~15-20 secondi (vs 30s originali)

#### Future Ottimizzazioni (TODO)

> 📝 Sezione WIP - Contributi benvenuti!

Per chi volesse sperimentare boot <10 secondi:

**Approcci da Testare**:
- **Init custom**: Sostituire systemd con init più leggero (runit, OpenRC)
- **Kernel minimale**: Compilare kernel Linux custom con solo driver necessari
- **Read-only root**: Root filesystem in sola lettura (più veloce, più stabile)
- **Initramfs ottimizzato**: Ridurre servizi caricati all'avvio

**Roadmap**:
1. Documentare procedura "Debian minimal" passo-passo
2. Script automatico per applicare ottimizzazioni boot
3. Immagine SD pre-configurata scaricabile

Se hai esperienza con embedded Linux e vuoi contribuire, apri una [issue](https://github.com/cyberpandino/cluster/issues)!

---

### 2. Installazione Node.js e npm

⚠️ **Importante**: `apt install nodejs` installa una versione obsoleta (v12-14). PandaOS richiede **Node.js 18+**.

**Metodo Consigliato: NodeSource**

```bash
# Installa Node.js 20 LTS
curl -fsSL https://deb.nodesource.com/setup_20.x | sudo -E bash -
sudo apt install -y nodejs

# Verifica
node --version   # v20.x.x
npm --version    # 10.x.x
```

**Alternativa: nvm** (se serve gestire più versioni)

```bash
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.0/install.sh | bash
source ~/.bashrc
nvm install 20 && nvm use 20
```

> 💡 NodeSource è più stabile con PM2/systemd (consigliato per produzione)

**Git e Build Tools**

```bash
sudo apt update && sudo apt upgrade -y
sudo apt install -y git build-essential python3
```

`build-essential` è necessario per compilare moduli nativi (SerialPort, onoff, i2c-bus).

**Verifica**

```bash
node --version    # >= v18.0.0
npm --version     # >= 9.0.0
gcc --version     # Verifica compiler
```

---

### 3. Configurazione Interfacce Hardware

```bash
sudo raspi-config
```

Abilita:
- **Interface Options** → **I2C** → Yes (per ADS1115)
- **Interface Options** → **Serial Port** → 
  - "Would you like a login shell...?" → **No**
  - "Would you like the serial port hardware...?" → **Yes**
- **Interface Options** → **1-Wire** → Yes (per DS18B20)

Riavvia:
```bash
sudo reboot
```

### 4. Permessi Utente

```bash
# Aggiungi utente ai gruppi necessari
sudo usermod -a -G dialout $USER    # Porta seriale
sudo usermod -a -G gpio $USER       # GPIO
sudo usermod -a -G i2c $USER        # I2C

# Logout e login per applicare
```

### 5. Verifica Configurazione

```bash
# Verifica I2C
ls -l /dev/i2c*
# Output atteso: /dev/i2c-1

# Verifica Serial
ls -l /dev/ttyUSB* /dev/ttyAMA*
# Output atteso: /dev/ttyUSB0 (con ELM327 collegato)

# Verifica 1-Wire
ls -l /dev/1-wire*
# O controlla: ls /sys/bus/w1/devices/
```

---

## 🔌 Porta Seriale OBD-II

### Hardware Setup

#### 1. Adattatore ELM327

**Specifiche richieste**:
- Protocollo: OBD-II (ISO 9141-2, ISO 14230-4)
- Connessione: USB (FTDI o CH340)
- Compatibilità: ELM327 v1.5 o superiore

**Connessione**:
1. Collegare ELM327 al connettore diagnostico Magneti Marelli IAW 4AF
2. Collegare ELM327 via USB al Raspberry Pi
3. Verificare LED acceso su ELM327

#### 2. Identificazione Porta

```bash
# Lista porte seriali disponibili
ls -l /dev/ttyUSB*
ls -l /dev/ttyAMA*

# Output tipico:
# /dev/ttyUSB0 → ELM327 USB
# /dev/ttyAMA0 → UART GPIO (alternativo)

# Info dettagliate
dmesg | grep tty
```

#### 3. Test Connessione

```bash
# Installa minicom
sudo apt install minicom

# Connetti alla porta (38400 baud)
minicom -D /dev/ttyUSB0 -b 38400

# Testa comandi ELM327:
ATZ          # Reset (risposta: ELM327 v1.5)
ATE0         # Echo off (risposta: OK)
0100         # PID supported (risposta: dati hex)

# Esci: CTRL+A poi X
```

### Configurazione Software

#### File: `services/OBDCommunicationService.js`

```javascript
constructor() {
  this.portPath = '/dev/ttyUSB0';  // ← MODIFICA QUI se diversa
  this.baudRate = 38400;            // Standard ELM327
}
```

#### Porte Alternative

Se l'ELM327 viene riconosciuto su porta diversa:

```javascript
// Porta USB alternativa
this.portPath = '/dev/ttyUSB1';

// UART GPIO (se cablato direttamente)
this.portPath = '/dev/ttyAMA0';

// Adattatore CH340 (alcuni cloni)
this.portPath = '/dev/ttyACM0';
```

#### Baudrate Alternativi

```javascript
// ELM327 standard
this.baudRate = 38400;

// Alcuni adattatori configurati diversamente
this.baudRate = 115200;
this.baudRate = 9600;
```

### Protocollo OBD

Il server supporta automaticamente i protocolli Fiat Panda 141 / Magneti Marelli IAW 4AF:
- **ISO 9141-2** (K-Line)
- **ISO 14230-4** (KWP2000)

Il comando `ATSP0` imposta auto-rilevamento del protocollo.

---

## 🔢 Configurazione GPIO

### File di Configurazione

**Path**: `server/config/gpio-mapping.js`

Questo file contiene tutta la mappatura GPIO per spie, sensori e quadro accensione.

### Schema Optoaccoppiatori

```
                  RASPBERRY PI
                  +------------+
    Spia 12V ---->|PC817   GPIO|----> Lettura software
                  |            |
            GND-->|GND         |
                  +------------+
```

**Logica**:
- Spia **ACCESA** (12V) → Optoaccoppiatore ON → GPIO **HIGH** (1)
- Spia **SPENTA** (0V) → Optoaccoppiatore OFF → GPIO **LOW** (0)

### Mappatura Pin Completa

#### Configurazione Globale

```javascript
config: {
  mode: 'BCM',              // Numerazione Broadcom (non fisica)
  pullMode: 'PUD_DOWN',     // Resistenza pull-down interna
  debounceTime: 50,         // Anti-rimbalzo (ms)
  pollingInterval: 100,     // Frequenza lettura (ms)
}
```

**Spiegazione parametri**:
- **mode**: `'BCM'` usa numerazione GPIO, non numeri fisici pin
- **pullMode**: `'PUD_DOWN'` assicura 0V quando optoaccoppiatore aperto
- **debounceTime**: Filtra segnali spurii (es. flickering LED)
- **pollingInterval**: Ogni 100ms controlla stato GPIO

#### Tabella Pin GPIO

| Spia/Funzione | GPIO (BCM) | Pin Fisico | Descrizione |
|---------------|------------|------------|-------------|
| **Illuminazione** |
| Abbaglianti | 5 | 29 | Fari abbaglianti |
| Anabbaglianti | 6 | 31 | Fari anabbaglianti |
| Fendinebbia | 13 | 33 | Fendinebbia posteriore |
| **Indicatori direzione** |
| Frecce | 17 | 11 | Indicatori direzione |
| Quattro frecce | 12 | 32 | Luci emergenza |
| **Sistema motore** |
| Temperatura raffreddamento | 16 | 36 | Liquido refrigerante |
| Pressione olio | 22 | 15 | Pressione olio motore |
| Iniettori | 24 | 18 | Sistema iniezione |
| **Sistema elettrico** |
| Alternatore/Batteria | 27 | 13 | Carica batteria |
| Chiave inserita (KEY) | 25 | 22 | Quadro acceso |
| **Altri sistemi** |
| Sistema frenante | 23 | 16 | Freni |
| Termoresistenza lunotto | 19 | 35 | Sbrinatore |
| Riserva carburante | 20 | 38 | Livello basso |
| **Power Management** |
| Ignition (quadro) | 21 | 40 | Rilevamento quadro ON/OFF |

### Schema Pinout Raspberry Pi 4B

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

### Esempio Mappatura

#### Codice nel File

```javascript
mapping: {
  highBeam: {
    pin: 5,                    // GPIO 5 (Pin fisico 29)
    name: 'Abbaglianti',
    description: 'Fari abbaglianti',
  },
  
  turnSignals: {
    pin: 17,                   // GPIO 17 (Pin fisico 11)
    name: 'Frecce',
    description: 'Indicatori di direzione',
  },
  
  // ... altre spie
}
```

#### Modificare Mappatura

Se hai cablato diversamente gli optoaccoppiatori:

```javascript
// ESEMPIO: Spostare abbaglianti da GPIO 5 a GPIO 26
highBeam: {
  pin: 26,  // ← Cambia solo questo numero
  name: 'Abbaglianti',
  description: 'Fari abbaglianti',
}
```

### Cablaggio Optoaccoppiatori

> 📘 **IMPORTANTE**: Prima di collegare gli optoaccoppiatori, consulta lo [Schema Elettrico Ufficiale Fiat Panda 141](http://www.bunkeringegnere.altervista.org/esplosi/FIAT%20PANDA/panda%20141/1100%20mpi/55%20IMPIANTO%20ELETTRICO%20-%20SCHEMI%20-%20GAMMA%202000.pdf) per identificare i cavi corretti con i codici colore (es. R=Rosso, BN=Bianco-Nero, GV=Giallo-Verde).

#### Schema Singolo Optoaccoppiatore (PC817)

```
        Veicolo (Lato Input)          |  Raspberry Pi (Lato Output)
                                       |
    +12V (da spia) ----[R 1kΩ]---+    |
                                  |    |
                            LED+ (1)   |
                                       |
                            LED- (2)---|--- GND comune
                                       |
                                  (3)--|--- GPIO (es. GPIO 5)
                                       |
                                  (4)--|--- GND
```

**Componenti**:
- **R**: Resistenza limitatrice LED (1kΩ - 2.2kΩ)
- **PC817**: Optoaccoppiatore standard
- **Pin 1-2**: LED interno (lato veicolo)
- **Pin 3-4**: Transistor di uscita (lato Raspberry)

**Come identificare i cavi**:
1. Consulta lo schema elettrico ufficiale (link sopra)
2. Trova il quadro strumenti (pagina "Collegamento quadro strumenti")
3. Identifica la spia desiderata (es. abbaglianti, frecce, etc.)
4. Annota il codice colore del cavo (es. "BN" = Bianco-Nero)
5. Verifica con multimetro la presenza di 12V quando la spia è accesa

#### Circuito Completo Multi-Spia

```
Spia 1 (12V) --[1kΩ]--+
                       |
                    PC817 #1 -----> GPIO 17
                       |
                      GND

Spia 2 (12V) --[1kΩ]--+
                       |
                    PC817 #2 -----> GPIO 27
                       |
                      GND

... (ripeti per ogni spia)
```

**Note**:
- Usa GND comune per tutti gli optoaccoppiatori
- Ogni spia ha il suo optoaccoppiatore dedicato
- Resistenze in serie per proteggere LED interno

### Test GPIO

#### Test Manuale con CLI

```bash
# Installa wiringpi (opzionale)
sudo apt install wiringpi

# Leggi stato GPIO 17
gpio -g read 17

# Output:
# 0 = LOW (spia spenta)
# 1 = HIGH (spia accesa)

# Modalità watch (aggiorna ogni secondo)
watch -n 1 'gpio -g read 17'
```

#### Test con Python

```python
#!/usr/bin/env python3
import RPi.GPIO as GPIO
import time

# Setup
GPIO.setmode(GPIO.BCM)
pin = 17  # GPIO 17 (frecce)
GPIO.setup(pin, GPIO.IN, pull_up_down=GPIO.PUD_DOWN)

try:
    print(f"Monitoraggio GPIO {pin} (CTRL+C per uscire)")
    while True:
        state = GPIO.input(pin)
        print(f"GPIO {pin}: {'HIGH (spia accesa)' if state else 'LOW (spia spenta)'}")
        time.sleep(0.5)
except KeyboardInterrupt:
    print("\nInterrotto")
finally:
    GPIO.cleanup()
```

Salva come `test_gpio.py` e esegui:
```bash
python3 test_gpio.py
```

---

## 🌡️ Sensore Temperatura DS18B20

### Hardware Setup

#### Specifiche

- **Tipo**: Sensore digitale temperatura 1-Wire
- **Range**: -55°C a +125°C
- **Precisione**: ±0.5°C
- **Alimentazione**: 3.0V - 5.5V
- **Protocollo**: Dallas 1-Wire

#### Schema Collegamento

```
DS18B20 (TO-92)                Raspberry Pi
                               
Pin 1 (GND)    ---------------- GND (Pin 6, 9, 14, 20, 25, 30, 34, 39)
Pin 2 (DATA)   ----[4.7kΩ]----- 3.3V (Pin 1, 17)
    |                           
    +-------------------------- GPIO 4 (Pin 7)
                               
Pin 3 (VDD)    ---------------- 3.3V (Pin 1, 17)
```

**Componenti**:
- **Resistenza pull-up**: 4.7kΩ tra DATA e 3.3V (essenziale!)
- **Cavo**: Lunghezza max consigliata 3 metri

#### Cablaggio

1. Collega **GND** (Pin 1) → Raspberry GND
2. Collega **VDD** (Pin 3) → Raspberry 3.3V
3. Collega **DATA** (Pin 2) → GPIO 4 (Pin 7)
4. Inserisci resistenza 4.7kΩ tra DATA e 3.3V

### Software Setup

#### 1. Abilitazione 1-Wire

```bash
# Via raspi-config
sudo raspi-config
# Interface Options → 1-Wire → Yes

# O manualmente in /boot/config.txt
sudo nano /boot/config.txt

# Aggiungi (se non presente):
dtoverlay=w1-gpio,gpiopin=4

# Salva e riavvia
sudo reboot
```

#### 2. Verifica Rilevamento

```bash
# Carica moduli kernel (automatico dopo riavvio)
sudo modprobe w1-gpio
sudo modprobe w1-therm

# Lista sensori rilevati
ls /sys/bus/w1/devices/

# Output atteso:
# 28-xxxxxxxxxxxx  w1_bus_master1
#
# "28-xxxxxxxxxxxx" è l'ID del sensore DS18B20
```

#### 3. Test Lettura

```bash
# Leggi temperatura (sostituisci ID sensore)
cat /sys/bus/w1/devices/28-xxxxxxxxxxxx/w1_slave

# Output:
# 7d 01 4b 46 7f ff 0c 10 57 : crc=57 YES
# 7d 01 4b 46 7f ff 0c 10 57 t=23812
#                              ^^^^^^
#                              23.812°C
```

### Configurazione Software

#### File: `config/gpio-mapping.js`

```javascript
temperature: {
  enabled: true,                    // Abilita/disabilita sensore
  sensorId: null,                   // null = auto-detect primo sensore
  basePath: '/sys/bus/w1/devices',
  readInterval: 5000,               // Leggi ogni 5 secondi
  pin: 4,                           // GPIO 4 (default 1-Wire)
}
```

#### Parametri

- **enabled**: `false` disabilita completamente il sensore
- **sensorId**: 
  - `null` → rileva automaticamente primo DS18B20
  - `'28-xxxxxxxxxxxx'` → forza ID specifico (multi-sensori)
- **readInterval**: Frequenza lettura in millisecondi (min 1000)
- **pin**: GPIO per 1-Wire (default 4, modificabile)

#### Sensori Multipli

Per usare più DS18B20 contemporaneamente:

```javascript
temperature: {
  enabled: true,
  sensorId: '28-0123456789ab',  // Specifica quale sensore usare
  // ... altre opzioni
}
```

E crea servizio duplicato per secondo sensore modificando `TemperatureSensorService.js`.

### Troubleshooting

#### Sensore Non Rilevato

```bash
# Controlla moduli caricati
lsmod | grep w1

# Output atteso:
# w1_therm
# w1_gpio

# Se mancano, carica manualmente
sudo modprobe w1-gpio
sudo modprobe w1-therm
```

#### Lettura CRC Errata

```bash
# Output con errore:
# xx xx xx xx xx xx xx xx xx : crc=xx NO

# Cause:
# 1. Resistenza pull-up mancante o valore errato
# 2. Cavo troppo lungo (>3m)
# 3. Interferenze elettriche
# 4. Sensore difettoso
```

**Soluzioni**:
1. Verifica resistenza 4.7kΩ presente
2. Accorcia cavi
3. Usa cavo schermato
4. Testa con altro sensore

---

## ⛽ Sensore Carburante ADS1115

### Hardware Setup

#### Specifiche ADS1115

- **Tipo**: ADC (Analog-to-Digital Converter) 16-bit I2C
- **Canali**: 4 single-ended o 2 differenziali
- **Risoluzione**: 16 bit (65536 livelli)
- **Range**: ±0.256V a ±6.144V (programmabile)
- **Interfaccia**: I2C (indirizzo default 0x48)
- **Sample rate**: 8-860 SPS

#### Schema Collegamento

```
ADS1115 Module               Raspberry Pi
                             
VDD    --------------------- 3.3V (Pin 1 o 17)
GND    --------------------- GND (Pin 6, 9, 14, ...)
SCL    --------------------- GPIO 3 / SCL (Pin 5)
SDA    --------------------- GPIO 2 / SDA (Pin 3)
ADDR   --------------------- GND (indirizzo 0x48)
A0     --------------------- Sensore carburante (partitore)
A1     --------------------- Non usato
A2     --------------------- Non usato
A3     --------------------- Non usato
```

#### Partitore Resistivo per Sensore Carburante

```
         +12V veicolo
              |
              R1 (100kΩ)
              |
              +-------> A0 (ADS1115)
              |
              R2 (33kΩ)
              |
             GND

Tensione uscita = Vin × (R2 / (R1 + R2))
                = 12V × (33kΩ / 133kΩ)
                = ~3.0V (max)
```

**Perché serve**:
- Sensore carburante originale varia 0-12V
- ADS1115 accetta max ±4.096V (gain 4096)
- Partitore riduce 12V → ~3V

**Calcolo custom**:
```
R2 / (R1 + R2) = Vout_max / Vin_max

Esempio per Vin_max=12V, Vout_max=3V:
R2 / (R1 + R2) = 3 / 12 = 0.25

Se R2 = 33kΩ:
33kΩ / (R1 + 33kΩ) = 0.25
R1 = 99kΩ ≈ 100kΩ
```

### Software Setup

#### 1. Abilitazione I2C

```bash
# Via raspi-config
sudo raspi-config
# Interface Options → I2C → Yes

# O manualmente
sudo nano /boot/config.txt

# Aggiungi:
dtparam=i2c_arm=on

# Salva e riavvia
sudo reboot
```

#### 2. Verifica Rilevamento

```bash
# Installa i2c-tools
sudo apt install i2c-tools

# Scansiona bus I2C
sudo i2cdetect -y 1

# Output atteso:
#      0  1  2  3  4  5  6  7  8  9  a  b  c  d  e  f
# 00:          -- -- -- -- -- -- -- -- -- -- -- -- --
# 10: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
# 20: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
# 30: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
# 40: -- -- -- -- -- -- -- -- 48 -- -- -- -- -- -- --
#                              ^^
#                         ADS1115 trovato!
# 50: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
# 60: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
# 70: -- -- -- -- -- -- -- --
```

#### 3. Test Lettura

```bash
# Installa libreria Python (per test veloce)
sudo apt install python3-pip
pip3 install adafruit-circuitpython-ads1x15

# Script test
python3 << EOF
import board
import busio
import adafruit_ads1x15.ads1115 as ADS
from adafruit_ads1x15.analog_in import AnalogIn

i2c = busio.I2C(board.SCL, board.SDA)
ads = ADS.ADS1115(i2c, gain=1)  # gain=1 → ±4.096V
chan = AnalogIn(ads, ADS.P0)    # Canale A0

print(f"Tensione: {chan.voltage:.3f}V")
print(f"Valore raw: {chan.value}")
EOF
```

### Configurazione Software

#### File: `config/gpio-mapping.js`

```javascript
fuel: {
  enabled: true,                // Abilita/disabilita sensore
  chip: 0,                      // 0=ADS1115 (16-bit) | 1=ADS1015 (12-bit)
  channel: 0,                   // Canale ADC: 0=A0, 1=A1, 2=A2, 3=A3
  gain: 4096,                   // ±4.096V full-scale
  sampleRate: 250,              // Sample rate (SPS)
  readInterval: 500,            // Lettura ogni 500ms
  
  // Partitore resistivo
  voltageDivider: {
    r1: 100000,                 // 100kΩ
    r2: 33000,                  // 33kΩ
  },
  
  // Calibrazione tensione → percentuale carburante
  calibration: {
    voltageEmpty: 0.5,          // Tensione serbatoio vuoto
    voltageFull: 4.0,           // Tensione serbatoio pieno
  },
  
  pins: {
    sda: 2,                     // GPIO 2 (SDA)
    scl: 3,                     // GPIO 3 (SCL)
  },
}
```

#### Parametri Gain

| gain | Range | Risoluzione (16-bit) |
|------|-------|----------------------|
| 256 | ±0.256V | 7.8 µV |
| 512 | ±0.512V | 15.6 µV |
| 1024 | ±1.024V | 31.2 µV |
| 2048 | ±2.048V | 62.5 µV |
| 4096 | ±4.096V | 125 µV |
| 6144 | ±6.144V | 187.5 µV |

**Scelta gain**:
- Usa gain più basso possibile che contiene il tuo range
- Per sensore 0-3V (con partitore): `gain: 4096`
- Se tensione max <2V: `gain: 2048` (più risoluzione)

#### Calibrazione

La calibrazione mappa tensione ADC → percentuale carburante (0-100%).

**Procedura**:

1. **Serbatoio vuoto**: 
   - Accendi quadro con serbatoio vuoto
   - Leggi tensione in console debug
   - Imposta `voltageEmpty`

2. **Serbatoio pieno**:
   - Riempi serbatoio
   - Accendi quadro
   - Leggi tensione
   - Imposta `voltageFull`

**Esempio**:
```javascript
calibration: {
  voltageEmpty: 0.8,   // Misurato: 0.8V a serbatoio vuoto
  voltageFull: 3.2,    // Misurato: 3.2V a serbatoio pieno
}
```

Il sistema calcolerà linearmente:
```
percentuale = ((V_misurata - V_empty) / (V_full - V_empty)) × 100
```

#### Sample Rate

| SPS | Utilizzo |
|-----|----------|
| 8 | Massima precisione, lento |
| 16-64 | Bilanciato |
| 128-250 | Standard (consigliato) |
| 475-860 | Alta velocità, meno preciso |

**Consiglio**: 250 SPS è ideale per sensore carburante (cambia lentamente).

### Troubleshooting

#### ADS1115 Non Rilevato

```bash
# Verifica modulo I2C caricato
lsmod | grep i2c

# Verifica device presente
ls /dev/i2c*
# Output atteso: /dev/i2c-1

# Scansiona bus
sudo i2cdetect -y 1
# Se non appare "48", controlla:
# - Alimentazione VDD e GND
# - Cablaggio SDA/SCL
# - Saldature (se modulo custom)
```

#### Lettura Sempre 0V o Valore Fisso

**Cause**:
1. A0 non collegato → legge 0V
2. Canale sbagliato in configurazione
3. Gain troppo basso (segnale satura)

**Soluzioni**:
```javascript
// Verifica canale corretto
channel: 0,  // 0=A0, 1=A1, 2=A2, 3=A3

// Prova gain più alto
gain: 6144,  // Se segnale >4.096V
```

#### Percentuale Carburante Errata

**Causa**: Calibrazione non corretta

**Soluzione**: Ricalibra con procedura sopra

---

## ⚡ Gestione Quadro Accensione (Ignition)

### Scopo

Il sistema rileva quando il quadro veicolo viene acceso/spento e esegue azioni automatiche:
- **Quadro spento** → Esegui script `low-power.sh` (risparmio energetico)
- **Quadro acceso** → Esegui script `wake.sh` (riattivazione)

### Hardware Setup

#### Collegamento

Collega un optoaccoppiatore dedicato al segnale "chiave inserita" (12V quando quadro acceso):

```
Quadro 12V (KEY) --[1kΩ]--+
                           |
                        PC817 -----> GPIO 21
                           |
                          GND
```

**Logica**:
- Quadro **acceso** (12V) → GPIO 21 **HIGH**
- Quadro **spento** (0V) → GPIO 21 **LOW**

O viceversa se usi optoaccoppiatore active-low (imposta `activeOn: 0`)

### Configurazione

#### File: `config/gpio-mapping.js`

```javascript
ignition: {
  enabled: true,                    // Abilita gestione quadro
  pin: 21,                          // GPIO 21 dedicato
  activeOn: 1,                      // 1=active high | 0=active low
  scripts: {
    lowPower: './scripts/low-power.sh',   // Quando quadro si SPEGNE
    wake: './scripts/wake.sh',             // Quando quadro si ACCENDE
  },
}
```

#### Parametri

- **enabled**: `false` disabilita completamente la funzione
- **pin**: GPIO dedicato (diverso da quelli usati per le spie)
- **activeOn**:
  - `1` → HIGH = quadro acceso (optoaccoppiatore active high)
  - `0` → LOW = quadro acceso (optoaccoppiatore active low)
- **scripts**: Path relativi a `server/` directory

### Script Power-Saving

#### File: `scripts/low-power.sh`

Eseguito quando il quadro viene **spento**.

```bash
#!/bin/bash
# Script eseguito quando il quadro si spegne

logger "PandaOS: Quadro spento - Avvio modalità risparmio energetico"

# Spegni display HDMI
vcgencmd display_power 0

# Riduci luminosità backlight (se GPIO controllato)
# echo 0 > /sys/class/backlight/rpi_backlight/brightness

# Ferma servizi non essenziali
# systemctl stop bluetooth

# Opzionale: Shutdown dopo X minuti
# sleep 600 && sudo shutdown -h now &

logger "PandaOS: Modalità risparmio energetico attivata"
```

Rendi eseguibile:
```bash
chmod +x server/scripts/low-power.sh
```

#### File: `scripts/wake.sh`

Eseguito quando il quadro viene **acceso**.

```bash
#!/bin/bash
# Script eseguito quando il quadro si accende

logger "PandaOS: Quadro acceso - Riattivazione sistema"

# Riaccendi display HDMI
vcgencmd display_power 1

# Ripristina luminosità backlight
# echo 255 > /sys/class/backlight/rpi_backlight/brightness

# Riavvia servizi
# systemctl start bluetooth

logger "PandaOS: Sistema riattivato"
```

Rendi eseguibile:
```bash
chmod +x server/scripts/wake.sh
```

### Customizzazione Script

Gli script possono eseguire qualsiasi comando bash, ad esempio:

#### Auto-Shutdown dopo 10 minuti

In `low-power.sh`:
```bash
# Shutdown se quadro rimane spento 10 minuti
(sleep 600 && sudo shutdown -h now) &
echo $! > /tmp/pandaos-shutdown.pid
```

In `wake.sh`:
```bash
# Annulla shutdown se quadro si riaccende
if [ -f /tmp/pandaos-shutdown.pid ]; then
  kill $(cat /tmp/pandaos-shutdown.pid) 2>/dev/null
  rm /tmp/pandaos-shutdown.pid
fi
```

#### Notifica Telegram

```bash
# In low-power.sh
curl -s -X POST "https://api.telegram.org/bot<TOKEN>/sendMessage" \
  -d chat_id=<CHAT_ID> \
  -d text="🚗 Quadro spento - PandaOS in standby"
```

---

## 🚨 Troubleshooting

### Server non si avvia

#### Errore: "Piattaforma non supportata"

```
❌ ERRORE: Dipendenze Raspberry Pi essenziali non disponibili
Piattaforma non supportata: linux x64 - richiesto Linux ARM
```

**Causa**: Esegui su sistema non-Raspberry Pi

**Soluzione**: Esegui server solo su Raspberry Pi, o disabilita check in `OBDServer.js` (non consigliato)

#### Errore: "Modulo GPIO non disponibile"

```
Modulo GPIO (onoff) non disponibile
```

**Causa**: Libreria `onoff` non installata o sistema non compatibile

**Soluzione**:
```bash
cd server
npm install onoff
```

### ELM327 non risponde

#### Sintomo: "Porta /dev/ttyUSB0 non trovata"

**Soluzione**:
1. Verifica ELM327 collegato: `lsusb`
2. Verifica porta: `ls -l /dev/ttyUSB*`
3. Controlla permessi: `groups` (deve includere `dialout`)

#### Sintomo: "Timeout - Nessuna risposta ricevuta"

**Causa**: ELM327 non comunica

**Soluzione**:
1. Verifica baudrate in `OBDCommunicationService.js`
2. Testa con minicom: `minicom -D /dev/ttyUSB0 -b 38400`
3. Prova reset: Scollega ELM327, attendi 10 secondi, ricollega

### GPIO non funziona

#### Sintomo: Spie non rilevate

**Diagnosi**:
```bash
# Test manuale GPIO
gpio -g mode 17 in
gpio -g read 17

# Accendi/spegni spia veicolo e controlla se valore cambia
```

**Soluzioni**:
1. Verifica cablaggio optoaccoppiatore
2. Testa con LED e resistenza al posto di optoaccoppiatore
3. Controlla numero pin in `gpio-mapping.js` (BCM vs fisico)

### Sensori non rilevano

#### DS18B20: "1-Wire non trovato"

```bash
# Verifica modulo caricato
lsmod | grep w1

# Se assente
sudo modprobe w1-gpio
sudo modprobe w1-therm

# Verifica config permanente
grep w1-gpio /boot/config.txt
```

#### ADS1115: "Indirizzo 0x48 non risponde"

```bash
# Scansiona bus I2C
sudo i2cdetect -y 1

# Se non appare "48":
# - Controlla alimentazione ADS1115
# - Verifica SDA/SCL non invertiti
# - Testa con cavo più corto
```

---

## 📚 Riferimenti Tecnici

### Datasheet e Documentazione

- **Raspberry Pi GPIO**: https://pinout.xyz/
- **DS18B20**: https://datasheets.maximintegrated.com/en/ds/DS18B20.pdf
- **ADS1115**: https://www.ti.com/lit/ds/symlink/ads1115.pdf
- **PC817 Optocoupler**: https://www.farnell.com/datasheets/73758.pdf
- **ELM327**: https://www.elmelectronics.com/wp-content/uploads/2017/01/ELM327DS.pdf
- **Fiat Panda 141 - Schema Elettrico Ufficiale**: http://www.bunkeringegnere.altervista.org/esplosi/FIAT%20PANDA/panda%20141/1100%20mpi/55%20IMPIANTO%20ELETTRICO%20-%20SCHEMI%20-%20GAMMA%202000.pdf

### Comandi Utili

```bash
# GPIO
gpio readall                    # Stato tutti pin
gpio -g read <pin>              # Leggi pin specifico

# I2C
sudo i2cdetect -y 1             # Scansione bus
sudo i2cget -y 1 0x48 0x00 w    # Leggi registro ADS1115

# 1-Wire
ls /sys/bus/w1/devices/         # Lista sensori
cat /sys/bus/w1/devices/28-*/w1_slave  # Leggi temperatura

# Serial
ls -l /dev/tty*                 # Lista porte
sudo minicom -D /dev/ttyUSB0    # Monitor seriale

# Logs
journalctl -u obd-server -f     # Log server (se PM2/systemd)
pm2 logs obd-server             # Log PM2
```

---

**Ultimo aggiornamento**: v0.9.0  
**Hardware target**: Raspberry Pi 4B + Fiat Panda 141

