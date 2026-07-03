package tools

func ExecuteTool(call ToolCall) string {

	switch call.Tool {

	case "get_pods":
		return `postgres-0
postgres-1
redis-0`

	case "describe_pod":
		return `Name: postgres-0
Status: Running
Ready: True
Restarts: 0`

	case "get_logs":
		return `Application started successfully`

	default:
		return "Unknown tool"
	}
}
