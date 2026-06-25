/*
 * PandaOS — Dashboard (raylib-go)
 * Copyright (C) 2025  Cyberpandino
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3.
 */

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math"
	"strings"
	"sync"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gorilla/websocket"
)

// ── Layout ────────────────────────────────────────────────────────────────────
// Mirrors Cockpit.scss: 1920×580, padding 10px 120px, gap 120px between 3×480px columns.
const (
	ScreenW = 1920
	ScreenH = 580

	CxLeft   = 360.0
	CxCenter = 960.0
	CxRight  = 1560.0
	CyGauge  = 290.0

	// Gauge is a rounded square matching the React design.
	// CSS: max-width 420px, border-radius 100px → roundness = 100/(420/2) ≈ 0.476
	GaugeSize        = 420.0
	gaugeHalf        = GaugeSize / 2 // 210
	gaugeRoundness   = 100.0 / gaugeHalf
	gaugeMidRound    = 80.0 / (GaugeSize * 0.75 / 2)  // .mid: 75% size, 80px radius
	gaugeInnerRound  = 50.0 / (GaugeSize * 0.50 / 2)  // .inner: 50% size, 50px radius

	MaxRPM  = 8000.0
	MaxKmh  = 160.0
	MaxTemp = 160.0
)

// ── Colour palette ────────────────────────────────────────────────────────────
var (
	colBg     = rl.Color{R: 7, G: 8, B: 8, A: 255}
	colBorder = rl.Color{R: 255, G: 255, B: 255, A: 30}
	colLabel  = rl.Color{R: 255, G: 255, B: 255, A: 80}
	colCyan   = rl.Color{R: 123, G: 212, B: 211, A: 153}
	colWhite  = rl.Color{R: 255, G: 255, B: 255, A: 153}
	colRed    = rl.Color{R: 255, G: 0, B: 0, A: 153}

	fontReg      rl.Font
	fontBold     rl.Font
	warnTextures []rl.Texture2D
	carTex       rl.Texture2D

	// Warning overlay: shows large icon for 5 s when a warning newly activates.
	overlayMu  sync.Mutex
	overlayKey string
	overlayEnd time.Time

	clockMu  sync.Mutex
	clockStr string
)

// ── Warning light definitions ─────────────────────────────────────────────────
// key matches server gpio-mapping.js; file is the PNG in assets/icons/warnings/.
var warningDefs = []struct {
	key   string
	file  string
	label string
	col   rl.Color
}{
	{"doors",        "doors",        "DOOR", rl.Color{R: 255, G: 165, B: 0, A: 255}},
	{"battery",      "battery",      "BAT",  rl.Color{R: 255, G: 0, B: 0, A: 255}},
	{"brakeSystem",  "brakeSystem",  "BRK",  rl.Color{R: 255, G: 0, B: 0, A: 255}},
	{"engineOil",    "engineOil",    "OIL",  rl.Color{R: 255, G: 0, B: 0, A: 255}},
	{"engineColant", "engineColant", "COOL", rl.Color{R: 255, G: 0, B: 0, A: 255}},
	{"fuel",         "fuel",         "FUEL", rl.Color{R: 255, G: 255, B: 0, A: 255}},
	{"injectors",    "injectors",    "INJ",  rl.Color{R: 255, G: 0, B: 0, A: 255}},
	{"warning",      "warning",      "WARN", rl.Color{R: 255, G: 165, B: 0, A: 255}},
	{"hazard",       "hazard",       "HAZ",  rl.Color{R: 255, G: 165, B: 0, A: 255}},
	{"turnSignals",  "turnSignals",  "TURN", rl.Color{R: 0, G: 255, B: 0, A: 255}},
	{"light",        "light",        "LITE", rl.Color{R: 255, G: 255, B: 0, A: 255}},
	{"lowBeam",      "lowBeam",      "LO",   rl.Color{R: 0, G: 255, B: 0, A: 255}},
	{"highBeam",     "highBeam",     "HI",   rl.Color{R: 0, G: 0, B: 255, A: 255}},
	{"fogLight",     "fogLight",     "FOG",  rl.Color{R: 0, G: 255, B: 0, A: 255}},
	{"hood",         "hood",         "HOOD", rl.Color{R: 255, G: 165, B: 0, A: 255}},
	{"rearDefrost",  "rearDefrost",  "DEF",  rl.Color{R: 0, G: 255, B: 0, A: 255}},
}

