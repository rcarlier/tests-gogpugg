package main

import (
	"gogpugg/internal/tools"

	"github.com/gogpu/gg"
)

var w = 512.
var h = 1024.
var radius = 50.

func main() {
	dc := gg.NewContext(400, 700)
	dc.ClearWithColor(gg.White)

	dc.SetLineWidth(5)
	dc.DrawLine(100, 0, 100, 800)
	dc.DrawLine(300, 0, 300, 800)
	dc.Stroke()

	dc.DrawCircle(100, 200, radius*2)
	dc.SetHexColor("#3498db")
	dc.Fill()

	dc.DrawCircle(100, 330, radius*2)
	dc.SetHexColor("#FF6F00aa")
	dc.Fill()

	dc.DrawCircle(100, 440, radius*2)
	dc.SetHexColor("#3498db")
	dc.Fill()

	stepY := 55.
	y := 50.
	dc.DrawCircle(300, y, radius)
	dc.SetRGBA(1.000, 0.627, 0.000, 0)
	dc.Fill()

	for i := 0.1; i <= 1; i += 0.1 {
		y += stepY
		dc.DrawCircle(300, y, radius)
		dc.SetRGBA(1.000, 0.627, 0.000, i)
		dc.Fill()
	}

	tools.Save(dc, "alpha", "alpha")
}
