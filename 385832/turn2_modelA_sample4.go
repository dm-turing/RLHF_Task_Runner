package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type Config struct {
	DBHost string `validate:"required"`
	DBPort int    `validate:"required,min=1,max=65535"`
	Debug  bool   `validate:"required"`
}

func main() {
	var config Config
	config.DBHost = os.Getenv("DB_HOST")
	config.DBPort, _ = strconv.Atoi(os.Getenv("DB_PORT"))
	config.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))

	validate := validator.New()
	err := validate.Struct(&config)
	if err != nil {
		fmt.Println("Error validating config:", err)
		os.Exit(1)
	}

	fmt.Println("DB_HOST:", config.DBHost)
	fmt.Println("DB_PORT:", config.DBPort)
	fmt.Println("DEBUG:", config.Debug)
}
