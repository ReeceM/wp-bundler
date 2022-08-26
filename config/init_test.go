package config

import (
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestLoadConfig(t *testing.T) {
	configPath = "config.test.toml"
	loadConfig()
	if Config.Name != "wp-bundler" {
		t.Fatalf("No config values loaded")
	}
}
