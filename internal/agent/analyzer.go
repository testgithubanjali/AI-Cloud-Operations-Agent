package agent

import (
	"strings"

	"ai-sre-agent/internal/llm"
)

func Analyze(userQuestion string, results []ToolResult) (string, error) {

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
You are an experienced Kubernetes Site Reliability Engineer.

The user asked:

` + userQuestion + `

Investigation Results:

` + report.String() + `

Your job is to perform a professional Root Cause Analysis.

Rules:

- Never invent Kubernetes information.
- Only use the investigation results.
- If information is missing, clearly say that more investigation is needed.
- If multiple tools provide conflicting information, explain the conflict instead of guessing.
- Keep the answer concise.

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
• If the pod name is known, use the actual pod name.
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
