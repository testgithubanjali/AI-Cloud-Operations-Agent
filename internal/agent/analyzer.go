package agent

import (
	"strings"

	"ai-sre-agent/internal/llm"
)

func Analyze(userQuestion string, context string, results []ToolResult) (string, error) {

	var report strings.Builder

	// Build investigation report
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

	// If no tools were executed
	if report.Len() == 0 {
		report.WriteString("No Kubernetes tools were executed because the question did not require live cluster investigation.\n\n")
	}

	// If no documentation was found
	if strings.TrimSpace(context) == "" {
		context = "No relevant internal documentation was found."
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
- If documentation does not answer the question, clearly state that.
- Never invent Kubernetes information.
- Keep the answer concise and professional.

Return exactly in this format:

🚨 Incident Analysis

Severity:
(LOW / MEDIUM / HIGH)

Root Cause:
(Explain the most likely reason.)

Evidence:
• Bullet point
• Bullet point

Impact:
(Explain what could happen.)

Recommended Fix:
• Bullet point
• Bullet point

Useful kubectl Commands:
• kubectl describe pod <pod-name>
• kubectl logs <pod-name> --previous
• kubectl top pod <pod-name>
`

	answer, err := llm.Ask(prompt)
	if err != nil {
		return "", err
	}

	return answer, nil
}
