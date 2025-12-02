[🇬🇧 English](CONFIGURAZIONE.en.md) | [🇮🇹 Italiano](CONFIGURAZIONE.md) | [🇩🇪 Deutsch](CONFIGURAZIONE.de.md)

---

# ⚙️ Configurazione Client - PandaOS Cluster

Guida completa alla configurazione del client Cyberpandino.

---

## 📁 File di Configurazione

Il file principale di configurazione si trova in:

```
client/src/config/environment.ts
```

---

## 🔧 Parametri di Configurazione

### 1. WebSocket

Configurazione della connessione al server backend.

```typescript
websocket: {
  url: 'http://127.0.0.1:3001',
  mock: true,
  reconnectionAttempts: 3,
  reconnectionDelay: 1000,
  timeout: 5000,
}
```

#### `url` (string)
- **Descrizione**: Indirizzo del server WebSocket
- **Default**: `'http://127.0.0.1:3001'`
- **Produzione Raspberry Pi**: `'http://127.0.0.1:3001'` (localhost)
- **Sviluppo remoto**: `'http://192.168.x.x:3001'` (IP del Raspberry)

**Esempi**:
```typescript
// Locale (sviluppo o produzione su stesso dispositivo)
url: 'http://127.0.0.1:3001'

// Raspberry Pi sulla rete locale
url: 'http://192.168.1.100:3001'

// Server su porta custom
url: 'http://127.0.0.1:8080'
```

#### `mock` (boolean)
- **Descrizione**: Modalità operativa del client
- **Valori**:
  - `true` = Modalità demo con dati simulati (non richiede server)
  - `false` = Connessione reale al server backend
- **Default**: `true`

**Quando usare**:
```typescript
// Sviluppo locale su Mac/Windows/Linux (senza Raspberry Pi)
mock: true

// Produzione su Raspberry Pi con server attivo
mock: false

// Testing dell'interfaccia senza hardware
mock: true
```

#### `reconnectionAttempts` (number)
- **Descrizione**: Numero massimo di tentativi di riconnessione
- **Default**: `3`
- **Range consigliato**: 2-10
- **Comportamento**: Dopo esaurimento tentativi, passa a modalità mock

#### `reconnectionDelay` (number)
- **Descrizione**: Ritardo tra tentativi di riconnessione (millisecondi)
- **Default**: `1000` (1 secondo)
- **Range consigliato**: 500-5000 ms

#### `timeout` (number)
- **Descrizione**: Timeout connessione WebSocket (millisecondi)
- **Default**: `5000` (5 secondi)
- **Range consigliato**: 3000-10000 ms

---

### 2. Splash Screen

Configurazione schermata di avvio.

```typescript
splashScreen: {
  path: '/splashscreen.mp4'
}
```

#### `path` (string)
- **Descrizione**: Percorso del video splash screen
- **Default**: `'/splashscreen.mp4'`
- **Formato supportato**: MP4, WebM
- **Posizione file**: `client/public/splashscreen.mp4`

**Personalizzazione**:
```typescript
// Video custom
path: '/custom-splash.mp4'

// Disabilita splash (usa immagine statica o nulla)
path: '' // Richiede modifica componente SplashScreen
```

---

### 3. Debug

Configurazione modalità debug e sviluppo.

```typescript
debug: {
  enabled: true,
  showConsoleViewer: true,
}
```

#### `enabled` (boolean)
- **Descrizione**: Abilita funzionalità di debug
- **Default**: `true`
- **Effetti**:
  - Mostra log dettagliati in console browser
  - Abilita pulsanti di debug
  - Mostra informazioni di connessione WebSocket
- **Produzione**: Impostare a `false` per performance ottimali

#### `showConsoleViewer` (boolean)
- **Descrizione**: Abilita console viewer integrata (tasto `d`)
- **Default**: `true`
- **Funzionalità**:
  - Visualizzazione log in tempo reale
  - Monitoraggio errori
  - Debug WebSocket
  - Stato sensori e GPIO

**Controlli Console Viewer**:
- Premi **`d`** per aprire
- Premi **`ESC`** per chiudere
- Pulsante "Clear" per pulire log

---

### 4. App

Configurazione generale dell'applicazione.

```typescript
app: {
  name: "PandaOS Cluster",
  version: "0.9.0",
  locale: "it",
  timezone: "Europe/Rome",
  timeFormat: "24h",
}
```

#### `name` (string)
- **Descrizione**: Nome applicazione visualizzato
- **Default**: `"PandaOS Cluster"`
- **Utilizzo**: Titolo finestra, splash screen, about

#### `version` (string)
- **Descrizione**: Versione applicazione
- **Default**: `"0.9.0"`
- **Formato**: Semantic versioning (MAJOR.MINOR.PATCH)

#### `locale` (string)
- **Descrizione**: Lingua applicazione
- **Default**: `"it"` (Italiano)
- **Supportate**: 
  - `"it"` - Italiano

#### `timezone` (string)
- **Descrizione**: Fuso orario per visualizzazione data/ora
- **Default**: `"Europe/Rome"`
- **Formato**: IANA timezone database
- **Altri esempi**:
  - `"Europe/London"`
  - `"America/New_York"`
  - `"Asia/Tokyo"`

#### `timeFormat` (string)
- **Descrizione**: Formato orario
- **Valori**:
  - `"24h"` - Formato 24 ore (es. 15:30)
  - `"12h"` - Formato 12 ore AM/PM (es. 3:30 PM)
