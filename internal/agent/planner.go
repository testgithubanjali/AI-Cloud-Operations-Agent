package agent

import (
	"encoding/json"

	"ai-sre-agent/internal/llm"
	"ai-sre-agent/internal/prompts"
	"ai-sre-agent/internal/tools"
)

func Plan(userMessage string) ([]tools.ToolCall, error) {

	prompt := prompts.PlannerPrompt(userMessage)

	response, err := llm.Ask(prompt)
	if err != nil {
		return nil, err
	}

	var plan []tools.ToolCall

	err = json.Unmarshal([]byte(response), &plan)
	if err != nil {
		return nil, err
	}

	return plan, nil
}
