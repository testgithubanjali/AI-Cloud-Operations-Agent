package rag

import "fmt"

func GetRelevantContext(question string) (string, error) {

	if GlobalVectorStore == nil {
		return "", fmt.Errorf("knowledge base not initialized")
	}

	results, err := Retrieve(
		GlobalVectorStore,
		question,
		3,
	)
	if err != nil {
		return "", err
	}

	return BuildContext(results), nil
}
