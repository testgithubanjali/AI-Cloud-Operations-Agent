package api

import (
	"net/http"

	"ai-sre-agent/internal/tools"

	"github.com/gin-gonic/gin"
)

func TestConnection(c *gin.Context) {

	_, err := tools.NewClient()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Connected to Kubernetes",
	})
}
