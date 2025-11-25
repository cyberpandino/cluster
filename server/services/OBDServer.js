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

const http = require('http');
const fs = require('fs');
const os = require('os');
const OBDCommunicationService = require('./OBDCommunicationService');
const WebSocketService = require('./WebSocketService');
const MonitoringService = require('./MonitoringService');
const GPIOService = require('./GPIOService');
const TemperatureSensorService = require('./TemperatureSensorService');
const FuelSensorService = require('./FuelSensorService');
const IgnitionService = require('./IgnitionService');

class OBDServer {
  constructor() {
    this.isReconnecting = false;
    this.reconnectAttempts = 0;
    this.totalReconnects = 0;
    this.periodicScanInterval = null;
    
    // Verifica dipendenze Raspberry Pi prima di inizializzare
    this.checkRaspberryPiDependencies();
    
    this.initializeServices();
    this.setupEventHandlers();
  }

  /**
   * Verifica che il sistema sia un Raspberry Pi con le dipendenze essenziali
   * I sensori (temperatura e carburante) sono opzionali e non bloccano l'avvio
   * Se le dipendenze essenziali non sono disponibili, termina il processo con errore
   * 
   * Per sviluppo su laptop/desktop: usa il client in modalità mock (websocket.mock = true)
   * oppure imposta variabile d'ambiente DEV_MODE=true per bypassare i check (non consigliato)
   */
  checkRaspberryPiDependencies() {
    const errors = [];
    
    // Verifica modulo GPIO onoff (essenziale per il funzionamento base)
    try {
      require('onoff');
    } catch (error) {
      errors.push('Modulo GPIO (onoff) non disponibile - richiesto per Raspberry Pi');
    }
    
    // Verifica architettura ARM Linux (essenziale)
    const arch = os.arch();
    const platform = os.platform();
    if (platform !== 'linux' || (arch !== 'arm' && arch !== 'arm64')) {
      errors.push(`Piattaforma non supportata: ${platform} ${arch} - richiesto Linux ARM (Raspberry Pi)`);
    }
    
    // Se ci sono errori E non siamo in DEV_MODE, termina il processo
    if (errors.length > 0) {
      const isDevMode = process.env.DEV_MODE === 'true';
      
      console.error('═══════════════════════════════════════════════════════════');
      console.error('❌ ERRORE: Dipendenze Raspberry Pi essenziali non disponibili');
      console.error('═══════════════════════════════════════════════════════════');
      errors.forEach((error, index) => {
        console.error(`${index + 1}. ${error}`);
      });
      console.error('═══════════════════════════════════════════════════════════');
      console.error('Il server richiede un Raspberry Pi con le seguenti dipendenze essenziali:');
      console.error('  - Modulo GPIO (onoff)');
      console.error('  - Sistema Linux ARM (Raspberry Pi)');
      console.error('═══════════════════════════════════════════════════════════');
      console.error('');
      console.error('💡 NOTA: Per sviluppo su laptop/desktop:');
      console.error('   1. Usa il CLIENT in modalità mock (websocket.mock = true)');
      console.error('   2. Non serve avviare il server - il client funziona standalone');
      console.error('   3. Vedi: client/src/config/environment.ts');
      console.error('');
      
      if (isDevMode) {
        console.warn('⚠️  DEV_MODE attivo - server avviato senza hardware (non funzionale)');
        console.warn('   Usa solo per test di sviluppo, non per produzione!');
        return; // Non uscire, continua senza hardware
      }
      
      console.error('Per installare dipendenze su sistemi non-Raspberry (es. per dev):');
      console.error('   npm install --ignore-scripts');
      console.error('');
      process.exit(1);
    }
    
    console.log('✅ Verifica dipendenze Raspberry Pi essenziali completata con successo');
    
    // Avvisi informativi per sensori opzionali (non bloccanti)
    const warnings = [];
    
    // Verifica percorso 1-Wire per sensore temperatura DS18B20 (opzionale)
    const w1Path = '/sys/bus/w1/devices';
    if (!fs.existsSync(w1Path)) {
      warnings.push('Sensore temperatura DS18B20 non disponibile (1-Wire non trovato)');
    }
    
    // Verifica modulo ADS1115 per sensore carburante (opzionale)
    try {
      require('ads1x15');
    } catch (error) {
      warnings.push('Sensore carburante ADS1115 non disponibile (modulo ads1x15 non trovato)');
    }
    
    if (warnings.length > 0) {
      console.warn('⚠️  Avvisi: Alcuni sensori opzionali non sono disponibili:');
      warnings.forEach((warning, index) => {
        console.warn(`   ${index + 1}. ${warning}`);
      });
      console.warn('   Il server continuerà senza questi sensori.');
    }
  }

