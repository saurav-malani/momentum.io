package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadPrompt(t *testing.T) {
	testPromptFile := "testdata/testPrompt.yaml"
	// Test case for a valid prompt key
	t.Run("Valid prompt key", func(t *testing.T) {
		prompt, err := loadPrompt("mock_transcript", testPromptFile)
		assert.NoError(t, err, "Expected no error for a valid prompt key")
		assert.NotEmpty(t, prompt, "Prompt should not be empty for a valid key")
		assert.Equal(t, "Generate a realistic sales transcript.", prompt)
	})

	// Test case for an invalid prompt key
	t.Run("Invalid prompt key", func(t *testing.T) {
		_, err := loadPrompt("non_existent_key", testPromptFile)
		assert.Error(t, err, "Expected an error for a non-existent prompt key")
		assert.Contains(t, err.Error(), "prompt key 'non_existent_key' not found", "Error message should mention the missing key")
	})
}
