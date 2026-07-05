package tools

import (
	"context"
	"fmt"
	"strings"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetPodMetrics(namespace string) (string, error) {

	client, err := NewMetricsClient()
	if err != nil {
		return "", err
	}

	metrics, err := client.
		MetricsV1beta1().
		PodMetricses(namespace).
		List(context.Background(), metav1.ListOptions{})

	if err != nil {
		return "", err
	}

	if len(metrics.Items) == 0 {
		return "No pod metrics available.", nil
	}

	var result strings.Builder

	result.WriteString("========== Pod Metrics ==========\n\n")

	result.WriteString(fmt.Sprintf(
		"Namespace: %s\n",
		namespace,
	))

	result.WriteString(fmt.Sprintf(
		"Collected At: %s\n\n",
		time.Now().Format(time.RFC1123),
	))

	for _, pod := range metrics.Items {

		result.WriteString("---------------------------------\n")

		result.WriteString(fmt.Sprintf(
			"Pod Name: %s\n",
			pod.Name,
		))

		for _, container := range pod.Containers {

			result.WriteString(fmt.Sprintf(
				"Container: %s\n",
				container.Name,
			))

			result.WriteString(fmt.Sprintf(
				"CPU Usage: %s\n",
				container.Usage.Cpu().String(),
			))

			result.WriteString(fmt.Sprintf(
				"Memory Usage: %s\n",
				container.Usage.Memory().String(),
			))

			result.WriteString("\n")
		}
	}

	return result.String(), nil
}
