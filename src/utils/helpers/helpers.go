package helpers

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func IsInvalid(value string) bool {
	if value == "" {
		return true
	}
	return false
}

func GoDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load("$GOPATH/src/github.com/rezwanul-haque/ID-Service/.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}