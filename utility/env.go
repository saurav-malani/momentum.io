package utility

import (
	"log"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/joho/godotenv"
)

var once sync.Once

// LoadEnv loads environment variables from the root directory's .env file.
func LoadEnv() {
	once.Do(func() {
		// Get the directory of the current file (config.go)
		_, file, _, _ := runtime.Caller(0)

		rootDir := filepath.Join(filepath.Dir(file), "..")

		// Construct the path to the .env file
		envPath := filepath.Join(rootDir, ".env")

		// Load the .env file
		err := godotenv.Load(envPath)
		if err != nil {
			log.Printf("Warning: .env file not found at %s or could not be loaded", envPath)
		}
	})
}
