package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AppEnv      string `mapstructure:"app_env"`
	DatabaseUrl string `mapstructure:"database_url"`
}

func GetConfig() *Config {
	viper.SetDefault("AppEnv", "prod")
	viper.SetDefault("AppEnv", "postgres://postgres@localhost:5432/dev?sslmode=disable")
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	var config Config
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatal(err)
	}
	return &config
}
