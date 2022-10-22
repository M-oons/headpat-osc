package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/m-oons/headpat-osc/data"
)

type Config struct {
	Osc     ConfigOsc     `json:"osc"`
	Headpat ConfigHeadpat `json:"headpat"`
}

type ConfigOsc struct {
	Host string `json:"host"`
	Port uint16 `json:"port"`
}

type ConfigHeadpat struct {
	Message string `json:"message"`
}

var Current Config

func init() {
	Load()
}

func Load() {
	config := Config{}

	dir := data.GetDataFolder()

	d, err := os.ReadFile(fmt.Sprintf("%s/config.json", dir))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Printf("Config file doesn't exist, creating default config")
			config = defaultConfig()
		} else {
			log.Fatalf("Error reading config: %v", err)
		}
	} else {
		err = json.Unmarshal(d, &config)
		if err != nil {
			log.Fatalf("Error unmarshalling config: %v", err)
		}
	}

	Save(config)

	Current = config
}

func Save(config Config) {
	d, err := json.MarshalIndent(config, "", "\t")
	if err != nil {
		log.Fatalf("Error marshalling config: %v", err)
	}

	dir := data.GetDataFolder()

	err = os.WriteFile(fmt.Sprintf("%s/config.json", dir), d, 0644)
	if err != nil {
		log.Fatalf("Error writing config to file: %v", err)
	}
}

func defaultConfig() Config {
	return Config{
		Osc: ConfigOsc{
			Host: "127.0.0.1",
			Port: 9001,
		},
		Headpat: ConfigHeadpat{
			Message: "{{count}} headpat{{plural}}",
		},
	}
}
