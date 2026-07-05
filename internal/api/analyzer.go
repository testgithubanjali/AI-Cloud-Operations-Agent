package api

import (
	"net/http"

	"ai-sre-agent/internal/agent"

	"github.com/gin-gonic/gin"
)

func TestAnalyzer(c *gin.Context) {

	results := []agent.ToolResult{
		{
			Tool:   "describe_pod",
			Output: "Restart Count: 2",
		},
		{
			Tool:   "get_logs",
			Output: "OOMKilled",
		},
		{
			Tool:   "get_metrics",
			Output: "CPU: 0m\nMemory: 980Mi",
		},
	}

	answer, err := agent.Analyze(
		"Why is nginx-pod restarting?",
		results,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"answer": answer,
	})
}
