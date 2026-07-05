package tools

import (
	"context"
	"fmt"
	"strings"

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

	var result strings.Builder

	result.WriteString("========== Pod Information ==========\n")

	result.WriteString(fmt.Sprintf("Name: %s\n", pod.Name))
	result.WriteString(fmt.Sprintf("Namespace: %s\n", pod.Namespace))
	result.WriteString(fmt.Sprintf("Status: %s\n", pod.Status.Phase))
	result.WriteString(fmt.Sprintf("Node: %s\n", pod.Spec.NodeName))
	result.WriteString(fmt.Sprintf("Pod IP: %s\n", pod.Status.PodIP))

	if len(pod.Status.ContainerStatuses) > 0 {

		container := pod.Status.ContainerStatuses[0]

		result.WriteString(fmt.Sprintf("Container: %s\n", container.Name))
		result.WriteString(fmt.Sprintf("Image: %s\n", pod.Spec.Containers[0].Image))
		result.WriteString(fmt.Sprintf("Ready: %t\n", container.Ready))
		result.WriteString(fmt.Sprintf("Restart Count: %d\n", container.RestartCount))

		if container.State.Running != nil {
			result.WriteString("Current State: Running\n")
		}

		if container.State.Waiting != nil {
			result.WriteString(fmt.Sprintf(
				"Current State: Waiting (%s)\n",
				container.State.Waiting.Reason,
			))
		}

		if container.State.Terminated != nil {

			result.WriteString("Current State: Terminated\n")
			result.WriteString(fmt.Sprintf(
				"Reason: %s\n",
				container.State.Terminated.Reason,
			))

			result.WriteString(fmt.Sprintf(
				"Exit Code: %d\n",
				container.State.Terminated.ExitCode,
			))
		}

		if container.LastTerminationState.Terminated != nil {

			last := container.LastTerminationState.Terminated

			result.WriteString("\n========== Last Termination ==========\n")

			result.WriteString(fmt.Sprintf(
				"Reason: %s\n",
				last.Reason,
			))

			result.WriteString(fmt.Sprintf(
				"Exit Code: %d\n",
				last.ExitCode,
			))

			result.WriteString(fmt.Sprintf(
				"Finished At: %s\n",
				last.FinishedAt,
			))
		}
	}

	result.WriteString("\n========== Conditions ==========\n")

	for _, condition := range pod.Status.Conditions {

		result.WriteString(fmt.Sprintf(
			"%s : %s\n",
			condition.Type,
			condition.Status,
		))
	}

	events, err := client.CoreV1().
		Events(namespace).
		List(context.Background(), metav1.ListOptions{
			FieldSelector: "involvedObject.name=" + podName,
		})

	if err == nil {

		result.WriteString("\n========== Recent Events ==========\n")

		if len(events.Items) == 0 {

			result.WriteString("No recent events\n")

		} else {

			for _, event := range events.Items {

				result.WriteString(fmt.Sprintf(
					"[%s] %s : %s\n",
					event.Type,
					event.Reason,
					event.Message,
				))
			}
		}
	}

	return result.String(), nil
}
