[🇬🇧 English](CONTRIBUTING.en.md) | [🇮🇹 Italiano](CONTRIBUTING.md) | [🇩🇪 Deutsch](CONTRIBUTING.de.md)

---

# 🤝 Contribuire a PandaOS

Grazie per l'interesse nel contribuire a PandaOS! Ogni contributo è benvenuto, che si tratti di codice, documentazione, bug report o suggerimenti.

## 📋 Indice

- [Codice di Condotta](#-codice-di-condotta)
- [Come Contribuire](#-come-contribuire)
- [Segnalare Bug](#-segnalare-bug)
- [Proporre Nuove Feature](#-proporre-nuove-feature)
- [Pull Request](#-pull-request)
- [Stile del Codice](#-stile-del-codice)
- [Documentazione](#-documentazione)
- [Licenza](#-licenza)

---

## 🤝 Codice di Condotta

Questo progetto aderisce a un codice di condotta implicito basato sul rispetto reciproco:

- Sii rispettoso verso tutti i contributori
- Accetta critiche costruttive con apertura mentale
- Focalizzati su ciò che è meglio per la community
- Mostra empatia verso altri membri della community

## 🎯 Come Contribuire

Ci sono molti modi per contribuire a PandaOS:

### 🐛 Segnalare Bug
Hai trovato un bug? Apri una [issue](https://github.com/cyberpandino/cluster/issues/new?template=bug_report.md) usando il template "Bug Report".

### ✨ Proporre Feature
Hai un'idea per migliorare PandaOS? Apri una [issue](https://github.com/cyberpandino/cluster/issues/new?template=feature_request.md) usando il template "Feature Request".

### 📚 Migliorare la Documentazione
- Correggere typo o errori
- Migliorare spiegazioni esistenti
- Aggiungere esempi e guide
- Tradurre documentazione

### 💻 Contribuire con Codice
- Risolvere bug aperti
- Implementare nuove feature
- Ottimizzare performance
- Aggiungere test

### 🧪 Testing
- Testare su hardware diverso
- Verificare compatibilità
- Segnalare problemi specifici del tuo setup

---

## 🐛 Segnalare Bug

Prima di segnalare un bug:

1. **Cerca issue esistenti**: Verifica che il bug non sia già stato segnalato
2. **Usa la modalità mock**: Testa in modalità mock per escludere problemi hardware
3. **Raccogli informazioni**: Prepara log, screenshot e dettagli ambiente

**Usa il template**: [Bug Report](https://github.com/cyberpandino/cluster/issues/new?template=bug_report.md)

### Informazioni da Includere

- **Descrizione chiara** del problema
- **Passi per riprodurre** il comportamento
- **Comportamento atteso** vs **comportamento attuale**
- **Ambiente**: OS, hardware, versioni software
- **Log completi** del server e/o client
- **Screenshot** se rilevante

---

## ✨ Proporre Nuove Feature

Prima di proporre una feature:

1. **Verifica che non esista già**: Cerca in issue e PR aperte/chiuse
2. **Considera lo scope**: La feature è in linea con gli obiettivi del progetto?
3. **Pensa all'implementazione**: Hai idee su come implementarla?

**Usa il template**: [Feature Request](https://github.com/cyberpandino/cluster/issues/new?template=feature_request.md)

### Discussione

Per feature complesse, è consigliato:
1. Aprire prima una issue per discutere
2. Aspettare feedback dai maintainer
3. Procedere con l'implementazione dopo approvazione

---

## 🔀 Pull Request

### Workflow

1. **Fork** il repository
2. **Crea un branch** dalla `main`:
   ```bash
   git checkout -b feature/nome-feature
   # oppure
   git checkout -b fix/nome-bug
   ```
3. **Fai le tue modifiche**
4. **Committa** con messaggi chiari:
   ```bash
   git commit -m "feat: aggiungi supporto per sensore XYZ"
   git commit -m "fix: correggi lettura GPIO pin 17"
   git commit -m "docs: aggiorna guida configurazione"
   ```
5. **Push** al tuo fork:
   ```bash
   git push origin feature/nome-feature
   ```
6. **Apri una Pull Request** verso `main`

### Convenzioni Commit

Usa [Conventional Commits](https://www.conventionalcommits.org/):

- `feat:` nuova feature
- `fix:` correzione bug
- `docs:` modifiche documentazione
- `style:` formattazione, missing semi colons, etc.
- `refactor:` refactoring codice
- `perf:` miglioramenti performance
- `test:` aggiunta test
- `chore:` aggiornamento dipendenze, config, etc.

**Esempi**:
```bash
feat(server): aggiungi supporto per protocollo CAN
fix(client): correggi rendering tachimetro su Safari
docs(readme): aggiorna istruzioni installazione Raspberry Pi
refactor(gpio): semplifica logica debouncing
```

### Checklist PR

Prima di aprire la PR, verifica:

- [ ] Il codice compila senza errori
- [ ] Hai testato le modifiche localmente
- [ ] Hai aggiunto/aggiornato la documentazione
- [ ] Hai aggiunto l'header GPL-3.0 ai nuovi file
- [ ] I commit sono atomici e ben descritti
- [ ] Hai fatto rebase su main aggiornato
- [ ] Non ci sono conflitti

### Review Process

1. Apri la PR compilando il template
2. I maintainer revisionano il codice
3. Applica le modifiche richieste (se necessario)
4. Una volta approvata, la PR viene mergiata

---

## 🎨 Stile del Codice

### JavaScript/Node.js (Server)

- **Indentazione**: 2 spazi
- **Quotes**: Single quotes `'`
- **Semicolons**: Sì
- **Naming**:
  - `camelCase` per variabili e funzioni
  - `PascalCase` per classi
  - `UPPER_CASE` per costanti

**Esempio**:
```javascript
const MAX_RETRIES = 3;

class OBDService {
  constructor() {
    this.retryCount = 0;
  }

  async readPID(pidCode) {
    // implementazione
  }
}
```

### TypeScript/React (Client)

- **Indentazione**: 2 spazi
- **Quotes**: Double quotes `"`
- **Semicolons**: Sì
- **Components**: PascalCase
- **Hooks**: camelCase con prefisso `use`

**Esempio**:
```typescript
interface OdometerProps {
  speed: number;
  rpm: number;
}

export const Odometer: React.FC<OdometerProps> = ({ speed, rpm }) => {
  const [isActive, setIsActive] = useState(false);
  
  return <div>{speed} km/h</div>;
};
```

### Commenti

- Commenta codice complesso o non ovvio
- Usa JSDoc per funzioni pubbliche
- Spiega il "perché", non il "cosa"

**Esempio**:
```javascript
/**
 * Legge un PID dalla centralina con retry automatico
 * @param {string} pid - Codice PID in formato hex (es. '010C')
 * @param {string} name - Nome descrittivo del parametro
 * @returns {Promise<Object>} Risultato lettura PID
 */
async readPID(pid, name) {
  // Risveglia ECU se in modalità sleep per evitare timeout
  await this.wakeupECU();
  
  // Implementazione...
}
```

---

## 📚 Documentazione

La documentazione è cruciale! Ogni modifica dovrebbe includere aggiornamenti alla documentazione rilevante.

### File da Aggiornare

| Modifica | Documentazione |
|----------|----------------|
| Nuova feature client | `client/CONFIGURAZIONE.md` |
| Nuova feature server | `server/CONFIGURAZIONE_SERVER.md` |
| Cambio architettura | `ARCHITETTURA.md` |
| Nuova configurazione | `README.md` + file specifico |
| Setup/installazione | `README.md` + `QUICK_START.md` |

### Stile Documentazione

- **Lingua**: Italiano
- **Formato**: Markdown
- **Tono**: Informale ma tecnico
- **Sezioni**: Ben strutturate con emoji
- **Esempi**: Sempre includere esempi pratici
- **Screenshot**: Quando utili per UI

---

## 🔐 Sicurezza

### Segnalare Vulnerabilità

**NON** aprire issue pubbliche per vulnerabilità di sicurezza. Contatta i maintainer privatamente.

### Checklist Sicurezza

- [ ] Non committare credenziali, token, password
- [ ] Non esporre endpoint senza autenticazione (se aggiunta)
- [ ] Validare input utente
- [ ] Non eseguire comandi shell non sanitizzati
- [ ] Considerare implicazioni hardware (GPIO, seriale)

---

## 📄 Licenza

Contribuendo a PandaOS, accetti che il tuo contributo sia rilasciato sotto la licenza [GNU General Public License v3.0 or later](../LICENSE).

### Header Licenza

Aggiungi questo header a ogni nuovo file sorgente:

```javascript
/*
 * PandaOS
 * Copyright (C) 2025  Cyberpandino
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU General Public License for more details.
 */
```

---

## 🎓 Risorse Utili

### Documentazione Progetto
- [README.md](../README.md) - Documentazione principale
- [QUICK_START.md](../QUICK_START.md) - Guida rapida
- [ARCHITETTURA.md](../ARCHITETTURA.md) - Architettura tecnica
- [DOCUMENTAZIONE.md](../DOCUMENTAZIONE.md) - Indice documentazione

### Setup Sviluppo
```bash
# Clone repository
git clone https://github.com/cyberpandino/cluster.git
cd cluster

# Installa dipendenze
npm run install:all

# Avvia in modalità sviluppo
npm start
```

### Testing
```bash
# Client con modalità mock
cd client
npm run dev

# Server (richiede Raspberry Pi)
cd server
node server.js
```

---

## 💡 Vuoi Contribuire ma Non Hai Idee?

Ecco una lista di feature e miglioramenti che ci piacerebbe implementare ma non abbiamo ancora trovato il tempo!

### 🚗 Feature Hardware

**Alta Priorità**:
- [ ] **Retrocamera integrata** - Visualizzazione camera posteriore nel cluster quando si innesta la retromarcia
- [ ] **Sensori di parcheggio** - Visualizzazione grafica distanza ostacoli con radar ultrasonici
- [ ] **Animazione portiere 3D** - Rappresentare sportelli aperti/chiusi sul modello 3D della Panda nel cluster
- [ ] **Luci sul modello 3D** - Mostrare luci accese (abbaglianti, frecce, etc.) direttamente sul modello 3D

**Media Priorità**:
- [ ] **Sensore pioggia** - Regolazione automatica tergicristalli
- [ ] **Sensore luminosità** - Auto-regolazione luminosità display
- [ ] **Pressione pneumatici (TPMS)** - Integrazione sensori pressione gomme
- [ ] **Supporto CAN Bus** - Oltre a OBD-II, supporto per protocollo CAN nativo (non previsto sulla Panda ma potrebbe essere utile per altre vetture)
- [ ] **Camera 360°** - Sistema multi-camera per visione completa (tipo auto moderne)

### 💻 Feature Software

**Alta Priorità**:
- [ ] **Sistema trip computer** - Log viaggi con consumo, distanza, tempo
- [ ] **Dashboard personalizzabili** - Layout multipli selezionabili dall'utente
- [ ] **Temi colore** - Dark mode, light mode, temi personalizzati
- [ ] **Calibrazione assistita** - Wizard per calibrare sensori carburante/temperatura
- [ ] **App mobile companion** - Statistiche veicolo su smartphone

**Media Priorità**:
- [ ] **Manutenzione programmata** - Alert per tagliandi, cambio olio, revisione
- [ ] **Integrazione meteo** - Temperatura esterna da API meteo se sensore non disponibile
- [ ] **Mode notte/giorno automatico** - Basato su ora o sensore luminosità

### 📚 Documentazione

**Alta Priorità**:
- [ ] **Tutorial cablaggio fotografico** - Guida passo-passo con foto reali del cablaggio optoaccoppiatori
- [ ] **Video guida installazione** - Tutorial video completo dal cablaggio al software
- [ ] **Internazionalizzazione (i18n)** - Traduzioni EN, ES, DE, FR
- [ ] **File di traduzione centralizzato** - Spostare tutti i testi in file JSON/i18n
- [ ] **Schema PCB custom** - Design PCB professionale per optoaccoppiatori (KiCad/Eagle)

**Media Priorità**:
- [ ] **FAQ estesa** - Domande frequenti con troubleshooting dettagliato
- [ ] **Case study installazioni** - Esempi reali di installazioni completate
- [ ] **Guide compatibilità** - Lista veicoli compatibili oltre Panda 141
- [ ] **Wiring diagram interattivo** - Schema elettrico navigabile online

### 🧪 Testing & Qualità

- [ ] **Unit tests** - Testing automatizzato per servizi backend
- [ ] **E2E tests** - Test interfaccia completi con Playwright/Cypress
- [ ] **Performance profiling** - Ottimizzazione rendering e memoria

### 🔧 Compatibilità & Estensioni

- [ ] **Supporto altri veicoli** - Uno, Tipo, Punto, Seicento...

---

### 🚀 Come Iniziare

1. **Scegli un task** dalla lista sopra che ti interessa
2. **Apri una issue** usando [Feature Request](https://github.com/cyberpandino/cluster/issues/new?template=feature_request.md)
3. **Discuti l'implementazione** con i maintainer
4. **Fork e sviluppa** seguendo questa guida
5. **Apri una PR** quando pronto

Anche implementazioni parziali sono benvenute! Non serve completare tutta la feature in una volta.

---

## ❓ Domande?

- **Issue**: Apri una [question issue](https://github.com/cyberpandino/cluster/issues/new?template=question.md)
- **Discussioni**: Partecipa alle [GitHub Discussions](https://github.com/cyberpandino/cluster/discussions) (se abilitato)

---

## 🙏 Riconoscimenti

Grazie a tutti i contributori che aiutano a rendere PandaOS migliore! 🚗💨

Tutti i contributori sono riconosciuti nel file [AUTHORS](../AUTHORS.md).

---

**Per informazioni su autori e maintainer, vedi [AUTHORS](../AUTHORS.md)**

**Ultimo aggiornamento**: Novembre 2025

