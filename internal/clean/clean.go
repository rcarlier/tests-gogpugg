package clean

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func RemovePNGFiles(directory string) error {
	png := 0
	entries, err := os.ReadDir(directory)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		if strings.HasSuffix(strings.ToLower(name), ".png") {
			path := filepath.Join(directory, name)
			if err := os.Remove(path); err != nil {
				return err
			}
			png += 1
		}
	}
	fmt.Printf("PNG deleted %d\n", png)
	return nil
}
