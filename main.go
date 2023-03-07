package main

import (
	"flag"
	"log"
	"wp-bundler/config"
	"wp-bundler/zipper"
)

var (
	configPath string
	verbose    bool
	source_dir string
)

func main() {
	flag.StringVar(&configPath, "config", "config.toml", "Path to config file")
	flag.BoolVar(&verbose, "vv", false, "Get Verbose")
	flag.StringVar(&source_dir, "source", "dist", "Provide the source directory, will be searched for config.toml")

	flag.Parse()

	config.Init(config.ConfigOptions{
		ConfigPath: configPath,
		Verbose:    verbose,
		SourceDir:  source_dir,
	})

	zipper.Init(config.Config)

	log.Println("Application is using source dir:", config.Config.Bundler.SourceDir)
	log.Println("Application is using ignore file:", config.Config.Bundler.IgnoreFile)

	zipper.ZipWriter.Write(config.Config.Bundler.SourceDir)
	zipper.ZipWriter.Close()
}
