package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func GetConfigDir() string {
	return filepath.Join(os.Getenv("HOME"), ".config", "devy")
}

func GetEnvironmentsDir() string {
	return filepath.Join(GetConfigDir(), "environments")
}

func CreateConfig() {
	configDir := GetConfigDir()
	environmentsDir := GetEnvironmentsDir()
	configFile := filepath.Join(configDir, "config.json")

	err := os.MkdirAll(configDir, 0755)
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll(environmentsDir, 0755)
	if err != nil {
		panic(err)
	}

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		defaultConfig := GlobalConfig{
			ConfigVersion:      "1.0.0",
			DefaultEnvironment: "",
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
	ConfigVersion      string `json:"configVersion"`
	DefaultEnvironment string `json:"defaultEnvironment,omitempty"`
}

func ReadGlobalConfig() (*GlobalConfig, error) {
	configFile := filepath.Join(GetConfigDir(), "config.json")
	var globalConfig GlobalConfig

	config, err := os.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	err = json.Unmarshal(config, &globalConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return &globalConfig, nil
}

func WriteGlobalConfig(config *GlobalConfig) error {
	configFile := filepath.Join(GetConfigDir(), "config.json")

	configData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	err = os.WriteFile(configFile, configData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}
