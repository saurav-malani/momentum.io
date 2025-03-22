package main

import (
	"bufio"
	"fmt"
	"os"

	service "github.com/saurav-malani/momentumio/service"
	utility "github.com/saurav-malani/momentumio/utility"
)

type CallAnalyzer interface {
	GenerateMockTranscript() string
	SummarizeTranscript(filePath string) string
	QueryTranscript(filePath, query string) string
}

func main() {
	utility.LoadEnv()
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Println("Error: OPENAI_API_KEY not found in environment")
		return
	}
	callAnalyzer := service.NewCallAnalyzerClient(apiKey)
	for {
		{
			fmt.Println("----Welcome to Sales Helper----")
			fmt.Println("What feature would you like to make use of?")
			fmt.Println("Type '1' to Generate Mock Sales Transcript")
			fmt.Println("Type '2' to Summarize a Transcript File")
			fmt.Println("Type '3' Query a Transcript File")
			fmt.Println("Type '4' to Exit")
		}

		var input int
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Errorf("Error while taking input, %w", err)
		}

		switch input {
		case 1:
			fmt.Println(*callAnalyzer.GenerateMockTranscript())

		case 2:
			fmt.Println("Please provide the absolute Transcript File path you would like to Summarize.")
			var filename string
			fmt.Scan(&filename)
			fmt.Println(*callAnalyzer.SummarizeTranscript(filename))

		case 3:
			fmt.Println("Please provide the absolute file path of the Transcript File you would like to query.")
			var filename string
			fmt.Scan(&filename)
			fmt.Println("Please provide your query?")
			query, err := bufio.NewReader(os.Stdin).ReadString('\n')
			if err != nil {
				fmt.Println("Error while reading query: ", err)
				break
			}
			fmt.Println(*callAnalyzer.QueryTranscript(filename, query))

		case 4:
			fmt.Println("Exiting Sales Helper. Goodbye!")

		default:
			fmt.Println("Invalid option. Please enter a number between 1 and 4.")
		}

		if input == 4 {
			break
		}
	}
}