// ── Live OBD state ────────────────────────────────────────────────────────────
var (
	rawMu sync.RWMutex
	raw   struct {
		RPM      float32
		Speed    float32
		Coolant  float32
		Battery  float32
		Fuel     float32
		Km       float32
		Warnings     map[string]bool
		PrevWarnings map[string]bool
	}
	connected bool

	disp struct {
		RPM     float32
		Speed   float32
		Coolant float32
		Fuel    float32
	}
)

// ── Entry point ───────────────────────────────────────────────────────────────

func main() {
	serverAddr := flag.String("server",
		"ws://127.0.0.1:3001/socket.io/?EIO=4&transport=websocket",
		"Socket.IO WebSocket URL")
	flag.Parse()

	raw.Battery = 14.2
	raw.Km = 138000
	raw.Warnings = make(map[string]bool)
	raw.PrevWarnings = make(map[string]bool)

	clockStr = time.Now().Format("15:04")
	go func() {
		for {
			now := time.Now()
			next := now.Truncate(time.Minute).Add(time.Minute)
			time.Sleep(time.Until(next))
			clockMu.Lock()
			clockStr = time.Now().Format("15:04")
			clockMu.Unlock()
		}
	}()

	go connectLoop(*serverAddr)

	rl.SetConfigFlags(rl.FlagMsaa4xHint | rl.FlagWindowHighdpi)
	rl.InitWindow(ScreenW, ScreenH, "PandaOS")
	rl.SetTargetFPS(60)

	fontReg = rl.LoadFontEx("assets/fonts/Orbitron-Regular.ttf", 128, nil, 0)
	fontBold = rl.LoadFontEx("assets/fonts/Orbitron-Bold.ttf", 128, nil, 0)
	rl.SetTextureFilter(fontReg.Texture, rl.FilterBilinear)
	rl.SetTextureFilter(fontBold.Texture, rl.FilterBilinear)
	defer rl.UnloadFont(fontReg)
	defer rl.UnloadFont(fontBold)

	warnTextures = make([]rl.Texture2D, len(warningDefs))
	for i, w := range warningDefs {
		warnTextures[i] = rl.LoadTexture("assets/icons/warnings/" + w.file + ".png")
		rl.SetTextureFilter(warnTextures[i], rl.FilterBilinear)
	}
	defer func() {
		for _, t := range warnTextures {
			rl.UnloadTexture(t)
		}
	}()

	// Load wireframe car image and invert so the white background becomes black.
	img := rl.LoadImage("assets/images/panda.png")
	rl.ImageColorInvert(img)
	carTex = rl.LoadTextureFromImage(img)
	rl.UnloadImage(img)
	rl.SetTextureFilter(carTex, rl.FilterBilinear)
	defer rl.UnloadTexture(carTex)

	for !rl.WindowShouldClose() {
		updateSmoothing(rl.GetFrameTime())
		rl.BeginDrawing()
		rl.ClearBackground(colBg)
		drawDashboard()
		rl.EndDrawing()
	}

	rl.CloseWindow()
}

// ── Smoothing ─────────────────────────────────────────────────────────────────

func updateSmoothing(dt float32) {
	rawMu.RLock()
	tRPM     := raw.RPM
	tSpeed   := raw.Speed
	tCoolant := raw.Coolant
	tFuel    := raw.Fuel
	rawMu.RUnlock()

	disp.RPM     = lerp(disp.RPM, tRPM, dt, 0.12)
	disp.Speed   = lerp(disp.Speed, tSpeed, dt, 0.15)
	disp.Coolant = lerp(disp.Coolant, tCoolant, dt, 0.50)
	disp.Fuel    = lerp(disp.Fuel, tFuel, dt, 2.00)
}

func lerp(current, target, dt, tau float32) float32 {
	factor := float32(1 - math.Exp(-float64(dt)/float64(tau)))
	return current + (target-current)*factor
}

// ── Dashboard render ──────────────────────────────────────────────────────────

