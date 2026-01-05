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

import { HelmetProvider } from "react-helmet-async";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import { useSnapshot } from "valtio";
import { state } from "./store/state";
import NotFoundPage from "./routes/NotFoundPage";
import { websocketService } from "./services/WebSocketService";
import SplashScreen from "./components/SplashScreen/SplashScreen";
import { useEffect, useState } from "react";
import { useConsoleCapture } from "./hooks/useConsoleCapture";
import ConsoleViewer from "./components/ConsoleViewer";

import moment from "moment-timezone";
import 'moment/locale/it';
import 'moment/dist/locale/it';
import Cockpit from "./routes/Cockpit/Cockpit";
import { app as appConfig, graphics } from "./config/environment";

export default function App() {
  moment.tz.setDefault(appConfig.timezone);
  moment.locale(appConfig.locale);

  // Console capture hook
  const { logs, clearLogs } = useConsoleCapture();
  const [isConsoleOpen, setIsConsoleOpen] = useState(false);

  // Inizializza la connessione WebSocket immediatamente
  useEffect(() => {
    websocketService.connect();

    // Cleanup alla chiusura dell'app
    return () => {
      websocketService.disconnect();
    };
  }, []);

  // Applica classe qualità grafica sul body
  useEffect(() => {
    document.body.classList.remove('quality-1', 'quality-2', 'quality-3');
    document.body.classList.add(`quality-${graphics.quality}`);
  }, []);

  // Gestione tasti di scelta rapida
  useEffect(() => {
    const handleKeyDown = (event: KeyboardEvent) => {
      // Apri console con tasto "d"
      if (event.key.toLowerCase() === 'd' && !isConsoleOpen) {
        setIsConsoleOpen(true);
      }
      // Chiudi console con ESC
      if (event.key === 'Escape' && isConsoleOpen) {
        setIsConsoleOpen(false);
      }
      // Ricarica applicazione con tasto "r"
      if (event.key.toLowerCase() === 'r' && !isConsoleOpen) {
        event.preventDefault();
        window.location.reload();
      }
    };

    document.addEventListener('keydown', handleKeyDown);
    return () => document.removeEventListener('keydown', handleKeyDown);
  }, [isConsoleOpen]);

  return (
    <SplashScreen>
      <HelmetProvider>
        <BrowserRouter>
            <Routes>
              <Route index element={
                <>
                  <Cockpit />
                  
                  {/* Console Viewer Modal - apertura con tasto "d" */}
                  <ConsoleViewer
                    isOpen={isConsoleOpen}
                    onClose={() => setIsConsoleOpen(false)}
                    logs={logs}
                    onClear={clearLogs}
                  />
                </>
              } />
              <Route path='*' element={<NotFoundPage />} />
            </Routes>
        </BrowserRouter>
      </HelmetProvider>
    </SplashScreen>
  );
}
