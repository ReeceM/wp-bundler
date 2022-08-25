package main

import (
	"log"
	"wp-bundler/config"
	"wp-bundler/zipper"
)

func main() {
	config.Init()

	log.Println("Application is using ignore file:", config.Config.IgnoreFile)
	zipper.Init(config.Config)
	zipper.ZipWriter.Write("dist")
	zipper.ZipWriter.Close()
}
