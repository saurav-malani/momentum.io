package main

import (
	"fmt"
)

/*
Upon running this 3 options will be displayed in console.
those 3 options will be corrsponding to the 3 features we provide that a user can make use of.
1. Generate Mock sales transcript.
Upon calling this function. A mock sales (very realistic looking) transcript will be generated, making use of chatgpt API.
This will be printed in console and also saved to a local file.

2. Given location of a transcript file as input, return summary (using chatgpt API) of important points that could be of use in context of understanding the clients requirement and closing the deal.
3. Given Transcript file & query as input, return answer for the query, making use of the chatgpt API.
*/

// Upon calling this function. A mock sales (very realistic looking) transcript will be generated, making use of chatgpt API.
func generateMockTranscript() {
	fmt.Println("generate mock transcript feature called.")
}

// Given location of a transcript file as input, return summary (using chatgpt API) of important points that could be of use in context
//
//	of understanding the clients requirement and closing the deal.
func summarizeTranscriptFile() {
	fmt.Println("summarize transcript feature called.")
}

// Given Transcript file & query as input, return answer for the query, making use of the chatgpt API.
func queryTranscriptFile() {
	fmt.Println("query transcript feature called.")
}

func main() {
	for {
		fmt.Println("----Welcome to Sales Helper----")
		fmt.Println("What feature would you like to make use of?")
		fmt.Println("Type '1' to Generate Mock Sales Transcript")
		fmt.Println("Type '2' to Summarize a Transcript File")
		fmt.Println("Type '3' Query a Transcript File")
		fmt.Println("Type '4' to Exit")
		var input int
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Errorf("Error while taking input, %w", err)
		}

		switch input {
		case 1:
			generateMockTranscript()

		case 2:
			summarizeTranscriptFile()

		case 3:
			queryTranscriptFile()

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
