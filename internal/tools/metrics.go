package tools

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetPodMetrics(namespace string) (string, error) {

	client, err := NewMetricsClient()
	if err != nil {
		return "", err
	}

	metrics, err := client.MetricsV1beta1().
		PodMetricses(namespace).
		List(context.Background(), metav1.ListOptions{})

	if err != nil {
		return "", err
	}

	var result string

	for _, pod := range metrics.Items {

		result += fmt.Sprintf(
			"Pod: %s\nCPU: %s\nMemory: %s\n\n",
			pod.Name,
			pod.Containers[0].Usage.Cpu().String(),
			pod.Containers[0].Usage.Memory().String(),
		)
	}

	return result, nil
}
