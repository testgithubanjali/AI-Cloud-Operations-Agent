package main

import (
	"fmt"

	"ai-sre-agent/internal/rag"
)

func main() {

	a := []float32{1, 2, 3}
	b := []float32{1, 2, 3}

	score := rag.CosineSimilarity(a, b)

	fmt.Println(score)
}
