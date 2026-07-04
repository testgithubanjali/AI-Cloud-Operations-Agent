package routes

import (
	"ai-sre-agent/internal/api"

	"github.com/gin-gonic/gin"
)

func TestRoutes(r *gin.Engine) {

	r.GET("/test", api.TestConnection)

}
