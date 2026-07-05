package rag

type VectorStore struct {
	Chunks []Chunk
}

func NewVectorStore() *VectorStore {

	return &VectorStore{
		Chunks: []Chunk{},
	}
}
func (v *VectorStore) Add(chunk Chunk) {

	v.Chunks = append(v.Chunks, chunk)
}
func (v *VectorStore) Count() int {

	return len(v.Chunks)
}
