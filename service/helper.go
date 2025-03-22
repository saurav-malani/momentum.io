package service

import (
	"bufio"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// PromptsConfig holds the structure for the prompts in the YAML file.
type PromptsConfig struct {
	Prompts map[string]string `yaml:"prompts"`
}

// loadPrompt loads a specific prompt key from the `prompt.yaml` file.
func loadPrompt(key, filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("failed to read prompt.yaml: %v", err)
	}

	var config PromptsConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return "", fmt.Errorf("failed to parse prompt.yaml: %v", err)
	}

	prompt, exists := config.Prompts[key]
	if !exists {
		return "", fmt.Errorf("prompt key '%s' not found in prompt.yaml", key)
	}

	return prompt, nil
}

// writeToFile creates a file with the given name and writes the content from the string pointer to it.
func writeToFile(content *string, fileName string) error {
	// Create or truncate the file
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close() // Ensure the file is closed when the function ends
	// Create a buffered writer
	writer := bufio.NewWriter(file)

	// Write the content to the file
	_, err = writer.WriteString(*content)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	// Flush the buffered writer to ensure all data is written
	err = writer.Flush()
	if err != nil {
		return fmt.Errorf("failed to flush writer: %w", err)
	}

	return nil
}

// generateRandomFileName creates a random file name using a secure random string
func generateRandomFileName() (string, error) {
	// Create a random byte slice
	randomBytes := make([]byte, 8) // 8 bytes = 16 hex characters
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", fmt.Errorf("failed to generate random bytes: %w", err)
	}

	// Convert bytes to a hexadecimal string
	randomString := hex.EncodeToString(randomBytes)
	return fmt.Sprintf("%s%s.%s", "generatedTranscript-", randomString, "txt"), nil
}
