package main

import (
	"log"

	"ai-sre-agent/internal/api"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	// Create Fiber app
	app := fiber.New()

	// Register routes
	app.Post("/chat", api.ChatHandler)

	// Start server
	log.Println("Server running on port 3000")
	log.Fatal(app.Listen(":3000"))
}
