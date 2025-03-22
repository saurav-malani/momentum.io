package api

import (
	"context"
	"fmt"
	"os"
	"strings"

	openai "github.com/sashabaranov/go-openai"

	// this will load openai api key in env variable
	_ "github.com/saurav-malani/momentumio/utility"
)

type OpenAIClient struct {
	Client *openai.Client
}

func NewOpenAIClient(apiKey string) *OpenAIClient {
	return &OpenAIClient{
		Client: openai.NewClient(apiKey),
	}
}

// Utility function to combine file contents
func combineFileContents(filePaths ...string) (string, error) {
	var fileContentBuilder strings.Builder
	for _, filePath := range filePaths {
		data, err := os.ReadFile(filePath)
		if err != nil {
			return "", fmt.Errorf("failed to read file %s: %v", filePath, err)
		}
		fileContentBuilder.WriteString(string(data) + "\n")
	}
	return fileContentBuilder.String(), nil
}

// Core function to call OpenAI API
func (o *OpenAIClient) generateChatCompletion(systemMessage string, prompt string, maxTokens int) (*string, error) {
	messages := []openai.ChatCompletionMessage{
		{Role: "system", Content: systemMessage},
	}
	if prompt != "" {
		messages = append(messages, openai.ChatCompletionMessage{Role: "user", Content: prompt})
	}

	resp, err := o.Client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: maxTokens,
		Messages:  messages,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to generate response: %v", err)
	}
	return &resp.Choices[0].Message.Content, nil
}

// Wrapper function to call OpenAI API with file support
func (o *OpenAIClient) GenerateChatCompletionWithFiles(prompt string, maxTokens int, filePaths ...string) (*string, error) {
	// Combine file contents
	fileContents, err := combineFileContents(filePaths...)
	if err != nil {
		return nil, err
	}

	// Call the OpenAI API
	return o.generateChatCompletion(prompt, fileContents, maxTokens)
}

// Wrapper function for generating chat completion without files
func (o *OpenAIClient) GenerateChatCompletionWithoutFiles(systemPrompt string, maxTokens int) (*string, error) {
	return o.generateChatCompletion(systemPrompt, "", maxTokens)
}
