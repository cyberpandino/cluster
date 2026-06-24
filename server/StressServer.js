/*
 * PandaOS — Node.js Stress Server
 * Copyright (C) 2025  Cyberpandino
 *
 * Mirrors cmd/stress/main.go so both servers are tested under identical conditions.
 *
 * Environment variables:
 *   TICK_MS  – interval between emits in ms; 0 = uncapped (default 50)
 *   PAYLOAD  – "small" (5 params) or "full" (17 params, default "small")
 *   PORT     – listen port (default 3001)
 */

'use strict';

const http   = require('http');
const sio    = require('socket.io');

const TICK_MS      = parseInt(process.env.TICK_MS  ?? '50', 10);
const PAYLOAD_MODE = process.env.PAYLOAD ?? 'small';
const PORT         = parseInt(process.env.PORT ?? '3001', 10);

// ── animated state (time-based, same formula as Go server) ───────────────────

const START = Date.now();

function values() {
  const t = (Date.now() - START) / 1000;

  const rpm      = 800  + 3700 * (0.5 + 0.5 * Math.sin(t * 2 * Math.PI / 10));
  const speed    = rpm  / 60;
  const load     = 20   + 60   * (0.5 + 0.5 * Math.sin(t * 2 * Math.PI / 7));
  const coolant  = 85   + 10   * Math.sin(t * 2 * Math.PI / 60);
  const throttle = 5    + 40   * (0.5 + 0.5 * Math.sin(t * 2 * Math.PI / 8));
  const manifold = 30   + 50   * (0.5 + 0.5 * Math.sin(t * 2 * Math.PI / 9));
  const timing   = 8    + 12   * (0.5 + 0.5 * Math.sin(t * 2 * Math.PI / 12));
  const airTemp  = 22   + 3    * Math.sin(t * 2 * Math.PI / 120);
  const stFuel   = -3   + 4    * Math.sin(t * 2 * Math.PI / 15);
  const ltFuel   = 2    + 2    * Math.sin(t * 2 * Math.PI / 90);

  const small = {
    rpm:             r1(rpm),
    speed:           Math.floor(speed),
    coolant_temp:    r1(coolant),
    engine_load:     r1(load),
    battery_voltage: 13.8,
  };

  if (PAYLOAD_MODE !== 'full') return small;

  return {
    ...small,
    throttle_pos:           r1(throttle),
    manifold_pressure:      Math.floor(manifold),
    timing_advance:         r1(timing),
    air_temp:               r1(airTemp),
    short_fuel_trim_1:      r2(stFuel),
    long_fuel_trim_1:       r2(ltFuel),
    monitor_status:         0,
    fuel_system_status:     'CL',
    oxygen_sensors_present: 3,
    oxygen_sensor_1_1:      r2(0.72 + 0.1 * Math.sin(t * 2 * Math.PI / 3)),
    oxygen_sensor_2_1:      r2(0.68 + 0.1 * Math.sin(t * 2 * Math.PI / 4)),
    obd_standards:          'OBD-II',
    distance_with_mil_on:   Math.floor(t * 0.5),
  };
}

function buildPacket() {
  const raw = values();
  const parameters = {};
  for (const [k, v] of Object.entries(raw)) {
    parameters[k] = { value: v, success: true };
  }
  return { parameters };
}

// ── server ────────────────────────────────────────────────────────────────────

let totalEmitted = 0;
const server = http.createServer((req, res) => {
  if (req.url === '/stats') {
    const mem = process.memoryUsage();
    res.setHeader('Content-Type', 'application/json');
    res.end(JSON.stringify({
      heap_used_kb:   Math.round(mem.heapUsed  / 1024),
      heap_total_kb:  Math.round(mem.heapTotal / 1024),
      rss_kb:         Math.round(mem.rss       / 1024),
      total_emitted:  totalEmitted,
    }));
    return;
  }
  res.writeHead(200);
  res.end('Node.js stress server');
});

const io = sio(server, { cors: { origin: '*' } });

io.on('connection', (socket) => {
  let timer = null;
  socket.on('error', () => {}); // prevent crash on write-buffer overflow

  function emit() {
    try {
      socket.emit('obd-live', buildPacket());
      totalEmitted++;
    } catch (_) {}
  }

  if (TICK_MS <= 0) {
    // Uncapped: schedule via setImmediate so the event loop stays responsive.
    let active = true;
    function loop() {
      if (!active) return;
      emit();
      setImmediate(loop);
    }
    loop();
    socket.on('disconnect', () => { active = false; });
  } else {
    timer = setInterval(emit, TICK_MS);
    socket.on('disconnect', () => clearInterval(timer));
  }
});

server.listen(PORT, () => {
  console.log(`Node.js stress server :${PORT}  tick=${TICK_MS}ms  payload=${PAYLOAD_MODE}`);
});

// ── helpers ───────────────────────────────────────────────────────────────────

function r1(v) { return Math.round(v * 10)  / 10;  }
function r2(v) { return Math.round(v * 100) / 100; }
