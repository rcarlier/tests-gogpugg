package main

import (
	"gogpugg/internal/grid"

	"github.com/gogpu/gg"
)

var w = 900.
var h = 500.

var col = "#37474F"

func drawRect(dc *gg.Context, x, y, lw float64) {
	dc.DrawRectangle(x, y, 100, 100)
	dc.SetHexColor(col)
	dc.SetLineWidth(lw)
	dc.Stroke()
}

func drawCirc(dc *gg.Context, x, y, lw float64) {
	dc.DrawCircle(x, y, 75)
	dc.SetHexColor(col)
	dc.SetLineWidth(lw)
	dc.Stroke()
}

func main() {
	dc := gg.NewContext(int(w), int(h))
	defer dc.Close()
	dc.ClearWithColor(gg.White)
	dc.SetHexColor("#000")
	grid.DrawGrid(dc, 0, w, 0, h, 10, "#FFE0B2", 100, "#FFB74D")

	drawRect(dc, 100, 100, 1)
	drawRect(dc, 300, 100, 5)
	drawRect(dc, 500, 100, 10)
	drawRect(dc, 700, 100, 15)

	drawCirc(dc, 150, 350, 1)
	drawCirc(dc, 350, 350, 5)
	drawCirc(dc, 550, 350, 10)
	drawCirc(dc, 750, 350, 15)

	dc.SavePNG("20260523/output.png")
}
