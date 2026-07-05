package api

import (
	"net/http"

	"ai-sre-agent/internal/agent"
	"ai-sre-agent/internal/tools"

	"github.com/gin-gonic/gin"
)

func TestExecutor(c *gin.Context) {

	plan := []tools.ToolCall{
		{
			Tool: "describe_pod",
			Pod:  "nginx-pod",
		},
		{
			Tool: "get_logs",
			Pod:  "nginx-pod",
		},
		{
			Tool: "get_metrics",
		},
	}

	results := agent.ExecutePlan(plan)

	c.JSON(http.StatusOK, results)
}
