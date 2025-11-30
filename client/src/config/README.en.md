[🇬🇧 English](README.en.md) | [🇮🇹 Italiano](README.md) | [🇩🇪 Deutsch](README.de.md)

---

# Environment Configuration

Centralized PandaOS Cockpit application configuration.

## 📁 Structure

```
client/src/config/
├── types.ts         # TypeScript interfaces
├── environment.ts   # Single configuration
├── index.ts        # Entry point
└── README.md       # Documentation
```

## ⚙️ Single Configuration

**One configuration** in `environment.ts` - directly modifiable in the file.

### 🔄 Mock Mode

To switch between demo mode and real WebSocket:

```typescript
// environment.ts
export const environment: EnvironmentConfig = {
  websocket: {
    url: 'http://cyberpandino.local:3001',
    mock: true,  // ← Change this value
    // ...
  },
};
```

- **`mock: true`** → Simulated demo animations (without server)
- **`mock: false`** → Real WebSocket connection

## 📖 Usage

### Direct Import

```typescript
import { websocket, splashScreen, app } from '@/config';

const url = websocket.url;
const isMock = websocket.mock;
```

### Complete Import

```typescript
import { environment } from '@/config';

console.log(environment.websocket.url);
console.log(environment.app.version);
```

### Import Types

```typescript
import type { WebSocketConfig, EnvironmentConfig } from '@/config';
```

## ✏️ Modifying Configuration

**1. Open** `client/src/config/environment.ts`

**2. Modify** desired values

**3. Save** and restart dev server

### Common Examples

#### Enable/Disable Mock

```typescript
websocket: {
  mock: false,  // Disable mock, use real WebSocket
}
```

#### Change WebSocket URL

```typescript
websocket: {
  url: 'http://192.168.1.100:3001',  // Custom IP
  mock: false,
}
```

#### Change URL for localhost

```typescript
websocket: {
  url: 'http://127.0.0.1:3001',  // Localhost
  mock: false,
}
```

#### Change Splash Duration

```typescript
splashScreen: {
  path: '/static-splash.png',
  duration: 3000,  // 3 seconds instead of 2
}
```

## 📝 Available Parameters

### WebSocket
| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `url` | string | `cyberpandino.local:3001` | WebSocket server address |
| `mock` | boolean | `true` | `true`=demo, `false`=real WebSocket |
| `reconnectionAttempts` | number | `3` | Reconnection attempts |
| `reconnectionDelay` | number | `1000` | Delay between reconnections (ms) |
| `timeout` | number | `5000` | Connection timeout (ms) |

### Splash Screen
| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `path` | string | `/static-splash.png` | Splash image path |
| `duration` | number | `2000` | Display duration (ms) |

### Debug
| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `enabled` | boolean | `true` | Enable debug functions |
| `showConsoleViewer` | boolean | `true` | Show console with "d" key |

### App
| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `name` | string | `PandaOS Cockpit` | Application name |
| `version` | string | `1.0.0` | Version |
| `locale` | string | `it` | Language (moment.js) |
| `timezone` | string | `Europe/Rome` | Timezone |

## 🔄 Files Using Config

- ✅ `services/websocket.ts` - WebSocket/Mock
- ✅ `components/SplashScreen/SplashScreen.tsx` - Splash screen
- ✅ `App.tsx` - Locale and timezone

## 🚀 Workflow

### Development with Mock (without server)

```typescript
// environment.ts
websocket: { mock: true }
```

```bash
npm run dev
```

### Testing with Real Server

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

### Production Build

```typescript
// environment.ts
websocket: { mock: false }  // Make sure it's false
```

```bash
npm run build
npm run preview
```

## 💡 Notes

- ⚠️ **No `.env` file** - configuration only in `environment.ts`
- ✅ **Single configuration** - easy to manage
- 🔄 **Restart dev server** after changes
- 📝 **Version `environment.ts`** - shared configuration in team
