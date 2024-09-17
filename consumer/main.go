package main

import (
	"fmt"
	"log"
	"os"

	"github.com/christoff-linde/pih-core-go/db"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	exampleEnvVar := os.Getenv("EXAMPLE_ENV_VAR")
	fmt.Println("Found EXAMPLE_ENV_VAR with value:", exampleEnvVar)

	db.InitDb("consumer")
}
