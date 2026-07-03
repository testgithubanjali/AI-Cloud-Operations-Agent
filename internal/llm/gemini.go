package llm

import (
	"context"
	"os"

	"google.golang.org/genai"
)

func Ask(userMessage string) (string, error) {

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

	// Step 4: Build a prompt that tells Gemini about the available tools
	fullPrompt := `
You are an AI Kubernetes SRE assistant.

You have access to the following tools:

1. get_pods
   - Returns all pods in a namespace.

2. describe_pod
   - Returns detailed information about a pod.

3. get_logs
   - Returns logs of a pod.

If the user's question requires one of these tools, respond ONLY with JSON.

Examples:

User: Show all pods in the default namespace

Response:
{
  "tool": "get_pods",
  "namespace": "default"
}

User: Describe postgres-0

Response:
{
  "tool": "describe_pod",
  "namespace": "default",
  "pod": "postgres-0"
}

User: Show logs of postgres-0

Response:
{
  "tool": "get_logs",
  "namespace": "default",
  "pod": "postgres-0"
}

If the user's question does NOT require a tool (for example, "What is Kubernetes?"),
answer normally.

User:
` + userMessage

	resp, err := client.Models.GenerateContent(
		context.Background(),
		"gemini-2.5-flash",
		genai.Text(fullPrompt),
		nil,
	)

	if err != nil {
		return "", err
	}

	return resp.Text(), nil
}
