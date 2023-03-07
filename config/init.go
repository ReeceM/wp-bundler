package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type bundler struct {
	IgnoreFile string `toml:"ignore_file"` // the .distignore file that should be used
	SourceDir  string `toml:"source_dir"`  // the input directory that has the plugin
}

type TomlConfig struct {
	Name    string  `toml:"name"`    // this is the plugin / theme name
	Bundler bundler `toml:"bundler"` // the bundler options
	Verbose bool
}

var (
	Config     TomlConfig
	configPath string
	source_dir string
)

type ConfigOptions struct {
	ConfigPath string
	SourceDir  string
	Verbose    bool
}

func loadConfig() {

	if _, err := toml.DecodeFile(configPath, &Config); err != nil {
		if _, err := toml.DecodeFile(source_dir+"/config.toml", &Config); err != nil {
			log.Fatalln("Reading config failed", err)
		}
	}
}

func Init(options ConfigOptions) {
	// Path to config file can be passed in.

	configPath = options.ConfigPath
	Config.Verbose = options.Verbose

	source_dir = options.SourceDir

	loadConfig()
}
