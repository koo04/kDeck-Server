package settings

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

type Base struct {
	Enabled bool `json:"enabled"`
}

var settings map[string]interface{} = make(map[string]interface{})

func Load() error {

	log.Println("Trying to get the working directory")
	workDir, err := os.Getwd()
	if err != nil {
		// log.Fatalf("Could not get work directory: %s", err)
		return err
	}

	log.Println("Checking to see if the settings file exists")
	settingsFile, err := os.Open(fmt.Sprintf("%s/settings.json", workDir))
	if err != nil {
		//TODO: handle this error nicely.
		SaveSettings()
	} else {
		log.Println("Found an existing settings file. Loading...")
		jsonParser := json.NewDecoder(settingsFile)
		if err = jsonParser.Decode(&settings); err != nil {
			log.Fatalf("Could not parse settings: %s", err)
		}

		j, err := json.MarshalIndent(&settings, "", "  ")
		if err != nil {
			log.Fatalf("Could not marshal default Settings: %s", err)
		}
		log.Printf("... Loaded: %s", j)
	}
	return nil
}

func GetSettings(plugin string) (map[string]interface{}, error) {
	if cfg, ok := settings[plugin]; ok {
		return cfg.(map[string]interface{}), nil
	}
	return nil, SettingsError_NO_SETTINGS_FOUND
}

func SetSettings(plugin string, pluginSettings interface{}) {
	settings[plugin] = pluginSettings
	SaveSettings()
}

func SaveSettings() {
	workDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Could not get work directory: %s", err)
	}

	j, err := json.MarshalIndent(settings, "", "  ")
	if err != nil {
		log.Fatalf("Could not marshal settings: %s", err)
	}

	defer os.WriteFile(fmt.Sprintf("%s/settings.json", workDir), j, 0777)
	if err != nil {
		log.Fatalf("Could not save settings file")
	}
}

func LoadSettings(plugin string, settings interface{}) {
	cfg, err := GetSettings(plugin)
	if err == nil {
		jcfg, err := json.Marshal(cfg)
		if err != nil {
			log.Fatalln(err)
		}

		err = json.Unmarshal(jcfg, &settings)
		if err != nil {
			log.Fatalln(err)
		}
	}

	SetSettings(plugin, settings)
}

type SettingsError error

var (
	SettingsError_NO_SETTINGS_FOUND = errors.New("Could not find the settings")
)
