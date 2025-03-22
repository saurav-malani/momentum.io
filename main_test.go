package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestMainFunction(t *testing.T) {
	// Create a buffer to capture the output
	var outputBuffer bytes.Buffer

	// Redirect standard output to the buffer
	originalStdout := os.Stdout
	rOut, wOut, _ := os.Pipe()
	os.Stdout = wOut

	// Create a pipe for simulating stdin
	originalStdin := os.Stdin
	rIn, wIn, _ := os.Pipe()
	os.Stdin = rIn

	// Write the simulated input into the stdin pipe
	go func() {
		defer wIn.Close()
		input := "1\n2\n3\n4\n" // Simulates input: "1" (generate mock transcript) and "4" (exit)
		wIn.Write([]byte(input))
	}()

	// Run the main function
	go func() {
		main()
		wOut.Close() // Close the write end of stdout pipe when main exits
	}()

	// Read captured output
	io.Copy(&outputBuffer, rOut)
	os.Stdout = originalStdout
	os.Stdin = originalStdin

	// Validate the captured output
	output := outputBuffer.String()

	// Check for expected messages
	expectedOutputs := []string{
		"----Welcome to Sales Helper----",
		"What feature would you like to make use of?",
		"Type '1' to Generate Mock Sales Transcript",
		"Type '2' to Summarize a Transcript File",
		"Type '3' Query a Transcript File",
		"Type '4' to Exit",
		"1",
		"generate mock transcript feature called.",
		"----Welcome to Sales Helper----",
		"What feature would you like to make use of?",
		"Type '1' to Generate Mock Sales Transcript",
		"Type '2' to Summarize a Transcript File",
		"Type '3' Query a Transcript File",
		"Type '4' to Exit",
		"2",
		"summarize transcript feature called.",
		"----Welcome to Sales Helper----",
		"What feature would you like to make use of?",
		"Type '1' to Generate Mock Sales Transcript",
		"Type '2' to Summarize a Transcript File",
		"Type '3' Query a Transcript File",
		"Type '4' to Exit",
		"3",
		"query transcript feature called.",
		"----Welcome to Sales Helper----",
		"What feature would you like to make use of?",
		"Type '1' to Generate Mock Sales Transcript",
		"Type '2' to Summarize a Transcript File",
		"Type '3' Query a Transcript File",
		"Type '4' to Exit",
		"4",
		"Exiting Sales Helper. Goodbye!",
	}

	for _, expected := range expectedOutputs {
		if !strings.Contains(output, expected) {
			t.Errorf("Expected output to contain: %q, but it was not found", expected)
		}
	}
}
