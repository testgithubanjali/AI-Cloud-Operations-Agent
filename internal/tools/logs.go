package tools

import (
	"context"
	"io"

	corev1 "k8s.io/api/core/v1"
)

func GetLogs(namespace, podName string) (string, error) {

	client, err := NewClient()
	if err != nil {
		return "", err
	}

	req := client.CoreV1().
		Pods(namespace).
		GetLogs(podName, &corev1.PodLogOptions{})

	stream, err := req.Stream(context.Background())
	if err != nil {
		return "", err
	}
	defer stream.Close()

	logs, err := io.ReadAll(stream)
	if err != nil {
		return "", err
	}

	return string(logs), nil
}
