package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigType("yaml") // Assuming you have a config.yaml file
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("APP")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	dbHost := viper.GetString("db.host")
	fmt.Println("DB_HOST:", dbHost)
}
