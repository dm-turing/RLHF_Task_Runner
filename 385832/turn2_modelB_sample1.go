package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		os.Exit(1)
	}

	fmt.Println("DB_HOST:", os.Getenv("DB_HOST"))
}
