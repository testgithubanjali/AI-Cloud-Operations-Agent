package llm

import (
	"context"
	"os"

	"google.golang.org/genai"
)

func Ask(prompt string) (string, error) {

	apiKey := os.Getenv("GEMINI_API_KEY")

	client, err := genai.NewClient(
		context.Background(),
		&genai.ClientConfig{
			APIKey: apiKey,
		},
	)

	if err != nil {
		return "", err
	}

	resp, err := client.Models.GenerateContent(
		context.Background(),
		"gemini-2.5-flash",
		genai.Text(prompt),
		nil,
	)

	if err != nil {
		return "", err
	}

	return resp.Text(), nil
}
