package tools

import (
	"context"
	"fmt"
	"io"

	corev1 "k8s.io/api/core/v1"
)

func GetLogs(namespace, podName string, tail int, previous bool) (string, error) {

	client, err := NewClient()
	if err != nil {
		return "", err
	}

	tailLines := int64(tail)

	req := client.CoreV1().
		Pods(namespace).
		GetLogs(
			podName,
			&corev1.PodLogOptions{
				TailLines: &tailLines,
				Previous:  previous,
			},
		)

	stream, err := req.Stream(context.Background())
	if err != nil {
		return "", err
	}
	defer stream.Close()

	logs, err := io.ReadAll(stream)
	if err != nil {
		return "", err
	}

	var result string

	if previous {

		result += "========== Previous Container Logs ==========\n\n"

	} else {

		result += "========== Current Container Logs ==========\n\n"
	}

	if len(logs) == 0 {

		result += "No logs available."

	} else {

		result += fmt.Sprintf(
			"Showing last %d log lines\n\n%s",
			tail,
			string(logs),
		)
	}

	return result, nil
}
