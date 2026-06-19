package routes

import (
	"ai-sre-agent/internal/api"
	"ai-sre-agent/internal/jwt"

	"github.com/gin-gonic/gin"
)

func ChatRoutes(r *gin.Engine) {

	protected := r.Group("/")

	protected.Use(jwt.AuthMiddleware())

	{
		protected.POST("/chat", api.ChatHandler)
	}
}
