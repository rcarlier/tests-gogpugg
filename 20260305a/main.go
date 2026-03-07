package main

import (
	"fmt"
	"time"

	"gogpugg/internal/clean"
	"gogpugg/internal/grid"
	"gogpugg/internal/html"

	"github.com/gogpu/gg"
	"github.com/gogpu/gg/text"
)

var OutputDir = "./20260305a/"

var ColorText = "#A1887F"
var ColorRTS = "#37474F"
var ColorAnchor = "#FFC107"

var w = 1000.
var h = 1000.
var MaxWrap = 320

var lorem = ` Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla in orci eu risus mollis ultrices non sed libero. Aenean sit amet feugiat mi. Quisque semper scelerisque dapibus. Suspendisse interdum congue nibh, eget pretium sem mollis vel. Aenean at gravida eros. Nam purus dolor, vehicula et congue ut, rhoncus at est. Proin at commodo eros, sed convallis quam. Fusce quis quam massa. Aliquam ultricies eget nunc eget eleifend. Integer neque libero, tempus in cursus ullamcorper, viverra quis nisi. Suspendisse potenti. `

func DrawWrapText(
	dc *gg.Context, content string, face text.Face,
	maxWidth float64,
	wrap text.WrapMode,
	x, y float64,
	ax, ay float64, // 0=left, 0.5=center, 1=right
	lgFactor float64, // linegap factor
	color string,
) {
	dc.SetHexColor(color)
	dc.SetFont(face)
	// lines := text.WrapText(
	// 	content,
	// 	face,
	// 	maxWidth,
	// 	wrap,
	// )
	// metrics := face.Metrics()
	// lineHeight := metrics.Ascent + metrics.Descent + metrics.LineGap/lgFactor
	// for _, line := range lines {
	// 	dc.DrawStringAnchored(line.Text, x, y, ax, ay)
	// 	y += lineHeight
	// }

	dc.DrawStringWrapped(content, x, y, ax, ay, maxWidth, lgFactor, gg.AlignCenter)
}

func drawAndSave(r, tx, ty, sx, sy float64) string {
	filename := "result"
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

	face := source.Face(10)
	dc.SetFont(face)

	DrawWrapText(
		dc, lorem, face,
		float64(MaxWrap), // maxwidth
		text.WrapWord,
		w/2, h/2, // x,y
		0.5, 0.5, // ax,ay
		1,
		ColorText,
	)

	if r != 0 {
		dc.Rotate(0.2)
		filename += fmt.Sprintf("-r%0.1f", r)
	}

	if tx != 0 && ty != 0 {
		dc.Translate(tx, ty)
		filename += fmt.Sprintf("_tx%0.1f", tx)
		filename += fmt.Sprintf("_ty%0.1f", ty)
	} else if tx != 0 {
		dc.Translate(tx, 0)
		filename += fmt.Sprintf("_tx%0.1f", tx)
	} else if ty != 0 {
		dc.Translate(0, ty)
		filename += fmt.Sprintf("_ty%0.1f", ty)
	}

	if sx != 0 && sy != 0 {
		dc.Scale(sx, sy)
		filename += fmt.Sprintf("_sx%0.1f", sx)
		filename += fmt.Sprintf("_sy%0.1f", sy)
	} else if sx != 0 {
		dc.Scale(sx, 0)
		filename += fmt.Sprintf("_sx%0.1f", sx)
	} else if sy != 0 {
		dc.Scale(0, sy)
		filename += fmt.Sprintf("_sy%0.1f", sy)
	}

	DrawWrapText(
		dc, lorem, face,
		float64(MaxWrap), // maxwidth
		text.WrapWord,
		w/2, h/2, // x,y
		0.5, 0.5, // ax,ay
		1,
		ColorRTS,
	)

	filename += ".png"
	dc.SavePNG(OutputDir + filename)
	fmt.Println(OutputDir + filename)

	return filename
}

func main() {
	start := time.Now()

	clean.RemovePNGFiles(OutputDir)

	tests := []map[string]float64{
		{"r": 0, "tx": 0, "ty": 0., "sx": 0, "sy": 0},
		//
		{"r": 0.2, "tx": 0, "ty": 0., "sx": 0, "sy": 0},
		{"r": 0, "tx": 100., "ty": 100., "sx": 0, "sy": 0},
		{"r": 0, "tx": 0, "ty": 0., "sx": 1.9, "sy": 1.9},
		//
		{"r": 0.2, "tx": -100., "ty": -100., "sx": 0.9, "sy": 0.9},
		{"r": 0.2, "tx": 150., "ty": -140., "sx": 1.2, "sy": 0.5},
	}

	// for _, t := range tests {
	// 	drawAndSave(OutputDir+"result", t["r"], t["tx"], t["ty"], t["sx"], t["sy"])
	// }
	var generatedFiles []string
	for _, t := range tests {
		filename := drawAndSave(t["r"], t["tx"], t["ty"], t["sx"], t["sy"])
		generatedFiles = append(generatedFiles, filename)
	}

	html.GenerateHTML(generatedFiles, OutputDir+"index.html")

	fmt.Printf("Tests completed in %v\n", time.Since(start))
}
