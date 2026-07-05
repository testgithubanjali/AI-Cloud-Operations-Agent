package main

import (
	"fmt"
	"log"

	"ai-sre-agent/internal/rag"
)

func main() {

	docs, err := rag.LoadDocuments("docs")
	if err != nil {
		log.Fatal(err)
	}

	chunks := rag.ChunkDocuments(docs, 300)

	chunks, err = rag.EmbedChunks(chunks)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Embedded Chunks:", len(chunks))
}
