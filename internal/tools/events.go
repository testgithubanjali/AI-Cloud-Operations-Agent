package tools

import (
	"context"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetEvents(namespace string) (string, error) {

	client, err := NewClient()
	if err != nil {
		return "", err
	}

	events, err := client.CoreV1().
		Events(namespace).
		List(context.Background(), metav1.ListOptions{})

	if err != nil {
		return "", err
	}

	var result strings.Builder

	for _, event := range events.Items {

		result.WriteString(event.Reason)
		result.WriteString(": ")
		result.WriteString(event.Message)
		result.WriteString("\n\n")
	}

	return result.String(), nil
}
