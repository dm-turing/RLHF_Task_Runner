package main

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DBHost string `envconfig:"DB_HOST" required:"true"`
	DBPort int    `envconfig:"DB_PORT" default:"3306"`
	Debug  bool   `envconfig:"DEBUG" default:"false"`
}

func main() {
	var config Config
	if err := envconfig.Process("", &config); err != nil {
		fmt.Println("Error processing env config:", err)
		return
	}

	fmt.Println("Config:", config)
}
