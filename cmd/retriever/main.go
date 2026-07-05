package main

import (
	"fmt"
	"log"

	"ai-sre-agent/internal/rag"
)

func main() {

	// Load documents
	docs, err := rag.LoadDocuments("docs")
	if err != nil {
		log.Fatal(err)
	}

	// Chunk them
	chunks := rag.ChunkDocuments(docs, 300)

	// Generate embeddings
	chunks, err = rag.EmbedChunks(chunks)
	if err != nil {
		log.Fatal(err)
	}

	// Create vector store
	store := rag.NewVectorStore()
	store.Load(chunks)

	// Ask a question
	results, err := rag.Retrieve(
		store,
		"How do I rollback a deployment?",
		3,
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Top Results\n")

	for _, result := range results {

		fmt.Println("==========================")
		fmt.Printf("Score : %.4f\n", result.Score)
		fmt.Println("File  :", result.Chunk.FileName)
		fmt.Println("--------------------------")
		fmt.Println(result.Chunk.Content)
		fmt.Println()
	}
}
