package config

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
)

type bundler struct {
	IgnoreFile string `toml:"ignore_file"`
	SourceDir  string `toml:"source_dir"`
}

type TomlConfig struct {
	APIAddr string  `toml:"api_addr"`
	Bundler bundler `toml:"bundler"`
	Verbose bool
}

var (
	Config     TomlConfig
	configPath string
)

func loadConfig() {

	if _, err := toml.DecodeFile(configPath, &Config); err != nil {
		log.Fatalln("Reading config failed", err)
	}
}

func Init() {
	// Path to config file can be passed in.
	flag.StringVar(&configPath, "config", "config.toml", "Path to config file")
	Config.Verbose = *flag.Bool("vv", false, "Get Verbose")

	flag.Parse()

	loadConfig()
}
