package agent

import (
	"ai-sre-agent/internal/tools"
)

func ExecutePlan(plan []tools.ToolCall) []ToolResult {

	var results []ToolResult

	for _, tool := range plan {

		output := tools.ExecuteTool(tool)

		results = append(results, ToolResult{
			Tool:   tool.Tool,
			Output: output,
		})
	}

	return results
}
