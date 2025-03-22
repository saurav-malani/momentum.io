package service

import (
	"fmt"
	"log"

	api "github.com/saurav-malani/momentumio/api/openai"
)

const (
	maxTokens                    = 1000
	generateMockTranscriptPrompt = "generate_mock_sales_transcript_prompt"
	summarizeTranscriptPrompt    = "sales_transcript_summary_prompt"
	queryTranscriptPrompt        = "sales_transcript_qa_prompt"
)

// LanguageModelService defines the contract for interacting with a language model API.
type LanguageModelService interface {
	// GenerateChatCompletion generates a chat completion for a given prompt.
	GenerateChatCompletion(prompt string, maxTokens int) (*string, error)

	// GenerateChatCompletionWithFiles generates a chat completion using a prompt and additional file inputs.
	GenerateChatCompletionWithFiles(prompt string, filePaths []string, maxTokens int) (*string, error)
}

type CallAnalyzerClient struct {
	api *api.OpenAIClient
}

func NewCallAnalyzerClient(apiKey string) *CallAnalyzerClient {
	return &CallAnalyzerClient{
		api: api.NewOpenAIClient(apiKey),
	}
}

// GenerateMockTranscript generates a mock sales transcript based on the prompt key.
func (ca *CallAnalyzerClient) GenerateMockTranscript() *string {
	filename := "prompt.yaml"
	prompt, err := loadPrompt(generateMockTranscriptPrompt, filename)
	if err != nil {
		log.Fatalf("Error loading prompt: %v", err)
	}
	generatedTranscript, err := ca.api.GenerateChatCompletionWithoutFiles(prompt, maxTokens)
	if err != nil {
		log.Fatalf("Error loading prompt: %v", err)
	}
	transcriptFileName, err := generateRandomFileName()
	if err != nil {
		log.Fatalf("Error while generating random file name: %v", err)
	}
	err = writeToFile(generatedTranscript, transcriptFileName)
	if err != nil {
		log.Fatalf("Error while writing to file: %v", err)
	}
	fmt.Println("Generated Transcript Successfully written to file: ", transcriptFileName)
	return generatedTranscript
}

// SummarizeTranscript summarizes a transcript file based on the prompt key.
func (ca *CallAnalyzerClient) SummarizeTranscript(filePath string) *string {
	filename := "prompt.yaml"
	prompt, err := loadPrompt(summarizeTranscriptPrompt, filename)
	if err != nil {
		log.Fatalf("Error loading prompt: %v", err)
	}
	filepaths := []string{filePath}
	summary, err := ca.api.GenerateChatCompletionWithFiles(prompt, maxTokens, filepaths...)
	if err != nil {
		log.Fatalf("Error while generating summary from transcript file: %v", err)
	}
	return summary
}

// QueryTranscript handles a query on the transcript file based on the prompt key.
func (ca *CallAnalyzerClient) QueryTranscript(filePath, query string) *string {
	filename := "prompt.yaml"
	prompt, err := loadPrompt(queryTranscriptPrompt, filename)
	if err != nil {
		log.Fatalf("Error loading prompt: %v", err)
	}
	combinedPrompt := prompt + "\n\n" + fmt.Sprint("Query Question: "+query)
	fmt.Println("combined Prompt: ", combinedPrompt)
	filepaths := []string{filePath}
	ans, err := ca.api.GenerateChatCompletionWithFiles(combinedPrompt, maxTokens, filepaths...)
	if err != nil {
		log.Fatalf("Error while generating answer for input query: %v", err)
	}

	return ans
}
