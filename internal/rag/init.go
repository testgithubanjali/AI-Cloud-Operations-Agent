package rag

func InitializeKnowledgeBase() error {

	// Load documents
	docs, err := LoadDocuments("docs")
	if err != nil {
		return err
	}

	// Chunk documents
	chunks := ChunkDocuments(docs, 300)

	// Generate embeddings
	chunks, err = EmbedChunks(chunks)
	if err != nil {
		return err
	}

	// Store everything
	GlobalVectorStore = NewVectorStore()
	GlobalVectorStore.Load(chunks)

	return nil
}
