package main

import (
	"gogpugg/internal/grid"

	"github.com/gogpu/gg"
	"github.com/gogpu/gg/text"
)

var w = 640.
var h = 480.

var ColorText = "#37474F"
var ColorAnchor = "#FFC107"

func main() {
	dc := gg.NewContext(int(w), int(h))
	defer dc.Close()
	dc.ClearWithColor(gg.White)
	dc.SetHexColor("#000")

	grid.DrawGrid(dc, 0, w, 0, h, 10, "#FFE0B2", 100, "#FFB74D")

	dc.DrawCircle(w/2, h/2, 6)
	dc.SetHexColor(ColorAnchor)
	dc.Fill()

	source, _ := text.NewFontSourceFromFile("Arial.ttf")
	defer source.Close()

	face := source.Face(45)
	dc.SetFont(face)

	dc.RotateAbout(.2, w/2, h/2)
	dc.SetHexColor(ColorText)
	dc.DrawString("Hello", w/2, h/2)

	dc.DrawStringAnchored("Hello", w/2, h/2, 0.5, 0.5)

	dc.SavePNG("20260305b/output.png")
}
