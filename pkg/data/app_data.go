package app_data

import (
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
