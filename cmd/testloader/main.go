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

	for _, doc := range docs {

		fmt.Printf("\n========== %s ==========\n\n", doc.FileName)
		fmt.Println(doc.Content)
	}
}
