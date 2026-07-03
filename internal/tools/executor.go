package tools

func ExecuteTool(call ToolCall) string {

	switch call.Tool {

	case "get_pods":
		return "Mock Result:\npostgres-0\npostgres-1"

	case "describe_pod":
		return "Mock Result:\nPod is Running"

	case "get_logs":
		return "Mock Result:\nApplication started successfully."

	default:
		return "Unknown tool requested."
	}
}
