[🇬🇧 English](CONFIGURAZIONE_SERVER.en.md) | [🇮🇹 Italiano](CONFIGURAZIONE_SERVER.md) | [🇩🇪 Deutsch](CONFIGURAZIONE_SERVER.de.md)

---

# 🔧 Server Configuration - PandaOS Cluster

Complete guide to backend server hardware and software configuration.

> ⚠️ **WARNING**: This guide describes electrical connections and hardware modifications. Any intervention on vehicle electrical systems involves risks. Read the [full Disclaimer](../README.md#️-disclaimer) and proceed only if you know what you're doing. The authors assume no responsibility for damages resulting from the use of this information.

---

## 📋 Table of Contents

> 💡 **Need to purchase components?** First consult [HARDWARE.md](../HARDWARE.md) for the complete list of everything needed.

1. [Hardware Requirements](#-hardware-requirements)
2. [Raspberry Pi Configuration](#-raspberry-pi-configuration)
3. [OBD-II Serial Port](#-obd-ii-serial-port)
4. [GPIO Configuration](#-gpio-configuration)
5. [DS18B20 Temperature Sensor](#-ds18b20-temperature-sensor)
6. [ADS1115 Fuel Sensor](#-ads1115-fuel-sensor)
7. [Ignition Management](#-ignition-management)
8. [Troubleshooting](#-troubleshooting)

---

## 🛠️ Hardware Requirements

### Essential Components

| Component | Model | Purpose |
|-----------|-------|---------|
| **SBC** | Raspberry Pi 4B (4GB+) | Main processing |
| **OBD Adapter** | ELM327 USB | ECU communication |
| **Optocouplers** | PC817 or similar | Electrical isolation for warning lights |
| **Power Supply** | 5V 3A USB-C | Raspberry power |
| **Display** | HDMI 1920x480+ | Cluster display |

### Optional Components

| Component | Model | Purpose |
|-----------|-------|---------|
| **Temperature sensor** | DS18B20 | External temperature |
| **ADC** | ADS1115 | Fuel sensor reading |
| **Resistors** | 4.7kΩ, 100kΩ, 33kΩ | Pull-up and divider |

---

## 🔧 Raspberry Pi Configuration

### 1. Operating System Installation

#### OS Choice

**TL;DR**: Raspberry Pi OS Lite (64-bit) Debian-based, stripped of unnecessary services.

**Recommended Distribution**: Raspberry Pi OS Lite (64-bit)
- **Download**: [raspberrypi.com/software](https://www.raspberrypi.com/software/)
- **Version**: Bookworm (Debian 12) or later
- **Architecture**: 64-bit (better performance for Node.js/Electron)

**Why Lite instead of Desktop?**
- ✅ Boot time ~30 seconds (vs ~60s with Desktop)
- ✅ Free RAM: ~200MB (vs ~500MB with Desktop Environment)
- ✅ No unnecessary background services
- ✅ Electron already provides UI, no Desktop Manager needed
- ❌ More complex to configure (no GUI, everything via SSH)

**Tested Alternatives**:
- **Raspberry Pi OS Desktop**: Works but slow boot (~60s) and RAM waste
- **DietPi**: Excellent for ultra-fast boot (~15-20s) but requires more manual configuration

#### Boot Time: The Reality

After various tests, with **stripped-down Raspberry Pi OS Lite** we achieved:

- **~30 seconds** full boot (POST → Login → PandaOS operational)
- **~20 seconds** if you disable non-essential services (see optimizations below)

**Is this too much?** Depends:
- ❌ If you turn ignition on/off frequently: yes, waiting is annoying
- ✅ If you use **always-on standby mode** (what we used): not a problem.

#### Always-On Standby Approach:

**How it works in our setup**:

1. Raspberry Pi **always powered** (direct battery)
2. GPIO 21 detects "key inserted" (see § Ignition)
3. When **ignition off**:
   - Script `low-power.sh` turns off HDMI display
   - System in standby: ~0.4W consumption (negligible for car battery)
4. When **ignition on**:
   - Script `wake.sh` turns display back on
   - System **immediately operational** (0 seconds boot!)

**Advantages**:
- ⚡ Cluster available instantly on ignition
- 🔋 Very low standby consumption (~30mA @ 12V)
- 🛡️ SD card protected (no abrupt shutdowns)
- 🕐 Boot time becomes irrelevant

**Actual Measured Consumption**:
- **Standby** (display off, CPU idle): ~0.3-0.5W
- **Operational** (display on, OBD data): ~6-8W
- **Battery impact**: Negligible (<0.01% charge/day)

⚠️ **Note**: If you leave the car parked for >2 weeks, consider manual switch or automatic shutdown after 7 days of inactivity.

#### Base Installation

```bash
# Download Raspberry Pi Imager
# https://www.raspberrypi.com/software/

# 1. Select OS: "Raspberry Pi OS Lite (64-bit)"
# 2. Configure (gear icon):
#    - Hostname: pandaos
#    - Enable SSH
#    - Username/Password: pi/your-password
#    - WiFi (SSID and password)
#    - Locale: en_US, timezone America/New_York (or your timezone)
# 3. Write to microSD
# 4. Insert in Raspberry and power on
```

#### Boot Time Optimizations (Advanced)

> 💡 **Note**: This section is for those who want the fastest boot possible. If you use **always-on standby**, you can safely skip it.

With these tweaks you can drop from 30s to ~15-20s:

**1. Disable Unnecessary Services**

```bash
# Bluetooth (if not needed)
sudo systemctl disable bluetooth.service
sudo systemctl disable hciuart.service

# ModemManager (if you don't have USB modem)
sudo systemctl disable ModemManager.service

# Printer services (not needed in a car)
sudo systemctl disable cups.service
sudo systemctl disable cups-browsed.service

# Triggerhappy (not needed)
sudo systemctl disable triggerhappy.service

# Reboot and verify boot time
sudo reboot
systemd-analyze  # Shows total time
systemd-analyze blame  # Shows slowest services
```

**2. Optimize Kernel Boot**

Edit `/boot/firmware/cmdline.txt`:

```bash
sudo nano /boot/firmware/cmdline.txt

# Add at the end of the line (everything on ONE line):
quiet splash fastboot noatime nodiratime
```

**3. Disable Network Wait**

```bash
# If using static IP or network not needed at boot
sudo systemctl disable NetworkManager-wait-online.service
sudo systemctl disable systemd-networkd-wait-online.service
```

**4. Reduce Boot Timeouts**

In `/etc/systemd/system.conf`:

```bash
sudo nano /etc/systemd/system.conf

# Uncomment and modify:
DefaultTimeoutStartSec=10s
DefaultTimeoutStopSec=5s
```

**5. Start OBD Server Before Desktop**

PM2 + systemd for parallel startup (see § PM2 Configuration in README.md).

**Expected Results**:
- **OS Boot**: ~8-12 seconds
- **Services startup**: ~5-8 seconds
- **Total**: ~15-20 seconds (vs 30s original)

#### Future Optimizations (TODO)

> 📝 WIP Section - Contributions welcome!

For those wanting to experiment with <10 second boot:

**Approaches to Test**:
- **Custom init**: Replace systemd with lighter init (runit, OpenRC)
- **Minimal kernel**: Compile custom Linux kernel with only necessary drivers
- **Read-only root**: Root filesystem in read-only mode (faster, more stable)
- **Optimized initramfs**: Reduce services loaded at boot

**Roadmap**:
1. Document step-by-step "Debian minimal" procedure
2. Automatic script to apply boot optimizations
3. Pre-configured downloadable SD image

If you have experience with embedded Linux and want to contribute, open an [issue](https://github.com/cyberpandino/cluster/issues)!

---

### 2. Node.js and npm Installation

⚠️ **Important**: `apt install nodejs` installs an obsolete version (v12-14). PandaOS requires **Node.js 18+**.

**Recommended Method: NodeSource**

```bash
# Install Node.js 20 LTS
curl -fsSL https://deb.nodesource.com/setup_20.x | sudo -E bash -
sudo apt install -y nodejs

# Verify
node --version   # v20.x.x
npm --version    # 10.x.x
```

**Alternative: nvm** (if you need to manage multiple versions)

```bash
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.0/install.sh | bash
source ~/.bashrc
nvm install 20 && nvm use 20
```

> 💡 NodeSource is more stable with PM2/systemd (recommended for production)

**Git and Build Tools**

```bash
sudo apt update && sudo apt upgrade -y
sudo apt install -y git build-essential python3
```

`build-essential` is necessary to compile native modules (SerialPort, onoff, i2c-bus).

**Verification**

```bash
node --version    # >= v18.0.0
npm --version     # >= 9.0.0
gcc --version     # Verify compiler
```

---

### 3. Hardware Interfaces Configuration

```bash
sudo raspi-config
```

Enable:
- **Interface Options** → **I2C** → Yes (for ADS1115)
- **Interface Options** → **Serial Port** → 
  - "Would you like a login shell...?" → **No**
  - "Would you like the serial port hardware...?" → **Yes**
- **Interface Options** → **1-Wire** → Yes (for DS18B20)

Reboot:
```bash
sudo reboot
```

### 4. User Permissions

```bash
# Add user to necessary groups
sudo usermod -a -G dialout $USER    # Serial port
sudo usermod -a -G gpio $USER       # GPIO
sudo usermod -a -G i2c $USER        # I2C

# Logout and login to apply
```

### 5. Configuration Verification

```bash
# Verify I2C
ls -l /dev/i2c*
# Expected output: /dev/i2c-1

# Verify Serial
ls -l /dev/ttyUSB* /dev/ttyAMA*
# Expected output: /dev/ttyUSB0 (with ELM327 connected)

# Verify 1-Wire
ls -l /dev/1-wire*
# Or check: ls /sys/bus/w1/devices/
```

---

## 🔌 OBD-II Serial Port

### Hardware Setup

#### 1. ELM327 Adapter

**Required specifications**:
- Protocol: OBD-II (ISO 9141-2, ISO 14230-4)
- Connection: USB (FTDI or CH340)
- Compatibility: ELM327 v1.5 or higher

**Connection**:
1. Connect ELM327 to Magneti Marelli IAW 4AF diagnostic connector
2. Connect ELM327 via USB to Raspberry Pi
3. Verify LED lit on ELM327

#### 2. Port Identification

```bash
# List available serial ports
ls -l /dev/ttyUSB*
ls -l /dev/ttyAMA*

# Typical output:
# /dev/ttyUSB0 → ELM327 USB
# /dev/ttyAMA0 → GPIO UART (alternative)

# Detailed info
dmesg | grep tty
```

#### 3. Connection Test

```bash
# Install minicom
sudo apt install minicom

# Connect to port (38400 baud)
minicom -D /dev/ttyUSB0 -b 38400

# Test ELM327 commands:
ATZ          # Reset (response: ELM327 v1.5)
ATE0         # Echo off (response: OK)
0100         # PIDs supported (response: hex data)

# Exit: CTRL+A then X
```

### Software Configuration

#### File: `services/OBDCommunicationService.js`

```javascript
constructor() {
  this.portPath = '/dev/ttyUSB0';  // ← MODIFY HERE if different
  this.baudRate = 38400;            // ELM327 standard
}
```

#### Alternative Ports

If ELM327 is recognized on different port:

```javascript
// Alternative USB port
this.portPath = '/dev/ttyUSB1';

// GPIO UART (if wired directly)
this.portPath = '/dev/ttyAMA0';

// CH340 adapter (some clones)
this.portPath = '/dev/ttyACM0';
```

#### Alternative Baudrates

```javascript
// Standard ELM327
this.baudRate = 38400;

// Some adapters configured differently
this.baudRate = 115200;
this.baudRate = 9600;
```

### OBD Protocol

The server automatically supports Fiat Panda 141 / Magneti Marelli IAW 4AF protocols:
- **ISO 9141-2** (K-Line)
- **ISO 14230-4** (KWP2000)

The command `ATSP0` sets auto-detection of protocol.

---

## 🔢 GPIO Configuration

### Configuration File

**Path**: `server/config/gpio-mapping.js`

This file contains all GPIO mapping for warning lights, sensors, and ignition.

### Optocoupler Schematic

```
                  RASPBERRY PI
                  +------------+
    12V Light --->|PC817   GPIO|----> Software read
                  |            |
            GND-->|GND         |
                  +------------+
```

**Logic**:
- Light **ON** (12V) → Optocoupler ON → GPIO **HIGH** (1)
- Light **OFF** (0V) → Optocoupler OFF → GPIO **LOW** (0)

### Complete Pin Mapping

#### Global Configuration

```javascript
config: {
  mode: 'BCM',              // Broadcom numbering (not physical)
  pullMode: 'PUD_DOWN',     // Internal pull-down resistor
  debounceTime: 50,         // Anti-bounce (ms)
  pollingInterval: 100,     // Read frequency (ms)
}
```

**Parameter explanation**:
- **mode**: `'BCM'` uses GPIO numbering, not physical pin numbers
- **pullMode**: `'PUD_DOWN'` ensures 0V when optocoupler open
- **debounceTime**: Filters spurious signals (e.g. LED flickering)
- **pollingInterval**: Every 100ms checks GPIO state

#### GPIO Pin Table

| Light/Function | GPIO (BCM) | Physical Pin | Description |
|----------------|------------|--------------|-------------|
| **Lighting** |
| High Beam | 5 | 29 | High beam headlights |
| Low Beam | 6 | 31 | Low beam headlights |
| Rear Fog | 13 | 33 | Rear fog light |
| **Direction Indicators** |
| Turn Signals | 17 | 11 | Direction indicators |
| Hazards | 12 | 32 | Emergency lights |
| **Engine System** |
| Coolant Temp | 16 | 36 | Coolant temperature |
| Oil Pressure | 22 | 15 | Engine oil pressure |
| Injectors | 24 | 18 | Injection system |
| **Electrical System** |
| Alternator/Battery | 27 | 13 | Battery charge |
| Key Inserted (KEY) | 25 | 22 | Ignition on |
| **Other Systems** |
| Brake System | 23 | 16 | Brakes |
| Rear Window Heater | 19 | 35 | Defroster |
| Fuel Reserve | 20 | 38 | Low level |
| **Power Management** |
| Ignition (dashboard) | 21 | 40 | Detect ignition ON/OFF |

### Raspberry Pi 4B Pinout Diagram

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

### Mapping Example

#### Code in File

```javascript
mapping: {
  highBeam: {
    pin: 5,                    // GPIO 5 (Physical pin 29)
    name: 'High Beam',
    description: 'High beam headlights',
  },
  
  turnSignals: {
    pin: 17,                   // GPIO 17 (Physical pin 11)
    name: 'Turn Signals',
    description: 'Direction indicators',
  },
  
  // ... other lights
}
```

#### Modify Mapping

If you wired optocouplers differently:

```javascript
// EXAMPLE: Move high beam from GPIO 5 to GPIO 26
highBeam: {
  pin: 26,  // ← Change only this number
  name: 'High Beam',
  description: 'High beam headlights',
}
```

### Optocoupler Wiring

> 📘 **IMPORTANT**: Before connecting optocouplers, consult the [Official Fiat Panda 141 Electrical Diagram](http://www.bunkeringegnere.altervista.org/esplosi/FIAT%20PANDA/panda%20141/1100%20mpi/55%20IMPIANTO%20ELETTRICO%20-%20SCHEMI%20-%20GAMMA%202000.pdf) to identify correct wires with color codes (e.g. R=Red, BN=White-Black, GV=Yellow-Green).

#### Single Optocoupler Schematic (PC817)

```
        Vehicle (Input Side)          |  Raspberry Pi (Output Side)
                                      |
    +12V (from light) ---[R 1kΩ]--+  |
                                   |  |
                            LED+ (1)  |
                                      |
                            LED- (2)--|--- Common GND
                                      |
                                 (3)--|--- GPIO (e.g. GPIO 5)
                                      |
                                 (4)--|--- GND
```

**Components**:
- **R**: LED current limiting resistor (1kΩ - 2.2kΩ)
- **PC817**: Standard optocoupler
- **Pin 1-2**: Internal LED (vehicle side)
- **Pin 3-4**: Output transistor (Raspberry side)

**How to identify wires**:
1. Consult official electrical diagram (link above)
2. Find instrument cluster (page "Instrument cluster connection")
3. Identify desired light (e.g. high beam, turn signals, etc.)
4. Note wire color code (e.g. "BN" = White-Black)
5. Verify with multimeter presence of 12V when light is on

#### Complete Multi-Light Circuit

```
Light 1 (12V) --[1kΩ]--+
                        |
                    PC817 #1 -----> GPIO 17
                        |
                       GND

Light 2 (12V) --[1kΩ]--+
                        |
                    PC817 #2 -----> GPIO 27
                        |
                       GND

... (repeat for each light)
```

**Notes**:
- Use common GND for all optocouplers
- Each light has its dedicated optocoupler
- Resistors in series to protect internal LED

### GPIO Testing

#### Manual CLI Test

```bash
# Install wiringpi (optional)
sudo apt install wiringpi

# Read GPIO 17 state
gpio -g read 17

# Output:
# 0 = LOW (light off)
# 1 = HIGH (light on)

# Watch mode (updates every second)
watch -n 1 'gpio -g read 17'
```

#### Python Test

```python
#!/usr/bin/env python3
import RPi.GPIO as GPIO
import time

# Setup
GPIO.setmode(GPIO.BCM)
pin = 17  # GPIO 17 (turn signals)
GPIO.setup(pin, GPIO.IN, pull_up_down=GPIO.PUD_DOWN)

try:
    print(f"Monitoring GPIO {pin} (CTRL+C to exit)")
    while True:
        state = GPIO.input(pin)
        print(f"GPIO {pin}: {'HIGH (light on)' if state else 'LOW (light off)'}")
        time.sleep(0.5)
except KeyboardInterrupt:
    print("\nInterrupted")
finally:
    GPIO.cleanup()
```

Save as `test_gpio.py` and run:
```bash
python3 test_gpio.py
```

---

## 🌡️ DS18B20 Temperature Sensor

### Hardware Setup

#### Specifications

- **Type**: Digital temperature sensor 1-Wire
- **Range**: -55°C to +125°C
- **Accuracy**: ±0.5°C
- **Power**: 3.0V - 5.5V
- **Protocol**: Dallas 1-Wire

#### Connection Diagram

```
DS18B20 (TO-92)                Raspberry Pi
                               
Pin 1 (GND)    ---------------- GND (Pin 6, 9, 14, 20, 25, 30, 34, 39)
Pin 2 (DATA)   ----[4.7kΩ]----- 3.3V (Pin 1, 17)
    |                           
    +-------------------------- GPIO 4 (Pin 7)
                               
Pin 3 (VDD)    ---------------- 3.3V (Pin 1, 17)
```

**Components**:
- **Pull-up resistor**: 4.7kΩ between DATA and 3.3V (essential!)
- **Cable**: Recommended max length 3 meters

#### Wiring

1. Connect **GND** (Pin 1) → Raspberry GND
2. Connect **VDD** (Pin 3) → Raspberry 3.3V
3. Connect **DATA** (Pin 2) → GPIO 4 (Pin 7)
4. Insert 4.7kΩ resistor between DATA and 3.3V

### Software Setup

#### 1. Enable 1-Wire

```bash
# Via raspi-config
sudo raspi-config
# Interface Options → 1-Wire → Yes

# Or manually in /boot/config.txt
sudo nano /boot/config.txt

# Add (if not present):
dtoverlay=w1-gpio,gpiopin=4

# Save and reboot
sudo reboot
```

#### 2. Verify Detection

```bash
# Load kernel modules (automatic after reboot)
sudo modprobe w1-gpio
sudo modprobe w1-therm

# List detected sensors
ls /sys/bus/w1/devices/

# Expected output:
# 28-xxxxxxxxxxxx  w1_bus_master1
#
# "28-xxxxxxxxxxxx" is the DS18B20 sensor ID
```

#### 3. Read Test

```bash
# Read temperature (replace sensor ID)
cat /sys/bus/w1/devices/28-xxxxxxxxxxxx/w1_slave

# Output:
# 7d 01 4b 46 7f ff 0c 10 57 : crc=57 YES
# 7d 01 4b 46 7f ff 0c 10 57 t=23812
#                              ^^^^^^
#                              23.812°C
```

### Software Configuration

#### File: `config/gpio-mapping.js`

```javascript
temperature: {
  enabled: true,                    // Enable/disable sensor
  sensorId: null,                   // null = auto-detect first sensor
  basePath: '/sys/bus/w1/devices',
  readInterval: 5000,               // Read every 5 seconds
  pin: 4,                           // GPIO 4 (default 1-Wire)
}
```

#### Parameters

- **enabled**: `false` completely disables sensor
- **sensorId**: 
  - `null` → automatically detect first DS18B20
  - `'28-xxxxxxxxxxxx'` → force specific ID (multi-sensors)
- **readInterval**: Read frequency in milliseconds (min 1000)
- **pin**: GPIO for 1-Wire (default 4, modifiable)

#### Multiple Sensors

To use multiple DS18B20 simultaneously:

```javascript
temperature: {
  enabled: true,
  sensorId: '28-0123456789ab',  // Specify which sensor to use
  // ... other options
}
```

And create duplicate service for second sensor by modifying `TemperatureSensorService.js`.

### Troubleshooting

#### Sensor Not Detected

```bash
# Check loaded modules
lsmod | grep w1

# Expected output:
# w1_therm
# w1_gpio

# If missing, load manually
sudo modprobe w1-gpio
sudo modprobe w1-therm
```

#### Wrong CRC Read

```bash
# Output with error:
# xx xx xx xx xx xx xx xx xx : crc=xx NO

# Causes:
# 1. Pull-up resistor missing or wrong value
# 2. Cable too long (>3m)
# 3. Electrical interference
# 4. Defective sensor
```

**Solutions**:
1. Verify 4.7kΩ resistor present
2. Shorten cables
3. Use shielded cable
4. Test with another sensor

---

## ⛽ ADS1115 Fuel Sensor

### Hardware Setup

#### ADS1115 Specifications

- **Type**: 16-bit I2C ADC (Analog-to-Digital Converter)
- **Channels**: 4 single-ended or 2 differential
- **Resolution**: 16 bit (65536 levels)
- **Range**: ±0.256V to ±6.144V (programmable)
- **Interface**: I2C (default address 0x48)
- **Sample rate**: 8-860 SPS

#### Connection Diagram

```
ADS1115 Module               Raspberry Pi
                             
VDD    --------------------- 3.3V (Pin 1 or 17)
GND    --------------------- GND (Pin 6, 9, 14, ...)
SCL    --------------------- GPIO 3 / SCL (Pin 5)
SDA    --------------------- GPIO 2 / SDA (Pin 3)
ADDR   --------------------- GND (address 0x48)
A0     --------------------- Fuel sensor (divider)
A1     --------------------- Not used
A2     --------------------- Not used
A3     --------------------- Not used
```

#### Resistive Divider for Fuel Sensor

```
         +12V vehicle
              |
              R1 (100kΩ)
              |
              +-------> A0 (ADS1115)
              |
              R2 (33kΩ)
              |
             GND

Output voltage = Vin × (R2 / (R1 + R2))
               = 12V × (33kΩ / 133kΩ)
               = ~3.0V (max)
```

**Why it's needed**:
- Original fuel sensor varies 0-12V
- ADS1115 accepts max ±4.096V (gain 4096)
- Divider reduces 12V → ~3V

**Custom calculation**:
```
R2 / (R1 + R2) = Vout_max / Vin_max

Example for Vin_max=12V, Vout_max=3V:
R2 / (R1 + R2) = 3 / 12 = 0.25

If R2 = 33kΩ:
33kΩ / (R1 + 33kΩ) = 0.25
R1 = 99kΩ ≈ 100kΩ
```

### Software Setup

#### 1. Enable I2C

```bash
# Via raspi-config
sudo raspi-config
# Interface Options → I2C → Yes

# Or manually
sudo nano /boot/config.txt

# Add:
dtparam=i2c_arm=on

# Save and reboot
sudo reboot
```

#### 2. Verify Detection

```bash
# Install i2c-tools
sudo apt install i2c-tools

# Scan I2C bus
sudo i2cdetect -y 1

# Expected output:
#      0  1  2  3  4  5  6  7  8  9  a  b  c  d  e  f
# 00:          -- -- -- -- -- -- -- -- -- -- -- -- --
# 10: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
# 20: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
# 30: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
# 40: -- -- -- -- -- -- -- -- 48 -- -- -- -- -- -- --
#                              ^^
#                         ADS1115 found!
# 50: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
# 60: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
# 70: -- -- -- -- -- -- -- --
```

#### 3. Read Test

```bash
# Install Python library (for quick test)
sudo apt install python3-pip
pip3 install adafruit-circuitpython-ads1x15

# Test script
python3 << EOF
import board
import busio
import adafruit_ads1x15.ads1115 as ADS
from adafruit_ads1x15.analog_in import AnalogIn

i2c = busio.I2C(board.SCL, board.SDA)
ads = ADS.ADS1115(i2c, gain=1)  # gain=1 → ±4.096V
chan = AnalogIn(ads, ADS.P0)    # Channel A0

print(f"Voltage: {chan.voltage:.3f}V")
print(f"Raw value: {chan.value}")
EOF
```

### Software Configuration

#### File: `config/gpio-mapping.js`

```javascript
fuel: {
  enabled: true,                // Enable/disable sensor
  chip: 0,                      // 0=ADS1115 (16-bit) | 1=ADS1015 (12-bit)
  channel: 0,                   // ADC channel: 0=A0, 1=A1, 2=A2, 3=A3
  gain: 4096,                   // ±4.096V full-scale
  sampleRate: 250,              // Sample rate (SPS)
  readInterval: 500,            // Read every 500ms
  
  // Resistive divider
  voltageDivider: {
    r1: 100000,                 // 100kΩ
    r2: 33000,                  // 33kΩ
  },
  
  // Calibration voltage → fuel percentage
  calibration: {
    voltageEmpty: 0.5,          // Voltage empty tank
    voltageFull: 4.0,           // Voltage full tank
  },
  
  pins: {
    sda: 2,                     // GPIO 2 (SDA)
    scl: 3,                     // GPIO 3 (SCL)
  },
}
```

#### Gain Parameters

| gain | Range | Resolution (16-bit) |
|------|-------|---------------------|
| 256 | ±0.256V | 7.8 µV |
| 512 | ±0.512V | 15.6 µV |
| 1024 | ±1.024V | 31.2 µV |
| 2048 | ±2.048V | 62.5 µV |
| 4096 | ±4.096V | 125 µV |
| 6144 | ±6.144V | 187.5 µV |

**Gain choice**:
- Use lowest gain possible that contains your range
- For sensor 0-3V (with divider): `gain: 4096`
- If max voltage <2V: `gain: 2048` (more resolution)

#### Calibration

Calibration maps ADC voltage → fuel percentage (0-100%).

**Procedure**:

1. **Empty tank**: 
   - Turn on ignition with empty tank
   - Read voltage in debug console
   - Set `voltageEmpty`

2. **Full tank**:
   - Fill tank
   - Turn on ignition
   - Read voltage
   - Set `voltageFull`

**Example**:
```javascript
calibration: {
  voltageEmpty: 0.8,   // Measured: 0.8V at empty tank
  voltageFull: 3.2,    // Measured: 3.2V at full tank
}
```

System will calculate linearly:
```
percentage = ((V_measured - V_empty) / (V_full - V_empty)) × 100
```

#### Sample Rate

| SPS | Usage |
|-----|-------|
| 8 | Maximum precision, slow |
| 16-64 | Balanced |
| 128-250 | Standard (recommended) |
| 475-860 | High speed, less precise |

**Recommendation**: 250 SPS is ideal for fuel sensor (changes slowly).

### Troubleshooting

#### ADS1115 Not Detected

```bash
# Verify I2C module loaded
lsmod | grep i2c

# Verify device present
ls /dev/i2c*
# Expected output: /dev/i2c-1

# Scan bus
sudo i2cdetect -y 1
# If "48" doesn't appear, check:
# - VDD and GND power
# - SDA/SCL wiring
# - Soldering (if custom module)
```

#### Always Reading 0V or Fixed Value

**Causes**:
1. A0 not connected → reads 0V
2. Wrong channel in configuration
3. Gain too low (signal saturates)

**Solutions**:
```javascript
// Verify correct channel
channel: 0,  // 0=A0, 1=A1, 2=A2, 3=A3

// Try higher gain
gain: 6144,  // If signal >4.096V
```

#### Wrong Fuel Percentage

**Cause**: Incorrect calibration

**Solution**: Recalibrate with procedure above

---

## ⚡ Ignition Management

### Purpose

System detects when vehicle ignition is turned on/off and executes automatic actions:
- **Ignition off** → Execute `low-power.sh` script (power saving)
- **Ignition on** → Execute `wake.sh` script (reactivation)

### Hardware Setup

#### Connection

Connect dedicated optocoupler to "key inserted" signal (12V when ignition on):

```
Ignition 12V (KEY) --[1kΩ]--+
                             |
                          PC817 -----> GPIO 21
                             |
                            GND
```

**Logic**:
- Ignition **on** (12V) → GPIO 21 **HIGH**
- Ignition **off** (0V) → GPIO 21 **LOW**

Or vice versa if using active-low optocoupler (set `activeOn: 0`)

### Configuration

#### File: `config/gpio-mapping.js`

```javascript
ignition: {
  enabled: true,                    // Enable ignition management
  pin: 21,                          // Dedicated GPIO 21
  activeOn: 1,                      // 1=active high | 0=active low
  scripts: {
    lowPower: './scripts/low-power.sh',   // When ignition TURNS OFF
    wake: './scripts/wake.sh',             // When ignition TURNS ON
  },
}
```

#### Parameters

- **enabled**: `false` completely disables function
- **pin**: Dedicated GPIO (different from those used for lights)
- **activeOn**:
  - `1` → HIGH = ignition on (active high optocoupler)
  - `0` → LOW = ignition on (active low optocoupler)
- **scripts**: Paths relative to `server/` directory

### Power-Saving Scripts

#### File: `scripts/low-power.sh`

Executed when ignition is **turned off**.

```bash
#!/bin/bash
# Script executed when ignition turns off

logger "PandaOS: Ignition off - Starting power saving mode"

# Turn off HDMI display
vcgencmd display_power 0

# Reduce backlight brightness (if GPIO controlled)
# echo 0 > /sys/class/backlight/rpi_backlight/brightness

# Stop non-essential services
# systemctl stop bluetooth

# Optional: Shutdown after X minutes
# sleep 600 && sudo shutdown -h now &

logger "PandaOS: Power saving mode activated"
```

Make executable:
```bash
chmod +x server/scripts/low-power.sh
```

#### File: `scripts/wake.sh`

Executed when ignition is **turned on**.

```bash
#!/bin/bash
# Script executed when ignition turns on

logger "PandaOS: Ignition on - Reactivating system"

# Turn on HDMI display
vcgencmd display_power 1

# Restore backlight brightness
# echo 255 > /sys/class/backlight/rpi_backlight/brightness

# Restart services
# systemctl start bluetooth

logger "PandaOS: System reactivated"
```

Make executable:
```bash
chmod +x server/scripts/wake.sh
```

### Script Customization

Scripts can execute any bash command, for example:

#### Auto-Shutdown after 10 minutes

In `low-power.sh`:
```bash
# Shutdown if ignition stays off 10 minutes
(sleep 600 && sudo shutdown -h now) &
echo $! > /tmp/pandaos-shutdown.pid
```

In `wake.sh`:
```bash
# Cancel shutdown if ignition turns back on
if [ -f /tmp/pandaos-shutdown.pid ]; then
  kill $(cat /tmp/pandaos-shutdown.pid) 2>/dev/null
  rm /tmp/pandaos-shutdown.pid
fi
```

#### Telegram Notification

```bash
# In low-power.sh
curl -s -X POST "https://api.telegram.org/bot<TOKEN>/sendMessage" \
  -d chat_id=<CHAT_ID> \
  -d text="🚗 Ignition off - PandaOS in standby"
```

---

## 🚨 Troubleshooting

### Server Won't Start

#### Error: "Unsupported platform"

```
❌ ERROR: Essential Raspberry Pi dependencies not available
Unsupported platform: linux x64 - required Linux ARM
```

**Cause**: Running on non-Raspberry Pi system

**Solution**: Run server only on Raspberry Pi, or disable check in `OBDServer.js` (not recommended)

#### Error: "GPIO module not available"

```
GPIO module (onoff) not available
```

**Cause**: `onoff` library not installed or incompatible system

**Solution**:
```bash
cd server
npm install onoff
```

### ELM327 Not Responding

#### Symptom: "Port /dev/ttyUSB0 not found"

**Solution**:
1. Verify ELM327 connected: `lsusb`
2. Verify port: `ls -l /dev/ttyUSB*`
3. Check permissions: `groups` (must include `dialout`)

#### Symptom: "Timeout - No response received"

**Cause**: ELM327 not communicating

**Solution**:
1. Verify baudrate in `OBDCommunicationService.js`
2. Test with minicom: `minicom -D /dev/ttyUSB0 -b 38400`
3. Try reset: Disconnect ELM327, wait 10 seconds, reconnect

### GPIO Not Working

#### Symptom: Lights not detected

**Diagnosis**:
```bash
# Manual GPIO test
gpio -g mode 17 in
gpio -g read 17

# Turn on/off vehicle light and check if value changes
```

**Solutions**:
1. Verify optocoupler wiring
2. Test with LED and resistor instead of optocoupler
3. Check pin number in `gpio-mapping.js` (BCM vs physical)

### Sensors Not Detecting

#### DS18B20: "1-Wire not found"

```bash
# Verify module loaded
lsmod | grep w1

# If missing
sudo modprobe w1-gpio
sudo modprobe w1-therm

# Verify permanent config
grep w1-gpio /boot/config.txt
```

#### ADS1115: "Address 0x48 not responding"

```bash
# Scan I2C bus
sudo i2cdetect -y 1

# If "48" doesn't appear:
# - Check ADS1115 power
# - Verify SDA/SCL not swapped
# - Test with shorter cable
```

---

## 📚 Technical References

### Datasheets and Documentation

- **Raspberry Pi GPIO**: https://pinout.xyz/
- **DS18B20**: https://datasheets.maximintegrated.com/en/ds/DS18B20.pdf
- **ADS1115**: https://www.ti.com/lit/ds/symlink/ads1115.pdf
- **PC817 Optocoupler**: https://www.farnell.com/datasheets/73758.pdf
- **ELM327**: https://www.elmelectronics.com/wp-content/uploads/2017/01/ELM327DS.pdf
- **Fiat Panda 141 - Official Electrical Diagram**: http://www.bunkeringegnere.altervista.org/esplosi/FIAT%20PANDA/panda%20141/1100%20mpi/55%20IMPIANTO%20ELETTRICO%20-%20SCHEMI%20-%20GAMMA%202000.pdf

### Useful Commands

```bash
# GPIO
gpio readall                    # All pins status
gpio -g read <pin>              # Read specific pin

# I2C
sudo i2cdetect -y 1             # Bus scan
sudo i2cget -y 1 0x48 0x00 w    # Read ADS1115 register

# 1-Wire
ls /sys/bus/w1/devices/         # List sensors
cat /sys/bus/w1/devices/28-*/w1_slave  # Read temperature

# Serial
ls -l /dev/tty*                 # List ports
sudo minicom -D /dev/ttyUSB0    # Serial monitor

# Logs
journalctl -u obd-server -f     # Server log (if PM2/systemd)
pm2 logs obd-server             # PM2 log
```

---

**Last update**: v0.9.0  
**Target hardware**: Raspberry Pi 4B + Fiat Panda 141
