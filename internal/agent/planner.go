package agent

import (
	"encoding/json"

	"ai-sre-agent/internal/llm"
	"ai-sre-agent/internal/tools"
)

func Process(userMessage string) (string, error) {

	// Step 1: Tell Gemini about all available tools
	toolPrompt := `
You are an AI Kubernetes SRE Assistant.

You have the following tools:

1. get_pods
   - List all Kubernetes pods.

2. describe_pod
   - Describe a Kubernetes pod.

3. get_logs
   - Get logs of a Kubernetes pod.

4. get_metrics
   - Get CPU and memory usage of Kubernetes pods.

Rules:

- If a tool is required, return ONLY valid JSON.
- Do not explain anything.
- Do not return Markdown.
- Do not wrap JSON inside code blocks.

Examples:

User: List all pods

{
	"tool":"get_pods",
	"namespace":"default"
}

User: Describe nginx-pod

{
	"tool":"describe_pod",
	"pod":"nginx-pod"
}

User: Show logs of nginx-pod

{
	"tool":"get_logs",
	"pod":"nginx-pod"
}

User: Show CPU usage

{
	"tool":"get_metrics"
}

Now answer this user request:

` + userMessage

	// Ask Gemini which tool to use
	response, err := llm.Ask(toolPrompt)
	if err != nil {
		return "", err
	}

	// Step 2: Convert Gemini's JSON response into Go struct
	var toolCall tools.ToolCall

	err = json.Unmarshal([]byte(response), &toolCall)

	// If Gemini returned normal text instead of JSON,
	// return it directly.
	if err != nil {
		return response, nil
	}

	// Step 3: Execute the selected tool
	toolResult := tools.ExecuteTool(toolCall)

	// Step 4: Ask Gemini to explain the tool result
	finalPrompt := `
You are an AI Kubernetes SRE Assistant.

The user asked:

` + userMessage + `

The tool returned:

` + toolResult + `

Generate a clear, helpful and concise answer.

If metrics are returned:
- Explain CPU usage.
- Explain memory usage.
- Tell whether the pod appears healthy.

If logs are returned:
- Summarize the important information.

If pod details are returned:
- Explain the pod status in simple language.

Do not mention JSON.
`

	finalAnswer, err := llm.Ask(finalPrompt)
	if err != nil {
		return "", err
	}

	return finalAnswer, nil
}
