package config

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
)

// type certsConfig struct {
// 	CrtFile string `toml:"crt_file"`
// 	KeyFile string `toml:"key_file"`
// }

type tomlConfig struct {
	APIAddr    string `toml:"api_addr"`
	IgnoreFile string `toml:"ignore_file"`
}

var (
	Config     tomlConfig
	configPath string
)

func loadConfig() {

	if _, err := toml.DecodeFile(configPath, &Config); err != nil {
		log.Fatalln("Reading config failed", err)
	}
}

func Init() {
	// Path to config file can be passed in.
	flag.StringVar(&configPath, "config", "config.dev.toml", "Path to config file")
	flag.Parse()

	loadConfig()
}
