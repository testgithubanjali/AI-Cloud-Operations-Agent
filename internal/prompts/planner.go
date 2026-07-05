package prompts

import (
	"strings"

	"ai-sre-agent/internal/tools"
)

func PlannerPrompt(userMessage string) string {

	var builder strings.Builder

	builder.WriteString(`
You are an AI Kubernetes Planner.

Available tools:

`)

	for _, tool := range tools.AvailableTools {

		builder.WriteString("- ")
		builder.WriteString(tool.Name)
		builder.WriteString("\n")

		builder.WriteString("  Description: ")
		builder.WriteString(tool.Description)
		builder.WriteString("\n\n")
	}

	builder.WriteString(`
Rules:

- Decide which tools are required.
- Return ONLY valid JSON.
- Return a JSON array.
- Do not explain anything.
- Do not use Markdown.

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
    "tool":"get_metrics",
    "pod":"nginx-pod"
  }
]

User:

`)

	builder.WriteString(userMessage)

	return builder.String()
}
