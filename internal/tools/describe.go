package tools

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func DescribePod(namespace string, podName string) (string, error) {

	client, err := NewClient()
	if err != nil {
		return "", err
	}

	pod, err := client.CoreV1().
		Pods(namespace).
		Get(context.Background(), podName, metav1.GetOptions{})

	if err != nil {
		return "", err
	}

	result := fmt.Sprintf(
		"Name: %s\nNamespace: %s\nStatus: %s\nNode: %s\nPod IP: %s\nRestart Count: %d\nImage: %s",
		pod.Name,
		pod.Namespace,
		pod.Status.Phase,
		pod.Spec.NodeName,
		pod.Status.PodIP,
		pod.Status.ContainerStatuses[0].RestartCount,
		pod.Spec.Containers[0].Image,
	)

	return result, nil
}