func drawDashboard() {
	// Large teal glow at centre-bottom (matches .pageCockpit::before).
	rl.DrawCircleV(rl.Vector2{X: CxCenter, Y: ScreenH + 300}, 700,
		rl.Color{R: 123, G: 212, B: 211, A: 12})

	// ── Car wireframe image ───────────────────────────────────────────────────
	// Drawn first (behind gauges). Inverted PNG: white→black, wireframe→light grey.
	if carTex.Width > 0 {
		aspect := float32(carTex.Width) / float32(carTex.Height)
		imgH := float32(ScreenH) * 0.68
		imgW := imgH * aspect
		rl.DrawTexturePro(carTex,
			rl.Rectangle{Width: float32(carTex.Width), Height: float32(carTex.Height)},
			rl.Rectangle{X: CxCenter - imgW/2, Y: float32(ScreenH)*0.52 - imgH/2, Width: imgW, Height: imgH},
			rl.Vector2{}, 0,
			rl.Color{R: 255, G: 255, B: 255, A: 55})
	}

	// ── Top bar: clock + external temperature ─────────────────────────────────
	clockMu.Lock()
	clk := clockStr
	clockMu.Unlock()
	drawTextLeft(fontReg, clk, 128, 18, 16, 1, rl.Color{R: 255, G: 255, B: 255, A: 200})
	drawTextRight(fontReg, "21 °C", float32(ScreenW)-128, 18, 16, 1, rl.Color{R: 255, G: 255, B: 255, A: 200})

	// ── Left column: Tachometer ───────────────────────────────────────────────
	drawRoundedGauge(CxLeft, CyGauge, disp.RPM, MaxRPM, "RPM",
		[4]cornerLabel{
			{"1",  CxLeft + gaugeHalf - 18, CyGauge - gaugeHalf + 18},
			{"3",  CxLeft + gaugeHalf - 18, CyGauge + gaugeHalf - 18},
			{"5",  CxLeft - gaugeHalf + 18, CyGauge + gaugeHalf - 18},
			{"7",  CxLeft - gaugeHalf + 18, CyGauge - gaugeHalf + 18},
		})

	rawMu.RLock()
	fuel := disp.Fuel
	rawMu.RUnlock()
	drawTextCentered(fontReg, fmt.Sprintf("FUEL  %.0f%%   250 km", fuel),
		CxLeft, float32(ScreenH)-20, 11, 2, colLabel)

	// ── Centre column: temp bar + warnings + odometer ─────────────────────────
	drawTemperatureBar(CxCenter, 28, 380)

	rawMu.RLock()
	warnings := raw.Warnings
	km       := raw.Km
	rawMu.RUnlock()

	drawWarningRow(CxCenter, 432, warnings)
	drawTextCentered(fontBold, fmt.Sprintf("%.0f km", km),
		CxCenter, float32(ScreenH)-20, 13, 2, colLabel)

	// ── Right column: Speedometer ─────────────────────────────────────────────
	drawRoundedGauge(CxRight, CyGauge, disp.Speed, MaxKmh, "km/h",
		[4]cornerLabel{
			{"20",  CxRight + gaugeHalf - 24, CyGauge - gaugeHalf + 18},
			{"70",  CxRight + gaugeHalf - 24, CyGauge + gaugeHalf - 18},
			{"100", CxRight - gaugeHalf + 24, CyGauge + gaugeHalf - 18},
			{"150", CxRight - gaugeHalf + 24, CyGauge - gaugeHalf + 18},
		})

	rawMu.RLock()
	bat := raw.Battery
	rawMu.RUnlock()
	drawTextCentered(fontReg, fmt.Sprintf("BATTERY  %.1fV   0V", bat),
		CxRight, float32(ScreenH)-20, 11, 2, colLabel)

	// ── Connection dot ────────────────────────────────────────────────────────
	rawMu.RLock()
	conn := connected
	rawMu.RUnlock()
	connCol := rl.Color{R: 80, G: 80, B: 80, A: 200}
	if conn {
		connCol = rl.Color{R: 123, G: 212, B: 211, A: 200}
	}
	rl.DrawCircleV(rl.Vector2{X: float32(ScreenW) - 20, Y: 20}, 5, connCol)

	// ── Warning overlay (on top of everything, shown for 5 s) ─────────────────
	overlayMu.Lock()
	oKey := overlayKey
	oEnd := overlayEnd
	overlayMu.Unlock()
	if oKey != "" && time.Now().Before(oEnd) {
		drawWarningOverlay(oKey)
	}
}

// ── Rounded-square conic-gradient gauge ───────────────────────────────────────
// Matches the React Tachometer / Odometer: a rounded square whose background is
// a conic-gradient that fills 0–80 % of the circle clockwise from the top.

type cornerLabel struct {
	text string
	x, y float32
}

