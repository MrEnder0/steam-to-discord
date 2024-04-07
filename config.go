package main

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	GroupName  string
	WebhookURL string
	CheckFreq  int
}

func LoadConfig() Config {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

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
	viper.SetConfigType("yaml")

	viper.SetDefault("GroupName", "steamgroupname")
	viper.SetDefault("WebhookURL", "discordwebhookurl")
	viper.SetDefault("CheckFreq", 180)

	err := viper.WriteConfigAs("config.yaml")
	if err != nil {
		log.Println("Error writing config file:", err)
		os.Exit(1)
	}

	log.Println("Config file generated, please fill in the details and re-run the program.")
	CreateDatabase()
	os.Exit(1)
}
