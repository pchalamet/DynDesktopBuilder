package core

import "archive/zip"
import "io/ioutil"
import "path/filepath"
import "fmt"
import "os"


func addFiles(w *zip.Writer, basePath string) {
    files, err := ioutil.ReadDir(basePath)
    if err != nil {
		panic("failed to read dir")
    }

    for _, file := range files {
        if !file.IsDir() {
            dat, err := ioutil.ReadFile(filepath.Join(basePath, file.Name()))
            if err != nil {
                fmt.Println(err)
            }
            // Add some files to the archive.
            f, err := w.Create(file.Name())
            if err != nil {
                fmt.Println(err)
            }
            _, err = f.Write(dat)
            if err != nil {
                fmt.Println(err)
            }
        }
    }
}


func GenZip(basename string, steps int) {
    // zip everything
	zipFile := fmt.Sprintf("%s.zip", basename)
	outFile, err := os.Create(zipFile)
    if err != nil {
		panic("failed to create zip file")
    }
    defer outFile.Close()

    // Create a new zip archive.
    zipArch := zip.NewWriter(outFile)
	defer zipArch.Close()

	addFiles(zipArch, basename)
}