func drawRoundedGauge(cx, cy, value, maxValue float32, unit string, corners [4]cornerLabel) {
	half := float32(gaugeHalf)
	rect := rl.Rectangle{X: cx - half, Y: cy - half, Width: GaugeSize, Height: GaugeSize}
	center := rl.Vector2{X: cx, Y: cy}

	pct       := clamp01(value / maxValue)
	fillAngle := pct * 0.8 * 360 // max 80 % of full circle → 288 ° at max value

	fillCol := gaugeColor(pct)

	// 1. Wrapper glow (replaces linear-gradient background on .wrapper)
	rl.DrawRectangleRounded(rect, gaugeRoundness, 120,
		rl.Color{R: 123, G: 212, B: 211, A: 6})

	// 2. Crosshair lines (.wrapper::before / ::after)
	lineCol := rl.Color{R: 255, G: 255, B: 255, A: 18}
	rl.DrawLineEx(rl.Vector2{X: cx - half, Y: cy}, rl.Vector2{X: cx + half, Y: cy}, 1, lineCol)
	rl.DrawLineEx(rl.Vector2{X: cx, Y: cy - half}, rl.Vector2{X: cx, Y: cy + half}, 1, lineCol)

	// 3. Filled rounded rect — the "colored" portion of the conic gradient.
	rl.DrawRectangleRounded(rect, gaugeRoundness, 120, fillCol)

	// 4. Mask sector in background colour — hides the "transparent" portion.
	//    Conic starts at top (−90°) clockwise; unfilled arc runs from
	//    (−90° + fillAngle) back to the top (270° = −90° + 360°).
	maskStart := float32(-90) + fillAngle
	maskEnd   := float32(270)
	if maskStart < maskEnd {
		segs := int32(maskEnd-maskStart) + 1
		if segs < 12 {
			segs = 12
		}
		bigR := half * 1.6 // exceeds corner diagonal (half × √2 ≈ half × 1.414)
		rl.DrawCircleSector(center, bigR, maskStart, maskEnd, segs, colBg)
	}

	// 5. Outer border (1.5 px, white @ 12 % opacity — matches CSS border)
	rl.DrawRectangleRoundedLinesEx(rect, gaugeRoundness, 120, 1.5, colBorder)

	// 6. Mid ring — .mid: 75 % size, border-radius 80 px, dashed white border
	midS    := float32(GaugeSize * 0.75)
	midHalf := midS / 2
	midRect := rl.Rectangle{X: cx - midHalf, Y: cy - midHalf, Width: midS, Height: midS}
	rl.DrawRectangleRoundedLinesEx(midRect, gaugeMidRound, 120, 1.5,
		rl.Color{R: 255, G: 255, B: 255, A: 20})

	// 7. Inner disc — .inner: 50 % size, solid dark background + border
	innerS    := float32(GaugeSize * 0.50)
	innerHalf := innerS / 2
	innerRect := rl.Rectangle{X: cx - innerHalf, Y: cy - innerHalf, Width: innerS, Height: innerS}
	rl.DrawRectangleRounded(innerRect, gaugeInnerRound, 120, colBg)
	rl.DrawRectangleRoundedLinesEx(innerRect, gaugeInnerRound, 120, 1.5, colBorder)

	// 8. Corner labels (counter1–4 in CSS: top-right, bottom-right, bottom-left, top-left)
	cornerCol := rl.Color{R: 255, G: 255, B: 255, A: 80}
	for _, c := range corners {
		drawTextCentered(fontReg, c.text, c.x, c.y, 18, 1, cornerCol)
	}

	// 9. Main value + unit label
	drawTextCentered(fontBold, fmt.Sprintf("%.0f", value), cx, cy-14, 52, 2, rl.White)
	drawTextCentered(fontReg, unit, cx, cy+30, 11, 2, colLabel)
}

// ── Temperature bar (top of centre column) ────────────────────────────────────

