package tools

import (
	"context"
	"fmt"
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

	// Debugging
	fmt.Println("===================================")
	fmt.Println("Namespace:", namespace)
	fmt.Println("Number of Events:", len(events.Items))
	fmt.Println("===================================")

	var result strings.Builder

	if len(events.Items) == 0 {
		return "No events found.", nil
	}

	for _, event := range events.Items {

		fmt.Println("Reason :", event.Reason)
		fmt.Println("Message:", event.Message)
		fmt.Println("-----------------------------------")

		result.WriteString("Reason: ")
		result.WriteString(event.Reason)
		result.WriteString("\n")

		result.WriteString("Message: ")
		result.WriteString(event.Message)
		result.WriteString("\n\n")
	}

	return result.String(), nil
}
