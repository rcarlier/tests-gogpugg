package main

import (
	"gogpugg/internal/tools"
	"log"

	"github.com/gogpu/gg"
	"github.com/gogpu/gg/text"
)

var w = 800
var h = 600

var fontPath = "/Users/richnou/htdocs/data/fonts/ComicMono.ttf"

func main() {
	source, err := text.NewFontSourceFromFile(fontPath)
	if err != nil {
		log.Fatalf("Failed to load font: %v", err)
	}
	defer func() { _ = source.Close() }()

	dc := gg.NewContext(w, h)
	dc.ClearWithColor(gg.White)

	x := 20.
	y := 0.

	/*
	* ORDER OF FUNC.
	 */

	colorStroke := "#FF8F00"
	colorFill := "#90CAF9"
	textToWrite := "RW"
	y = 210.
	dc.SetFont(source.Face(285))
	dc.SetLineJoin(gg.LineJoinRound)

	// dc.SetLineCap(gg.LineCapSquare)

	dc.SetLineWidth(10.0)
	dc.SetHexColor(colorStroke)
	dc.StrokeString(textToWrite, x, y)
	dc.SetHexColor(colorFill)
	dc.DrawString(textToWrite, x, y)

	drawLegend(dc, source, "Stroke > Draw ", x, y+30)

	x = 400
	dc.SetFont(source.Face(285))
	dc.SetHexColor(colorFill)
	dc.DrawString(textToWrite, x, y)
	dc.SetHexColor(colorStroke)
	dc.StrokeString(textToWrite, x, y)
	drawLegend(dc, source, "Draw > Stroke", x, y+30)

	/*
	* BORDER... PIXEL PERFECT (soon)
	 */
	dc.SetLineWidth(15.0)
	x = 20
	y = 550
	dc.SetFont(source.Face(300))
	dc.SetHexColor("#FF8F00")
	dc.SetLineJoin(gg.LineJoinMiter)
	dc.StrokeString("R", x, y)

	dc.SetFont(source.Face(300))
	dc.SetHexColor("#FF8F00")
	dc.SetLineJoin(gg.LineJoinRound)
	dc.StrokeString("R", x+200, y)

	dc.SetFont(source.Face(300))
	dc.SetHexColor("#FF8F00")
	dc.SetLineJoin(gg.LineJoinBevel)
	dc.StrokeString("R", x+400, y)

	drawLegend(dc, source, "LineJoinMiter", x, y+30)
	drawLegend(dc, source, "LineJoinRound", x+200, y+30)
	drawLegend(dc, source, "LineJoinBevel", x+400, y+30)

	drawLegend(dc, source, "TextModeAliased", x+600, y+30)

	dc.SetTextMode(gg.TextModeAliased)
	dc.SetFont(source.Face(300))
	dc.SetLineJoin(gg.LineJoinRound)
	dc.SetHexColor("#607D8B")
	dc.DrawString("R", x+600, y)

	tools.Save(dc, "font2", "font2")
}

func drawLegend(dc *gg.Context, source *text.FontSource, tex string, x, y float64) {
	dc.SetLineJoin(gg.LineJoinMiter)
	dc.SetFont(source.Face(17))
	dc.SetHexColor("#607D8B")
	dc.DrawString(tex, x, y)

}
