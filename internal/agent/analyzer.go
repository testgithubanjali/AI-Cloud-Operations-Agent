package agent

import (
	"strings"

	"ai-sre-agent/internal/llm"
)

func Analyze(userQuestion string, context string, results []ToolResult) (string, error) {

	var report strings.Builder

	for _, result := range results {

		report.WriteString("========== ")
		report.WriteString(result.Tool)
		report.WriteString(" ==========\n")

		if result.Successful {
			report.WriteString(result.Output)
		} else {
			report.WriteString("ERROR: ")
			report.WriteString(result.Error)
		}

		report.WriteString("\n\n")
	}
	prompt := `
You are an expert Kubernetes Site Reliability Engineer.

You have two sources of information:

1. Live Kubernetes investigation results (highest priority)
2. Internal documentation (supporting information)

If they conflict, ALWAYS trust the live investigation results.

Relevant Documentation:

` + context + `

Investigation Results:

` + report.String() + `

User Question:

` + userQuestion + `

Instructions:

- Use the live investigation results as the primary source of truth.
- Use the documentation only to explain concepts or recommend fixes.
- If documentation does not answer the question, say so.
- Never invent Kubernetes information.
- Keep the answer concise.

Return exactly in this format:

🚨 Incident Analysis

Severity:

Root Cause:

Evidence:
• ...

Impact:

Recommended Fix:
• ...

Useful kubectl Commands:
• ...
`

	answer, err := llm.Ask(prompt)
	if err != nil {
		return "", err
	}

	return answer, nil
}
