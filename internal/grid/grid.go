package grid

import "github.com/gogpu/gg"

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
