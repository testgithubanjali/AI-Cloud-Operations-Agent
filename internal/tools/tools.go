package tools

type Tool struct {
	Name        string
	Description string
}

var AvailableTools = []Tool{
	{
		Name:        "get_pods",
		Description: "Get all pods in a namespace",
	},
	{
		Name:        "describe_pod",
		Description: "Describe a Kubernetes pod",
	},
	{
		Name:        "get_logs",
		Description: "Return logs of a pod",
	},
}
