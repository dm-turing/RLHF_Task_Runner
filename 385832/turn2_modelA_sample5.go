package main

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DBHost string `envconfig:"APP_DB_HOST" required:"true"`
	DBPort int    `envconfig:"APP_DB_PORT" default:"5432"`
	Debug  bool   `envconfig:"APP_DEBUG" default:"false"`
}

func main() {
	var config Config
	err := envconfig.Process("APP", &config)
	if err != nil {
		fmt.Println("Error processing env config:", err)
		os.Exit(1)
	}

	fmt.Println("DB_HOST:", config.DBHost)
	fmt.Println("DB_PORT:", config.DBPort)
	fmt.Println("DEBUG:", config.Debug)
}
