package util

import "math"

func CosineSimilarity(v1, v2 []float64) float64 {
	if len(v1) != len(v2) {
		return 0
	}
	var dotProduct float64
	var magnitudeV1 float64
	var magnitudeV2 float64
	for i := 0; i < len(v1); i++ {
		dotProduct += v1[i] * v2[i]  // Dot product of v1 and v2
		magnitudeV1 += v1[i] * v1[i] // Sum of v1^2 or math.pow(v1[i], 2)
		magnitudeV2 += v2[i] * v2[i] // Sum of v2^2 or math.pow(v2[i], 2)
	}
	magnitudeV1 = math.Sqrt(magnitudeV1) // Square root of sum of v1^2
	magnitudeV2 = math.Sqrt(magnitudeV2) // Square root of sum of v2^2
	if magnitudeV1 == 0 || magnitudeV2 == 0 {
		return 0
	}
	return dotProduct / (magnitudeV1 * magnitudeV2)
}
