package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"wp-bundler/config"
	"wp-bundler/zipper"
)

func main() {
	config.Init()

	zipper.Init()
	zipper.ZipWriter.Write("dist")
	zipper.ZipWriter.Writer.Close()
	zipper.ZipWriter.Archive.Close()

	log.Println("Application is running at", config.Config.APIAddr)
	log.Println("Application is using ignore file:", config.Config.IgnoreFile)
}

func test() {
	fmt.Println("creating zip archive...")
	archive, err := os.Create("archive.zip")
	if err != nil {
		panic(err)
	}
	defer archive.Close()
	zipWriter := zip.NewWriter(archive)

	fmt.Println("opening first file...")
	f1, err := os.Open("dist/test.csv")
	if err != nil {
		panic(err)
	}
	defer f1.Close()

	fmt.Println("writing first file to archive...")
	w1, err := zipWriter.Create("csv/test.csv")
	if err != nil {
		panic(err)
	}
	if _, err := io.Copy(w1, f1); err != nil {
		panic(err)
	}

	fmt.Println("opening second file")
	f2, err := os.Open("dist/test.txt")
	if err != nil {
		panic(err)
	}
	defer f2.Close()

	fmt.Println("writing second file to archive...")
	w2, err := zipWriter.Create("txt/test.txt")
	if err != nil {
		panic(err)
	}
	if _, err := io.Copy(w2, f2); err != nil {
		panic(err)
	}
	fmt.Println("closing zip archive...")
	zipWriter.Close()
}
