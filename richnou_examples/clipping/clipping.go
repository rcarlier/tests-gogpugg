package main

import (
	"gogpugg/internal/tools"

	"github.com/gogpu/gg"
)

var w = 512.
var h = 1024.

func reversePolygon(poly []float64) []float64 {
	n := len(poly)
	result := make([]float64, n)
	for i := 0; i < n; i += 2 {
		result[i] = poly[n-2-i]
		result[i+1] = poly[n-1-i]
	}
	return result
}

func DrawPolygonWithHoles(dc *gg.Context, ring []float64, holes [][]float64, fill string) {
	dc.MoveTo(ring[0], ring[1])
	for i := 2; i < len(ring); i += 2 {
		dc.LineTo(ring[i], ring[i+1])
	}
	dc.ClosePath()
	for _, hole := range holes {
		if len(hole) >= 6 {
			dc.MoveTo(hole[0], hole[1])
			for i := 2; i < len(hole); i += 2 {
				dc.LineTo(hole[i], hole[i+1])
			}
			dc.ClosePath()
		}
	}
	dc.SetHexColor(fill)
	dc.Fill()
}
func gradient(dc *gg.Context, color1, color2 string) {
	h := float64(dc.Height())
	g := gg.NewLinearGradientBrush(0, 0, 0, h)
	g.AddColorStop(0, gg.Hex(color1))
	g.AddColorStop(1, gg.Hex(color2))
	dc.SetFillBrush(g)
	dc.DrawRectangle(0, 0, float64(dc.Width()), h)
	dc.Fill()
}

func main() {
	dc := gg.NewContext(400, 400)
	dc.SetRGB(0.95, 0.95, 0.95)
	dc.Clear()
	gradient(dc, "#E3F2FD", "#0D47A1")

	ring := []float64{50, 50, 350, 50, 350, 350, 50, 350}
	hole1 := []float64{150, 150, 100, 200, 150, 250, 200, 200}
	hole2 := []float64{280, 150, 250, 200, 280, 250, 310, 200}
	hole3 := []float64{130, 280, 140, 310, 145, 315, 150, 320, 155, 322, 200, 325, 245, 322, 250, 320, 255, 315, 260, 310, 270, 280}
	DrawPolygonWithHoles(dc, ring, [][]float64{hole1, hole2, hole3}, "#FF6F00")
	tools.Save(dc, "clipping", "clipping")
}
