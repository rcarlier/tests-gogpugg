package main

import (
	"fmt"
	"time"

	"github.com/gogpu/gg"
	"github.com/gogpu/gg/text"
)

var OutputFile = "20260304/result.png"
var ColorText = "#37474F"
var MaxWrap = 320
var lorem = ` Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla in orci eu risus mollis ultrices non sed libero. Aenean sit amet feugiat mi. Quisque semper scelerisque dapibus. Suspendisse interdum congue nibh, eget pretium sem mollis vel. Aenean at gravida eros. Nam purus dolor, vehicula et congue ut, rhoncus at est. Proin at commodo eros, sed convallis quam. Fusce quis quam massa. Aliquam ultricies eget nunc eget eleifend. Integer neque libero, tempus in cursus ullamcorper, viverra quis nisi. Suspendisse potenti. `

func DrawWrapText(
	dc *gg.Context, content string, face text.Face,
	maxWidth float64,
	wrap text.WrapMode,
	x, y float64,
	ax, ay float64, // 0=left, 0.5=center, 1=right
	lgFactor float64, // linegap factor
) {
	dc.SetFont(face)
	lines := text.WrapText(
		content,
		face,
		maxWidth,
		wrap,
	)
	metrics := face.Metrics()
	lineHeight := metrics.Ascent + metrics.Descent + metrics.LineGap/lgFactor
	for _, line := range lines {
		dc.DrawStringAnchored(line.Text, x, y, ax, ay)
		y += lineHeight
	}
}
func main() {
	w := 5000.
	h := 2000.

	start := time.Now()
	dc := gg.NewContext(int(w), int(h))
	defer dc.Close()
	dc.ClearWithColor(gg.White)
	dc.SetHexColor("#000")

	source, _ := text.NewFontSourceFromFile("Arial.ttf")
	defer source.Close()
	face := source.Face(20)
	dc.SetFont(face)

	dc.Rotate(0.2)
	// dc.Translate(1500, -1400)
	dc.Translate(150, -140)
	dc.Scale(2.8, 0.5)

	face = source.Face(10)
	dc.SetFont(face)
	dc.SetHexColor(ColorText)
	DrawWrapText(
		dc, lorem, face,
		float64(MaxWrap), // maxwidth
		text.WrapWord,
		w/2, h/2, // x,y
		0.5, 0.0, // ax,ay
		1,
	)
	dc.SavePNG(OutputFile)
	fmt.Printf("Tests completed in %v\n", time.Since(start))
}
