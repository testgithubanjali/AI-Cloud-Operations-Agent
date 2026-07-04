package tools

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetPods(namespace string) ([]string, error) {

	client, err := NewClient()
	if err != nil {
		return nil, err
	}

	podList, err := client.CoreV1().
		Pods(namespace).
		List(context.Background(), metav1.ListOptions{})

	if err != nil {
		return nil, err
	}

	var pods []string

	for _, pod := range podList.Items {
		pods = append(pods, pod.Name)
	}

	return pods, nil
}
