package api

import (
	"net/http"

	"ai-sre-agent/internal/tools"

	"github.com/gin-gonic/gin"
)

func GetEvents(c *gin.Context) {

	events, err := tools.GetEvents("default")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"events": events,
	})
}
