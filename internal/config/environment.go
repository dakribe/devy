package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type EnvironmentConfig struct {
	Name       string            `json:"name"`
	ProjectDir string            `json:"projectDir"`
	Commands   map[string]string `json:"commands"`
}

func CreateEnvironment(envConfig *EnvironmentConfig) error {
	environmentsDir := GetEnvironmentsDir()
	envFile := filepath.Join(environmentsDir, envConfig.Name+".json")

	if _, err := os.Stat(envFile); err == nil {
		return fmt.Errorf("environment '%s' already exists", envConfig.Name)
	}

	envData, err := json.MarshalIndent(envConfig, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal environment config: %w", err)
	}

	err = os.WriteFile(envFile, envData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write environment config: %w", err)
	}

	return nil
}

func ReadEnvironment(name string) (*EnvironmentConfig, error) {
	envFile := filepath.Join(GetEnvironmentsDir(), name+".json")
	var envConfig EnvironmentConfig

	data, err := os.ReadFile(envFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read environment '%s': %w", name, err)
	}

	err = json.Unmarshal(data, &envConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to parse environment '%s': %w", name, err)
	}

	return &envConfig, nil
}

func ListEnvironments() ([]string, error) {
	environmentsDir := GetEnvironmentsDir()

	entries, err := os.ReadDir(environmentsDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read environments directory: %w", err)
	}

	var environments []string
	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".json" {
			name := entry.Name()[:len(entry.Name())-5] // Remove .json extension
			environments = append(environments, name)
		}
	}

	return environments, nil
}

func DeleteEnvironment(name string) error {
	envFile := filepath.Join(GetEnvironmentsDir(), name+".json")

	if _, err := os.Stat(envFile); os.IsNotExist(err) {
		return fmt.Errorf("environment '%s' does not exist", name)
	}

	err := os.Remove(envFile)
	if err != nil {
		return fmt.Errorf("failed to delete environment '%s': %w", name, err)
	}

	return nil
}
