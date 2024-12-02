package main

import (
	"envvar"
	"fmt"
)

func main() {
	apiKey := envvar.GetEnvWithDefault("API_KEY", "your_default_api_key")
	appName := envvar.GetEnvWithDefault("APP_NAME", "Demo App")

	fmt.Printf("API Key: %s\n", apiKey)
	fmt.Printf("App Name: %s\n", appName)
}