  initializeServices() {
    this.obdService = new OBDCommunicationService();
    
    this.httpServer = http.createServer();
    
    this.webSocketService = new WebSocketService(this.httpServer);
    
    // Servizio GPIO per rilevamento spie fisiche
    this.gpioService = new GPIOService(this.webSocketService);
    
    // Servizio quadro accensione (ignition)
    this.ignitionService = new IgnitionService(this.webSocketService);
    
    // Servizio sensore temperatura esterna DS18B20
    this.temperatureSensorService = new TemperatureSensorService(this.webSocketService);
    
    // Servizio sensore livello carburante ADS1115
    this.fuelSensorService = new FuelSensorService(this.webSocketService);
    
    this.monitoringService = new MonitoringService(
      this.obdService,
      this.webSocketService
    );
  }

  setupEventHandlers() {
    this.webSocketService.io.on('connection', (socket) => {
      this.webSocketService.sendInitialStatus(socket, {
        connected: this.obdService.isConnected,
        monitoring: this.monitoringService.isMonitoringActive(),
        reconnectAttempts: this.reconnectAttempts,
        totalReconnects: this.totalReconnects
      });
      
      socket.on('force-restart', () => {
        console.log('Riavvio forzato richiesto dal client');
        this.forceRestart();
      });
    });

    this.obdService.port?.on('error', () => {
      this.webSocketService.emitError(new Error('Errore porta seriale'));
      this.forceReconnect();
    });

    this.obdService.port?.on('close', () => {
      this.webSocketService.emitStatus({ 
        connected: false, 
        message: 'Connessione OBD persa' 
      });
      this.forceReconnect();
    });
  }

  async start() {
    console.log('Avvio servizio OBD-II per Fiat Panda 141');
    
    const PORT = process.env.PORT || 3001;
    
    this.httpServer.listen(PORT, () => {
      console.log(`Server OBD in ascolto su porta ${PORT}`);
      console.log(`WebSocket endpoint: ws://localhost:${PORT}`);
    });

    // Avvia polling GPIO per rilevamento spie fisiche
    this.gpioService.startPolling();
    console.log('Servizio GPIO avviato per rilevamento spie veicolo');

    // Avvia lettura sensore temperatura esterna
    this.temperatureSensorService.startReading();
    console.log('Servizio sensore temperatura DS18B20 avviato');

    // Avvia lettura sensore livello carburante
    this.fuelSensorService.startReading();
    console.log('Servizio sensore carburante ADS1115 avviato');

    // Start OBD connection with retry
    await this.startWithRetry();
  }

  async startWithRetry() {
    if (this.isReconnecting) {
      console.warn('Riconnessione già in corso, richiesta ignorata');
      return;
    }
    
    this.isReconnecting = true;
    this.reconnectAttempts = 0;
    
    while (this.isReconnecting) {
      try {
        this.reconnectAttempts++;
        console.log(this.reconnectAttempts === 1 ? 
          'Tentativo di connessione al connettore OBD...' : 
          `Tentativo di connessione #${this.reconnectAttempts}...`);
        
        this.webSocketService.emitStatus({
          connected: false,
          message: `Tentativo connessione #${this.reconnectAttempts}...`,
          reconnectAttempts: this.reconnectAttempts
        });
        
        // Connect to OBD device
        await this.obdService.connect();
        
        console.log('Connessione OBD stabilita con successo');
        this.isReconnecting = false;
        this.totalReconnects++;
        
        this.webSocketService.emitStatus({
          connected: true,
          message: 'ELM327 connesso'
        });
        
        // Initialize OBD
        await this.obdService.initialize();
        
        // Scan all PIDs
        const workingPIDs = await this.scanAllPIDs();
        
        // Start monitoring
        if (workingPIDs.length > 0) {
          await this.monitoringService.startMonitoring(workingPIDs);
          
          this.startPeriodicPIDScan();
        }
        
        console.warn('Monitoraggio terminato, riavvio sequenza di connessione');
        await this.sleep(2000);
        
      } catch (error) {
        console.error(`Connessione fallita (tentativo ${this.reconnectAttempts}): ${error.message}`);
        
        this.webSocketService.emitStatus({
          connected: false,
          message: `Connessione fallita (tentativo ${this.reconnectAttempts}). Riavvio tra 5 secondi...`,
          error: error.message,
          reconnectAttempts: this.reconnectAttempts
        });
        
        if (this.reconnectAttempts >= 20) {
          console.error('Troppi tentativi falliti - eseguo riavvio completo del processo');
          this.forceRestart();
          return;
        }
        
        console.log('Nuovo tentativo di connessione tra 5 secondi');
        await this.sleep(5000);
      }
    }
  }

