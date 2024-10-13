package open_ai

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/storozhenko98/Q/pkg/data/app_data"
)

func getApiKey() (string, error) {
	configPath, err := app_data.GetConfigPath()
	if err != nil {
		return "", fmt.Errorf("failed to get config path: %w", err)
	}
	// read the file
	file, err := os.ReadFile(configPath)
	if err != nil {
		return "", fmt.Errorf("failed to read config file: %w", err)
	}
	// parse the json
	var config struct {
		ApiKey string `json:"apiKey"`
	}
	if err := json.Unmarshal(file, &config); err != nil {
		return "", fmt.Errorf("failed to parse config file: %w", err)
	}
	return config.ApiKey, nil
}
