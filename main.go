package main

import "fmt"
import "flag"
import "os"
import "strings"
import "path/filepath"
import "github.com/pchalamet/DynDesktopBuilder/core"
import "errors"


func purgeWorkDir(workDir string) {
	err := os.RemoveAll(workDir)
	if err != nil {
		fmt.Println("WARNING: failed to purge temporary folder ", workDir)
	}
}


func parseArgs() (steps int, imageFileName string, err error) {
	// parse command line
	vsteps := flag.Int("steps", 8, "number of images to generate - must be between 1 and 24")
	flag.Parse()

	if *vsteps <= 0 || *vsteps > 24 {
		return steps, imageFileName, errors.New("steps has invalid value")
	}

	if len(flag.Args()) != 1 {
		return steps, imageFileName, errors.New("missing image file")
	}

	return *vsteps, flag.Arg(0), nil
}

func process() {
	steps, imageFileName, err := parseArgs()
	core.CheckError(err, "invalid arguments")

	fname := filepath.Base(imageFileName)
	ext := filepath.Ext(fname)
	basename := strings.TrimSuffix(fname, ext)

	// create temporary folder first & defer its deletion
	tmpDir := os.TempDir()
	workDir := filepath.Join(tmpDir, basename)
	err = os.Mkdir(workDir, os.ModeDir | os.ModePerm)
	core.CheckError(err, "failed to create output folder")
	defer purgeWorkDir(workDir)

	core.ProcessImage(workDir, imageFileName, steps)
	core.WriteTheme(workDir, steps)
	ddwFileName := core.GenTheme(workDir, steps)

	fmt.Println("Theme successfully built:", ddwFileName)
}


func main() {
	defer func() {
        if r := recover(); r != nil {
			fmt.Println("ERROR:", r)
			os.Exit(5)
        }
    }()

	process()
}