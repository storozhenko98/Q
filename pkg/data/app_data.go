package app_data

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	appName    = "YourAppName"
	configFile = "qConfig.json"
)

func SetupAppData() error {
	// Get the Application Support directory
	appSupportDir, err := getAppSupportDir()
	if err != nil {
		return fmt.Errorf("failed to get Application Support directory: %w", err)
	}

	// Create the app's directory if it doesn't exist
	appDir := filepath.Join(appSupportDir, appName)
	if err := os.MkdirAll(appDir, 0755); err != nil {
		return fmt.Errorf("failed to create app directory: %w", err)
	}

	// Check if qConfig.json exists, create it if not
	configPath := filepath.Join(appDir, configFile)
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		file, err := os.Create(configPath)
		if err != nil {
			return fmt.Errorf("failed to create config file: %w", err)
		}
		file.Close()
	}

	return nil
}

func getAppSupportDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, "Library", "Application Support"), nil
}

func GetConfigPath() (string, error) {
	appSupportDir, err := getAppSupportDir()
	if err != nil {
		return "", fmt.Errorf("failed to get Application Support directory: %w", err)
	}

	appDir := filepath.Join(appSupportDir, appName)
	configPath := filepath.Join(appDir, configFile)

	return configPath, nil
}

func CheckIfConfigFileEmpty() (bool, error) {
	configPath, err := GetConfigPath()
	if err != nil {
		return false, fmt.Errorf("failed to get config path: %w", err)
	}

	file, err := os.ReadFile(configPath)
	if err != nil {
		return false, fmt.Errorf("failed to read config file: %w", err)
	}
	return len(file) == 0, nil
}

func WriteApiKeyToConfig(apiKey string) error {
	configPath, err := GetConfigPath()
	if err != nil {
		return fmt.Errorf("failed to get config path: %w", err)
	}

	config := struct {
		ApiKey string `json:"apiKey"`
	}{
		ApiKey: apiKey,
	}

	file, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	if err := os.WriteFile(configPath, file, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}

func UpdateApiKeyInConfig(apiKey string) error {
	configPath, err := GetConfigPath()
	if err != nil {
		return fmt.Errorf("failed to get config path: %w", err)
	}

	// Delete existing config file
	if err := os.Remove(configPath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to delete existing config file: %w", err)
	}

	// Write new API key to config
	if err := WriteApiKeyToConfig(apiKey); err != nil {
		return fmt.Errorf("failed to write new API key to config: %w", err)
	}

	return nil
}
