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


func writeImage(workDir string, step int, img image.Image) {
	fname := fmt.Sprintf("image_%d.png", step)
	outFile := filepath.Join(workDir, fname)

	fmt.Println("Writing to file ", outFile)

	outputFile, err := os.Create(outFile)
	CheckError(err, "failed to write output file")
	defer outputFile.Close()

	err = png.Encode(outputFile, img)
	CheckError(err, "failed to write output file")
}

func ProcessImage(workDir string, sourceImage string, steps int) {
	inFile, err := os.Open(sourceImage)
	CheckError(err, "missing image file")
	defer inFile.Close()

	imageData, _, err := image.Decode(inFile)
	CheckError(err, "failed to decode provided image")

	// generate all images
	writeImage(workDir, 1, imageData)
	for step:=2; step<=steps; step++ {
		dark := -50.0 * math.Sin(float64(step) / float64(steps))
		darkImg := imaging.AdjustBrightness(imageData, dark)
		writeImage(workDir, step, darkImg)
	}
}