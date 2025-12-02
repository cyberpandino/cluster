[🇬🇧 English](CONTRIBUTING.en.md) | [🇮🇹 Italiano](CONTRIBUTING.md) | [🇩🇪 Deutsch](CONTRIBUTING.de.md)

---

# 🤝 Contributing to PandaOS

Thank you for your interest in contributing to PandaOS! Every contribution is welcome, whether it's code, documentation, bug reports, or suggestions.

## 📋 Table of Contents

- [Code of Conduct](#-code-of-conduct)
- [How to Contribute](#-how-to-contribute)
- [Reporting Bugs](#-reporting-bugs)
- [Proposing New Features](#-proposing-new-features)
- [Pull Requests](#-pull-requests)
- [Code Style](#-code-style)
- [Documentation](#-documentation)
- [License](#-license)

---

## 🤝 Code of Conduct

This project adheres to an implicit code of conduct based on mutual respect:

- Be respectful towards all contributors
- Accept constructive criticism with an open mind
- Focus on what's best for the community
- Show empathy towards other community members

## 🎯 How to Contribute

There are many ways to contribute to PandaOS:

### 🐛 Report Bugs
Found a bug? Open an [issue](https://github.com/cyberpandino/cluster/issues/new?template=bug_report.md) using the "Bug Report" template.

### ✨ Propose Features
Have an idea to improve PandaOS? Open an [issue](https://github.com/cyberpandino/cluster/issues/new?template=feature_request.md) using the "Feature Request" template.

### 📚 Improve Documentation
- Fix typos or errors
- Improve existing explanations
- Add examples and guides
- Translate documentation

### 💻 Contribute Code
- Fix open bugs
- Implement new features
- Optimize performance
- Add tests

### 🧪 Testing
- Test on different hardware
- Verify compatibility
- Report issues specific to your setup

---

## 🐛 Reporting Bugs

Before reporting a bug:

1. **Search existing issues**: Verify the bug hasn't already been reported
2. **Use mock mode**: Test in mock mode to rule out hardware issues
3. **Gather information**: Prepare logs, screenshots, and environment details

**Use the template**: [Bug Report](https://github.com/cyberpandino/cluster/issues/new?template=bug_report.md)

### Information to Include

- **Clear description** of the problem
- **Steps to reproduce** the behavior
- **Expected behavior** vs **actual behavior**
- **Environment**: OS, hardware, software versions
- **Complete logs** from server and/or client
- **Screenshots** if relevant

---

## ✨ Proposing New Features

Before proposing a feature:

1. **Verify it doesn't already exist**: Search in open/closed issues and PRs
2. **Consider the scope**: Is the feature aligned with project goals?
3. **Think about implementation**: Do you have ideas on how to implement it?

**Use the template**: [Feature Request](https://github.com/cyberpandino/cluster/issues/new?template=feature_request.md)

### Discussion

For complex features, it's recommended to:
1. Open an issue first to discuss
2. Wait for feedback from maintainers
3. Proceed with implementation after approval

---

## 🔀 Pull Requests

### Workflow

1. **Fork** the repository
2. **Create a branch** from `main`:
   ```bash
   git checkout -b feature/feature-name
   # or
   git checkout -b fix/bug-name
   ```
3. **Make your changes**
4. **Commit** with clear messages:
   ```bash
   git commit -m "feat: add support for XYZ sensor"
   git commit -m "fix: correct GPIO pin 17 reading"
   git commit -m "docs: update configuration guide"
   ```
5. **Push** to your fork:
   ```bash
   git push origin feature/feature-name
   ```
6. **Open a Pull Request** towards `main`

### Commit Conventions

Use [Conventional Commits](https://www.conventionalcommits.org/):

- `feat:` new feature
- `fix:` bug fix
- `docs:` documentation changes
- `style:` formatting, missing semi colons, etc.
- `refactor:` code refactoring
- `perf:` performance improvements
- `test:` adding tests
- `chore:` dependency updates, config, etc.

**Examples**:
```bash
feat(server): add support for CAN protocol
fix(client): correct tachometer rendering on Safari
docs(readme): update Raspberry Pi installation instructions
refactor(gpio): simplify debouncing logic
```

### PR Checklist

Before opening the PR, verify:

- [ ] Code compiles without errors
- [ ] You've tested changes locally
- [ ] You've added/updated documentation
- [ ] You've added GPL-3.0 header to new files
- [ ] Commits are atomic and well described
- [ ] You've rebased on updated main
- [ ] There are no conflicts

### Review Process

1. Open PR filling out the template
2. Maintainers review the code
3. Apply requested changes (if needed)
4. Once approved, PR is merged

---

## 🎨 Code Style

### JavaScript/Node.js (Server)

- **Indentation**: 2 spaces
- **Quotes**: Single quotes `'`
- **Semicolons**: Yes
- **Naming**:
  - `camelCase` for variables and functions
  - `PascalCase` for classes
  - `UPPER_CASE` for constants

**Example**:
```javascript
const MAX_RETRIES = 3;

class OBDService {
  constructor() {
    this.retryCount = 0;
  }

  async readPID(pidCode) {
    // implementation
  }
}
```

### TypeScript/React (Client)

- **Indentation**: 2 spaces
- **Quotes**: Double quotes `"`
- **Semicolons**: Yes
- **Components**: PascalCase
- **Hooks**: camelCase with `use` prefix

**Example**:
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

### Comments

- Comment complex or non-obvious code
- Use JSDoc for public functions
- Explain the "why", not the "what"

**Example**:
```javascript
/**
 * Reads a PID from the ECU with automatic retry
 * @param {string} pid - PID code in hex format (e.g. '010C')
 * @param {string} name - Descriptive parameter name
 * @returns {Promise<Object>} PID reading result
 */
async readPID(pid, name) {
  // Wake up ECU if in sleep mode to avoid timeout
  await this.wakeupECU();
  
  // Implementation...
}
```

---

## 📚 Documentation

Documentation is crucial! Every change should include updates to relevant documentation.

### Files to Update

| Change | Documentation |
|--------|---------------|
| New client feature | `client/CONFIGURAZIONE.en.md` |
| New server feature | `server/CONFIGURAZIONE_SERVER.en.md` |
| Architecture change | `ARCHITETTURA.en.md` |
| New configuration | `README.en.md` + specific file |
| Setup/installation | `README.en.md` + `QUICK_START.en.md` |

### Documentation Style

- **Language**: English
- **Format**: Markdown
- **Tone**: Informal but technical
- **Sections**: Well structured with emojis
- **Examples**: Always include practical examples
- **Screenshots**: When useful for UI

---

## 🔐 Security

### Reporting Vulnerabilities

**DO NOT** open public issues for security vulnerabilities. Contact maintainers privately.

### Security Checklist

- [ ] Don't commit credentials, tokens, passwords
- [ ] Don't expose endpoints without authentication (if added)
- [ ] Validate user input
- [ ] Don't execute unsanitized shell commands
- [ ] Consider hardware implications (GPIO, serial)

---

## 📄 License

By contributing to PandaOS, you agree that your contribution will be released under the [GNU General Public License v3.0 or later](../LICENSE).

### License Header

Add this header to every new source file:

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

## 🎓 Useful Resources

### Project Documentation
- [README.en.md](../README.en.md) - Main documentation
- [QUICK_START.en.md](../QUICK_START.en.md) - Quick guide
- [ARCHITETTURA.en.md](../ARCHITETTURA.en.md) - Technical architecture
- [DOCUMENTAZIONE.en.md](../DOCUMENTAZIONE.en.md) - Documentation index

### Development Setup
```bash
# Clone repository
git clone https://github.com/cyberpandino/cluster.git
cd cluster

# Install dependencies
npm run install:all

# Start in development mode
npm start
```

### Testing
```bash
# Client with mock mode
cd client
npm run dev

# Server (requires Raspberry Pi)
cd server
node server.js
```

---

## 💡 Want to Contribute but Have No Ideas?

Here's a list of features and improvements we'd like to implement but haven't found time yet!

### 🚗 Hardware Features

**High Priority**:
- [ ] **Integrated backup camera** - Display rear camera feed in cluster when reverse gear engaged
- [ ] **Parking sensors** - Graphical display of obstacle distance with ultrasonic radars
- [ ] **3D door animation** - Represent open/closed doors on Panda 3D model in cluster
- [ ] **Lights on 3D model** - Show active lights (high beam, turn signals, etc.) directly on 3D model

**Medium Priority**:
- [ ] **Rain sensor** - Automatic windshield wiper adjustment
- [ ] **Light sensor** - Auto-brightness adjustment for display
- [ ] **Tire pressure (TPMS)** - Integration of tire pressure sensors
- [ ] **CAN Bus support** - Beyond OBD-II, support for native CAN protocol
- [ ] **360° camera** - Multi-camera system for complete view

### 💻 Software Features

**High Priority**:
- [ ] **Trip computer system** - Trip logs with consumption, distance, time
- [ ] **Customizable dashboards** - Multiple layouts selectable by user
- [ ] **Color themes** - Dark mode, light mode, custom themes
- [ ] **Assisted calibration** - Wizard for calibrating fuel/temperature sensors
- [ ] **Mobile companion app** - Vehicle statistics on smartphone

**Medium Priority**:
- [ ] **Scheduled maintenance** - Alerts for service, oil changes, inspection
- [ ] **Weather integration** - External temperature from weather API if sensor unavailable
- [ ] **Automatic night/day mode** - Based on time or light sensor

### 📚 Documentation

**High Priority**:
- [ ] **Photographic wiring tutorial** - Step-by-step guide with real photos of optocoupler wiring
- [ ] **Video installation guide** - Complete video tutorial from wiring to software
- [ ] **Internationalization (i18n)** - EN, ES, DE, FR translations
- [ ] **Centralized translation file** - Move all texts to JSON/i18n files
- [ ] **Custom PCB schematic** - Professional PCB design for optocouplers (KiCad/Eagle)

**Medium Priority**:
- [ ] **Extended FAQ** - Frequently asked questions with detailed troubleshooting
- [ ] **Installation case studies** - Real examples of completed installations
- [ ] **Compatibility guides** - List of compatible vehicles beyond Panda 141
- [ ] **Interactive wiring diagram** - Navigable electrical schematic online

### 🧪 Testing & Quality

- [ ] **Unit tests** - Automated testing for backend services
- [ ] **E2E tests** - Complete interface tests with Playwright/Cypress
- [ ] **Performance profiling** - Rendering and memory optimization

### 🔧 Compatibility & Extensions

- [ ] **Support for other vehicles** - Uno, Tipo, Punto, Seicento...

---

### 🚀 How to Start

1. **Choose a task** from the list above that interests you
2. **Open an issue** using [Feature Request](https://github.com/cyberpandino/cluster/issues/new?template=feature_request.md)
3. **Discuss implementation** with maintainers
4. **Fork and develop** following this guide
5. **Open a PR** when ready

Even partial implementations are welcome! You don't need to complete the entire feature at once.

---

## ❓ Questions?

- **Issue**: Open a [question issue](https://github.com/cyberpandino/cluster/issues/new?template=question.md)
- **Discussions**: Participate in [GitHub Discussions](https://github.com/cyberpandino/cluster/discussions) (if enabled)

---

## 🙏 Acknowledgments

Thanks to all contributors who help make PandaOS better! 🚗💨

All contributors are recognized in the [AUTHORS](../AUTHORS.en.md) file.

---

**For information about authors and maintainers, see [AUTHORS](../AUTHORS.en.md)**

**Last update**: November 2025