- **Default**: `"24h"`

---

## 📋 Esempi di Configurazione

### Configurazione Sviluppo Locale (Mac/Windows)

```typescript
export const environment: EnvironmentConfig = {
  websocket: {
    url: 'http://127.0.0.1:3001',
    mock: true,                    // ← Modalità demo
    reconnectionAttempts: 3,
    reconnectionDelay: 1000,
    timeout: 5000,
  },
  splashScreen: {
    path: '/splashscreen.mp4'
  },
  debug: {
    enabled: true,                 // ← Debug attivo
    showConsoleViewer: true,       // ← Console viewer disponibile
  },
  app: {
    name: "PandaOS Cluster [DEV]",
    version: "0.9.0",
    locale: "it",
    timezone: "Europe/Rome",
    timeFormat: "24h",
  },
};
```

### Configurazione Produzione (Raspberry Pi)

```typescript
export const environment: EnvironmentConfig = {
  websocket: {
    url: 'http://127.0.0.1:3001',
    mock: false,                   // ← Connessione reale
    reconnectionAttempts: 10,      // ← Più tentativi per stabilità
    reconnectionDelay: 2000,
    timeout: 10000,
  },
  splashScreen: {
    path: '/splashscreen.mp4'
  },
  debug: {
    enabled: false,                // ← Debug disabilitato per performance
    showConsoleViewer: false,
  },
  app: {
    name: "PandaOS Cluster",
    version: "0.9.0",
    locale: "it",
    timezone: "Europe/Rome",
    timeFormat: "24h",
  },
};
```

### Configurazione Testing Remoto

```typescript
export const environment: EnvironmentConfig = {
  websocket: {
    url: 'http://192.168.1.100:3001', // ← IP Raspberry Pi
    mock: false,                       // ← Connessione reale
    reconnectionAttempts: 5,
    reconnectionDelay: 1500,
    timeout: 7000,
  },
  splashScreen: {
    path: '/splashscreen.mp4'
  },
  debug: {
    enabled: true,                     // ← Debug per troubleshooting
    showConsoleViewer: true,
  },
  app: {
    name: "PandaOS Cluster [TEST]",
    version: "0.9.0",
    locale: "it",
    timezone: "Europe/Rome",
    timeFormat: "24h",
  },
};
```

---

## 🔄 Workflow di Sviluppo

### 1. Sviluppo Interfaccia (senza hardware)

```typescript
websocket.mock = true
debug.enabled = true
```

Avvia solo il client:
```bash
cd client
npm run dev
```

### 2. Test Integrazione (con Raspberry Pi)

```typescript
websocket.mock = false
websocket.url = 'http://192.168.1.100:3001'
debug.enabled = true
```

Avvia server su Raspberry Pi e client locale

### 3. Produzione (su Raspberry Pi)

```typescript
websocket.mock = false
websocket.url = 'http://127.0.0.1:3001'
debug.enabled = false
```

Avvia stack completo:
```bash
npm start
```

---

## 🔍 Modalità Mock vs Reale

### Modalità Mock (`mock: true`)

**Caratteristiche**:
- ✅ Nessun server richiesto
- ✅ Dati simulati realistici
- ✅ Animazioni fluide
- ✅ Funziona su qualsiasi OS
- ❌ Non legge dati OBD reali
- ❌ Non legge GPIO/sensori

**Uso**: Sviluppo UI, demo, testing visuale

### Modalità Reale (`mock: false`)

**Caratteristiche**:
- ✅ Dati OBD reali dalla centralina
- ✅ Lettura GPIO per spie veicolo
- ✅ Sensori temperatura e carburante
- ❌ Richiede server attivo
- ❌ Richiede hardware (Raspberry Pi + ELM327)

**Uso**: Produzione, testing hardware, integrazione completa

---

## 🚨 Troubleshooting

### Client non si connette al server

**Sintomo**: Console mostra errori di connessione

**Soluzioni**:
1. Verifica che `mock: false`
2. Controlla URL server: `websocket.url`
3. Verifica che il server sia avviato: `npm run server`
4. Controlla firewall/porta 3001 aperta
5. Testa connessione: `curl http://127.0.0.1:3001`

### Splash screen non si carica

**Sintomo**: Schermata bianca all'avvio

**Soluzioni**:
1. Verifica file esiste: `client/public/splashscreen.mp4`
2. Controlla path in configurazione
3. Verifica formato video (MP4 preferito)
4. Controlla console browser per errori

### Debug console non si apre

**Sintomo**: Premendo `d` non succede nulla

**Soluzioni**:
1. Verifica `debug.showConsoleViewer: true`
2. Controlla che `debug.enabled: true`
3. Verifica console browser per errori
4. Prova `CTRL+D` o `CMD+D`

### Timezone/Ora sbagliata

**Sintomo**: Orario non corretto nel cluster

**Soluzioni**:
1. Verifica `app.timezone` corretto per la tua zona
2. Lista timezone: https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
3. Verifica ora sistema Raspberry Pi: `timedatectl`
4. Sincronizza NTP: `sudo timedatectl set-ntp true`

---

## 📚 Riferimenti

- **Codice sorgente**: `client/src/config/environment.ts`
- **Tipi TypeScript**: `client/src/config/types.ts`
- **Documentazione principale**: [README.md](../../README.md)
- **Configurazione server**: [CONFIGURAZIONE_SERVER.md](../server/CONFIGURAZIONE_SERVER.md)

---

**Ultimo aggiornamento**: v0.9.0

