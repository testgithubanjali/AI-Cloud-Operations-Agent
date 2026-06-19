package api

import (
	"net/http"

	"ai-sre-agent/internal/jwt"

	"github.com/gin-gonic/gin"
)

func GetToken(c *gin.Context) {

	token, err := jwt.GenerateToken()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
