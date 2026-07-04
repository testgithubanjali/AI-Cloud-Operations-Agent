package api

import (
	"net/http"

	"ai-sre-agent/internal/tools"

	"github.com/gin-gonic/gin"
)

func GetPods(c *gin.Context) {

	pods, err := tools.GetPods("default")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"pods": pods,
	})
}
