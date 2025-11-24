# 🛒 Hardware Necessario - PandaOS Cluster

Lista completa dei componenti hardware necessari per realizzare il progetto PandaOS.

> ⚠️ **ATTENZIONE**: Questa lista è fornita solo a scopo informativo. Acquista e installa hardware solo se hai competenze tecniche adeguate. Vedi il [Disclaimer](README.md#️-disclaimer).

---

## 📋 Indice

1. [Componenti Essenziali](#-componenti-essenziali)
2. [Sensori Opzionali](#-sensori-opzionali)
3. [Display](#-display)
4. [Accessori e Cablaggio](#-accessori-e-cablaggio)
5. [Note Tecniche](#-note-tecniche)

---

## 🔧 Componenti Essenziali

### 1. Single Board Computer

**Raspberry Pi 4 Model B** (consigliato) o **Raspberry Pi 5**

- **RAM**: 4GB minimo (8GB consigliato per performance ottimali)
- **Storage**: MicroSD da 32GB minimo (Classe 10 o superiore)
- **Interfacce richieste**:
  - GPIO 40 pin
  - USB 2.0/3.0 (per ELM327)
  - HDMI (per display)
  - Ethernet/WiFi (per sviluppo)

**Perché Raspberry Pi 4B/5**:
- GPIO completi per optoaccoppiatori
- Performance adeguate per React + Electron
- Supporto I2C e 1-Wire nativi
- Comunità ampia e documentazione estesa

**Raspberry Pi 5 vs 4B**:
- ✅ **Performance molto superiori**, specialmente per le animazioni 3D del modello Panda
- ✅ Rendering più fluido e reattivo dell'interfaccia
- ⚠️ **Alimentazione più rognosa**: richiede **5 ampere** invece dei 3A del 4B
- ⚠️ Necessario alimentatore più potente e cablaggio più robusto
- 💰 Costo superiore

**Raccomandazione**: Il **Raspberry Pi 4B con 4GB** è il miglior compromesso prestazioni/facilità di installazione. Il Pi 5 è consigliato solo se vuoi animazioni 3D ultra-fluide e hai un alimentatore adeguato.

---

### 2. Adattatore OBD-II

**ELM327 USB**

- **Tipo**: USB (non Bluetooth/WiFi)
- **Chipset**: ELM327 v1.5 o superiore
- **Protocolli supportati**: 
  - ISO 9141-2 (K-Line)
  - ISO 14230-4 (KWP2000)
- **Connettore**: USB Type-A
- **Compatibilità**: Magneti Marelli IAW 4AF (Fiat Panda 141)

**Specifiche tecniche**:
- Baudrate: 38400 bps
- Tensione: 12V da veicolo
- Interfaccia: Serial USB (/dev/ttyUSB0)

---

### 3. Alimentatore

**Alimentatore per Display**

- **Tensione input**: 12V DC (da veicolo)
- **Tensione output**: 12V DC
- **Corrente**: Minimo 2A (dipende dal display)
- **Connettore**: DC Jack o cavi diretti
- **Protezioni**: Sovratensione, cortocircuito

**Alimentatore Raspberry Pi** (opzionale se alimentato via USB)

- **Tensione**: 5V DC
- **Corrente**: 3A minimo (Raspberry Pi 4B/5)
- **Connettore**: USB-C
- **Tipo**: Step-down DC-DC converter 12V→5V per veicolo

---

## 📺 Display

### Pannello LCD

**Specifiche Display Utilizzato nel Progetto**:

- **Dimensioni**: 12.6 pollici
- **Risoluzione**: 1920×480 pixel
- **Aspect Ratio**: 4:1 (ultra-wide)
- **Tipo pannello**: IPS
- **Refresh rate**: 60Hz
- **Interfaccia**: HDMI
- **Curvatura**: Piatto (non curvo)
- **Luminosità**: Adeguata per uso automotive
- **Tensione**: 12V DC

**Caratteristiche Importanti**:
- Formato ultra-wide ideale per dashboard automotive
- IPS per ampi angoli di visione
- Risoluzione ottimizzata per strumentazione (1920×480)

**Alternative**:
- Display 10-14 pollici con risoluzione 1920×480 o simili
- Pannelli automotive certificati per temperatura estesa
- Touch screen opzionale (non utilizzato nel progetto base)

---

## 🔌 Sensori Opzionali

### 1. Sensore Temperatura Esterna

**DS18B20 - Sensore Temperatura Digitale 1-Wire**

- **Tipo**: Sensore digitale sigillato (waterproof)
- **Range temperatura**: -55°C a +125°C
- **Precisione**: ±0.5°C
- **Protocollo**: 1-Wire (Dallas)
- **Alimentazione**: 3.0V - 5.5V
- **Package**: TO-92 o versione waterproof con cavo
- **Cavo**: 1-2 metri (più lungo possibile per installazione esterna)

**Componenti aggiuntivi**:
- Resistenza pull-up 4.7kΩ (essenziale)

---

### 2. Sensore Livello Carburante

**ADS1115 - ADC 16-bit I2C**

- **Tipo**: Convertitore Analogico-Digitale
- **Risoluzione**: 16 bit
- **Canali**: 4 single-ended o 2 differenziali
- **Range tensione**: ±0.256V a ±6.144V (programmabile)
- **Interfaccia**: I2C
- **Indirizzo**: 0x48 (default)
- **Sample rate**: 8-860 SPS
- **Alimentazione**: 2.0V - 5.5V

**Componenti aggiuntivi**:
- 2× Resistenze per partitore (es. 100kΩ + 33kΩ)
- Cavi per collegamento al sensore carburante originale

---

### 3. Optoaccoppiatori per Spie

**PC817 o equivalenti**

- **Quantità**: 13 unità (uno per ogni spia da rilevare)
- **Tipo**: Fotoaccoppiatore a transistor
- **Tensione isolamento**: 5000V
- **Corrente LED**: 20mA (tipico)
- **Tensione uscita**: 5V compatibile con GPIO Raspberry
- **Package**: DIP-4

**Componenti aggiuntivi**:
- 13× Resistenze limitatrici LED (1kΩ - 2.2kΩ)
- Breadboard o PCB custom per montaggio

> 📘 **Schema Elettrico Ufficiale**: Per identificare i cavi corretti delle spie sul veicolo, consulta lo [Schema Elettrico Fiat Panda 141 - Manuale Officina](http://www.bunkeringegnere.altervista.org/esplosi/FIAT%20PANDA/panda%20141/1100%20mpi/55%20IMPIANTO%20ELETTRICO%20-%20SCHEMI%20-%20GAMMA%202000.pdf). Include tutti gli schemi dell'impianto elettrico gamma 2000 con codici colore cavi.

---

## 🔗 Accessori e Cablaggio

### Cavi e Connettori

- **Cavo HDMI**: Per collegamento Raspberry Pi → Display
- **Cavo USB Type-A**: Per ELM327 → Raspberry Pi
- **Cavi Dupont**: Maschio-Femmina per GPIO (set da 40 pezzi)
- **Cavo Ethernet**: Per configurazione iniziale (opzionale)

### Materiale Elettrico

- **Cavetti automotive**: 0.5-1.0 mm² per collegamenti 12V
- **Connettori faston**: Per connessioni spie veicolo
- **Guaina termorestringente**: Protezione collegamenti
- **Fascette**: Organizzazione cavi

### Protezioni

**⚡ FONDAMENTALE: Proteggere TUTTO con fusibili separati**

- **Fusibile Raspberry Pi**:
  - Valore: **2A** per Raspberry Pi 4B (3A per Pi 5)
  - Tipo: Fusibile blade automotive o inline
  - Posizione: Sul cavo 12V prima del convertitore DC-DC
  - Motivo: Protegge da cortocircuiti nel convertitore o Raspberry

- **Fusibile Display**:
  - Valore: **3A** (verifica consumo specifico del tuo display)
  - Tipo: Fusibile blade automotive o inline
  - Posizione: Sul cavo 12V di alimentazione display
  - Motivo: Protegge da cortocircuiti nel display o cablaggio

- **Fusibile Circuiti GPIO** (opzionale ma consigliato):
  - Valore: **1A**
  - Tipo: Fusibile inline o su breadboard
  - Posizione: Sul 12V comune agli optoaccoppiatori
  - Motivo: Protegge optoaccoppiatori da errori di cablaggio

**Perché fusibili separati**:
- ✅ Isolamento guasti: un problema al display non spegne il Raspberry
- ✅ Diagnosi più facile: capisci subito quale circuito ha problemi
- ✅ Protezione mirata: ogni fusibile dimensionato sul suo carico
- ✅ Sicurezza: eviti di bruciare componenti o peggio, causare incendi

**Componenti necessari**:
- **Portafusibili**: Inline per circuiti 12V (3 unità minimo)
- **Fusibili di ricambio**: Averne sempre qualcuno in più
- **Diodi**: 1N4001 o simili per protezione inversione polarità (uno per linea 12V)

---

## 🧰 Strumenti Necessari

### Per Installazione Hardware

- Multimetro digitale
- Saldatore a stagno (per PCB/connessioni permanenti)
- Spelafili / pinze
- Cacciaviti set
- Tester circuiti

### Per Sviluppo Software

- Computer di sviluppo (Mac/Windows/Linux)
- Lettore microSD USB
- Cavo Ethernet (setup iniziale Raspberry)

---

## 💰 Stima Costi Indicativi

| Componente | Costo Indicativo |
|------------|------------------|
| Raspberry Pi 4B (4GB) | €50-70 |
| ELM327 USB | €15-30 |
| Display LCD 12.6" | €80-150 |
| Convertitore DC-DC 12V→5V 3A | €8-15 |
| DS18B20 Waterproof | €3-8 |
| ADS1115 | €5-10 |
| PC817 (set 10pz) | €2-5 |
| Fusibili + Portafusibili (×3) | €5-10 |
| Cavi e accessori | €20-40 |
| **TOTALE STIMATO** | **€188-338** |

> 💡 I prezzi sono indicativi e variano in base a fornitori, disponibilità e qualità dei componenti.

---

## 📦 Kit Consigliati

Per semplificare l'acquisto, considera questi kit:

### Kit Base (Solo Sviluppo Software)
- Raspberry Pi 4B/5
- Alimentatore USB-C
- MicroSD 32GB
- Case Raspberry Pi
- ELM327 USB

**Per**: Sviluppo e test in modalità mock

### Kit Completo (Produzione)
- Tutto del Kit Base
- Display LCD 12.6" 1920×480
- Convertitore DC-DC 12V→5V (3A per Pi 4B, 5A per Pi 5)
- **3× Fusibili automotive** (2A, 3A, 1A)
- **3× Portafusibili inline**
- Set optoaccoppiatori PC817
- Resistenze assortite
- Cavi e connettori
- Diodi protezione (1N4001 o simili)

**Per**: Installazione completa su veicolo

### Kit Sensori (Opzionale)
- DS18B20 waterproof
- ADS1115
- Resistenze (4.7kΩ, 100kΩ, 33kΩ)

**Per**: Funzionalità avanzate (temperatura, carburante)

---

## 🔍 Note Tecniche

### Compatibilità Display

Il progetto è ottimizzato per display **ultra-wide 1920×480** ma può essere adattato a:
- 1280×480 (minor risoluzione)
- 1920×720 (16:9 standard, richiede adattamento UI)
- 1024×600 (7" tablet, richiede resize interfaccia)

### Raspberry Pi: 4B vs 5

| Caratteristica | Raspberry Pi 4B | Raspberry Pi 5 |
|----------------|-----------------|----------------|
| CPU | Quad-core ARM Cortex-A72 1.5GHz | Quad-core ARM Cortex-A76 2.4GHz |
| RAM | 2/4/8 GB | 4/8 GB |
| GPIO | 40 pin | 40 pin |
| Performance | Adeguate | Migliori |
| Consumo | Inferiore | Superiore |
| Costo | Inferiore | Superiore |
| **Consiglio** | ✅ Ottimo rapporto qualità/prezzo | ✅ Per performance massime |

**Raccomandazione**: Raspberry Pi 4B con 4GB è più che sufficiente per il progetto.

### Alimentazione da Veicolo

**Schema Alimentazione Consigliato**:

```
Batteria 12V Veicolo
    │
    ├─→ [FUSIBILE 2A] → Convertitore DC-DC 12V→5V → Raspberry Pi 4B
    │
    ├─→ [FUSIBILE 3A] → Display LCD 12V
    │
    └─→ [FUSIBILE 1A] → Circuito Optoaccoppiatori 12V
```

**Specifiche**:
- Input: 12V auto (tolleranza 9-16V per compensare variazioni motore)
- Output Raspberry: 5V 3A USB-C (5A per Pi 5)
- Output Display: 12V 2-3A (verifica specifiche display)

**Convertitore DC-DC Consigliato**:
- Input: 9-30V DC
- Output: 5V 3A (o 5A per Pi 5)
- Tipo: Step-down buck converter automotive
- Protezioni: Sovratensione, sovracorrente, cortocircuito, inversione polarità
- Efficienza: >85%

### Installazione Pratica Alimentazione

#### Da Dove Prelevare i 12V nella Panda 141

Hai diverse opzioni per prelevare l'alimentazione 12V:

**Opzione 1: Dalla Scatola Fusibili (CONSIGLIATA)**
- **Posizione**: Sotto il cofano, lato sinistro vicino alla batteria
- **Fusibili consigliati da cui derivare**:
  - **F15 (10A)** - Quadro strumenti: Si attiva con chiave ON, perfetto per sincronizzare accensione sistema
  - **F16 (7.5A)** - Accessori: Sempre alimentato, utile se vuoi il sistema sempre attivo
  - **F17 (15A)** - Presa accendisigari: Comodo per test, sempre alimentato
- **Vantaggi**: Fusibili già presenti, facile derivazione con spinotti faston, protezione esistente
- **Come fare**: Usa spinotti faston a "Y" per derivare senza tagliare cavi originali

**Opzione 2: Batteria Diretta (per installazioni permanenti)**
- **Posizione**: Sotto il cofano, terminale positivo batteria
- **Vantaggi**: Alimentazione sempre disponibile, nessuna interferenza con fusibili esistenti
- **Svantaggi**: Richiede fusibile dedicato in linea subito dopo il terminale batteria
- **Come fare**: 
  1. Collegare cavo rosso direttamente al polo positivo batteria
  2. Inserire IMMEDIATAMENTE fusibile 5A in linea (max 10cm dal terminale)
  3. Far passare il cavo protetto fino al cruscotto

**Opzione 3: Dietro il Quadro Strumenti (più pulita)**
- **Posizione**: Dietro il quadro, centralina di derivazione
- **Vantaggi**: Cavi più corti, installazione più pulita, già nel cruscotto
- **Come fare**: Consulta lo [Schema Elettrico Ufficiale](http://www.bunkeringegnere.altervista.org/esplosi/FIAT%20PANDA/panda%20141/1100%20mpi/55%20IMPIANTO%20ELETTRICO%20-%20SCHEMI%20-%20GAMMA%202000.pdf) per identificare i cavi corretti

⚠️ **IMPORTANTE**: Indipendentemente dall'opzione scelta, usa SEMPRE fusibili separati per ogni componente!

#### Disposizione Fisica dei Fusibili

**Layout Consigliato nel Veicolo**:

```
Punto di prelievo 12V (batteria o fusibili)
    │
    ├─── Cavo rosso 1.0mm² (30cm) ──→ [PORTAFUSIBILI INLINE 2A] ──→ Convertitore DC-DC
    │                                                                       │
    ├─── Cavo rosso 1.0mm² (50cm) ──→ [PORTAFUSIBILI INLINE 3A] ──────────┼──→ Display 12V
    │                                                                       │
    └─── Cavo rosso 0.5mm² (20cm) ──→ [PORTAFUSIBILI INLINE 1A] ──────────┼──→ Optoaccoppiatori
                                                                            │
                                                                            ▼
                                                                      Convertitore DC-DC
                                                                            │
                                                                            ├──→ USB-C Raspberry Pi 5V 3A
```

**Posizionamento fisico**:
- **Fusibile Raspberry**: Vicino al convertitore DC-DC (sotto cruscotto)
- **Fusibile Display**: Dietro il display stesso (facilita sostituzione)
- **Fusibile Optoaccoppiatori**: Vicino alla breadboard/PCB optoaccoppiatori
- **Massa (GND)**: Collegare a massa carrozzeria con occhiello (vite telaio cruscotto)

#### Cablaggio Convertitore DC-DC

**Passaggi Installazione**:

1. **Posizionamento**: Fissa il convertitore sotto il cruscotto con fascette o velcro
2. **Input 12V**:
   - Collega cavo rosso da fusibile 2A → terminale INPUT+ convertitore
   - Collega cavo nero da massa carrozzeria → terminale INPUT- convertitore
3. **Output 5V**:
   - Usa cavo USB-C da OUTPUT+ / OUTPUT- convertitore → Raspberry Pi
   - Oppure saldare direttamente ai pin USB-C (più stabile)
4. **Regolazione tensione**:
   - Prima di collegare il Raspberry, verifica con multimetro l'output
   - Regola trimmer sul convertitore fino a leggere esattamente **5.0V - 5.2V**
   - ⚠️ NON superare 5.3V (danneggerebbe il Raspberry!)
5. **Test carico**:
   - Collega il Raspberry spento
   - Misura tensione sotto carico (deve rimanere 5.0-5.2V)
   - Se scende sotto 4.8V, il convertitore è sottodimensionato

#### Sezione Cavi Consigliata

| Componente | Corrente Max | Lunghezza | Sezione Cavo |
|------------|--------------|-----------|--------------|
| Raspberry Pi (12V→5V) | 2A @ 12V | <1m | 1.0 mm² |
| Display 12V | 3A @ 12V | <1m | 1.0 mm² |
| Optoaccoppiatori | 0.5A @ 12V | <0.5m | 0.5 mm² |
| Massa (GND comune) | 5A totali | <1m | 1.5 mm² |

**Nota**: Sezioni maggiorate rispetto al minimo per compensare perdite e vibrazioni automotive.

#### Connessione a Quadro ON/OFF

Per sincronizzare l'accensione del sistema con il quadro:

**Metodo 1: Preleva da F15 (Quadro Strumenti)**
- Pro: Si accende/spegne automaticamente con chiave
- Pro: Non scarica batteria quando veicolo spento
- Contro: Nessun ritardo, spegnimento immediato alla rimozione chiave

**Metodo 2: Batteria Diretta + GPIO Ignition (CONSIGLIATO)**
- Pro: Alimentazione continua, controllo software dello shutdown
- Pro: Puoi implementare timer di spegnimento ritardato
- Pro: Eviti corruzione SD card con shutdown controllato
- Richiede: GPIO 21 collegato a segnale "chiave inserita" (vedi CONFIGURAZIONE_SERVER.md)

**Procedura Metodo 2**:
1. Alimenta Raspberry da batteria diretta (sempre ON)
2. Collega GPIO 21 a spia "chiave inserita" tramite optoaccoppiatore
3. Script `low-power.sh` viene eseguito quando togli la chiave
4. Implementa shutdown ritardato di 5-10 minuti (vedi CONFIGURAZIONE_SERVER.md § Ignition)

#### Checklist Pre-Accensione

Prima di accendere il sistema per la prima volta:

- [ ] Verifica polarità con multimetro (rosso=12V, nero=0V)
- [ ] Controlla tensione batteria: deve essere 12.0-14.5V
- [ ] Verifica continuità massa (GND) verso carrozzeria
- [ ] Fusibili correttamente inseriti (2A, 3A, 1A)
- [ ] Convertitore DC-DC regolato a 5.0-5.2V (senza carico)
- [ ] Cavi ben isolati (guaina termorestringente)
- [ ] Connessioni salde (non volanti)
- [ ] Raspberry Pi NON ancora collegato (test convertitore a vuoto prima)

#### Troubleshooting Alimentazione

**Raspberry non si accende**:
1. Verifica tensione OUTPUT convertitore con multimetro (deve essere 5.0-5.2V)
2. Controlla LED alimentazione Raspberry (rosso fisso = alimentato)
3. Verifica fusibile 2A non bruciato
4. Controlla connessione USB-C ben inserita

**Raspberry si riavvia casualmente**:
- Causa: Tensione instabile o picchi sotto carico
- Soluzione: Usa condensatore elettrolitico 1000µF 16V sull'output convertitore
- Oppure: Convertitore sottodimensionato, passa a modello 5A

**Display non si accende**:
1. Verifica alimentazione 12V ai pin display
2. Controlla fusibile 3A
3. Verifica cavo HDMI collegato al Raspberry

**Sistema si scarica la batteria**:
- Verifica corrente totale a quadro spento: <50mA è accettabile
- Se >200mA: problema di dispersione, controlla cablaggio
- Usa interruttore manuale o preleva da F15 (si spegne con chiave)

⚠️ **IMPORTANTE**: Usa sempre fusibili PRIMA di ogni componente, non dopo!



### Protezione GPIO

⚠️ **IMPORTANTE**: I GPIO del Raspberry Pi sono 3.3V e **NON** tolleranti a 5V. Gli optoaccoppiatori servono proprio a proteggere i GPIO dai 12V del veicolo.

**Schema protezione**:
```
12V veicolo → Resistenza → LED optoaccoppiatore
                            ↓ (isolamento ottico)
Transistor optoaccoppiatore → GPIO (3.3V)
```

---

## 📚 Riferimenti

### Datasheet Componenti

- **Raspberry Pi 4B**: [raspberrypi.com/products/raspberry-pi-4-model-b](https://www.raspberrypi.com/products/raspberry-pi-4-model-b/)
- **DS18B20**: Maxim Integrated DS18B20 Datasheet
- **ADS1115**: Texas Instruments ADS1115 Datasheet
- **PC817**: Sharp PC817 Datasheet
- **ELM327**: ELM Electronics ELM327 Datasheet

### Schema Elettrico Veicolo

- **Fiat Panda 141 - Impianto Elettrico Gamma 2000**: [Manuale Officina - Schemi Elettrici](http://www.bunkeringegnere.altervista.org/esplosi/FIAT%20PANDA/panda%20141/1100%20mpi/55%20IMPIANTO%20ELETTRICO%20-%20SCHEMI%20-%20GAMMA%202000.pdf)
  - Schema completo impianto elettrico 1100 MPI
  - Codici colore cavi (es. R=Rosso, N=Nero, BN=Bianco-Nero, etc.)
  - Posizione spie e connettori sul quadro strumenti
  - Identificazione centralina di derivazione e fusibili
  - **Essenziale per identificare i cavi corretti da collegare agli optoaccoppiatori**

### Guide Configurazione

- [Configurazione Raspberry Pi](server/CONFIGURAZIONE_SERVER.md#-configurazione-raspberry-pi)
- [Setup GPIO](server/CONFIGURAZIONE_SERVER.md#-configurazione-gpio)
- [Sensori](server/CONFIGURAZIONE_SERVER.md#-sensore-temperatura-ds18b20)

---

## ⚠️ Disclaimer Hardware

L'acquisto e l'installazione di hardware è sotto la tua completa responsabilità. Assicurati di:

- ✅ Avere competenze elettriche/elettroniche adeguate
- ✅ **Utilizzare SEMPRE fusibili separati** per ogni componente alimentato
- ✅ Utilizzare cavi con sezione adeguata alla corrente da trasportare
- ✅ Rispettare le normative locali sulla sicurezza automotive
- ✅ Non compromettere la sicurezza del veicolo
- ✅ Testare tutto su banco prima dell'installazione
- ✅ Verificare polarità prima di collegare alimentazioni
- ✅ Non bypassare MAI le protezioni (fusibili, diodi)

Vedi il [Disclaimer completo](README.md#️-disclaimer) per maggiori informazioni.

---

**Ultimo aggiornamento**: Novembre 2025  
**Hardware testato**: Fiat Panda 141 1100 mpi

