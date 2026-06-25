/*
 * PandaOS — Tachometer Prototype (raylib-go)
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

package main

import (
	"fmt"
	"math"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	W      = 420
	H      = 420
	MaxRPM = 8000
	ArcDeg = 288.0 // 80% of 360° — same cap as the React component
)

// Design tokens matching the React/SCSS palette
var (
	colBg    = rl.Color{R: 7, G: 8, B: 8, A: 255}
	colTrack = rl.Color{R: 30, G: 30, B: 40, A: 200}
	colBorder = rl.Color{R: 255, G: 255, B: 255, A: 30}
	colLabel  = rl.Color{R: 255, G: 255, B: 255, A: 80}
	colCyan   = rl.Color{R: 123, G: 212, B: 211, A: 153}
	colWhite  = rl.Color{R: 255, G: 255, B: 255, A: 153}
	colRed    = rl.Color{R: 255, G: 0, B: 0, A: 153}
	colInner  = rl.Color{R: 7, G: 8, B: 8, A: 220}

	// Fonts — loaded after InitWindow (GPU texture upload requires an active context).
	fontReg  rl.Font
	fontBold rl.Font
)

func main() {
	// MSAA 4× smooths all ring / circle geometry edges.
	// HighDPI renders at 2× pixel density on Retina screens — free extra AA for text.
	// Both must be set before InitWindow.
	rl.SetConfigFlags(rl.FlagMsaa4xHint | rl.FlagWindowHighdpi)
	rl.InitWindow(W, H, "PandaOS — Tachometer")
	rl.SetTargetFPS(60)

	// Load Orbitron at 128 px — the atlas is high-res enough for any render size
	// up to 128 px, and downscaling gives excellent quality with bilinear filtering.
	fontReg = rl.LoadFontEx("assets/fonts/Orbitron-Regular.ttf", 128, nil, 0)
	fontBold = rl.LoadFontEx("assets/fonts/Orbitron-Bold.ttf", 128, nil, 0)
	rl.SetTextureFilter(fontReg.Texture, rl.FilterBilinear)
	rl.SetTextureFilter(fontBold.Texture, rl.FilterBilinear)
	defer rl.UnloadFont(fontReg)
	defer rl.UnloadFont(fontBold)

	start := time.Now()

	for !rl.WindowShouldClose() {
		t := time.Since(start).Seconds()
		// Same sine-wave formula as the mock server
		rpm := float32(800 + 3700*(0.5+0.5*math.Sin(t*2*math.Pi/10)))

		rl.BeginDrawing()
		rl.ClearBackground(colBg)
		drawTachometer(rl.Vector2{X: W / 2, Y: H / 2}, 192, rpm)
		rl.DrawFPS(8, 8)
		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func drawTachometer(center rl.Vector2, outerR float32, rpm float32) {
	ringW := outerR * 0.20   // ~38 px ring band
	innerR := outerR - ringW // inner edge of the ring band

	midR  := outerR * 0.76 // .mid circle (75% from CSS)
	coreR := outerR * 0.50 // .inner circle (50% from CSS)

	pct     := rpm / MaxRPM
	fillDeg := pct * ArcDeg

	// ── fill color (same thresholds as Tachometer.tsx) ──────────────────────
	var fillCol rl.Color
	switch {
	case pct < 0.70:
		fillCol = colCyan
	case pct < 0.90:
		fillCol = colWhite
	default:
		fillCol = colRed
	}

	// ── 1. background track arc ──────────────────────────────────────────────
	rl.DrawRing(center, innerR, outerR, -90, -90+ArcDeg, 360, colTrack)

	// ── 2. colored fill arc ──────────────────────────────────────────────────
	if fillDeg > 0.5 {
		segs := int32(fillDeg)*2 + 1
		if segs < 12 {
			segs = 12
		}
		rl.DrawRing(center, innerR, outerR, -90, -90+fillDeg, segs, fillCol)
	}

	// ── 3. leading-edge dot ──────────────────────────────────────────────────
	if rpm > 200 {
		midRingR := innerR + ringW/2
		a := toRad(-90 + fillDeg)
		rl.DrawCircleV(
			rl.Vector2{X: center.X + midRingR*cos32(a), Y: center.Y + midRingR*sin32(a)},
			5, rl.White,
		)
	}

	// ── 4. outer ring border ─────────────────────────────────────────────────
	rl.DrawRingLines(center, innerR, outerR, 0, 360, 360, colBorder)

	// ── 5. mid dashed ring ───────────────────────────────────────────────────
	drawDashedCircle(center, midR, 128, colBorder)

	// ── 6. inner glass disc ──────────────────────────────────────────────────
	rl.DrawCircleSector(center, coreR, 0, 360, 360, colInner)
	rl.DrawCircleSectorLines(center, coreR, 0, 360, 360, colBorder)

	// ── 7. tick marks + labels at 1k / 3k / 5k / 7k RPM ────────────────────
	for _, m := range []struct {
		rpm   float32
		label string
	}{
		{1000, "1"}, {3000, "3"}, {5000, "5"}, {7000, "7"},
	} {
		angleDeg := -90 + (m.rpm/MaxRPM)*ArcDeg
		a := toRad(angleDeg)

		tickStart := rl.Vector2{X: center.X + (outerR-2)*cos32(a), Y: center.Y + (outerR-2)*sin32(a)}
		tickEnd   := rl.Vector2{X: center.X + (outerR+10)*cos32(a), Y: center.Y + (outerR+10)*sin32(a)}
		rl.DrawLineEx(tickStart, tickEnd, 1.5, colBorder)

		const labelFontSize = 13
		labelR := outerR + 24
		lx := center.X + labelR*cos32(a)
		ly := center.Y + labelR*sin32(a)
		drawTextCentered(fontReg, m.label, lx, ly, labelFontSize, 1, colLabel)
	}

	// ── 8. center RPM number (Bold, letter-spaced like the React h2) ─────────
	const rpmFontSize = 50
	drawTextCentered(fontBold, fmt.Sprintf("%d", int(rpm)), center.X, center.Y-14, rpmFontSize, 2, rl.White)

	// ── 9. "RPM" sub-label ───────────────────────────────────────────────────
	const unitFontSize = 12
	drawTextCentered(fontReg, "RPM", center.X, center.Y+rpmFontSize/2+2, unitFontSize, 2, colLabel)
}

// drawTextCentered draws text with the given font, centered on (cx, cy).
// spacing is the extra gap between glyphs in pixels (analogous to letter-spacing).
func drawTextCentered(font rl.Font, text string, cx, cy float32, fontSize, spacing float32, col rl.Color) {
	size := rl.MeasureTextEx(font, text, fontSize, spacing)
	rl.DrawTextEx(font, text,
		rl.Vector2{X: cx - size.X/2, Y: cy - size.Y/2},
		fontSize, spacing, col)
}

// drawDashedCircle approximates the CSS dashed border of the .mid element.
func drawDashedCircle(center rl.Vector2, radius float32, dashes int, col rl.Color) {
	step := 2 * math.Pi / float64(dashes)
	for i := 0; i < dashes; i++ {
		a1 := float64(i) * step
		a2 := a1 + step*0.45
		rl.DrawLineV(
			rl.Vector2{X: center.X + radius*float32(math.Cos(a1)), Y: center.Y + radius*float32(math.Sin(a1))},
			rl.Vector2{X: center.X + radius*float32(math.Cos(a2)), Y: center.Y + radius*float32(math.Sin(a2))},
			col,
		)
	}
}

func toRad(deg float32) float64 { return float64(deg) * math.Pi / 180 }
func cos32(rad float64) float32 { return float32(math.Cos(rad)) }
func sin32(rad float64) float32 { return float32(math.Sin(rad)) }
