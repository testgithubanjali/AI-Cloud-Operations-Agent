package agent

import "ai-sre-agent/internal/rag"

func Process(userMessage string) (string, error) {

	// Step 1: Retrieve relevant documentation
	context, err := rag.GetRelevantContext(userMessage)
	if err != nil {
		return "", err
	}

	// Step 2: Create investigation plan
	plan, err := Plan(userMessage)
	if err != nil {
		return "", err
	}

	// Step 3: Execute tools
	results := ExecutePlan(plan)

	// Step 4: Analyze using tool results + documentation
	answer, err := Analyze(
		userMessage,
		context,
		results,
	)

	if err != nil {
		return "", err
	}

	return answer, nil
}
