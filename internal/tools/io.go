package tools

import (
	"log"

	"github.com/gogpu/gg"
)

func Save(dc *gg.Context, folder, filename string) {
	if err := dc.SavePNG("richnou_examples/" + folder + "/" + filename + ".png"); err != nil {
		log.Fatalf("Failed to save PNG: %v", err)
	}
	log.Println("DONE")
}
