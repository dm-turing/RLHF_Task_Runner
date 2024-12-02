package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

// Config struct to hold configuration settings
type Config struct {
	DB DBConfig `toml:"db"`
}

type DBConfig struct {
	Host  string `toml:"host"`
	Port  int    `toml:"port"`
	Debug bool   `toml:"debug"`
}

func main() {
	var config Config
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		fmt.Println("Error decoding TOML file:", err)
		os.Exit(1)
	}

	fmt.Println("DB_HOST:", config.DB.Host)
	fmt.Println("DB_PORT:", config.DB.Port)
	fmt.Println("DEBUG:", config.DB.Debug)
}
