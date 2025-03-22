package api

import (
	"fmt"
	"os"
	"testing"

	"github.com/saurav-malani/momentumio/utility"
	"github.com/stretchr/testify/assert"
)

// TestCombineFileContents tests the CombineFileContents function.
func TestCombineFileContents(t *testing.T) {
	utility.LoadEnv()
	// Create sample test files
	file1Path := "../testdata/file1.txt"
	file2Path := "../testdata/file2.txt"
	err := os.WriteFile(file1Path, []byte("This is content of file 1."), 0644)
	assert.NoError(t, err, "Failed to create file1.txt")
	err = os.WriteFile(file2Path, []byte("This is content of file 2."), 0644)
	assert.NoError(t, err, "Failed to create file2.txt")

	// Cleanup after the test
	defer func() {
		os.Remove(file1Path)
		os.Remove(file2Path)
	}()

	// Combine the file contents
	filePaths := []string{file1Path, file2Path}
	combinedContent, err := combineFileContents(filePaths...)

	// Assert no error and correct combined content
	assert.NoError(t, err, "Error combining file contents")
	expectedContent := "This is content of file 1.\nThis is content of file 2.\n"
	assert.Equal(t, expectedContent, combinedContent, "Combined content does not match expected")
}

// TestGenerateChatCompletionWithoutFiles tests GenerateChatCompletion without files.
func TestGenerateChatCompletionWithoutFiles(t *testing.T) {
	utility.LoadEnv()
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		t.Fatalf("Error: OPENAI_API_KEY not found in environment")
	}

	client := NewOpenAIClient(apiKey)
	prompt := "Generate a realistic sales transcript in the format:\n\n" +
		"00:00:00 Sam (openai.com): Hello!\n" +
		"00:00:02 Satya (microsoft.com): Hi, how can I help you?\n" +
		"NOTE: we would later need to summarize it to get insights from the call that could be helpful in understanding the client, their requirements, and places where they are confused. So, generate keeping this in mind."
	maxTokens := 100

	result, err := client.GenerateChatCompletionWithoutFiles(prompt, maxTokens)
	assert.NoError(t, err, "Error generating chat completion without files")
	assert.NotEmpty(t, result, "The response should not be empty")
	fmt.Println("Generated Response: ", *result)
}

// TestGenerateChatCompletionWithFiles tests GenerateChatCompletionWithFiles.
func TestGenerateChatCompletionWithFiles(t *testing.T) {
	utility.LoadEnv()
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		t.Fatalf("Error: OPENAI_API_KEY not found in environment")
	}

	client := NewOpenAIClient(apiKey)

	// Create sample test files
	file1Path := "../testdata/file1.txt"
	file2Path := "../testdata/file2.txt"
	err := os.WriteFile(file1Path, []byte("This is content of file 1."), 0644)
	assert.NoError(t, err, "Failed to create file1.txt")
	err = os.WriteFile(file2Path, []byte("This is content of file2."), 0644)
	assert.NoError(t, err, "Failed to create file2.txt")

	// Cleanup after the test
	defer func() {
		os.Remove(file1Path)
		os.Remove(file2Path)
	}()

	prompt := "Generate a summary of the following transcript files."
	filePaths := []string{file1Path, file2Path}
	maxTokens := 100

	result, err := client.GenerateChatCompletionWithFiles(prompt, maxTokens, filePaths...)
	assert.NoError(t, err, "Error generating chat completion with files")
	assert.NotEmpty(t, result, "The response should not be empty")
	fmt.Println("Generated Response With Files: ", *result)
}
