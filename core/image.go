package core

import "fmt"
import "image"
import _ "image/gif"
import "image/png"
import _ "image/jpeg"
import "path/filepath"
import "math"
import "os"
import "github.com/disintegration/imaging"


func writeImage(basename string, step int, img image.Image) {
	fname := fmt.Sprintf("image_%d.png", step)
	outFile := filepath.Join(basename, fname)

	fmt.Println("Writing to file ", outFile)

	outputFile, err := os.Create(outFile)
	if err != nil {
		panic("failed to write output file")
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, img)
	if err != nil {
		panic("failed to write output file")
	}
}

func ProcessImage(filename string, basename string, steps int) {
	inFile, err := os.Open(filename)
	if err != nil {
		panic("missing image file")
	}
	defer inFile.Close()

	imageData, _, err := image.Decode(inFile)
	if err != nil {
		panic("failed to decode provided image")
	}

	// generate all images
	writeImage(basename, 1, imageData)
	for step:=2; step<=steps; step++ {
		dark := -60.0 * math.Sin(float64(step) / float64(steps))
		darkImg := imaging.AdjustBrightness(imageData, dark)
		writeImage(basename, step, darkImg)
	}
}