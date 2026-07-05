package rag

import (
	"sort"

	"ai-sre-agent/internal/llm"
)

type SearchResult struct {
	Chunk Chunk
	Score float32
}

func Retrieve(
	store *VectorStore,
	query string,
	topK int,
) ([]SearchResult, error) {

	// Generate embedding for the user's question
	queryEmbedding, err := llm.GenerateEmbedding(query)
	if err != nil {
		return nil, err
	}

	var results []SearchResult

	// Compare with every chunk
	for _, chunk := range store.GetAll() {

		score := CosineSimilarity(
			queryEmbedding,
			chunk.Embedding,
		)

		results = append(results, SearchResult{
			Chunk: chunk,
			Score: score,
		})
	}

	// Sort from highest similarity to lowest
	sort.Slice(results, func(i, j int) bool {
		return results[i].Score > results[j].Score
	})

	// Return only Top K
	if topK > len(results) {
		topK = len(results)
	}

	return results[:topK], nil
}
