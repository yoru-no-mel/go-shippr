package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Environment     string `mapstructure:"ENVIRONMENT"`
	DBHost          string `mapstructure:"DB_HOST"`
	DBPort          int    `mapstructure:"DB_PORT"`
	DBUser          string `mapstructure:"DB_USER"`
	DBPassword      string `mapstructure:"DB_PASS"`
	DBName          string `mapstructure:"DB_NAME"`
	DBRecreate      bool   `mapstructure:"DB_RECREATE"`
	DBMigrationPath string `mapstructure:"DB_MIGRATION_PATH"`
	DBUrl           string
	Port            int `mapstructure:"PORT"`
}

func LoadConfig(name string, path string) (config Config) {
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("config: %v", err)
		return
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("config: %v", err)
		return
	}

	return
}
