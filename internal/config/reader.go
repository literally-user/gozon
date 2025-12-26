package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type API struct {
	Host      string `toml:"host"`
	Port      int    `toml:"port"`
	SecretKey string `toml:"secret_key"`
}

type Config struct {
	API API `toml:"API"`
}

type Reader struct {
	ConfigPath string
}

func NewReader(configPath string) Reader {
	return Reader{
		ConfigPath: configPath,
	}
}

func (r *Reader) Read() Config {
	var config Config
	if _, err := toml.DecodeFile(r.ConfigPath, &config); err != nil {
		log.Fatalf("Error decoding TOML file: %v", err)
	}

	return config
}
