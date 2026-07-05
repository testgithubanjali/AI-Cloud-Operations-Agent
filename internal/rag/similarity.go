package rag

import "math"

func CosineSimilarity(a, b []float32) float32 {

	if len(a) != len(b) {
		return 0
	}

	var dotProduct float64
	var magnitudeA float64
	var magnitudeB float64

	for i := range a {

		dotProduct += float64(a[i] * b[i])

		magnitudeA += float64(a[i] * a[i])

		magnitudeB += float64(b[i] * b[i])
	}

	if magnitudeA == 0 || magnitudeB == 0 {
		return 0
	}

	return float32(
		dotProduct /
			(math.Sqrt(magnitudeA) * math.Sqrt(magnitudeB)),
	)
}
