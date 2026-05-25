package grid

import (
	"fmt"
	"log"

	"github.com/gogpu/gg"
	"github.com/gogpu/gg/text"
)

var fontPath = "Cabin-Regular.ttf"

func DrawQuickGrid(dc *gg.Context, xmax, ymax float64) {
	xmin := 0.
	ymin := 0.
	stepMinor := 10.
	stepMajor := 100.

	DrawGrid(dc, xmin, xmax, ymin, ymax, stepMinor, "#FFE0B2", stepMajor, "#FFB74D")

	source, err := text.NewFontSourceFromFile(fontPath)
	if err != nil {
		log.Fatalf("Failed to load font: %v", err)
	}
	defer func() { _ = source.Close() }()
	face48 := source.Face(10)
	dc.SetFont(face48)

	dc.SetHexColor("#37474F")
	for x := xmin; x <= xmax; x += stepMajor {
		dc.DrawString(fmt.Sprintf("%.0f", x), x, 10)
		dc.Stroke()
	}
	for y := stepMajor; y <= ymax; y += stepMajor {
		dc.DrawString(fmt.Sprintf("%.0f", y), 2, y-1)
		dc.Stroke()
	}

}

func DrawGrid(dc *gg.Context, xmin, xmax, ymin, ymax float64, stepMinor float64, colorMinor string, stepMajor ...any) {
	if stepMinor <= 0 {
		return
	}

	dc.SetHexColor(colorMinor)
	for x := xmin; x <= xmax; x += stepMinor {
		dc.DrawLine(x, ymin, x, ymax)
	}
	for y := ymin; y <= ymax; y += stepMinor {
		dc.DrawLine(xmin, y, xmax, y)
	}
	dc.Stroke()

	if len(stepMajor) >= 2 {
		var majorStep float64
		var majorColor string

		switch v := stepMajor[0].(type) {
		case float64:
			majorStep = v
		case int:
			majorStep = float64(v)
		}

		if s, ok := stepMajor[1].(string); ok {
			majorColor = s
		}

		if majorStep > 0 && majorColor != "" {
			dc.SetHexColor(majorColor)
			dc.SetLineWidth(1)
			for x := xmin; x <= xmax; x += majorStep {
				dc.DrawLine(x, ymin, x, ymax)
			}
			for y := ymin; y <= ymax; y += majorStep {
				dc.DrawLine(xmin, y, xmax, y)
			}
			dc.Stroke()
			dc.SetLineWidth(1)
		}
	}
}
