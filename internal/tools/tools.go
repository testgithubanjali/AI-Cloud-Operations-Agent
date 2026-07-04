package tools

type Tool struct {
	Name        string
	Description string
}
type ToolCall struct {
	Tool string `json:"tool"`

	Namespace string `json:"namespace,omitempty"`

	Pod string `json:"pod,omitempty"`
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
	{
		Name:        "get_metrics",
		Description: "Get CPU and memory usage of Kubernetes pods",
	},
}
