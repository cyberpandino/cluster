[🇬🇧 English](DOCUMENTAZIONE.en.md) | [🇮🇹 Italiano](DOCUMENTAZIONE.md) | [🇩🇪 Deutsch](DOCUMENTAZIONE.de.md)

---

# 📚 Indice Documentazione - PandaOS Cluster

Guida completa alla documentazione del progetto.

---

## 🎯 Da Dove Iniziare?

### 👋 Nuovo al Progetto?
**Inizia qui**: [QUICK_START.md](QUICK_START.md)  
Guida rapida per avviare il progetto in 5 minuti.

### 📖 Vuoi Capire Tutto?
**Leggi**: [README.md](README.md)  
Documentazione principale completa con setup, configurazione e troubleshooting.

### 🛒 Devi Acquistare Componenti?
**Consulta**: [HARDWARE.md](HARDWARE.md)  
Lista completa hardware necessario con specifiche e costi.

### 🔧 Devi Configurare l'Hardware?
**Vai a**: [server/CONFIGURAZIONE_SERVER.md](server/CONFIGURAZIONE_SERVER.md)  
Setup GPIO, OBD-II, sensori e configurazione Raspberry Pi.

### 💻 Vuoi Personalizzare il Client?
**Consulta**: [client/CONFIGURAZIONE.md](client/CONFIGURAZIONE.md)  
Configurazione completa dell'interfaccia e modalità operative.

### 🏗️ Vuoi Estendere il Sistema?
**Studia**: [ARCHITETTURA.md](ARCHITETTURA.md)  
Architettura tecnica, flussi dati e guide per aggiungere funzionalità.

### 💡 Vuoi Contribuire ma Non Hai Idee?
**Vedi**: [ROADMAP.md](ROADMAP.md)  
Feature future, wishlist e idee per contribuire al progetto.

---

## 📋 Struttura Documentazione

### 1. [README.md](README.md) - Documentazione Principale
**Contenuto**:
- 📋 Descrizione progetto e caratteristiche
- ⚠️ **Disclaimer importante** (responsabilità e sicurezza)
- 🏗️ Architettura generale
- 🤔 **Scelta dello stack tecnologico** (perché React/Electron)
- ⚙️ Requisiti di sistema
- 🚀 Setup completo step-by-step
- 🎯 Avvio del progetto (locale e Raspberry)
- 🔌 Configurazione GPIO (panoramica)
- 🔧 Setup PM2 per produzione
- 🛠️ Troubleshooting generale
- 📦 Build per produzione

**Per Chi**: Tutti - punto di partenza essenziale

---

### 2. [QUICK_START.md](QUICK_START.md) - Avvio Rapido
**Contenuto**:
- ⚡ Installazione in 3 comandi
- 🎛️ Configurazione minima
- 🚀 Avvio veloce
- 📋 Checklist hardware
- 🐛 Troubleshooting rapido
- 🎯 Prossimi passi

**Per Chi**: Sviluppatori che vogliono iniziare velocemente

---

### 3. [client/CONFIGURAZIONE.md](client/CONFIGURAZIONE.md) - Configurazione Client
**Contenuto**:
- 📁 File `environment.ts` spiegato in dettaglio
- 🔧 Parametri WebSocket (URL, mock mode, timeout)
- 🎬 Configurazione splash screen
- 🐛 Debug mode e console viewer
- 🌍 Locale, timezone, formato orario
- 📋 Esempi configurazione (sviluppo, produzione, testing)
- 🔄 Workflow di sviluppo
- 🔍 Modalità mock vs reale
- 🚨 Troubleshooting client

**Per Chi**: Frontend developers, chi personalizza l'interfaccia

---

### 4. [HARDWARE.md](HARDWARE.md) - Lista Hardware
**Contenuto**:
- 🛒 Componenti essenziali (Raspberry Pi, ELM327, Display)
- 🔌 Sensori opzionali (DS18B20, ADS1115)
- 📺 Specifiche display LCD utilizzato
- 🔗 Accessori e cablaggio necessari
- 💰 Stima costi indicativi
- 📦 Kit consigliati
- 🔍 Note tecniche e compatibilità
- ⚠️ Disclaimer hardware

**Per Chi**: Chi deve acquistare componenti, hardware builders

---

### 5. [ROADMAP.md](ROADMAP.md) - Feature Future e Wishlist
**Contenuto**:
- 🚗 Feature hardware (retrocamera, sensori parcheggio, animazioni 3D)
- 💻 Feature software (trip computer, dashboard custom, app mobile)
- 📚 Documentazione (tutorial foto/video, i18n, PCB design)
- 🧪 Testing & qualità
- 🔧 Compatibilità altri veicoli Fiat
- 🎨 UI/UX miglioramenti
- 🌐 Integrazioni servizi (Maps, Spotify, meteo)
- 📊 Stato sviluppo corrente
- 🤝 Come contribuire alle feature

