package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
	_ "go.uber.org/zap"

	"github.com/yoru-no-mel/go-shippr/api"
	conf "github.com/yoru-no-mel/go-shippr/config"
	db "github.com/yoru-no-mel/go-shippr/db/sqlc"
	"github.com/yoru-no-mel/go-shippr/logger"
	_ "github.com/yoru-no-mel/go-shippr/logger"
)

func main() {
	var (
		log    logger.Logger
		config conf.Config
	)
	env := os.Getenv("ENVIRONMENT")

	if env == "" || env == "dev" {
		env = "dev"
		logger, _ := zap.NewDevelopment()
		defer logger.Sync()
		log = logger.Sugar()
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

	server := api.NewServer(config, log)

	server.MountHandlers()

	addr := fmt.Sprintf(":%d", config.Port)

	srv := &http.Server{
		Addr:    addr,
		Handler: server.Router(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("Shutdown Server ...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Info("Server exiting")
}
