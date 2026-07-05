package agent

func Process(userMessage string) (string, error) {

	// Step 1: Create an investigation plan
	plan, err := Plan(userMessage)
	if err != nil {
		return "", err
	}

	// Step 2: Execute all the planned tools
	results := ExecutePlan(plan)

	// Step 3: Analyze the collected results
	answer, err := Analyze(userMessage, results)
	if err != nil {
		return "", err
	}

	return answer, nil
}
