package tools

import (
	"context"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetPodEvents(namespace, podName string) (string, error) {

	client, err := NewClient()
	if err != nil {
		return "", err
	}

	events, err := client.CoreV1().
		Events(namespace).
		List(context.Background(), metav1.ListOptions{
			FieldSelector: "involvedObject.kind=Pod,involvedObject.name=" + podName,
		})

	if err != nil {
		return "", err
	}

	if len(events.Items) == 0 {
		return "No events found for pod " + podName, nil
	}

	var result strings.Builder

	for _, event := range events.Items {
		result.WriteString("Reason: ")
		result.WriteString(event.Reason)
		result.WriteString("\n")

		result.WriteString("Message: ")
		result.WriteString(event.Message)
		result.WriteString("\n\n")
	}

	return result.String(), nil
}
