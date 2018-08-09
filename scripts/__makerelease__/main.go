package main

import (
	"archive/zip"
	"compress/flate"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// this package creates the release zip file

func main() {

	fullpath, _ := filepath.Abs("../build/auto_win_10_conf.zip")
	outFile, err := os.Create(fullpath)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	// Create a new zip
	w := zip.NewWriter(outFile)

	// compress the zip
	w.RegisterCompressor(zip.Deflate, func(out io.Writer) (io.WriteCloser, error) {
		return flate.NewWriter(out, flate.BestCompression)
	})

	// files to write
	var files = []struct {
		Name, Body string
	}{
		{"config.json", "../build/config.json"},
		{"setup.exe", "../build/setup.exe"},
	}

	// write the files to the zip
	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		FileContents, err := ioutil.ReadFile(file.Body)
		_, err = f.Write(FileContents)
		if err != nil {
			log.Fatal(err)
		}
	}

	// close the zip
	err = w.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created release zip file")
}
