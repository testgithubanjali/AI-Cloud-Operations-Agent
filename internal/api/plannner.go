package api

import (
	"net/http"

	"ai-sre-agent/internal/agent"

	"github.com/gin-gonic/gin"
)

func TestPlanner(c *gin.Context) {

	plan, err := agent.Plan("Why is nginx-pod restarting?")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, plan)
}
