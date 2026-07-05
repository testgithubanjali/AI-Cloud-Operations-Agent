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

func (v *VectorStore) Load(chunks []Chunk) {

	v.Chunks = chunks
}

func (v *VectorStore) GetAll() []Chunk {

	return v.Chunks
}

func (v *VectorStore) Count() int {

	return len(v.Chunks)
}
