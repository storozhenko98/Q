package open_ai

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	app_data "github.com/storozhenko98/Q/pkg/data"
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

func GetCompletion(prompt string) error {
	apiKey, err := getApiKey()
	if err != nil {
		return fmt.Errorf("failed to get api key: %w", err)
	}

	url := "https://api.openai.com/v1/chat/completions"
	contentType := "application/json"
	authHeader := fmt.Sprintf("Bearer %s", apiKey)
	body := map[string]interface{}{
		"model":    "gpt-4",
		"messages": []map[string]string{{"role": "system", "content": "You are a helpful assistant who lived in a CLI."}, {"role": "user", "content": prompt}},
		"stream":   true,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Authorization", authHeader)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	reader := bufio.NewReader(resp.Body)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("error reading stream: %w", err)
		}

		if string(line) == "data: [DONE]" {
			break
		}

		if strings.HasPrefix(string(line), "data: ") {
			data := strings.TrimPrefix(string(line), "data: ")
			var streamResp struct {
				Choices []struct {
					Delta struct {
						Content string `json:"content"`
					} `json:"delta"`
				} `json:"choices"`
			}

			if err := json.Unmarshal([]byte(data), &streamResp); err != nil {
				return fmt.Errorf("error unmarshaling stream data: %w", err)
			}

			if len(streamResp.Choices) > 0 && streamResp.Choices[0].Delta.Content != "" {
				fmt.Print(streamResp.Choices[0].Delta.Content)
			}
		}
	}

	return nil
}
