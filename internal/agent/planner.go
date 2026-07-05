package agent

import (
	"encoding/json"

	"ai-sre-agent/internal/llm"
	"ai-sre-agent/internal/tools"
)

func Plan(userMessage string) ([]tools.ToolCall, error) {

	plannerPrompt := `
You are an AI Kubernetes Planner.

Your job is NOT to answer the user.

Your job is ONLY to decide which tools are needed.

Available tools:

- get_pods
- describe_pod
- get_logs
- get_metrics

Rules:

- Return ONLY a JSON array.
- Do not explain.
- Do not use markdown.

Example:

User:
Why is nginx-pod restarting?

Return:

[
  {
    "tool":"describe_pod",
    "pod":"nginx-pod"
  },
  {
    "tool":"get_logs",
    "pod":"nginx-pod"
  },
  {
    "tool":"get_metrics"
  }
]

User:

` + userMessage

	response, err := llm.Ask(plannerPrompt)
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
