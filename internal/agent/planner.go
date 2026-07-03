package agent

import (
	"encoding/json"

	"ai-sre-agent/internal/llm"
	"ai-sre-agent/internal/tools"
)

func Process(userMessage string) (string, error) {

	// Ask Gemini what tool should be used
	response, err := llm.Ask(userMessage)
	if err != nil {
		return "", err
	}

	// Try to parse Gemini's response as JSON
	var toolCall tools.ToolCall

	err = json.Unmarshal([]byte(response), &toolCall)

	// If it's not JSON, Gemini answered normally.
	if err != nil {
		return response, nil
	}

	// Execute the requested tool
	result := tools.ExecuteTool(toolCall)

	return result, nil
}
