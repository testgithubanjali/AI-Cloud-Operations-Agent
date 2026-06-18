package llm

import (
	"context"
	"fmt"
)

// Ask sends a message to the Gemini AI and returns the response
func Ask(message string) (string, error) {
	if message == "" {
		return "", fmt.Errorf("message cannot be empty")
	}

	// TODO: Implement actual Gemini API call
	// This is a placeholder implementation
	ctx := context.Background()
	_ = ctx // Use context if needed in actual implementation

	return fmt.Sprintf("Response to: %s", message), nil
}
