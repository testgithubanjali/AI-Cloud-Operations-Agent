package agent

import (
	"strings"

	"ai-sre-agent/internal/llm"
)

func Analyze(userQuestion string, results []ToolResult) (string, error) {

	var report strings.Builder

	for _, result := range results {

		report.WriteString(result.Tool)
		report.WriteString("\n")

		report.WriteString(result.Output)
		report.WriteString("\n\n")
	}

	prompt := `
You are an experienced Kubernetes Site Reliability Engineer.

The user asked:

` + userQuestion + `

Investigation Results:

` + report.String() + `

Analyze everything carefully.

Return:

1. Root Cause

2. Evidence

3. Recommended Fix

Keep the answer concise.
`

	answer, err := llm.Ask(prompt)
	if err != nil {
		return "", err
	}

	return answer, nil
}
