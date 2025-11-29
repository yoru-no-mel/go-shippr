package main

import (
	"fmt"
	"os"

	_ "go.uber.org/zap"

	conf "github.com/yoru-no-mel/go-shippr/config"
	db "github.com/yoru-no-mel/go-shippr/db/sqlc"
	_ "github.com/yoru-no-mel/go-shippr/logger"
)

func main() {
	var (
		// log  logger.Logger
		config conf.Config
	)
	env := os.Getenv("ENVIRONMENT")

	if env == "" || env == "dev" {
		env = "dev"
		// logger, _ := zap.NewDevelopment()
		// defer logger.Sync()
		// log = logger.Sugar()
		config = conf.LoadConfig(env, "./env")
	}
	fmt.Println(config.DBHost)

	config.DBUrl = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)

	dbConn := db.Connect(config)
	defer db.Close(dbConn)

	db.AutoMigrate(config)
}
