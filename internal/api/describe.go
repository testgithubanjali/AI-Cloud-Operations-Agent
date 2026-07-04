package api

import (
	"net/http"

	"ai-sre-agent/internal/tools"

	"github.com/gin-gonic/gin"
)

func DescribePod(c *gin.Context) {

	// Get pod name from the URL
	podName := c.Param("name")

	// Call the Kubernetes tool
	result, err := tools.DescribePod("default", podName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Return the result
	c.JSON(http.StatusOK, gin.H{
		"description": result,
	})
}
