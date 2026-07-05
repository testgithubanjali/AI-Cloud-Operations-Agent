package tools

import "fmt"

func ExecuteTool(call ToolCall) (string, error) {

	switch call.Tool {

	case "get_pods":

		pods, err := GetPods("default")
		if err != nil {
			return "", err
		}

		result := ""

		for _, pod := range pods {
			result += pod + "\n"
		}

		return result, nil

	case "describe_pod":

		result, err := DescribePod("default", call.Pod)
		if err != nil {
			return "", err
		}

		return result, nil

	case "get_logs":

		// First try current container logs
		result, err := GetLogs(
			"default",
			call.Pod,
			100,
			false,
		)

		// If that fails, try previous container logs
		if err != nil {

			result, err = GetLogs(
				"default",
				call.Pod,
				100,
				true,
			)

			if err != nil {
				return "", err
			}
		}

		return result, nil

	case "get_metrics":

		result, err := GetPodMetrics("default")
		if err != nil {
			return "", err
		}

		return result, nil
	case "query_prometheus":

		result, err := QueryPrometheus(call.Query)
		if err != nil {
			return "", err
		}

		return result, nil

	default:
		return "", fmt.Errorf("unknown tool: %s", call.Tool)
	}
}
