package api

import (
	"net/http"

	"ai-sre-agent/internal/tools"

	"github.com/gin-gonic/gin"
)

func GetMetrics(c *gin.Context) {

	metrics, err := tools.GetPodMetrics("default")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"metrics": metrics,
	})
}
