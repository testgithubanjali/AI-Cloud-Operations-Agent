package rag

type Chunk struct {
	ID       int
	FileName string
	Content  string
}

func ChunkDocuments(
	documents []Document,
	chunkSize int,
) []Chunk {

	var chunks []Chunk

	id := 1

	for _, doc := range documents {

		content := doc.Content

		for len(content) > 0 {

			if len(content) <= chunkSize {

				chunks = append(chunks, Chunk{
					ID:       id,
					FileName: doc.FileName,
					Content:  content,
				})

				id++

				break
			}

			chunks = append(chunks, Chunk{
				ID:       id,
				FileName: doc.FileName,
				Content:  content[:chunkSize],
			})

			id++

			content = content[chunkSize:]
		}
	}

	return chunks
}
