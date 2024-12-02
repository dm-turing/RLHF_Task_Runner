package main

import (
	"envvar"
	"fmt"
)

func main() {
	dbHost := envvar.GetEnv("DB_HOST", "localhost")
	fmt.Println("DB_HOST:", dbHost)

	port, err := envvar.GetEnv("PORT", "8080")
	if err != nil {
		fmt.Println("Error getting PORT:", err)
		return
	}
	fmt.Println("PORT:", port)
}
