package routes

import (
	"ai-sre-agent/internal/api"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	r.GET("/token", api.GetToken)
}
