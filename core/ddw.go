package core

import "archive/zip"
import "io/ioutil"
import "path/filepath"
import "fmt"
import "os"


func addFiles(w *zip.Writer, basePath string) {
    files, err := ioutil.ReadDir(basePath)
    CheckError(err, "failed to read dir")

    for _, file := range files {
        if !file.IsDir() {
            dat, err := ioutil.ReadFile(filepath.Join(basePath, file.Name()))
            CheckError(err, "failed to read file")

            f, err := w.Create(file.Name())
            CheckError(err, "failed to create file")

            _, err = f.Write(dat)
            CheckError(err, "failed to write file")
        }
    }
}


func GenTheme(workDir string, steps int) string {
    // zip everything
    ddwFile := fmt.Sprintf("%s.ddw", workDir)

    outFile, err := os.Create(ddwFile)
    CheckError(err, "failed to create theme file")
    defer outFile.Close()

    // Create a new zip archive.
    zipArch := zip.NewWriter(outFile)
	defer zipArch.Close()

    addFiles(zipArch, workDir)
    
    return ddwFile
}