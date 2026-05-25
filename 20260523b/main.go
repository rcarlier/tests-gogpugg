package main

import (
	"gogpugg/internal/grid"
	"math"

	"github.com/gogpu/gg"
)

var w = 1100.
var h = 200.

var col = "#37474F"

func drawRect(dc *gg.Context, x, y, lw float64) {
	dc.DrawRectangle(x, y, 100, 100)
	dc.SetHexColor(col)
	dc.SetLineWidth(lw)
	dc.Stroke()
}

func closeShape(dc *gg.Context, lw float64) {
	dc.SetHexColor(col)
	dc.SetLineWidth(lw)
	dc.Stroke()
}

func main() {
	dc := gg.NewContext(int(w), int(h))
	defer dc.Close()
	dc.ClearWithColor(gg.White)
	dc.SetHexColor("#000")

	grid.DrawQuickGrid(dc, w, h)

	dc.DrawRegularPolygon(6, 100, 100, 75, 0)
	closeShape(dc, 20)

	dc.DrawRoundedRectangle(200, 50, 150, 100, 20)
	closeShape(dc, 20)

	dc.DrawCircle(450, 100, 60)
	closeShape(dc, 20)

	dc.DrawEllipse(650, 100, 80, 50)
	closeShape(dc, 20)

	dc.DrawArc(820, 100, 50, 0, math.Pi*1.5)
	closeShape(dc, 20)

	dc.DrawRectangle(900, 50, 150, 100)
	closeShape(dc, 20)

	dc.SavePNG("20260523b/output.png")
}

/*

func drawShapes(dc *gg.Context) {
	// Lines
	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth(3)
	dc.DrawLine(50, 450, 750, 450)
	dc.Stroke()

	// Arc
	dc.SetRGB(0.8, 0, 0.8)
	dc.SetLineWidth(5)
	dc.DrawArc(650, 300, 60, 0, math.Pi*1.5)
	dc.Stroke()

	// Transformed shapes
	dc.Push()
	dc.Translate(400, 500)
	dc.Rotate(math.Pi / 4)
	dc.SetRGB(0.2, 0.6, 0.8)
	dc.DrawRectangle(-40, -40, 80, 80)
	dc.Fill()
	dc.Pop()
*/
