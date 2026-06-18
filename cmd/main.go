package main

import (
	"log"

	"ai-sre-agent/internal/api"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	router := gin.Default()

	router.POST("/chat", api.ChatHandler)

	log.Println("Server running on :3000")

	err = router.Run(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
