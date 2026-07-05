package rag

type SearchResult struct {
	Chunk Chunk
	Score float32
}

func Retrieve(
	store *VectorStore,
	query string,
	topK int,
) ([]SearchResult, error) {

	// Generate embedding of question

	// Compare with every chunk

	// Sort

	// Return top K

	return nil, nil
}
