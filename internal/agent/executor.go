package agent

import (
	"sync"

	"ai-sre-agent/internal/tools"
)

func ExecutePlan(plan []tools.ToolCall) []ToolResult {

	var wg sync.WaitGroup

	results := make([]ToolResult, len(plan))

	for i, tool := range plan {

		wg.Add(1)

		go func(index int, tool tools.ToolCall) {
			defer wg.Done()

			output, err := tools.ExecuteTool(tool)

			result := ToolResult{
				Tool: tool.Tool,
			}

			if err != nil {
				result.Successful = false
				result.Error = err.Error()
			} else {
				result.Successful = true
				result.Output = output
			}

			results[index] = result

		}(i, tool)
	}

	wg.Wait()

	return results
}
