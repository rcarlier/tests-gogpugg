package main

import (
	"gogpugg/internal/grid"
	"log"

	"github.com/gogpu/gg"
)

var w = 512.
var h = 1024.

func main() {
	dc := gg.NewContext(int(w), int(h))
	dc.ClearWithColor(gg.Hex("#FFF8E1"))
	grid.DrawGrid(dc, 0, w, 0, h, 10, "#FFE0B2", 100, "#FFB74D")

	x := 10.
	y := 10.
	w := 492.
	stepy := 50.

	rect(dc, x, y, w, stepy-10)
	y += stepy

	dc.SetDash(10, 5)
	rect(dc, x, y, w, stepy-10)
	y += stepy

	dc.SetDash(10, 5, 5, 10)
	rect(dc, x, y, w, stepy-10)
	y += stepy

	for i := range 10 {
		dc.SetDashOffset(float64(i)) // no effect = must be after SetDash?
		dc.SetDash(10, 5)
		line(dc, y, "#2196F3")
		y += 5
	}

	for i := range 10 {
		dc.SetDash(10, 5)
		dc.SetDashOffset(float64(i))
		line(dc, y, "#FF6F00")
		y += 5
	}

	y += 20
	dc.SetStroke(gg.DefaultStroke().WithWidth(10).WithCap(gg.LineCapRound))
	line(dc, y, "#FF6F00")
	y += 20
	dc.SetStroke(gg.DefaultStroke().WithWidth(10).WithCap(gg.LineCapButt))
	line(dc, y, "#FF6F00")
	y += 20
	dc.SetStroke(gg.DefaultStroke().WithWidth(10).WithCap(gg.LineCapSquare))
	line(dc, y, "#FF6F00")

	y += 20
	dc.SetLineJoin(gg.LineJoinBevel)
	lineMulti(dc, y, "#2196F3")
	y += 20
	dc.SetLineJoin(gg.LineJoinMiter)
	lineMulti(dc, y, "#2196F3")
	y += 20
	dc.SetLineJoin(gg.LineJoinRound)
	lineMulti(dc, y, "#2196F3")

	y += 40

	gradientLinear := gg.NewLinearGradientBrush(x, 0, 512, 0).
		AddColorStop(0, gg.Hex("#FF6F00")).
		AddColorStop(0.5, gg.Hex("#0D47A1")).
		AddColorStop(1, gg.Hex("#4CAF50"))
	dc.SetFillBrush(gradientLinear)
	dc.DrawRectangle(x, y, w, stepy-10)
	dc.Fill()

	y += stepy
	gradientExtendRepeat := gg.NewLinearGradientBrush(x, 0, 100, 0).
		SetExtend(gg.ExtendRepeat).
		AddColorStop(0, gg.Hex("#FF6F00")).
		AddColorStop(0.5, gg.Hex("#0D47A1")).
		AddColorStop(1, gg.Hex("#4CAF50"))
	dc.SetFillBrush(gradientExtendRepeat)
	dc.DrawRectangle(x, y, w, stepy-10)
	dc.Fill()

	y += stepy
	gradientExtendReflect := gg.NewLinearGradientBrush(x, 0, 100, 0).
		SetExtend(gg.ExtendReflect).
		AddColorStop(0, gg.Hex("#FF6F00")).
		AddColorStop(0.5, gg.Hex("#0D47A1")).
		AddColorStop(1, gg.Hex("#4CAF50"))
	dc.SetFillBrush(gradientExtendReflect)
	dc.DrawRectangle(x, y, w, stepy-10)
	dc.Fill()

	y += stepy
	gradientExtendPad := gg.NewLinearGradientBrush(x, 0, 512, 0).
		SetExtend(gg.ExtendPad).
		AddColorStop(0, gg.Hex("#FF6F00")).
		AddColorStop(0.2, gg.Hex("#0D47A1")).
		AddColorStop(0.8, gg.Hex("#0D47A1")).
		AddColorStop(1, gg.Hex("#FF6F00"))
	dc.SetFillBrush(gradientExtendPad)
	dc.DrawRectangle(x, y, w, stepy-10)
	dc.Fill()

	y += stepy
	gradientVert := gg.NewLinearGradientBrush(0, y, 0, y+stepy).
		SetExtend(gg.ExtendPad).
		AddColorStop(0, gg.Hex("#FFFDE7")).
		AddColorStop(0.2, gg.Hex("#F57F17")).
		AddColorStop(1, gg.Hex("#00796B"))
	dc.SetFillBrush(gradientVert)
	dc.DrawRectangle(x, y, w, stepy)
	dc.Fill()

	save(dc, "dashes")
}

func lineMulti(dc *gg.Context, y float64, color string) {
	dc.SetHexColor(color)

	dc.MoveTo(10., y)
	dc.LineTo(100., y)
	dc.LineTo(130., y+20)
	dc.LineTo(180., y+10)
	dc.LineTo(250., y)
	dc.LineTo(450., y)

	dc.Stroke()
}
func line(dc *gg.Context, y float64, color string) {
	dc.SetHexColor(color)
	dc.DrawLine(10., y, 492, y)
	dc.Stroke()
}

func rect(dc *gg.Context, x, y, w, h float64) {
	dc.SetRGB(0, 0, 1)
	dc.DrawRectangle(x, y, w, h)
	dc.Stroke()
}
func save(dc *gg.Context, filename string) {
	if err := dc.SavePNG("richnou_examples/" + filename + "/" + filename + ".png"); err != nil {
		log.Fatalf("Failed to save PNG: %v", err)
	}
	log.Println("DONE")
}
