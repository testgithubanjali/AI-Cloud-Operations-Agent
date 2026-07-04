package tools

func ExecuteTool(call ToolCall) string {

	switch call.Tool {

	case "get_pods":

		pods, err := GetPods("default")
		if err != nil {
			return err.Error()
		}

		result := ""

		for _, pod := range pods {
			result += pod + "\n"
		}

		return result

	case "describe_pod":

		result, err := DescribePod("default", call.Pod)
		if err != nil {
			return err.Error()
		}

		return result

	case "get_logs":

		result, err := GetLogs("default", call.Pod)
		if err != nil {
			return err.Error()
		}

		return result

	case "get_metrics":

		result, err := GetPodMetrics("default")
		if err != nil {
			return err.Error()
		}

		return result

	default:
		return "Unknown tool"
	}
}
