package main

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
)

// Config struct to hold environment variables
type Config struct {
	DBHost string `envconfig:"DB_HOST" required:"true"`
	DBPort int    `envconfig:"DB_PORT" default:"5432"`
	Debug  bool   `envconfig:"DEBUG" default:"false"`
}

func main() {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		fmt.Println("Error processing env config:", err)
		os.Exit(1)
	}

	fmt.Println("DB_HOST:", config.DBHost)
	fmt.Println("DB_PORT:", config.DBPort)
	fmt.Println("DEBUG:", config.Debug)
}
