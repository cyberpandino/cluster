[🇬🇧 English](QUICK_START.en.md) | [🇮🇹 Italiano](QUICK_START.md) | [🇩🇪 Deutsch](QUICK_START.de.md)

---

# ⚡ Quick Start - PandaOS Cluster

Guida rapida per iniziare in 5 minuti.

> ⚠️ **ATTENZIONE**: PandaOS è un progetto sperimentale hobbistico. L'installazione su veicoli reali è a tuo rischio e pericolo. Leggi il [Disclaimer completo](README.md#️-disclaimer) prima di procedere.

---

## 🚀 Avvio Rapido

### 0. Prerequisiti

Node.js 18+, npm 9+, Git 2.0+

```bash
node --version && npm --version && git --version
```

> 📘 Raspberry Pi: vedi [CONFIGURAZIONE_SERVER.md](server/CONFIGURAZIONE_SERVER.md#2-installazione-nodejs-e-npm)

---

### 1. Installazione

```bash
# Clona repository
git clone git@github.com:cyberpandino/cluster.git
cd cluster

# Installa tutte le dipendenze
npm run install:all
```

### 2. Configurazione Base

#### Per Sviluppo Locale (Mac/Windows/Linux)

Modifica `client/src/config/environment.ts`:

```typescript
websocket: {
  mock: true,  // ← Modalità demo
  // ...
}
```

#### Per Raspberry Pi (Produzione)

```typescript
websocket: {
  mock: false,  // ← Connessione reale al server
  // ...
}
```

### 3. Avvio

#### Sviluppo Locale (Solo Client)

```bash
# Solo interfaccia con dati simulati
cd client
npm run dev
```

Apri browser: `http://localhost:5173`

#### Raspberry Pi (Stack Completo)

```bash
# Server + Client + Electron
npm start
```

---

## 📋 Checklist Hardware (Raspberry Pi)

> 💡 **Hai bisogno di acquistare componenti?** Consulta [HARDWARE.md](HARDWARE.md) per la lista completa di tutto il necessario.

Prima di avviare in produzione, verifica:

- [ ] ELM327 collegato a porta USB (`/dev/ttyUSB0`)
- [ ] Optoaccoppiatori cablati ai pin GPIO
- [ ] Permessi utente configurati:
  ```bash
  sudo usermod -a -G dialout,gpio,i2c $USER
  ```
- [ ] Interfacce abilitate via `raspi-config`:
  - [ ] I2C (se usi ADS1115)
  - [ ] 1-Wire (se usi DS18B20)
  - [ ] Serial Port
- [ ] Riavvio dopo configurazione: `sudo reboot`

---

## 🎛️ Configurazione Minima

### Client (`client/src/config/environment.ts`)

```typescript
export const environment = {
  websocket: {
    url: 'http://127.0.0.1:3001',
    mock: true,  // true=demo | false=reale
  },
  debug: {
    enabled: true,
  },
};
```

### Server (`server/config/gpio-mapping.js`)

```javascript
module.exports = {
  // Porta seriale OBD (modifica se necessario)
  // In OBDCommunicationService.js:
  portPath: '/dev/ttyUSB0',
  
  // Pin GPIO per spie (vedi tabella completa)
  mapping: {
    turnSignals: { pin: 17 },
    battery: { pin: 27 },
    highBeam: { pin: 5 },
    // ... altri
  },
  
  // Sensori opzionali
  temperature: { enabled: true },
  fuel: { enabled: true },
  ignition: { enabled: true },
};
```

---

## 🔑 Controlli Tastiera

Una volta avviata l'applicazione:

- **`d`** → Apri console di debug
- **`ESC`** → Chiudi console di debug
- **`r`** → Ricarica applicazione

---

## 🐛 Troubleshooting Rapido

### "Server non si avvia"

**Su Mac/Windows**: Normale! Il server richiede Raspberry Pi.  
**Soluzione**: Usa `mock: true` nel client.

### "ELM327 non trovato"

```bash
# Verifica porta
ls -l /dev/ttyUSB*

# Dai permessi
sudo usermod -a -G dialout $USER
# Logout e login
```

### "Client non si connette"

1. Verifica server avviato: `npm run server`
2. Controlla URL: `websocket.url` in `environment.ts`
3. Verifica porta 3001 libera: `lsof -i :3001`

### "GPIO non funziona"

```bash
# Test pin GPIO 17
gpio -g read 17

# Se errore:
sudo usermod -a -G gpio $USER
# Riavvia
```

---

## 📚 Documentazione Completa

- **[README.md](README.md)** → Documentazione principale completa
- **[client/CONFIGURAZIONE.md](client/CONFIGURAZIONE.md)** → Configurazione client dettagliata
- **[server/CONFIGURAZIONE_SERVER.md](server/CONFIGURAZIONE_SERVER.md)** → Setup hardware e server
- **[ARCHITETTURA.md](ARCHITETTURA.md)** → Architettura sistema

---

## 🎯 Prossimi Passi

1. **Sviluppo UI**: Modifica componenti in `client/src/components/`
2. **Personalizza GPIO**: Adatta `server/config/gpio-mapping.js` al tuo cablaggio
3. **Aggiungi PID**: Estendi `server/services/PIDParserService.js`
4. **Styling**: Modifica `client/src/assets/scss/`
5. **Produzione**: Setup PM2 (vedi README.md)

---

## 📞 Aiuto

Hai problemi? 

1. **Consulta la documentazione**:
   - [README.md](README.md) - Troubleshooting generale
   - [server/CONFIGURAZIONE_SERVER.md](server/CONFIGURAZIONE_SERVER.md) - Problemi hardware
   - [client/CONFIGURAZIONE.md](client/CONFIGURAZIONE.md) - Problemi client

2. **Apri una issue**:
   - [🐛 Bug Report](.github/ISSUE_TEMPLATE/bug_report.md) - Segnala un problema
   - [❓ Domanda](.github/ISSUE_TEMPLATE/question.md) - Fai una domanda

3. **Contribuisci**:
   - [✨ Feature Request](.github/ISSUE_TEMPLATE/feature_request.md) - Proponi miglioramenti
   - [CONTRIBUTING.md](.github/CONTRIBUTING.md) - Guida per contribuire

---

**Buon coding e non fate danni! 🚗💨**

