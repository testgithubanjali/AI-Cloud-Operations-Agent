package llm

import (
	"context"
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

	// We'll implement the API call in the next step.
	_ = client
	_ = text

	return nil, nil
}
