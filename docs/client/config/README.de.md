[🇬🇧 English](README.en.md) | [🇮🇹 Italiano](README.md) | [🇩🇪 Deutsch](README.de.md)

---

# Konfiguration Environment

Zentrale Anwendungskonfiguration von PandaOS Cockpit.

## 📁 Struktur

```
client/src/config/
├── types.ts         # TypeScript-Schnittstellen
├── environment.ts   # Einzige Konfiguration
├── index.ts        # Einstiegspunkt
└── README.md       # Dokumentation
```

## ⚙️ Einzige Konfiguration

**Eine einzige Konfiguration** in `environment.ts` - direkt in der Datei änderbar.

### 🔄 Mock-Modus

Um zwischen Demo-Modus und echtem WebSocket zu wechseln:

```typescript
// environment.ts
export const environment: EnvironmentConfig = {
  websocket: {
    url: 'http://cyberpandino.local:3001',
    mock: true,  // ← Diesen Wert ändern
    // ...
  },
};
```

- **`mock: true`** → Simulierte Demo-Animationen (ohne Server)
- **`mock: false`** → Echte WebSocket-Verbindung

## 📖 Verwendung

### Direkter Import

```typescript
import { websocket, splashScreen, app } from '@/config';

const url = websocket.url;
const isMock = websocket.mock;
```

### Vollständiger Import

```typescript
import { environment } from '@/config';

console.log(environment.websocket.url);
console.log(environment.app.version);
```

### Typen importieren

```typescript
import type { WebSocketConfig, EnvironmentConfig } from '@/config';
```

## ✏️ Konfiguration ändern

**1. Öffnen** `client/src/config/environment.ts`

**2. Werte ändern** wie gewünscht

**3. Speichern** und Dev-Server neu starten

### Häufige Beispiele

#### Mock aktivieren/deaktivieren

```typescript
websocket: {
  mock: false,  // Mock deaktivieren, echtes WebSocket verwenden
}
```

#### WebSocket-URL ändern

```typescript
websocket: {
  url: 'http://192.168.1.100:3001',  // Benutzerdefinierte IP
  mock: false,
}
```

#### URL für localhost ändern

```typescript
websocket: {
  url: 'http://127.0.0.1:3001',  // Localhost
  mock: false,
}
```

#### Splash-Dauer ändern

```typescript
splashScreen: {
  path: '/static-splash.png',
  duration: 3000,  // 3 Sekunden statt 2
}
```

## 📝 Verfügbare Parameter

### WebSocket
| Parameter | Typ | Standard | Beschreibung |
|-----------|-----|----------|--------------|
| `url` | string | `cyberpandino.local:3001` | WebSocket-Server-Adresse |
| `mock` | boolean | `true` | `true`=Demo, `false`=Echtes WebSocket |
| `reconnectionAttempts` | number | `3` | Wiederverbindungsversuche |
| `reconnectionDelay` | number | `1000` | Verzögerung zwischen Wiederverbindungen (ms) |
| `timeout` | number | `5000` | Verbindungs-Timeout (ms) |

### Splash Screen
| Parameter | Typ | Standard | Beschreibung |
|-----------|-----|----------|--------------|
| `path` | string | `/static-splash.png` | Splash-Bildpfad |
| `duration` | number | `2000` | Anzeigedauer (ms) |

### Debug
| Parameter | Typ | Standard | Beschreibung |
|-----------|-----|----------|--------------|
| `enabled` | boolean | `true` | Debug-Funktionen aktivieren |
| `showConsoleViewer` | boolean | `true` | Konsole mit Taste "d" anzeigen |

### App
| Parameter | Typ | Standard | Beschreibung |
|-----------|-----|----------|--------------|
| `name` | string | `PandaOS Cockpit` | Anwendungsname |
| `version` | string | `1.0.0` | Version |
| `locale` | string | `de` | Sprache (moment.js) |
| `timezone` | string | `Europe/Berlin` | Zeitzone |

## 🔄 Dateien, die die Konfiguration verwenden

- ✅ `services/websocket.ts` - WebSocket/Mock
- ✅ `components/SplashScreen/SplashScreen.tsx` - Splash-Screen
- ✅ `App.tsx` - Locale und Zeitzone

## 🚀 Workflow

### Entwicklung mit Mock (ohne Server)

```typescript
// environment.ts
websocket: { mock: true }
```

```bash
npm run dev
```

### Test mit echtem Server

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

### Produktions-Build

```typescript
// environment.ts
websocket: { mock: false }  // Sicherstellen, dass es false ist
```

```bash
npm run build
npm run preview
```

## 💡 Hinweise

- ⚠️ **Keine `.env`-Datei** - Konfiguration nur in `environment.ts`
- ✅ **Eine einzige Konfiguration** - einfach zu verwalten
- 🔄 **Dev-Server neu starten** nach Änderungen
- 📝 **`environment.ts` versionieren** - geteilte Konfiguration im Team
