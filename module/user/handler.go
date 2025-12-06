package user

import (
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"github.com/yoru-no-mel/go-shippr/logger"
)

type userRegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Type     string `json:"type" binding:"required"`
}

type Handler struct {
	log logger.Logger
}

func NewHandler(log logger.Logger) *Handler {
	return &Handler{
		log: log,
	}
}

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	userGroup := router.Group("/user")
	{
		userGroup.GET("", h.GetUsers)
		userGroup.POST("/register", h.RegisterUser)
	}
}

func (h *Handler) GetUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get Users",
	})
}

func (h *Handler) RegisterUser(ctx *gin.Context) {
	var req userRegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if !valid(req.Email) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email",
		})
		return
	}
	//response okay and the whole request body with custom message register success
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Register success",
		"user":    req,
	})
}

func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
