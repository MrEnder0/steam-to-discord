package main

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func LoadConfig() Config {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("toml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Error reading config file, generating new one...")
		genConfig()
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Println("Error unmarshalling config file, generating new one...")
		genConfig()
	}

	return config
}

func genConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("toml")

	viper.SetDefault("GroupName", "steamgroupname")
	viper.SetDefault("WebhookURL", "https://discord.com/api/webhooks/id")
	viper.SetDefault("CheckFreq", 60)
	viper.SetDefault("ShowSteamPrefix", true)

	err := viper.WriteConfigAs("config.toml")
	if err != nil {
		log.Println("Error writing config file:", err)
		os.Exit(1)
	}

	log.Println("Config file generated, please fill in the details and re-run the program.")
	CreateDatabase()
	os.Exit(1)
}