  startPeriodicPIDScan() {
    console.log('Avvio scansione periodica PID (ogni 30 secondi)');
    
    this.periodicScanInterval = setInterval(async () => {
      try {
        if (!this.obdService.isConnected || this.isReconnecting) {
          console.warn('OBD disconnesso, scansione periodica saltata');
          return;
        }

        console.log('Scansione periodica PID in corso');
        this.webSocketService.emitStatus({ message: 'Scansione periodica PID...' });
        
        // Esegui la scansione
        const newWorkingPIDs = await this.periodicPIDScan();

        // Aggiorna il MonitoringService con la nuova lista
        if (newWorkingPIDs.length > 0) {
          this.monitoringService.updateWorkingPIDs(newWorkingPIDs);
        }
        
      } catch (error) {
        console.error(`Errore durante scansione periodica: ${error.message}`);
      }
    }, 30000); // 30 secondi
  }

  stopPeriodicPIDScan() {
    if (this.periodicScanInterval) {
      clearInterval(this.periodicScanInterval);
      this.periodicScanInterval = null;
      console.log('Scansione periodica PID fermata');
    }
  }

  async periodicPIDScan() {
    const pidList = this.obdService.getPIDDefinitions();
    const workingPIDs = [];
    
    let scannedCount = 0;
    let foundNewPIDs = 0;

    for (const {pid, name, key} of pidList) {
      const result = await this.obdService.readPID(pid, name);
      scannedCount++;

      if (result.success) {
        const wasAlreadyWorking = this.monitoringService.isPIDCurrentlyMonitored(key);
        if (!wasAlreadyWorking) {
          foundNewPIDs++;
          console.log(`Nuovo PID disponibile: ${name}: ${result.value} ${result.unit}`);
        }
        
        workingPIDs.push({...result, key});

      }

      // Delay più breve per la scansione periodica
      await this.sleep(200);
    }

    const statusMessage = foundNewPIDs > 0 
      ? `Scansione periodica: ${foundNewPIDs} nuovi PID trovati (${workingPIDs.length}/${pidList.length} totali)`
      : `Scansione periodica: ${workingPIDs.length}/${pidList.length} PID funzionanti`;

    this.webSocketService.emitStatus({
      message: statusMessage,
      working_pids: workingPIDs.length,
      total_pids: pidList.length,
      new_pids: foundNewPIDs
    });

    console.log(statusMessage);
    return workingPIDs;
  }

  async scanAllPIDs() {
    console.log('Scansione PID iniziale avviata');
    this.webSocketService.emitStatus({ message: 'Scansione PID in corso...' });

    const pidList = this.obdService.getPIDDefinitions();
    const workingPIDs = [];
    
    const data = {
      timestamp: new Date().toISOString(),
      vehicle: {
        brand: 'Fiat',
        model: 'Panda 141',
        ecu: 'Magneti Marelli IAW 4AF'
      },
      parameters: {}
    };

    for (const {pid, name, key} of pidList) {
      const result = await this.obdService.readPID(pid, name);

      data.parameters[key] = {
        name: result.name,
        value: result.value,
        unit: result.unit,
        success: result.success,
        error: result.error,
        raw: result.raw,
        timestamp: result.timestamp
      };

      if (result.success) {
        console.log(`PID funzionante: ${name}: ${result.value} ${result.unit}`);
        workingPIDs.push({...result, key});
      } else {
        console.warn(`PID non funzionante: ${name}: ${result.value}`);
      }

      await this.sleep(500);
    }

    // Invio dati completi solo via WebSocket, nessuna persistenza
    this.webSocketService.emitOBDData(data);
    this.webSocketService.emitStatus({
      message: `Scansione completata: ${workingPIDs.length}/${pidList.length} PID funzionanti`,
      working_pids: workingPIDs.length,
      total_pids: pidList.length
    });

    console.log(`${workingPIDs.length}/${pidList.length} PID funzionanti`);
    return workingPIDs;
  }

  forceReconnect() {
    console.log('Richiesta di riconnessione forzata');
    this.stopPeriodicPIDScan();
    this.isReconnecting = false;
    this.monitoringService.stopMonitoring();
    this.reconnectAttempts = 0;
    
    this.obdService.disconnect();
    
    setTimeout(() => {
      this.startWithRetry();
    }, 2000);
  }

  forceRestart() {
    console.log('Riavvio completo del processo richiesto');
    this.webSocketService.emitStatus({ 
      message: 'Riavvio completo in corso...', 
      restarting: true 
    });
    
    this.stop();
    
    setTimeout(() => {
      process.exit(1);
    }, 1000);
  }

  stop() {
    console.log('Arresto del servizio OBD in corso');
    this.stopPeriodicPIDScan();
    this.monitoringService.stopMonitoring();
    this.monitoringService.stopWatchdog();
    this.gpioService.cleanup();
    if (this.ignitionService) {
      this.ignitionService.cleanup();
    }
    this.temperatureSensorService.stopReading();
    this.fuelSensorService.stopReading();
    this.obdService.disconnect();
    console.log('Servizio OBD arrestato correttamente');
  }

  sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
  }
}

module.exports = OBDServer; 