func drawTemperatureBar(cx, y, width float32) {
	rawMu.RLock()
	temp := raw.Coolant
	rawMu.RUnlock()

	pct     := clamp01(temp / MaxTemp)
	fillCol := temperatureColor(temp)

	const (
		barH    = 6.0
		padSide = 28.0
	)
	barX := cx - width/2 + padSide
	barW := width - 2*padSide

	// Min / max labels
	drawTextCentered(fontReg, "0°", cx-width/2+10, y+barH/2, 11, 1, colLabel)
	drawTextCentered(fontReg, fmt.Sprintf("%.0f°", MaxTemp), cx+width/2-10, y+barH/2, 11, 1, colLabel)

	// Track
	rl.DrawRectangle(int32(barX), int32(y), int32(barW), barH,
		rl.Color{R: 30, G: 30, B: 40, A: 200})
	rl.DrawRectangleLinesEx(
		rl.Rectangle{X: barX, Y: y, Width: barW, Height: barH}, 1, colBorder)

	// Fill
	if fillW := pct * barW; fillW > 1 {
		rl.DrawRectangle(int32(barX), int32(y), int32(fillW), barH, fillCol)
	}

	// Reading label below the fill edge
	if temp > 1 {
		labelX := barX + pct*barW
		if labelX > barX+barW-55 {
			labelX = barX + barW - 55
		}
		drawTextCentered(fontReg, fmt.Sprintf("↓ %.0f °", temp),
			labelX+22, y+16, 10, 1, colLabel)
	}
}

// ── Warning icon row ──────────────────────────────────────────────────────────

func drawWarningRow(cx, y float32, warnings map[string]bool) {
	const (
		iconSize = 20.0
		gap      = 14.0
	)
	n      := len(warningDefs)
	totalW := float32(n)*iconSize + float32(n-1)*gap
	startX := cx - totalW/2

	for i, w := range warningDefs {
		if i >= len(warnTextures) {
			break
		}
		x := startX + float32(i)*(iconSize+gap) + iconSize/2
		active := warnings[w.key]

		tex := warnTextures[i]
		src  := rl.Rectangle{Width: float32(tex.Width), Height: float32(tex.Height)}
		dest := rl.Rectangle{X: x - iconSize/2, Y: y - iconSize/2, Width: iconSize, Height: iconSize}

		if active {
			// Glow background behind icon
			rl.DrawRectangle(
				int32(x-iconSize/2-4), int32(y-iconSize/2-4),
				int32(iconSize+8), int32(iconSize+8),
				rl.Color{R: w.col.R, G: w.col.G, B: w.col.B, A: 30})
			rl.DrawTexturePro(tex, src, dest, rl.Vector2{}, 0,
				rl.Color{R: w.col.R, G: w.col.G, B: w.col.B, A: 220})
		} else {
			// Inactive: white tint at low alpha (CSS: brightness(0) invert + opacity 0.5)
			rl.DrawTexturePro(tex, src, dest, rl.Vector2{}, 0,
				rl.Color{R: 255, G: 255, B: 255, A: 45})
		}
	}
}

// ── Warning overlay (large centred icon when a warning newly activates) ────────

func drawWarningOverlay(key string) {
	texIdx := -1
	for i, w := range warningDefs {
		if w.key == key {
			texIdx = i
			break
		}
	}
	if texIdx < 0 || texIdx >= len(warnTextures) {
		return
	}

	// Dim backdrop
	rl.DrawRectangle(0, 0, ScreenW, ScreenH, rl.Color{A: 120})

	// Large icon
	const iconSize = 80.0
	tex := warnTextures[texIdx]
	w   := warningDefs[texIdx]
	rl.DrawTexturePro(tex,
		rl.Rectangle{Width: float32(tex.Width), Height: float32(tex.Height)},
		rl.Rectangle{
			X: CxCenter - iconSize/2, Y: float32(ScreenH)/2 - iconSize/2,
			Width: iconSize, Height: iconSize,
		},
		rl.Vector2{}, 0,
		rl.Color{R: w.col.R, G: w.col.G, B: w.col.B, A: 255})
}

// ── Socket.IO client ──────────────────────────────────────────────────────────

func connectLoop(addr string) {
	backoff := time.Second
	for {
		if err := connectOnce(addr); err != nil {
			log.Printf("ws: disconnected (%v) — retry in %v", err, backoff)
		}
		rawMu.Lock()
		connected = false
		rawMu.Unlock()
		time.Sleep(backoff)
		if backoff < 30*time.Second {
			backoff *= 2
		} else {
			backoff = 30 * time.Second
		}
	}
}