**Per Chi**: Chi vuole contribuire ma cerca idee, chi vuole sapere il futuro del progetto

---

### 6. [server/CONFIGURAZIONE_SERVER.md](server/CONFIGURAZIONE_SERVER.md) - Setup Hardware e Server
**Contenuto**:
- 🛠️ Requisiti hardware completi
- 🔧 Configurazione Raspberry Pi passo-passo
- 🔌 **Porta Seriale OBD-II**:
  - Setup hardware ELM327
  - Identificazione porta (`/dev/ttyUSB0`)
  - Test connessione
  - Configurazione baudrate
  - Protocollo OBD-II
- 🔢 **Configurazione GPIO**:
  - Mappatura completa pin (BCM)
  - Schema pinout Raspberry Pi 4B
  - Tabella tutte le spie (17 GPIO)
  - Cablaggio optoaccoppiatori
  - Schema elettrico
  - Test GPIO
- 🌡️ **Sensore Temperatura DS18B20**:
  - Specifiche tecniche
  - Schema collegamento (resistenza pull-up)
  - Setup 1-Wire
  - Verifica rilevamento
  - Configurazione software
  - Sensori multipli
- ⛽ **Sensore Carburante ADS1115**:
  - Specifiche ADC
  - Schema collegamento I2C
  - Partitore resistivo (calcolo)
  - Setup I2C
  - Test lettura
  - Configurazione gain e sample rate
  - Procedura calibrazione
- ⚡ **Gestione Quadro Accensione**:
  - Hardware setup
  - Script power-saving
  - Customizzazione (auto-shutdown, notifiche)
- 🚨 Troubleshooting hardware dettagliato
- 📚 Datasheet e comandi utili

**Per Chi**: Hardware engineers, chi configura il Raspberry Pi, sysadmin

---

### 7. [ARCHITETTURA.md](ARCHITETTURA.md) - Documentazione Tecnica
**Contenuto**:
- 📊 Diagramma architettura completo
- 🔄 Flussi dati:
  - Avvio sistema
  - Lettura dati OBD
  - Rilevamento spie GPIO
  - Lettura sensori
- 📦 **Moduli Server** (descrizione dettagliata):
  - OBDServer (orchestratore)
  - OBDCommunicationService (ELM327)
  - PIDParserService (parsing hex)
  - MonitoringService (polling)
  - GPIOService (spie)
  - IgnitionService (quadro)
  - TemperatureSensorService (DS18B20)
  - FuelSensorService (ADS1115)
  - WebSocketService (comunicazione)
- 🎨 **Moduli Client**:
  - State Management (Valtio)
  - WebSocketService client
  - MockAnimationService
- 🔐 Sicurezza e permessi
- 🧪 Testing (locale, integrazione, Electron)
- 📈 Performance e ottimizzazioni
- 🔄 **Guide Estendibilità**:
  - Aggiungere nuovo PID OBD
  - Aggiungere nuova spia GPIO
  - Aggiungere nuovo sensore (esempio BMP280)
- 📚 Tabella file principali

**Per Chi**: Developers avanzati, chi vuole contribuire, chi vuole estendere il sistema

---

## 🗂️ Organizzazione File

```
cockpit/
├── README.md                          ← 📖 Documentazione principale
├── QUICK_START.md                     ← ⚡ Guida rapida
├── HARDWARE.md                        ← 🛒 Lista completa hardware
├── ROADMAP.md                         ← 🗺️ Feature future e wishlist
├── DOCUMENTAZIONE.md                  ← 📚 Questo file (indice)
├── ARCHITETTURA.md                    ← 🏗️ Architettura tecnica
│
├── client/
│   ├── CONFIGURAZIONE.md              ← 💻 Configurazione client
│   └── src/config/
│       └── environment.ts             ← ⚙️ File configurazione
│
└── server/
    ├── CONFIGURAZIONE_SERVER.md       ← 🔧 Setup hardware
    ├── config/
    │   └── gpio-mapping.js            ← 🔢 Mappatura GPIO
    ├── services/                      ← 📦 Servizi backend
    └── scripts/                       ← ⚡ Script power-saving
```
---

## 🔍 Trova Rapidamente

