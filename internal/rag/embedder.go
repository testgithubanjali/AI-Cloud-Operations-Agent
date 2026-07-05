package rag

import "ai-sre-agent/internal/llm"

func EmbedChunks(chunks []Chunk) ([]Chunk, error) {

	for i := range chunks {

		vector, err := llm.GenerateEmbedding(
			chunks[i].Content,
		)

		if err != nil {
			return nil, err
		}

		chunks[i].Embedding = vector
	}

	return chunks, nil
}
