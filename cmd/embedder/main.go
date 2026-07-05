package main

import (
	"fmt"
	"log"

	"ai-sre-agent/internal/llm"
)

func main() {

	vector, err := llm.GenerateEmbedding(
		"Kubernetes Deployment"
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Embedding Dimension:", len(vector))

	fmt.Println("First 10 values:")

	for i := 0; i < 10 && i < len(vector); i++ {
		fmt.Printf("%.5f\n", vector[i])
	}
}