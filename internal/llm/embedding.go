package llm

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

func GenerateEmbedding(text string) ([]float32, error) {

	// Load .env
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("failed to load .env: %w", err)
	}

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("GEMINI_API_KEY is empty")
	}

	client, err := genai.NewClient(
		context.Background(),
		&genai.ClientConfig{
			APIKey: apiKey,
		},
	)
	if err != nil {
		return nil, err
	}

	// Generate embedding
	resp, err := client.Models.EmbedContent(
		context.Background(),
		"gemini-embedding-001",
		genai.Text(text),
		nil,
	)
	if err != nil {
		return nil, err
	}

	// Safety checks
	if resp == nil {
		return nil, fmt.Errorf("nil response from Gemini")
	}

	if len(resp.Embeddings) == 0 {
		return nil, fmt.Errorf("no embeddings returned")
	}

	return resp.Embeddings[0].Values, nil
}
