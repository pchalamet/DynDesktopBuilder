package main

import "fmt"
import "flag"
import "os"
import "strings"
import "path/filepath"
import "github.com/pchalamet/DynDesktopBuilder/core"


func purgeWorkDir(workDir string) {
	err := os.RemoveAll(workDir)
	if err != nil {
		fmt.Println("WARNING: failed to purge temporary folder ", workDir)
	}
}

func process() {
	// parse command line
	steps := flag.Int("steps", 8, "number of images to generate - must be between 1 and 24")
	flag.Parse()

	if *steps <= 0 || *steps > 24 {
		flag.Usage()
		panic("invalid arguments")
	}

	if len(flag.Args()) != 1 {
		flag.Usage()
		panic("invalid arguments")
	}

	imageFileName := flag.Arg(0)
	fname := filepath.Base(imageFileName)
	ext := filepath.Ext(fname)
	basename := strings.TrimSuffix(fname, ext)

	// create temporary folder first & defer its deletion
	tmpDir := os.TempDir()
	workDir := filepath.Join(tmpDir, basename)
	err := os.Mkdir(workDir, os.ModeDir | os.ModePerm)
	defer purgeWorkDir(workDir)

	if err != nil {
		panic("failed to create output folder")
	}

	core.ProcessImage(workDir, imageFileName, *steps)
	core.WriteTheme(workDir, *steps)
	ddwFileName := core.GenTheme(workDir, *steps)

	fmt.Println("Theme successfully built:", ddwFileName)
}


func main() {
	defer func() {
        if r := recover(); r != nil {
            fmt.Println("ERROR:", r)
        }
    }()

	process()
}