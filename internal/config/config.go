package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func CreateConfig() {
	configDir := filepath.Join(os.Getenv("HOME"), ".config", "devy")
	configFile := filepath.Join(configDir, "config.json")

	err := os.MkdirAll(configDir, 0755)
	if err != nil {
		panic(err)
	}

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		defaultConfig := GlobalConfig{
			ConfigVersion: "1.0.0",
		}

		configData, err := json.MarshalIndent(defaultConfig, "", "  ")
		if err != nil {
			panic(err)
		}

		err = os.WriteFile(configFile, configData, 0644)
		if err != nil {
			panic(err)
		}
	}
}

type GlobalConfig struct {
	ConfigVersion string `json:"configVersion"`
}

func ReadGlobalConfig() GlobalConfig {
	configDir := filepath.Join(os.Getenv("HOME"), ".config", "devy")
	var globalConfig GlobalConfig

	config, err := os.ReadFile(configDir + "config.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(config, &globalConfig)

	return globalConfig
}

type EnvironmentConfig struct {
	Name       string            `json:"name"`
	ProjectDir string            `json:"projectDir"`
	Commands   map[string]string `json:"commands"`
}
