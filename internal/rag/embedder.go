package rag

import "ai-sre-agent/internal/llm"

func EmbedChunks(chunks []Chunk) ([]Chunk, error) {

	for i := range chunks {

		embedding, err := llm.GenerateEmbedding(chunks[i].Content)
		if err != nil {
			return nil, err
		}

		chunks[i].Embedding = embedding
	}

	return chunks, nil
}
