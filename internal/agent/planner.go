package agent

import (
	"encoding/json"

	"ai-sre-agent/internal/llm"
	"ai-sre-agent/internal/tools"
)

func Process(userMessage string) (string, error) {

	// Step 1: Ask Gemini which tool to use
	response, err := llm.Ask(userMessage)
	if err != nil {
		return "", err
	}

	// Step 2: Convert Gemini's JSON response into a Go struct
	var toolCall tools.ToolCall

	err = json.Unmarshal([]byte(response), &toolCall)

	// If Gemini returned normal text instead of JSON,
	// just return the response.
	if err != nil {
		return response, nil
	}

	// Step 3: Execute the requested tool
	toolResult := tools.ExecuteTool(toolCall)

	// Step 4: Ask Gemini to explain the tool output
	finalPrompt := `
You are an AI Kubernetes SRE assistant.

The user asked:

` + userMessage + `

The tool returned:

` + toolResult + `

Generate a clear, helpful and concise answer for the user.
`

	finalAnswer, err := llm.Ask(finalPrompt)
	if err != nil {
		return "", err
	}

	return finalAnswer, nil
}
