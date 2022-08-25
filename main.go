package main

import (
	"log"
	"wp-bundler/config"
)

func main() {
	config.Init()

	// zipper.Init()
	// zipper.ZipWriter.Write(".distignore", "dist/.distignore")
	// zipper.ZipWriter.Write("go.sum", "dist/go.sum")
	// zipper.ZipWriter.CloseAll()

	log.Println("Application is running at", config.Config.APIAddr)
	log.Println("Application is using ignore file:", config.Config.IgnoreFile)
}
