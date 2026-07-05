package rag

import "strings"

type Chunk struct {
	ID        int
	FileName  string
	Content   string
	Embedding []float32
}

func ChunkDocuments(documents []Document, chunkSize int) []Chunk {

	var chunks []Chunk
	id := 1

	for _, doc := range documents {

		paragraphs := strings.Split(doc.Content, "\n\n")

		var currentChunk strings.Builder

		for _, paragraph := range paragraphs {

			// Skip empty paragraphs
			if strings.TrimSpace(paragraph) == "" {
				continue
			}

			// Save the current chunk only if it already has content
			// and adding the next paragraph would exceed the limit.
			if currentChunk.Len() > 0 &&
				currentChunk.Len()+len(paragraph)+2 > chunkSize {

				chunks = append(chunks, Chunk{
					ID:       id,
					FileName: doc.FileName,
					Content:  strings.TrimSpace(currentChunk.String()),
				})

				id++
				currentChunk.Reset()
			}

			currentChunk.WriteString(paragraph)
			currentChunk.WriteString("\n\n")
		}

		// Save the last chunk
		if currentChunk.Len() > 0 {

			chunks = append(chunks, Chunk{
				ID:       id,
				FileName: doc.FileName,
				Content:  strings.TrimSpace(currentChunk.String()),
			})

			id++
		}
	}

	return chunks
}
