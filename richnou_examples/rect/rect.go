package main

import (
	"fmt"
	"gogpugg/internal/tools"

	"github.com/gogpu/gg"
)

var w = 600
var h = 300
var radius = 50.

var source1 = "/Users/richnou/htdocs/data/sunflower2.png"
var source2 = "/Users/richnou/htdocs/data/sunflower.png"

func main() {
	dc := gg.NewContext(w, h)
	dc.ClearWithColor(gg.White)

	size := 200.

	img, err := gg.LoadImage(source1)

	if err != nil {
		fmt.Println(err)
	}

	dc.DrawImageEx(img, gg.DrawImageOptions{
		X:             50,
		Y:             50,
		DstWidth:      size,
		DstHeight:     size,
		Interpolation: gg.InterpBilinear,
		Opacity:       1.0,
		BlendMode:     gg.BlendNormal,
	})

	// dc.DrawRoundedRectangle(350, 50, rectW, rectH, 30)
	// dc.SetHexColor("#FF6D00")
	// dc.Fill()

	tools.Save(dc, "rect", "rect")
}
