package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yoru-no-mel/go-shippr/logger"
)

type Handler struct {
	log logger.Logger
}

func NewHandler(log logger.Logger) *Handler {
	return &Handler{
		log: log,
	}
}

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	userGroup := router.Group("/users")
	{
		userGroup.GET("", h.GetUsers)
	}
}

func (h *Handler) GetUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get Users",
	})
}
