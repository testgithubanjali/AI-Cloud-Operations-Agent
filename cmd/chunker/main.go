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

	fmt.Println("Total Chunks:", len(chunks))

	for _, chunk := range chunks {

		fmt.Println("===================================")
		fmt.Println("Chunk ID:", chunk.ID)
		fmt.Println("File:", chunk.FileName)
		fmt.Println("-----------------------------------")
		fmt.Println(chunk.Content)
		fmt.Println()
	}
}
