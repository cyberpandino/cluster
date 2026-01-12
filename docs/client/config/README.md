[🇬🇧 English](README.en.md) | [🇮🇹 Italiano](README.md) | [🇩🇪 Deutsch](README.de.md)

---

# Configurazione Environment

Configurazione centralizzata dell'applicazione PandaOS Cockpit.

## 📁 Struttura

```
client/src/config/
├── types.ts         # Interfacce TypeScript
├── environment.ts   # Configurazione unica
├── index.ts        # Entry point
└── README.md       # Documentazione
```

## ⚙️ Configurazione Unica

**Una sola configurazione** in `environment.ts` - modificabile direttamente nel file.

### 🔄 Modalità Mock

Per passare tra modalità demo e WebSocket reale:

```typescript
// environment.ts
export const environment: EnvironmentConfig = {
  websocket: {
    url: 'http://cyberpandino.local:3001',
    mock: true,  // ← Cambia questo valore
    // ...
  },
};
```

- **`mock: true`** → Animazioni demo simulate (senza server)
- **`mock: false`** → Connessione WebSocket reale

## 📖 Utilizzo

### Import Diretto

```typescript
import { websocket, splashScreen, app } from '@/config';

const url = websocket.url;
const isMock = websocket.mock;
```

### Import Completo

```typescript
import { environment } from '@/config';

console.log(environment.websocket.url);
console.log(environment.app.version);
```

### Import Tipi

```typescript
import type { WebSocketConfig, EnvironmentConfig } from '@/config';
```

## ✏️ Modificare la Configurazione

**1. Apri** `client/src/config/environment.ts`

**2. Modifica** i valori desiderati

**3. Salva** e riavvia il dev server

### Esempi Comuni

#### Attivare/Disattivare Mock

```typescript
websocket: {
  mock: false,  // Disattiva mock, usa WebSocket reale
}
```

#### Cambiare URL WebSocket

```typescript
websocket: {
  url: 'http://192.168.1.100:3001',  // IP personalizzato
  mock: false,
}
```

#### Cambiare URL per localhost

```typescript
websocket: {
  url: 'http://127.0.0.1:3001',  // Localhost
  mock: false,
}
```

#### Cambiare Durata Splash

```typescript
splashScreen: {
  path: '/static-splash.png',
  duration: 3000,  // 3 secondi invece di 2
}
```

## 📝 Parametri Disponibili

### WebSocket
| Parametro | Tipo | Default | Descrizione |
|-----------|------|---------|-------------|
| `url` | string | `cyberpandino.local:3001` | Indirizzo server WebSocket |
| `mock` | boolean | `true` | `true`=demo, `false`=WebSocket reale |
| `reconnectionAttempts` | number | `3` | Tentativi di riconnessione |
| `reconnectionDelay` | number | `1000` | Ritardo tra riconnessioni (ms) |
| `timeout` | number | `5000` | Timeout connessione (ms) |

### Splash Screen
| Parametro | Tipo | Default | Descrizione |
|-----------|------|---------|-------------|
| `path` | string | `/static-splash.png` | Percorso immagine splash |
| `duration` | number | `2000` | Durata visualizzazione (ms) |

### Debug
| Parametro | Tipo | Default | Descrizione |
|-----------|------|---------|-------------|
| `enabled` | boolean | `true` | Abilita funzioni debug |
| `showConsoleViewer` | boolean | `true` | Mostra console con tasto "d" |

### App
| Parametro | Tipo | Default | Descrizione |
|-----------|------|---------|-------------|
| `name` | string | `PandaOS Cockpit` | Nome applicazione |
| `version` | string | `1.0.0` | Versione |
| `locale` | string | `it` | Lingua (moment.js) |
| `timezone` | string | `Europe/Rome` | Fuso orario |

### Graphics (Qualità Grafica)

> 🔗 Vedi [Issue #25](https://github.com/cyberpandino/cluster/issues/25) per dettagli sull'ottimizzazione delle performance.

| Parametro | Tipo | Default | Descrizione |
|-----------|------|---------|-------------|
| `quality` | 1 \| 2 \| 3 | `3` | Livello qualità grafica |

#### Livelli di Qualità

| Livello | Modello 3D | Effetti Blur | Consigliato per |
|---------|------------|--------------|-----------------|
| **3 (max)** | ✅ Canvas WebGL | ✅ Abilitati | Hardware moderno (PC, Mac, Raspberry Pi 5) |
| **2 (medium)** | ❌ Immagine statica | ✅ Abilitati | Hardware medio |
| **1 (min)** | ❌ Immagine statica | ❌ Gradienti statici | Hardware datato/limitato (Raspberry Pi 4B e precedenti) |

#### Ottimizzazioni per Hardware Datato

Per hardware con risorse limitate (come Raspberry Pi 4B o precedenti), imposta `quality: 1`:

```typescript
graphics: {
  quality: 1,  // Disabilita blur e modello 3D
}
```

Questo disabilita:
- **Modello 3D WebGL** → sostituito con immagine PNG statica
- **`backdrop-filter: blur()`** → sostituito con background opaco
- **`filter: blur()`** sul glow → sostituito con gradiente radiale

## 🔄 File che Usano la Config

- ✅ `services/websocket.ts` - WebSocket/Mock
- ✅ `components/SplashScreen/SplashScreen.tsx` - Splash screen
- ✅ `App.tsx` - Locale, timezone e qualità grafica
- ✅ `components/ModelViewer/ModelViewer.tsx` - Modello 3D/fallback
- ✅ `App.scss` - Stili condizionali per qualità grafica

## 🚀 Workflow

### Sviluppo con Mock (senza server)

```typescript
// environment.ts
websocket: { mock: true }
```

```bash
npm run dev
```

### Test con Server Reale

```typescript
// environment.ts
websocket: { 
  url: 'http://cyberpandino.local:3001',
  mock: false 
}
```

```bash
npm run dev
```

### Build Produzione

```typescript
// environment.ts
websocket: { mock: false }  // Assicurati che sia false
```

```bash
npm run build
npm run preview
```

## 💡 Note

- ⚠️ **Nessun file `.env`** - configurazione solo in `environment.ts`
- ✅ **Una sola configurazione** - facile da gestire
- 🔄 **Riavvia il dev server** dopo le modifiche
- 📝 **Versiona `environment.ts`** - configurazione condivisa nel team
