package api

import (
	"net/http"

	"ai-sre-agent/internal/llm"

	"github.com/gin-gonic/gin"
)

func ChatHandler(c *gin.Context) {

	answer, _ := llm.Ask("Hello")

	c.JSON(http.StatusOK, gin.H{
		"answer": answer,
	})
}
