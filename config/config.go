package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// type OBS struct {
// 	Enabled bool   `json:"enabled"`
// 	Host    string `json:"host"`
// 	Port    int    `json:"port"`
// }

type OBS struct {
	Enabled bool   `json:"enabled"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
}

type Config struct {
	Obs OBS `json:"obs"`
}

var config *Config

func LoadConfig() *Config {
	log.Println("Trying to get the working directory")
	workDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Could not get work directory: %s", err)
	}

	log.Println("Checking to see if the config file exists")
	configFile, err := os.Open(fmt.Sprintf("%s/config.json", workDir))
	if err != nil {
		log.Printf("Could not load config: %s", err)
		log.Println("Creating default Config")
		config = &Config{Obs: OBS{Enabled: false, Host: "localhost", Port: 4444}}
		SaveConfig()
	} else {
		log.Println("Found an existing config file. Loading...")
		jsonParser := json.NewDecoder(configFile)
		if err = jsonParser.Decode(&config); err != nil {
			log.Fatalf("Could not parse config: %s", err)
		}

		j, err := json.MarshalIndent(&config, "", "  ")
		if err != nil {
			log.Fatalf("Could not marshal default Config: %s", err)
		}
		log.Printf("... Loaded: %s", j)
	}
	return config
}

func SaveConfig() {
	workDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Could not get work directory: %s", err)
	}

	j, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		log.Fatalf("Could not marshal config: %s", err)
	}

	defer os.WriteFile(fmt.Sprintf("%s/config.json", workDir), j, 0777)
	if err != nil {
		log.Fatalf("Could not save config file")
	}
}

func GetConfig() *Config {
	return config
}
