# 🗺️ Roadmap & Wishlist - PandaOS Cluster

Feature e miglioramenti che vorremmo implementare nel progetto.

> 💡 **Vuoi contribuire?** Scegli una feature dalla lista e apri una [Feature Request](https://github.com/cyberpandino/cluster/issues/new?template=feature_request.md)!

---

## 🚗 Feature Hardware

### Alta Priorità

#### 📹 Retrocamera Integrata
**Descrizione**: Visualizzazione camera posteriore nel cluster quando si innesta la retromarcia  
**Complessità**: Media  
**Componenti**: Camera USB/CSI, rilevamento retromarcia via GPIO  
**Benefici**: Sicurezza parcheggio, sostituzione specchietto retrovisore danneggiato

#### 📡 Sensori di Parcheggio
**Descrizione**: Visualizzazione grafica distanza ostacoli con radar ultrasonici  
**Complessità**: Media  
**Componenti**: 4-8 sensori ultrasonici, Arduino/ESP32 per processing  
**Benefici**: Assistenza parcheggio tipo auto moderne

### Media Priorità

- **Sensore pioggia** - Regolazione automatica tergicristalli
- **Sensore luminosità** - Auto-regolazione luminosità display (day/night mode)
- **Pressione pneumatici (TPMS)** - Integrazione sensori pressione gomme
- **Supporto CAN Bus** - Oltre a OBD-II, supporto per protocollo CAN nativo
- **Camera 360°** - Sistema multi-camera per visione completa

---

## 💻 Feature Software

### Alta Priorità

#### 🛣️ Sistema Trip Computer
**Descrizione**: Log viaggi con consumo, distanza, tempo, percorso  
**Complessità**: Media  
**Tecnologie**: Database locale (SQLite), geolocalizzazione  
**Benefici**: Statistiche dettagliate, analisi consumi, storico viaggi

#### 🎨 Dashboard Personalizzabili
**Descrizione**: Layout multipli selezionabili dall'utente (sport, eco, minimal, full)  
**Complessità**: Alta  
**Tecnologie**: React layout system, persistent storage  
**Benefici**: Esperienza personalizzata, adattabilità a preferenze

#### 🌓 Temi Colore
**Descrizione**: Dark mode, light mode, temi personalizzati (retro, futuristico, etc.)  
**Complessità**: Bassa  
**Tecnologie**: CSS variables, theme provider  
**Benefici**: Adattabilità luminosità ambiente, estetica personalizzata

#### 🎯 Calibrazione Assistita
**Descrizione**: Wizard step-by-step per calibrare sensori carburante/temperatura  
**Complessità**: Media  
**Tecnologie**: UI wizard, storage calibrazione  
**Benefici**: Setup più facile per utenti non tecnici

#### 📱 App Mobile Companion
**Descrizione**: App smartphone per statistiche veicolo, notifiche, controllo remoto  
**Complessità**: Alta  
**Tecnologie**: React Native / Flutter, API REST/WebSocket  
**Benefici**: Accesso dati anche lontano dal veicolo

#### 🚪 Animazione Portiere 3D
**Descrizione**: Rappresentare sportelli aperti/chiusi sul modello 3D della Panda nel cluster  
**Complessità**: Bassa  
**Tecnologie**: Three.js, animazioni modello 3D, integrazione GPIO  
**Benefici**: Feedback visivo immediato, UI più immersiva

#### 💡 Luci sul Modello 3D
**Descrizione**: Mostrare luci accese (abbaglianti, frecce, fendinebbia) direttamente sul modello 3D  
**Complessità**: Bassa  
**Tecnologie**: Three.js materials, emissive textures, dati GPIO esistenti  
**Benefici**: Visualizzazione intuitiva stato luci, UI più coinvolgente

### Media Priorità

- **Manutenzione programmata** - Alert per tagliandi, cambio olio, revisione basati su km
- **Integrazione meteo** - Temperatura esterna da API se sensore non disponibile
- **Mode notte/giorno automatico** - Basato su ora/GPS o sensore luminosità
- **Gesture control** - Controlli gestuali (se display touch)
- **Voice control** - Comandi vocali per funzioni principali (Annyang.js)
- **Multi-profilo utente** - Statistiche separate per guidatori diversi
- **Export dati** - CSV/JSON/Excel per analisi esterna
- **Sistema notifiche push** - Alert sonori/visivi per anomalie/manutenzione

---

## 📚 Documentazione

### Alta Priorità

#### 📸 Tutorial Cablaggio Fotografico
**Descrizione**: Guida passo-passo con foto reali del cablaggio optoaccoppiatori  
**Complessità**: Bassa (ma richiede installazione reale)  
**Contenuto**: Foto dettagliate ogni fase, identificazione cavi, connessioni  
**Benefici**: Riduce drasticamente errori installazione

#### 🎥 Video Guida Installazione
**Descrizione**: Tutorial video completo dal cablaggio al software  
**Complessità**: Media  
**Contenuto**: Video montato con voice-over, sottotitoli, capitoli  
**Benefici**: Formato più accessibile per utenti meno tecnici

#### 🌍 Internazionalizzazione (i18n)
**Descrizione**: Traduzioni EN, ES, DE, FR dell'interfaccia  
**Complessità**: Media  
**Tecnologie**: react-i18next (già presente), file JSON traduzioni  
**Benefici**: Accessibilità internazionale, community più ampia

#### 📝 File Traduzione Centralizzato
**Descrizione**: Spostare tutti i microcopy hardcoded in file JSON/i18n  
**Complessità**: Bassa  
**Tecnologie**: i18next, JSON  
**Benefici**: Manutenzione facile, traduzioni community-driven

#### 🔌 Schema PCB Custom
**Descrizione**: Design PCB professionale per optoaccoppiatori (KiCad/Eagle)  
**Complessità**: Alta  
**Tecnologie**: KiCad, Gerber export  
**Benefici**: Installazione pulita, niente breadboard, ordine PCB economico

### Media Priorità

- **FAQ estesa** - Domande frequenti con troubleshooting dettagliato
- **Case study installazioni** - Esempi reali con foto e log
- **Guide compatibilità** - Lista veicoli compatibili (Uno, Seicento, etc.)
- **Wiring diagram interattivo** - Schema elettrico navigabile online (SVG/HTML)

---

## 🧪 Testing & Qualità

- **Unit tests** - Testing automatizzato servizi backend (Jest)
- **E2E tests** - Test interfaccia completi (Playwright/Cypress)
- **CI/CD Pipeline** - GitHub Actions per build e deploy automatici
- **Performance profiling** - Ottimizzazione rendering e memoria
- **Hardware-in-the-loop testing** - Test automatizzati con hardware simulato (mock GPIO/Serial)
- **Stress testing** - Test stabilità long-running
- **Code coverage** - Coverage >80% per codice critico

---

## 🔧 Compatibilità & Estensioni

### Veicoli

- **Fiat Uno** (1983-1995) - OBD-I/II simile
- **Fiat Seicento** (1998-2010) - OBD-II compatibile
- **Fiat Punto** (prima serie) - Stessa centralina Magneti Marelli
- **Lancia Y** (prima serie) - ECU simili
- **Fiat Tipo** - Meccanica correlata

### Protocolli

- **J1850 PWM/VPW** - Protocolli americani
- **CAN Bus (ISO 15765)** - Veicoli più recenti
- **LIN Bus** - Accessori automotive

### Hardware

- **Raspberry Pi Zero 2W** - Versione compatta a basso consumo
- **Raspberry Pi CM4** - Per integrazioni custom
- **Orange Pi / Banana Pi** - Alternative SBC economiche
- **Android tablets** - Porta app su Android nativo

### Integrazioni

- **Android Auto / CarPlay** - Integrazione con sistemi moderni
- **Head Unit aftermarket** - Compatibilità con stereo cinesi Android
- **Backup automatico** - Sistema cloud/USB restore/backup impostazioni
- **OTA Updates** - Aggiornamenti software over-the-air

---

## 🎨 UI/UX

- **Animazioni transizioni** - Transizioni fluide GSAP tra stati/schermate
- **Modalità sport/eco** - Visualizzazioni diverse per stile guida (rosso aggressive / verde efficiente)
- **Widget personalizzabili** - Drag & drop componenti nel cluster
- **Skin/temi community** - Marketplace per condividere temi custom
- **Modalità minimal** - UI essenziale per minor distrazione alla guida
- **Screensaver** - Animazioni quando veicolo in sosta
- **Startup animation** - Animazione boot personalizzabile
- **Easter eggs** - Animazioni speciali per eventi (Natale, compleanno auto, etc.)

---

## 🌐 Integrazione Servizi

- **Google Maps / OpenStreetMap** - Navigazione integrata nel cluster
- **Spotify / Apple Music** - Controllo musica dal cluster
- **Telegram Bot** - Notifiche e controllo remoto
- **IFTTT / Home Assistant** - Automazioni smart home
- **Fuel prices API** - Prezzi carburante in tempo reale
- **Traffic data** - Informazioni traffico
- **Weather API** - Previsioni meteo integrate

---

## 🔐 Sicurezza & Privacy

- **Autenticazione utente** - Login per accesso statistiche/configurazione
- **Crittografia dati** - Protezione dati sensibili salvati
- **VPN client** - Connessione sicura per dati cloud
- **Privacy mode** - Disabilitazione tracking GPS/statistiche

---

## 🎯 Performance & Ottimizzazione

- **Lazy loading migliorato** - Caricamento componenti on-demand
- **Service Workers** - PWA per cache e offline
- **WebGL optimization** - Rendering 3D ottimizzato
- **Memory management** - Riduzione footprint memoria
- **Boot time** - Riduzione tempo avvio <10 secondi

---

## 📊 Stato Attuale

### Versione 0.9.0 (Corrente)

**Implementato**:
- ✅ Lettura dati OBD-II via ELM327
- ✅ Rilevamento spie GPIO con optoaccoppiatori
- ✅ Sensore temperatura DS18B20
- ✅ Sensore carburante ADS1115
- ✅ Dashboard 3D con modello Panda interattivo
- ✅ Modalità mock per sviluppo
- ✅ Console debug integrata
- ✅ Gestione quadro accensione (power-saving)
- ✅ WebSocket real-time
- ✅ Documentazione completa

**In Development**:
- 🔄 Nessuna feature attualmente in sviluppo attivo

**Planned**:
- 📋 Vedi roadmap sopra

### Versione 1.0.0 (Obiettivo)

**Target Feature**:
- Retrocamera funzionante
- Sensori parcheggio
- Animazioni 3D complete
- Dashboard personalizzabili
- Internazionalizzazione
- Testing completo

---

## 🤝 Come Contribuire a Queste Feature

### 1. Scegli una Feature
Scorri la lista e trova qualcosa che:
- Ti appassiona
- Hai competenze per implementare
- Risolve un tuo problema/bisogno

### 2. Apri una Issue
Usa [Feature Request](https://github.com/cyberpandino/cluster/issues/new?template=feature_request.md) per discutere:
- Approccio implementativo
- Tecnologie da usare
- Compatibilità con esistente
- Timeline stimata

### 3. Sviluppa
- Fork il repository
- Crea branch dedicato
- Implementa seguendo [stile codice](.github/CONTRIBUTING.md#-stile-del-codice)
- Testa approfonditamente

### 4. Pull Request
- Apri PR con [template](.github/PULL_REQUEST_TEMPLATE.md)
- Descrivi implementazione
- Allega screenshot/video
- Attendi review

---

## 💬 Discussione

Vuoi discutere feature, proporre alternative, o condividere idee?
- Apri una [Discussione](https://github.com/cyberpandino/cluster/discussions) (se abilitato)
- Oppure una [Question Issue](https://github.com/cyberpandino/cluster/issues/new?template=question.md)

---

## 🛠️ Rewrite su Stack Nativo?

### Il Grande Refactoring

Una feature "speciale" che vale la pena menzionare a parte:

#### ⚙️ Port a C++/Qt Native

**Descrizione**: Riscrivere PandaOS con stack tecnologico automotive professionale  
**Complessità**: Molto Alta (praticamente un rewrite completo)  
**Tecnologie**:
- **C++17/20** per backend e business logic
- **Qt 6 / QML** per interfaccia grafica (standard automotive)
- **Qt 3D** per modello Panda interattivo
- **Yocto/Buildroot** per Linux embedded ottimizzato
- **systemd** per gestione servizi
- **D-Bus** per IPC tra processi

**Benefici**:
- 🚀 **Boot time** <3 secondi (vs ~15s attuale)
- 💾 **Memoria** ~50MB totali (vs ~500MB attuali)
- ⚡ **Performance** rendering 60fps garantiti anche su Pi Zero
- 🔋 **Consumo energetico** ridotto del 60-70%
- 🏭 **Approccio professionale** degno di produzione

**Perché non l'abbiamo fatto?**

Perché dopo 2 ore a bestemmiare con CMake abbiamo detto "fanculo, usiamo npm" e non ci siamo più guardati indietro. 😅

Lo stack web ci ha permesso di:
- Avere qualcosa di funzionante in settimane, non anni
- Non impazzire con cross-compilation ARM
- Usare librerie fighe tipo Three.js senza dover scrivere shader a mano
- Divertirci invece di piangere su segfault misteriosi

**Ma se sei un masochista con tempo libero...**

Saremmo **entusiasti** se qualcuno volesse fare un port nativo! Potremmo avere:
- **PandaOS-Web** (attuale) - Per chi vuole divertirsi
- **PandaOS-Native** (futuro?) - Per chi vuole soffrire (ma andare veloce)

Se ti interessa, apri una discussione. Bonus points se riesci a compilare Qt senza bestemmie. 🚀

---

## 📅 Timeline

Non ci sono timeline fisse. Il progetto è hobbistico e procede quando c'è tempo e passione.  
**I contributi della community sono fondamentali per accelerare lo sviluppo!**

---

## 🤷 Ma Seriamente, Perché Questa Stack Assurda?

**Risposta breve**: Perché è divertente e non vogliamo impazzire.

**Risposta lunga**: [README - Ma React + Electron su un Automotive?!](../README.md#-ma-react--electron-su-un-automotive-siete-pazzi)

**Risposta onesta**: Se dovessimo rifarlo "per bene" useremmo C++/Qt. Ma poi non sarebbe più hobbistico, sarebbe masochismo. 😅

---

**Ultima revisione**: Novembre 2025  
**Prossimo aggiornamento**: Quando abbiamo nuove idee brillanti 💡

