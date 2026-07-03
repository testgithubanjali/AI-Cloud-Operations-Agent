package main

import (
	"log"

	"ai-sre-agent/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
	r := gin.Default()

	routes.AuthRoutes(r)
	routes.ChatRoutes(r)

	r.Run(":8080")

	log.Println("Server running on :8080")

	r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
