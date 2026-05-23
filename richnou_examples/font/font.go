package main

import (
	"gogpugg/internal/tools"
	"log"

	"github.com/gogpu/gg"
	"github.com/gogpu/gg/text"
)

var w = 500
var h = 500

var fontPath = "/Users/richnou/htdocs/data/fonts/ComicMono.ttf"

func main() {
	source, err := text.NewFontSourceFromFile(fontPath)
	if err != nil {
		log.Fatalf("Failed to load font: %v", err)
	}
	defer func() { _ = source.Close() }()

	dc := gg.NewContext(w, h)
	// dc.ClearWithColor(gg.White)

	x := 20.
	y := 0.

	/*
	* ORDER OF FUNC.
	 */
	y = 100.
	dc.SetFont(source.Face(85))
	dc.SetLineWidth(2.0)
	dc.SetHexColor("#FF8F00")
	dc.StrokeString("R", x, y) // STROKE
	dc.SetHexColor("#90CAF9")
	dc.DrawString("R", x, y) // FILL
	dc.SetFont(source.Face(17))
	dc.SetHexColor("#607D8B")
	dc.DrawString("StrokeString > DrawString", x+60, y) // STROKE

	y = 200.
	dc.SetFont(source.Face(85))
	dc.SetLineWidth(2.0)
	dc.SetHexColor("#FF8F00")
	dc.DrawString("R", x, y) // STROKE
	dc.SetHexColor("#90CAF9")
	dc.StrokeString("R", x, y) // FILL
	dc.SetFont(source.Face(17))
	dc.SetHexColor("#607D8B")
	dc.DrawString("DrawString > StrokeString", x+60, y) // STROKE

	/*
	* BORDER... PIXEL PERFECT (soon)
	 */
	dc.SetLineWidth(15.0)

	dc.SetFont(source.Face(300))
	dc.SetHexColor("#FF8F00")
	dc.StrokeString("R", x, 450) // STROKE

	dc.SetTextMode(gg.TextModeAliased) // pixel-perfect text

	dc.SetFont(source.Face(300))
	dc.SetHexColor("#FF8F00")
	dc.StrokeString("R", x+200, 450) // STROKE
	dc.DrawString("R", x+300, 450)   //

	y = 480

	dc.SetFont(source.Face(17))
	dc.SetHexColor("#607D8B")
	dc.DrawString("SetLineWidth(15.0)", x, y)
	dc.DrawString("gg.TextModeAliased (soon)", x+200, y)

	tools.Save(dc, "font", "font")
}
