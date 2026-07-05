package api

import (
	"net/http"

	"ai-sre-agent/internal/tools"

	"github.com/gin-gonic/gin"
)

func PrometheusTest(c *gin.Context) {

	query := c.Query("query")

	result, err := tools.QueryPrometheus(query)
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}
