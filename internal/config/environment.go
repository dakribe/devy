package config

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
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

func UpdateEnvironment(envConfig *EnvironmentConfig) error {
	environmentsDir := GetEnvironmentsDir()
	envFile := filepath.Join(environmentsDir, envConfig.Name+".json")

	if _, err := os.Stat(envFile); os.IsNotExist(err) {
		return fmt.Errorf("environment '%s' does not exist", envConfig.Name)
	}

	envData, err := json.MarshalIndent(envConfig, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal environment config: %w", err)
	}

	err = os.WriteFile(envFile, envData, 0644)
	if err != nil {
		return fmt.Errorf("failed to update environment config: %w", err)
	}

	return nil
}

func AddEnvVariable(envName, key, value string) error {
	env, err := ReadEnvironment(envName)
	if err != nil {
		return err
	}

	envFilePath := filepath.Join(env.ProjectDir, ".env")

	envVars := make(map[string]string)

	if _, err := os.Stat(envFilePath); err == nil {
		file, err := os.Open(envFilePath)
		if err != nil {
			return fmt.Errorf("failed to open .env file: %w", err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if line == "" || strings.HasPrefix(line, "#") {
				continue
			}
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				envVars[parts[0]] = parts[1]
			}
		}
	}

	envVars[key] = value

	file, err := os.Create(envFilePath)
	if err != nil {
		return fmt.Errorf("failed to create .env file: %w", err)
	}
	defer file.Close()

	for k, v := range envVars {
		_, err := file.WriteString(fmt.Sprintf("%s=%s\n", k, v))
		if err != nil {
			return fmt.Errorf("failed to write to .env file: %w", err)
		}
	}

	return nil
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
