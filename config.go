package main

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

type Config struct {
	FirebaseTokenPath string
	NgrokToken        string
	NgrokPort         string
}

func NewConfig(path string) *Config {
	file, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	viper.SetConfigType("yaml")

	err = viper.ReadConfig(file)
	if err != nil {
		log.Panicln(err)
	}
	return &Config{
		FirebaseTokenPath: viper.GetString("firebase_token_path"),
		NgrokToken:        viper.GetString("ngrok.token"),
		NgrokPort:         viper.GetString("ngrok.port"),
	}
}
