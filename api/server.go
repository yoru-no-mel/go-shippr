package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yoru-no-mel/go-shippr/config"
	"github.com/yoru-no-mel/go-shippr/logger"
	"github.com/yoru-no-mel/go-shippr/module/user"
)

type Server struct {
	config config.Config
	router *gin.Engine
	log    logger.Logger
}

func NewServer(config config.Config, log logger.Logger) *Server {
	var engine *gin.Engine
	engine = gin.Default()

	server := &Server{
		config: config,
		router: engine,
		log:    log,
	}

	return server
}

func (s *Server) MountHandlers() {
	api := s.router.Group("/api")

	api.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	userHandler := user.NewHandler(s.log)
	userHandler.RegisterRoutes(api)
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}

func (s *Server) Router() *gin.Engine {
	return s.router
}
