package agent

import (
	"ai-sre-agent/internal/tools"
)

func ExecutePlan(plan []tools.ToolCall) string {

	var allResults string

	for _, tool := range plan {

		result := tools.ExecuteTool(tool)

		allResults += result + "\n\n"
	}

	return allResults
}
