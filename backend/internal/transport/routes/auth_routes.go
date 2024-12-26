package routes

import (
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/transport/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.RouterGroup, authHandler *handlers.AuthHandler) {
	auth := router.Group("/auth")
	{
		auth.POST("/signup", authHandler.Signup)
	}
}
