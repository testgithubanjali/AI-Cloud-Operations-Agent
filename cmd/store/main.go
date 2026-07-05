package main

import (
	"fmt"

	"ai-sre-agent/internal/rag"
)

func main() {

	store := rag.NewVectorStore()

	store.Add(rag.Chunk{
		ID:      1,
		Content: "Deployment Guide",
	})

	store.Add(rag.Chunk{
		ID:      2,
		Content: "Monitoring",
	})

	fmt.Println("Chunks:", store.Count())
}
