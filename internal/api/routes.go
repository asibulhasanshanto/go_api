package api

import (
	"github.com/asibulhasanshanto/go_api/internal/api/handlers"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SetupRoutes(r *gin.Engine, authHandler *handlers.AuthHandler, log *zap.Logger) *gin.RouterGroup {
	root := r.Group("/")
	root.GET("healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/signup", authHandler.Signup)
			auth.POST("/login", authHandler.Login)
			auth.GET("/refresh-access-token", authHandler.RefreshAccessToken)
			auth.POST("/logout", authHandler.Logout)
			auth.GET("/me", authHandler.Me)
		}

	}
	return root
}
