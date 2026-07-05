package rag

import "fmt"

func GetRelevantContext(question string) (string, error) {

	// Load documents
	docs, err := LoadDocuments("docs")
	if err != nil {
		return "", err
	}

	// Split into chunks
	chunks := ChunkDocuments(docs, 300)

	// Generate embeddings
	chunks, err = EmbedChunks(chunks)
	if err != nil {
		return "", err
	}

	// Store chunks
	store := NewVectorStore()
	store.Load(chunks)

	// Retrieve relevant chunks
	results, err := Retrieve(store, question, 3)
	if err != nil {
		return "", err
	}

	// Convert to text
	context := BuildContext(results)

	fmt.Println("RAG Context Loaded")

	return context, nil
}