func connectOnce(addr string) error {
	conn, _, err := websocket.DefaultDialer.Dial(addr, nil)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, msg, err := conn.ReadMessage()
	if err != nil {
		return fmt.Errorf("read open: %w", err)
	}
	if len(msg) == 0 || msg[0] != '0' {
		return fmt.Errorf("unexpected EIO open: %s", msg)
	}
	if err := conn.WriteMessage(websocket.TextMessage, []byte("40")); err != nil {
		return fmt.Errorf("send connect: %w", err)
	}
	if _, _, err := conn.ReadMessage(); err != nil {
		return fmt.Errorf("read connect ack: %w", err)
	}

	rawMu.Lock()
	connected = true
	rawMu.Unlock()
	log.Printf("ws: connected to %s", addr)

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return fmt.Errorf("read: %w", err)
		}
		handlePacket(string(msg), conn)
	}
}

func handlePacket(msg string, conn *websocket.Conn) {
	if msg == "2" {
		_ = conn.WriteMessage(websocket.TextMessage, []byte("3"))
		return
	}
	if !strings.HasPrefix(msg, "42") {
		return
	}
	var arr []json.RawMessage
	if err := json.Unmarshal([]byte(msg[2:]), &arr); err != nil || len(arr) < 2 {
		return
	}
	var name string
	if err := json.Unmarshal(arr[0], &name); err != nil {
		return
	}

	switch name {
	case "obd-live":
		var payload struct {
			Parameters map[string]struct {
				Value   float64 `json:"value"`
				Success bool    `json:"success"`
			} `json:"parameters"`
		}
		if json.Unmarshal(arr[1], &payload) != nil {
			return
		}
		rawMu.Lock()
		if p, ok := payload.Parameters["rpm"]; ok && p.Success {
			raw.RPM = float32(p.Value)
		}
		if p, ok := payload.Parameters["speed"]; ok && p.Success {
			raw.Speed = float32(p.Value)
		}
		if p, ok := payload.Parameters["coolant_temp"]; ok && p.Success {
			raw.Coolant = float32(p.Value)
		}
		if p, ok := payload.Parameters["battery_voltage"]; ok && p.Success {
			raw.Battery = float32(p.Value)
		}
		rawMu.Unlock()

	case "fuel-level":
		var level float64
		if json.Unmarshal(arr[1], &level) == nil {
			rawMu.Lock()
			raw.Fuel = float32(level)
			rawMu.Unlock()
		}

	case "gpio-warnings":
		var warnings map[string]bool
		if json.Unmarshal(arr[1], &warnings) != nil {
			return
		}
		rawMu.Lock()
		prev := raw.PrevWarnings
		raw.PrevWarnings = raw.Warnings
		raw.Warnings = warnings
		rawMu.Unlock()

		// Show overlay for the first newly-active warning.
		overlayMu.Lock()
		for key, active := range warnings {
			if active && !prev[key] {
				overlayKey = key
				overlayEnd = time.Now().Add(5 * time.Second)
				break
			}
		}
		overlayMu.Unlock()
	}
}

// ── Colour helpers ────────────────────────────────────────────────────────────

func gaugeColor(pct float32) rl.Color {
	switch {
	case pct < 0.70:
		return colCyan
	case pct < 0.90:
		return colWhite
	default:
		return colRed
	}
}

func temperatureColor(t float32) rl.Color {
	if t > 100 {
		return rl.Color{R: 255, G: 0, B: 0, A: 153}
	}
	if t < 0.7*MaxTemp {
		return colCyan
	}
	return colWhite
}

// ── Drawing helpers ───────────────────────────────────────────────────────────

func drawTextCentered(font rl.Font, text string, cx, cy, fontSize, spacing float32, col rl.Color) {
	sz := rl.MeasureTextEx(font, text, fontSize, spacing)
	rl.DrawTextEx(font, text, rl.Vector2{X: cx - sz.X/2, Y: cy - sz.Y/2}, fontSize, spacing, col)
}

func drawTextLeft(font rl.Font, text string, x, cy, fontSize, spacing float32, col rl.Color) {
	sz := rl.MeasureTextEx(font, text, fontSize, spacing)
	rl.DrawTextEx(font, text, rl.Vector2{X: x, Y: cy - sz.Y/2}, fontSize, spacing, col)
}

func drawTextRight(font rl.Font, text string, rightX, cy, fontSize, spacing float32, col rl.Color) {
	sz := rl.MeasureTextEx(font, text, fontSize, spacing)
	rl.DrawTextEx(font, text, rl.Vector2{X: rightX - sz.X, Y: cy - sz.Y/2}, fontSize, spacing, col)
}

// ── Math helpers ──────────────────────────────────────────────────────────────

func clamp01(v float32) float32 {
	if v < 0 {
		return 0
	}
	if v > 1 {
		return 1
	}
	return v
}
