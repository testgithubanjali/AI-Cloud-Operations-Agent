package rag

import (
	"os"
	"path/filepath"
)

type Document struct {
	FileName string
	Content  string
}

func LoadDocuments(path string) ([]Document, error) {

	var docs []Document

	files, err := filepath.Glob(filepath.Join(path, "*.md"))
	if err != nil {
		return nil, err
	}

	for _, file := range files {

		content, err := os.ReadFile(file)
		if err != nil {
			return nil, err
		}

		doc := Document{
			FileName: filepath.Base(file),
			Content:  string(content),
		}

		docs = append(docs, doc)
	}

	return docs, nil
}
