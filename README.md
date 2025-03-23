# Momentum.io: AI-Powered Sales Call Transcript Generator & Analyzer

Momentum.io is a command-line tool that leverages **OpenAI GPT APIs** to generate, summarize, and analyze mock sales call transcripts. This application is designed to help businesses streamline the management of sales calls, extract key insights, and answer specific questions related to the transcript content.

---

## **Features**

1. **Generate Call Transcripts**

   - Create realistic mock sales call transcripts using OpenAI GPT APIs.
   - Transcripts are output to the console and saved to a file for later use.

   Example transcript:

   ```text
   00:00:00 Sam (openai.com): Hey there Staya.
   00:00:02 Satya (microsoft.com): Hi Sam, how are you?
   00:00:05 Sam (openai.com): I'm doing good. Do you think you can give us 10000 more GPUs?
   ```

2. **Summarize Call Transcripts**

   - Generate a concise summary of the key points from a given transcript file.
   - Output the summary directly to the console.

3. **Answer Questions**

   - Take a user query and provide relevant answers based on the transcript content.
   - Example question: _“What product was the customer interested in?”_

4. **Bonus Features** (Optional but implemented for demonstration):
   - Support for multiple languages (e.g., Spanish and French).
   - Unit tests for key functionality (e.g., transcript generation, summary, and Q&A).

---

## **Directory Structure**

```plaintext
.
├── api/
│   ├── openai/
│   │   ├── openai.go               # Handles interaction with OpenAI GPT APIs
│   │   └── openai_test.go          # Unit tests for OpenAI API interactions
│   └── testdata/                   # Sample test transcripts
│       ├── transcript1.txt
│       └── transcript2.txt
├── main.go                         # Entry point for CLI functionality
├── prompt.yaml                     # Config file for OpenAI prompts
├── service/
│   ├── helper.go                   # Helper functions for parsing and formatting transcripts
│   ├── service.go                  # Core business logic (e.g., generation, summarization, Q&A)
│   ├── service_test.go             # Unit tests for service functionality
│   └── testdata/
│       └── testPrompt.yaml         # Sample prompts for testing
├── utility/
│   └── env.go                      # Utility to load environment variables
├── go.mod                          # Go module definition
├── go.sum                          # Go dependencies
└── README.md                       # Project documentation
```

---

## **Setup Instructions**

### **Prerequisites**

- **Go** (version 1.18 or higher)
- **OpenAI API Key**: Obtain your API key from the [OpenAI platform](https://platform.openai.com/).
- Ensure the following environment variable is set:
  ```bash
  export OPENAI_API_KEY=<your-api-key>
  ```

### **Installation**

1. Clone this repository:
   ```bash
   git clone <repository_url>
   cd momentum.io
   ```
2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Set up the `.env` file:
   Create a `.env` file in the project root with the following content:
   ```plaintext
   OPENAI_API_KEY=<your-api-key>
   ```

---

## **Usage Guide**

### **1. Generate Call Transcripts**

Run the following command to generate a mock sales call transcript:

```bash
go run main.go -generate
```

Output:

- The transcript is printed to the console and saved to `generatedTranscript-<unique_id>.txt` in the project root.

---

### **2. Summarize Call Transcripts**

Run the following command to summarize a transcript:

```bash
go run main.go -summarize <path-to-transcript>
```

Example:

```bash
go run main.go -summarize "generatedTranscript-5862b118d6958dda.txt"
```

Output:

- A concise summary of the transcript printed to the console.

---

### **3. Answer Questions**

Run the following command to answer a question based on a transcript:

```bash
go run main.go question --file <path-to-transcript> --query "<user-question>"
```

Example:

```bash
go run main.go -query "generatedTranscript-f68099a7252fff71.txt" -question "What specific features or workflows does your team currently use in project management, and how do you envision an ideal solution improving these processes?"
```

Output:

- A relevant answer to the query, printed to the console.

---

## **Design Decisions**

1. **Modular Codebase**  
   The project is split into modules:

   - `api`: Manages interactions with OpenAI APIs.
   - `service`: Implements core business logic for transcript generation, summarization, and Q&A.
   - `utility`: Provides utility functions like environment variable loading.

2. **Config-Driven Design**

   - Prompts for OpenAI are stored in `prompt.yaml`, making it easy to tweak or extend without modifying the code.

3. **Error Handling**
   - Comprehensive error handling is implemented to handle missing files, invalid inputs, or API errors gracefully.

---

## **Testing**

Unit tests are written for core functionalities using Go's built-in `testing` package. To run the tests:

```bash
go test ./...
```

Test coverage includes:

- OpenAI API interactions (`api/openai_test.go`)
- Transcript summarization and Q&A functionality (`service/service_test.go`)

---
