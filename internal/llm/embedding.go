package llm

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/genai"
)

func GenerateEmbedding(text string) ([]float32, error) {

	apiKey := os.Getenv("GEMINI_API_KEY")

	client, err := genai.NewClient(
		context.Background(),
		&genai.ClientConfig{
			APIKey: apiKey,
		},
	)
	if err != nil {
		return nil, err
	}

	resp, err := client.Models.EmbedContent(
		context.Background(),
		"gemini-embedding-001",
		genai.Text(text),
		nil,
	)
	if err != nil {
		return nil, err
	}

	if resp == nil || len(resp.Embeddings) == 0 {
		return nil, fmt.Errorf("no embedding returned")
	}

	return resp.Embeddings[0].Values, nil
}
