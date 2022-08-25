package main

import (
	"log"
	"wp-bundler/config"
	"wp-bundler/zipper"
)

func main() {
	config.Init()

	log.Println("Application is using source dir:", config.Config.Bundler.SourceDir)
	log.Println("Application is using ignore file:", config.Config.Bundler.IgnoreFile)

	zipper.Init(config.Config)
	zipper.ZipWriter.Write(config.Config.Bundler.SourceDir)
	zipper.ZipWriter.Close()
}
