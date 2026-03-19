package main

import (
	"gogpugg/internal/grid"
	"gogpugg/internal/tools"

	"github.com/gogpu/gg"
)

var w = 512.
var h = 512.

func lineMulti(dc *gg.Context, y float64, join gg.LineJoin, color string) {
	dc.SetLineJoin(join)
	dc.SetHexColor(color)
	dc.MoveTo(10., y)
	dc.LineTo(100., y)
	dc.LineTo(130., y+50)
	dc.LineTo(130., y+10)
	dc.LineTo(250., y)
	dc.LineTo(450., y)
	dc.Stroke()
}
func line(dc *gg.Context, y float64, color string) {
	dc.SetHexColor(color)
	dc.DrawLine(10., y, 492, y)
	dc.Stroke()
}

func main() {
	dc := gg.NewContext(int(w), int(h))
	dc.ClearWithColor(gg.Hex("#FFF8E1"))
	grid.DrawGrid(dc, 0, w, 0, h, 10, "#FFE0B2", 100, "#FFB74D")

	dc.SetLineWidth(10)

	// Sharp corners (default)
	dc.SetHexColor("#F5D216")
	dc.SetLineJoin(gg.LineJoinMiter)
	dc.DrawRectangle(50, 50, 100, 80)
	dc.Stroke()

	// Rounded corners
	dc.SetHexColor("#0177E1")
	dc.SetLineJoin(gg.LineJoinRound)
	dc.DrawRectangle(200, 50, 100, 80)
	dc.Stroke()

	// Beveled (flat-cut) corners
	dc.SetHexColor("#FD3502")
	dc.SetLineJoin(gg.LineJoinBevel)
	dc.DrawRectangle(350, 50, 100, 80)
	dc.Stroke()

	lineMulti(dc, 200., gg.LineJoinMiter, "#F5D216")
	lineMulti(dc, 300., gg.LineJoinRound, "#0177E1")
	lineMulti(dc, 400., gg.LineJoinBevel, "#FD3502")

	tools.Save(dc, "join", "join2")
}
