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

	// server := api.NewServer(
	// 	config,
	// 	store,
	// 	tsHandler,
	// 	log,
	// )

	// Create a Gin router with default middleware (logger and recovery)
	// r := gin.Default()

	// r.GET("/", func(c *gin.Context) {
	// 	// Return JSON response
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "DAMBASS",
	// 	})
	// })

	// // Define a simple GET endpoint
	// r.GET("/ping", func(c *gin.Context) {
	// 	// Return JSON response
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong masa siiii",
	// 	})
	// })

	// r.GET("/auth", func(c *gin.Context) {
	// 	c.JSON(http.StatusUnauthorized, gin.H{
	// 		"message": "auth missing hehe hoho",
	// 	})
	// })

	// if err := r.Run(); err != nil {
	// 	log.Fatalf("failed to run server: %v", err)
	// }
}
