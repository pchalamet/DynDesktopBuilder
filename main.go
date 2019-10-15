package main

import "flag"
import "os"
import "strings"
import "path/filepath"
import "github.com/pchalamet/DynDesktopBuilder/core"


func main() {
	// parse command line
	steps := flag.Int("steps", 10, "number of images to generate")
	flag.Parse()

	if *steps <= 0 || *steps > 24 {
		panic("Steps must be between 1 and 24")
	}

	filename := flag.Arg(0)
	fname := filepath.Base(filename)
	ext := filepath.Ext(fname)
	basename := strings.TrimSuffix(fname, ext)

	// generate output folder first
	err := os.Mkdir(basename, os.ModeDir | os.ModePerm)
	if err != nil {
		panic("failed to create output folder")
	}

	core.ProcessImage(filename, basename, *steps)
	core.WriteTheme(basename, *steps)
	core.GenZip(basename, *steps)

	err = os.RemoveAll(basename)
	if err != nil {
		panic("failed to purge temporary folder")
	}
}
