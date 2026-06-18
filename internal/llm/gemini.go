package api

import (
	"net/http"

	"ai-sre-agent/internal/llm"
	"ai-sre-agent/internal/structs"

	"github.com/gin-gonic/gin"
)

func ChatHandler(c *gin.Context) {

	var req structs.ChatRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	answer, err := llm.Ask(req.Message)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, structs.ChatResponse{
		Answer: answer,
	})
}
