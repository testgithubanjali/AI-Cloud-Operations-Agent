package rag

import "strings"

func BuildContext(results []SearchResult) string {

	var context strings.Builder

	for _, result := range results {

		context.WriteString("Source: ")
		context.WriteString(result.Chunk.FileName)
		context.WriteString("\n\n")

		context.WriteString(result.Chunk.Content)
		context.WriteString("\n\n----------------------\n\n")
	}

	return context.String()
}
