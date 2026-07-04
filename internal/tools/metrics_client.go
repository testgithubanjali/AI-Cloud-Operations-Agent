package tools

import (
	"os"
	"path/filepath"

	"k8s.io/client-go/tools/clientcmd"
	metricsclient "k8s.io/metrics/pkg/client/clientset/versioned"
)

func NewMetricsClient() (*metricsclient.Clientset, error) {

	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	kubeconfig := filepath.Join(home, ".kube", "config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}

	client, err := metricsclient.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return client, nil
}
