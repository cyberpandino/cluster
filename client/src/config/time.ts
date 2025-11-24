/**
 * Funzioni di utilità per la gestione dell'ora
 */

import { app as appConfig } from "./environment";

export type TimeFormat = "24h" | "12h";

/**
 * Restituisce l'ora formattata in base al formato configurato
 */
export function formatTime(date: Date, format: TimeFormat = appConfig.timeFormat): string {
  const hours = date.getHours();
  const minutes = date.getMinutes();

  if (format === "12h") {
    const period = hours >= 12 ? "PM" : "AM";
    const hours12 = hours % 12 || 12;
    return `${padNumber(hours12)}:${padNumber(minutes)} ${period}`;
  }

  // Formato 24h
  return `${padNumber(hours)}:${padNumber(minutes)}`;
}

function padNumber(value: number): string {
  return value.toString().padStart(2, "0");
}