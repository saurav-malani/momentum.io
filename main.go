package main

import (
	"flag"
	"fmt"
	"os"

	service "github.com/saurav-malani/momentumio/service"
	utility "github.com/saurav-malani/momentumio/utility"
)

func main() {
	utility.LoadEnv()

	// Verify that OPENAI_API_KEY is set
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Println("Error: OPENAI_API_KEY not found in environment")
		os.Exit(1)
	}

	// Initialize the call analyzer client
	callAnalyzer := service.NewCallAnalyzerClient(apiKey)

	// Define CLI flags
	generateFlag := flag.Bool("generate", false, "Generate a mock sales call transcript")
	summarizeFlag := flag.String("summarize", "", "Summarize a transcript file (provide file path)")
	queryFlag := flag.String("query", "", "Query a transcript file (provide file path)")
	queryText := flag.String("question", "", "The question to ask when querying a transcript")
	flag.Parse()

	// Handle `generate` flag
	if *generateFlag {
		transcript := callAnalyzer.GenerateMockTranscript()
		fmt.Println(*transcript)
		os.Exit(0)
	}

	// Handle `summarize` flag
	if *summarizeFlag != "" {
		summary := callAnalyzer.SummarizeTranscript(*summarizeFlag)
		fmt.Println(*summary)
		os.Exit(0)
	}

	// Handle `query` flag
	if *queryFlag != "" && *queryText != "" {
		answer := callAnalyzer.QueryTranscript(*queryFlag, *queryText)
		fmt.Println(*answer)
		os.Exit(0)
	}

	// Default behavior if no valid flags are provided
	fmt.Println("Usage:")
	fmt.Println("  -generate                  Generate a mock sales call transcript")
	fmt.Println("  -summarize <file_path>     Summarize a transcript file")
	fmt.Println("  -query <file_path>         Query a transcript file (requires -question)")
	fmt.Println("  -question <query_text>     The question to ask when querying a transcript")
	fmt.Println("Example:")
	fmt.Println("  ./main -generate")
	fmt.Println("  ./main -summarize transcript1.txt")
	fmt.Println("  ./main -query transcript1.txt -question \"What product was discussed?\"")
	os.Exit(1)
}