### Configurazione
| Cosa cerchi | Dove trovarlo |
|-------------|---------------|
| **Disclaimer e responsabilità** | **[README.md](README.md) § Disclaimer** |
| Setup iniziale progetto | [README.md](README.md) § Setup |
| Configurazione client | [client/CONFIGURAZIONE.md](client/CONFIGURAZIONE.md) |
| Configurazione server | [server/CONFIGURAZIONE_SERVER.md](server/CONFIGURAZIONE_SERVER.md) |
| Modalità mock vs reale | [client/CONFIGURAZIONE.md](client/CONFIGURAZIONE.md) § Mock vs Reale |
| GPIO pin mapping | [server/CONFIGURAZIONE_SERVER.md](server/CONFIGURAZIONE_SERVER.md) § GPIO |
| OBD porta seriale | [server/CONFIGURAZIONE_SERVER.md](server/CONFIGURAZIONE_SERVER.md) § OBD |

### Hardware
| Componente | Documentazione |
|------------|----------------|
| **Lista completa componenti** | **[HARDWARE.md](HARDWARE.md)** |
| Raspberry Pi setup | [server/CONFIGURAZIONE_SERVER.md](server/CONFIGURAZIONE_SERVER.md) § Raspberry |
| ELM327 OBD | [server/CONFIGURAZIONE_SERVER.md](server/CONFIGURAZIONE_SERVER.md) § Porta Seriale |
| Optoaccoppiatori GPIO | [server/CONFIGURAZIONE_SERVER.md](server/CONFIGURAZIONE_SERVER.md) § GPIO |
| DS18B20 temperatura | [server/CONFIGURAZIONE_SERVER.md](server/CONFIGURAZIONE_SERVER.md) § Temperatura |
| ADS1115 carburante | [server/CONFIGURAZIONE_SERVER.md](server/CONFIGURAZIONE_SERVER.md) § Carburante |
| Ignition/Power saving | [server/CONFIGURAZIONE_SERVER.md](server/CONFIGURAZIONE_SERVER.md) § Ignition |

### Troubleshooting
| Problema | Soluzione |
|----------|-----------|
| Server non si avvia | [README.md](README.md) § Troubleshooting |
| ELM327 non trovato | [server/CONFIGURAZIONE_SERVER.md](server/CONFIGURAZIONE_SERVER.md) § Troubleshooting |
| GPIO non funziona | [server/CONFIGURAZIONE_SERVER.md](server/CONFIGURAZIONE_SERVER.md) § Troubleshooting |
| Sensori non rilevano | [server/CONFIGURAZIONE_SERVER.md](server/CONFIGURAZIONE_SERVER.md) § Troubleshooting |
| Client non si connette | [client/CONFIGURAZIONE.md](client/CONFIGURAZIONE.md) § Troubleshooting |

### Sviluppo
| Task | Guida |
|------|-------|
| Aggiungere PID OBD | [ARCHITETTURA.md](ARCHITETTURA.md) § Estendibilità |
| Aggiungere spia GPIO | [ARCHITETTURA.md](ARCHITETTURA.md) § Estendibilità |
| Aggiungere sensore | [ARCHITETTURA.md](ARCHITETTURA.md) § Estendibilità |
| Modificare UI | [client/CONFIGURAZIONE.md](client/CONFIGURAZIONE.md) + [ARCHITETTURA.md](ARCHITETTURA.md) |
| Testing | [ARCHITETTURA.md](ARCHITETTURA.md) § Testing |

---

## 📞 Supporto

### Documentazione Non Chiara?
Apri una [issue](https://github.com/cyberpandino/cluster/issues/new?template=question.md) per migliorare la documentazione.

### Bug o Problema?
1. Controlla [README.md](README.md) § Troubleshooting
2. Controlla documentazione specifica del modulo
3. Apri un [Bug Report](https://github.com/cyberpandino/cluster/issues/new?template=bug_report.md) con:
   - Sistema operativo
   - Hardware utilizzato
   - Log completi
   - Passaggi per riprodurre

### Vuoi Contribuire?
1. Leggi [CONTRIBUTING.md](.github/CONTRIBUTING.md) - Guida completa
2. Cerca idee nella [Roadmap & Wishlist](.github/CONTRIBUTING.md#-vuoi-contribuire-ma-non-hai-idee)
3. Scegli un [template issue](.github/ISSUE_TEMPLATE/) appropriato
4. Studia [ARCHITETTURA.md](ARCHITETTURA.md) per modifiche al codice
5. Usa il [template PR](.github/PULL_REQUEST_TEMPLATE.md) per contributi

### Template Disponibili
- [🐛 Bug Report](.github/ISSUE_TEMPLATE/bug_report.md)
- [✨ Feature Request](.github/ISSUE_TEMPLATE/feature_request.md)
- [❓ Domanda](.github/ISSUE_TEMPLATE/question.md)
- [🔀 Pull Request](.github/PULL_REQUEST_TEMPLATE.md)

---

## 🔄 Aggiornamenti Documentazione

**Versione**: 0.9.0  
**Ultimo aggiornamento**: Novembre 2025 




