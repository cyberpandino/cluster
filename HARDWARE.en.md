[🇬🇧 English](HARDWARE.en.md) | [🇮🇹 Italiano](HARDWARE.md) | [🇩🇪 Deutsch](HARDWARE.de.md)

---

# 🛒 Required Hardware - PandaOS Cluster

Complete list of hardware components needed to build the PandaOS project.

> ⚠️ **WARNING**: This list is provided for informational purposes only. Purchase and install hardware only if you have adequate technical skills. See the [Disclaimer](README.en.md#️-disclaimer).

---

## 📋 Index

1. [Essential Components](#-essential-components)
2. [Optional Sensors](#-optional-sensors)
3. [Display](#-display)
4. [Accessories and Wiring](#-accessories-and-wiring)
5. [Technical Notes](#-technical-notes)

---

## 🔧 Essential Components

### 1. Single Board Computer

**Raspberry Pi 4 Model B** (recommended) or **Raspberry Pi 5**

- **RAM**: 4GB minimum (8GB recommended for optimal performance)
- **Storage**: 32GB microSD minimum (Class 10 or higher)
- **Required interfaces**:
  - 40-pin GPIO
  - USB 2.0/3.0 (for ELM327)
  - HDMI (for display)
  - Ethernet/WiFi (for development)

**Why Raspberry Pi 4B/5**:
- Complete GPIO for optocouplers
- Adequate performance for React + Electron
- Native I2C and 1-Wire support
- Large community and extensive documentation

**Raspberry Pi 5 vs 4B**:
- ✅ **Much superior performance**, especially for 3D Panda model animations
- ✅ Smoother and more responsive interface rendering
- ⚠️ **Trickier power supply**: requires **5 amperes** instead of 3A for 4B
- ⚠️ Need more powerful power supply and more robust wiring
- 💰 Higher cost

**Recommendation**: The **Raspberry Pi 4B with 4GB** is the best performance/ease of installation compromise. The Pi 5 is recommended only if you want ultra-fluid 3D animations and have an adequate power supply.

---

### 2. OBD-II Adapter

**ELM327 USB**

- **Type**: USB (not Bluetooth/WiFi)
- **Chipset**: ELM327 v1.5 or higher
- **Supported protocols**: 
  - ISO 9141-2 (K-Line)
  - ISO 14230-4 (KWP2000)
- **Connector**: USB Type-A
- **Compatibility**: Magneti Marelli IAW 4AF (Fiat Panda 141)

**Technical specifications**:
- Baudrate: 38400 bps
- Voltage: 12V from vehicle
- Interface: Serial USB (/dev/ttyUSB0)

---

### 3. Power Supply

**Display Power Supply**

- **Input voltage**: 12V DC (from vehicle)
- **Output voltage**: 12V DC
- **Current**: Minimum 2A (depends on display)
- **Connector**: DC Jack or direct wires
- **Protections**: Overvoltage, short circuit

**Raspberry Pi Power Supply** (optional if powered via USB)

- **Voltage**: 5V DC
- **Current**: 3A minimum (Raspberry Pi 4B/5)
- **Connector**: USB-C
- **Type**: Step-down DC-DC converter 12V→5V for vehicle

---

## 📺 Display

### LCD Panel

**Display Specifications Used in Project**:

- **Size**: 12.6 inches
- **Resolution**: 1920×480 pixels
- **Aspect Ratio**: 4:1 (ultra-wide)
- **Panel type**: IPS
- **Refresh rate**: 60Hz
- **Interface**: HDMI
- **Curvature**: Flat (not curved)
- **Brightness**: Adequate for automotive use
- **Voltage**: 12V DC

**Important Characteristics**:
- Ultra-wide format ideal for automotive dashboard
- IPS for wide viewing angles
- Resolution optimized for instrumentation (1920×480)

**Alternatives**:
- 10-14 inch displays with 1920×480 or similar resolution
- Automotive panels certified for extended temperature
- Optional touch screen (not used in base project)

---

## 🔌 Optional Sensors

### 1. External Temperature Sensor

**DS18B20 - Digital 1-Wire Temperature Sensor**

- **Type**: Sealed (waterproof) digital sensor
- **Temperature range**: -55°C to +125°C
- **Accuracy**: ±0.5°C
- **Protocol**: 1-Wire (Dallas)
- **Power supply**: 3.0V - 5.5V
- **Package**: TO-92 or waterproof version with cable
- **Cable**: 1-2 meters (longer for external installation)

**Additional components**:
- 4.7kΩ pull-up resistor (essential)

---

### 2. Fuel Level Sensor

**ADS1115 - 16-bit I2C ADC**

- **Type**: Analog-to-Digital Converter
- **Resolution**: 16 bit
- **Channels**: 4 single-ended or 2 differential
- **Voltage range**: ±0.256V to ±6.144V (programmable)
- **Interface**: I2C
- **Address**: 0x48 (default)
- **Sample rate**: 8-860 SPS
- **Power supply**: 2.0V - 5.5V

**Additional components**:
- 2× Resistors for voltage divider (e.g. 100kΩ + 33kΩ)
- Wires to connect to original fuel sensor

---

### 3. Optocouplers for Warning Lights

**PC817 or equivalent**

- **Quantity**: 13 units (one for each warning light to detect)
- **Type**: Transistor photocoupler
- **Isolation voltage**: 5000V
- **LED current**: 20mA (typical)
- **Output voltage**: 5V compatible with Raspberry GPIO
- **Package**: DIP-4

**Additional components**:
- 13× LED limiting resistors (1kΩ - 2.2kΩ)
- Breadboard or custom PCB for mounting

> 📘 **Official Electrical Schematic**: To identify the correct wires for warning lights on the vehicle, consult the [Fiat Panda 141 Electrical Schematic - Workshop Manual](http://www.bunkeringegnere.altervista.org/esplosi/FIAT%20PANDA/panda%20141/1100%20mpi/55%20IMPIANTO%20ELETTRICO%20-%20SCHEMI%20-%20GAMMA%202000.pdf). Includes all electrical system schematics for 2000 range with wire color codes.

---

## 🔗 Accessories and Wiring

### Cables and Connectors

- **HDMI cable**: For Raspberry Pi → Display connection
- **USB Type-A cable**: For ELM327 → Raspberry Pi
- **Dupont wires**: Male-Female for GPIO (set of 40 pieces)
- **Ethernet cable**: For initial setup (optional)

### Electrical Material

- **Automotive wires**: 0.5-1.0 mm² for 12V connections
- **Faston connectors**: For vehicle warning light connections
- **Heat shrink tubing**: Connection protection
- **Cable ties**: Cable organization

### Protections

**⚡ FUNDAMENTAL: Protect EVERYTHING with separate fuses**

- **Raspberry Pi Fuse**:
  - Value: **2A** for Raspberry Pi 4B (3A for Pi 5)
  - Type: Automotive blade fuse or inline
  - Position: On 12V wire before DC-DC converter
  - Reason: Protects from short circuits in converter or Raspberry

- **Display Fuse**:
  - Value: **3A** (verify specific display consumption)
  - Type: Automotive blade fuse or inline
  - Position: On 12V display power wire
  - Reason: Protects from short circuits in display or wiring

- **GPIO Circuits Fuse** (optional but recommended):
  - Value: **1A**
  - Type: Inline fuse or on breadboard
  - Position: On common 12V for optocouplers
  - Reason: Protects optocouplers from wiring errors

**Why separate fuses**:
- ✅ Fault isolation: a display problem won't turn off the Raspberry
- ✅ Easier diagnosis: understand immediately which circuit has problems
- ✅ Targeted protection: each fuse sized for its load
- ✅ Safety: avoid burning components or worse, causing fires

**Required components**:
- **Fuse holders**: Inline for 12V circuits (3 units minimum)
- **Spare fuses**: Always have some extras
- **Diodes**: 1N4001 or similar for reverse polarity protection (one per 12V line)

---

## 🧰 Required Tools

### For Hardware Installation

- Digital multimeter
- Soldering iron (for PCB/permanent connections)
- Wire strippers / pliers
- Screwdriver set
- Circuit tester

### For Software Development

- Development computer (Mac/Windows/Linux)
- USB microSD card reader
- Ethernet cable (initial Raspberry setup)

---

## 💰 Indicative Cost Estimate

| Component | Indicative Cost |
|-----------|-----------------|
| Raspberry Pi 4B (4GB) | €50-70 |
| ELM327 USB | €15-30 |
| LCD Display 12.6" | €80-150 |
| DC-DC Converter 12V→5V 3A | €8-15 |
| DS18B20 Waterproof | €3-8 |
| ADS1115 | €5-10 |
| PC817 (10pcs set) | €2-5 |
| Fuses + Fuse Holders (×3) | €5-10 |
| Cables and accessories | €20-40 |
| **ESTIMATED TOTAL** | **€188-338** |

> 💡 Prices are indicative and vary based on suppliers, availability, and component quality.

---

## 📦 Recommended Kits

To simplify purchasing, consider these kits:

### Base Kit (Software Development Only)
- Raspberry Pi 4B/5
- USB-C power supply
- 32GB microSD
- Raspberry Pi case
- ELM327 USB

**For**: Development and testing in mock mode

### Complete Kit (Production)
- Everything from Base Kit
- LCD Display 12.6" 1920×480
- DC-DC Converter 12V→5V (3A for Pi 4B, 5A for Pi 5)
- **3× Automotive fuses** (2A, 3A, 1A)
- **3× Inline fuse holders**
- PC817 optocoupler set
- Assorted resistors
- Cables and connectors
- Protection diodes (1N4001 or similar)

**For**: Complete vehicle installation

### Sensor Kit (Optional)
- DS18B20 waterproof
- ADS1115
- Resistors (4.7kΩ, 100kΩ, 33kΩ)

**For**: Advanced features (temperature, fuel)

---

## 🔍 Technical Notes

### Display Compatibility

The project is optimized for **ultra-wide 1920×480** displays but can be adapted to:
- 1280×480 (lower resolution)
- 1920×720 (standard 16:9, requires UI adaptation)
- 1024×600 (7" tablet, requires interface resize)

### Raspberry Pi: 4B vs 5

| Feature | Raspberry Pi 4B | Raspberry Pi 5 |
|---------|-----------------|----------------|
| CPU | Quad-core ARM Cortex-A72 1.5GHz | Quad-core ARM Cortex-A76 2.4GHz |
| RAM | 2/4/8 GB | 4/8 GB |
| GPIO | 40 pin | 40 pin |
| Performance | Adequate | Better |
| Consumption | Lower | Higher |
| Cost | Lower | Higher |
| **Recommendation** | ✅ Great price/performance ratio | ✅ For maximum performance |

**Recommendation**: Raspberry Pi 4B with 4GB is more than sufficient for the project.

### Vehicle Power Supply

**Recommended Power Supply Schematic**:

```
Vehicle 12V Battery
    │
    ├─→ [2A FUSE] → DC-DC Converter 12V→5V → Raspberry Pi 4B
    │
    ├─→ [3A FUSE] → LCD Display 12V
    │
    └─→ [1A FUSE] → Optocoupler Circuit 12V
```

**Specifications**:
- Input: 12V auto (tolerance 9-16V to compensate for engine variations)
- Raspberry Output: 5V 3A USB-C (5A for Pi 5)
- Display Output: 12V 2-3A (verify display specifications)

**Recommended DC-DC Converter**:
- Input: 9-30V DC
- Output: 5V 3A (or 5A for Pi 5)
- Type: Automotive step-down buck converter
- Protections: Overvoltage, overcurrent, short circuit, reverse polarity
- Efficiency: >85%

### Practical Power Supply Installation

#### Where to Get 12V from Panda 141

You have several options to get 12V power:

**Option 1: From Fuse Box (RECOMMENDED)**
- **Location**: Under hood, left side near battery
- **Recommended fuses to tap**:
  - **F15 (10A)** - Instrument cluster: Activates with key ON, perfect to synchronize system startup
  - **F16 (7.5A)** - Accessories: Always powered, useful if you want system always active
  - **F17 (15A)** - Cigarette lighter: Convenient for testing, always powered
- **Advantages**: Fuses already present, easy tapping with Y faston connectors, existing protection
- **How to**: Use Y faston connectors to tap without cutting original wires

**Option 2: Direct Battery (for permanent installations)**
- **Location**: Under hood, battery positive terminal
- **Advantages**: Always available power, no interference with existing fuses
- **Disadvantages**: Requires dedicated inline fuse immediately after battery terminal
- **How to**: 
  1. Connect red wire directly to battery positive pole
  2. Insert 5A fuse IMMEDIATELY inline (max 10cm from terminal)
  3. Run protected wire to dashboard

**Option 3: Behind Instrument Cluster (cleaner)**
- **Location**: Behind cluster, distribution junction
- **Advantages**: Shorter wires, cleaner installation, already in dashboard
- **How to**: Consult [Official Electrical Schematic](http://www.bunkeringegnere.altervista.org/esplosi/FIAT%20PANDA/panda%20141/1100%20mpi/55%20IMPIANTO%20ELETTRICO%20-%20SCHEMI%20-%20GAMMA%202000.pdf) to identify correct wires

⚠️ **IMPORTANT**: Regardless of option chosen, ALWAYS use separate fuses for each component!

#### Physical Fuse Layout

**Recommended Vehicle Layout**:

```
12V tap point (battery or fuses)
    │
    ├─── Red 1.0mm² wire (30cm) ──→ [INLINE 2A FUSE HOLDER] ──→ DC-DC Converter
    │                                                                     │
    ├─── Red 1.0mm² wire (50cm) ──→ [INLINE 3A FUSE HOLDER] ────────────┼──→ Display 12V
    │                                                                     │
    └─── Red 0.5mm² wire (20cm) ──→ [INLINE 1A FUSE HOLDER] ────────────┼──→ Optocouplers
                                                                          │
                                                                          ▼
                                                                    DC-DC Converter
                                                                          │
                                                                          ├──→ USB-C Raspberry Pi 5V 3A
```

**Physical positioning**:
- **Raspberry Fuse**: Near DC-DC converter (under dashboard)
- **Display Fuse**: Behind display itself (facilitates replacement)
- **Optocoupler Fuse**: Near optocoupler breadboard/PCB
- **Ground (GND)**: Connect to chassis ground with eyelet (dashboard frame screw)

#### DC-DC Converter Wiring

**Installation Steps**:

1. **Positioning**: Fix converter under dashboard with cable ties or velcro
2. **12V Input**:
   - Connect red wire from 2A fuse → converter INPUT+ terminal
   - Connect black wire from chassis ground → converter INPUT- terminal
3. **5V Output**:
   - Use USB-C cable from converter OUTPUT+ / OUTPUT- → Raspberry Pi
   - Or solder directly to USB-C pins (more stable)
4. **Voltage adjustment**:
   - Before connecting Raspberry, verify output with multimeter
   - Adjust trimmer on converter until reading exactly **5.0V - 5.2V**
   - ⚠️ DO NOT exceed 5.3V (would damage Raspberry!)
5. **Load test**:
   - Connect Raspberry turned off
   - Measure voltage under load (must stay 5.0-5.2V)
   - If drops below 4.8V, converter is undersized

#### Recommended Wire Gauge

| Component | Max Current | Length | Wire Gauge |
|-----------|-------------|--------|------------|
| Raspberry Pi (12V→5V) | 2A @ 12V | <1m | 1.0 mm² |
| Display 12V | 3A @ 12V | <1m | 1.0 mm² |
| Optocouplers | 0.5A @ 12V | <0.5m | 0.5 mm² |
| Ground (common GND) | 5A total | <1m | 1.5 mm² |

**Note**: Oversized sections compared to minimum to compensate for losses and automotive vibrations.

#### Ignition ON/OFF Connection

To synchronize system startup with ignition:

**Method 1: Tap F15 (Instrument Cluster)**
- Pro: Turns on/off automatically with key
- Pro: Doesn't drain battery when vehicle off
- Con: No delay, immediate shutdown when key removed

**Method 2: Direct Battery + GPIO Ignition (RECOMMENDED)**
- Pro: Continuous power, software-controlled shutdown
- Pro: Can implement delayed shutdown timer
- Pro: Avoid SD card corruption with controlled shutdown
- Requires: GPIO 21 connected to "key inserted" signal (see CONFIGURAZIONE_SERVER.en.md)

**Method 2 Procedure**:
1. Power Raspberry from direct battery (always ON)
2. Connect GPIO 21 to "key inserted" warning light via optocoupler
3. `low-power.sh` script executes when key removed
4. Implement delayed shutdown of 5-10 minutes (see CONFIGURAZIONE_SERVER.en.md § Ignition)

#### Pre-Startup Checklist

Before powering the system for the first time:

- [ ] Verify polarity with multimeter (red=12V, black=0V)
- [ ] Check battery voltage: must be 12.0-14.5V
- [ ] Verify ground (GND) continuity to chassis
- [ ] Fuses correctly inserted (2A, 3A, 1A)
- [ ] DC-DC converter adjusted to 5.0-5.2V (no load)
- [ ] Well-insulated wires (heat shrink tubing)
- [ ] Solid connections (not loose)
- [ ] Raspberry Pi NOT yet connected (test converter empty first)

#### Power Supply Troubleshooting

**Raspberry won't turn on**:
1. Verify converter OUTPUT voltage with multimeter (must be 5.0-5.2V)
2. Check Raspberry power LED (solid red = powered)
3. Verify 2A fuse not blown
4. Check USB-C connection well inserted

**Raspberry randomly reboots**:
- Cause: Unstable voltage or spikes under load
- Solution: Use 1000µF 16V electrolytic capacitor on converter output
- Or: Undersized converter, switch to 5A model

**Display won't turn on**:
1. Verify 12V power at display pins
2. Check 3A fuse
3. Verify HDMI cable connected to Raspberry

**System drains battery**:
- Verify total current with ignition off: <50mA is acceptable
- If >200mA: dispersion problem, check wiring
- Use manual switch or tap F15 (turns off with key)

⚠️ **IMPORTANT**: Always use fuses BEFORE each component, not after!

### GPIO Protection

⚠️ **IMPORTANT**: Raspberry Pi GPIOs are 3.3V and **NOT** 5V tolerant. Optocouplers serve precisely to protect GPIOs from vehicle 12V.

**Protection schematic**:
```
12V vehicle → Resistor → LED optocoupler
                          ↓ (optical isolation)
Optocoupler transistor → GPIO (3.3V)
```

---

## 📚 References

### Component Datasheets

- **Raspberry Pi 4B**: [raspberrypi.com/products/raspberry-pi-4-model-b](https://www.raspberrypi.com/products/raspberry-pi-4-model-b/)
- **DS18B20**: Maxim Integrated DS18B20 Datasheet
- **ADS1115**: Texas Instruments ADS1115 Datasheet
- **PC817**: Sharp PC817 Datasheet
- **ELM327**: ELM Electronics ELM327 Datasheet

### Vehicle Electrical Schematic

- **Fiat Panda 141 - 2000 Range Electrical System**: [Workshop Manual - Electrical Schematics](http://www.bunkeringegnere.altervista.org/esplosi/FIAT%20PANDA/panda%20141/1100%20mpi/55%20IMPIANTO%20ELETTRICO%20-%20SCHEMI%20-%20GAMMA%202000.pdf)
  - Complete 1100 MPI electrical system schematic
  - Wire color codes (e.g. R=Red, N=Black, BN=White-Black, etc.)
  - Warning light and connector positions on instrument cluster
  - Distribution junction and fuse identification
  - **Essential to identify correct wires to connect to optocouplers**

### Configuration Guides

- [Raspberry Pi Configuration](server/CONFIGURAZIONE_SERVER.en.md#-raspberry-pi-configuration)
- [GPIO Setup](server/CONFIGURAZIONE_SERVER.en.md#-gpio-configuration)
- [Sensors](server/CONFIGURAZIONE_SERVER.en.md#-ds18b20-temperature-sensor)

---

## ⚠️ Hardware Disclaimer

Purchase and installation of hardware is under your complete responsibility. Make sure to:

- ✅ Have adequate electrical/electronic skills
- ✅ **ALWAYS use separate fuses** for each powered component
- ✅ Use wires with adequate gauge for current to carry
- ✅ Respect local automotive safety regulations
- ✅ Don't compromise vehicle safety
- ✅ Test everything on bench before installation
- ✅ Verify polarity before connecting power supplies
- ✅ NEVER bypass protections (fuses, diodes)

See the [Complete Disclaimer](README.en.md#️-disclaimer) for more information.

---

**Last update**: November 2025  
**Tested hardware**: Fiat Panda 141 1100 mpi
