package api

import (
	"net/http"

	"ai-sre-agent/internal/llm"

	"github.com/gin-gonic/gin"
)

func ChatHandler(c *gin.Context) {

	answer, err := llm.Ask("Hello")

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
