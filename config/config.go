package config

import {
	"log"
	"time"

	"github.com/spf13/viper"
}

type Config struct {
	Environment string `mapstructure:"ENVIRONMENT"`
	
	DB struct {
		Host     string `mapstructure:"DB_HOST"`
		Port     int    `mapstructure:"DB_PORT"`
		User     string `mapstructure:"DB_USER"`
		Password string `mapstructure:"DB_PASSWORD"`
		Name     string `mapstructure:"DB_NAME"`
		Url      string
	}
}

func LoadConfig(name string, path string) (config Config) {
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("config: $v", err)
		return
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("config: $v", err)
		return
	}

	return
}
	