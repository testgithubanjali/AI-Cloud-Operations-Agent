package main

import (
	"log"

	"ai-sre-agent/internal/rag"
	"ai-sre-agent/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize RAG knowledge base
	if err := rag.InitializeKnowledgeBase(); err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	routes.AuthRoutes(r)
	routes.ChatRoutes(r)
	routes.TestRoutes(r)

	log.Println("Server running on :8081")

	if err := r.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
