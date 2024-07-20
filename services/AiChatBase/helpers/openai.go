package helpers

import (
	"fmt"
	"math"
	"regexp"
)

func TokenCount(text string) float64 {
	characters := float64(len(text))
	return characters * 0.75
}

func ComputeCosineSimilarity(vector_a []float64, vector_b []float64) (float64, error) {
	if len(vector_a) != len(vector_b) {
		return 0, fmt.Errorf("Unmaching vectors shape")
	}

	var numerator float64 = 0
	var square_a float64 = 0
	var square_b float64 = 0

	for h := 0; h < len(vector_a); h++ {
		numerator += vector_a[h] * vector_b[h]
		square_a += math.Pow(vector_a[h], 2)
		square_b += math.Pow(vector_b[h], 2)
	}

	return numerator / (math.Sqrt(square_a) * math.Sqrt(square_b)), nil
}

func MessageLooksLikeContactInfo(message_content string) bool {
	var email_verifier *regexp.Regexp = regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	var phone_verifier *regexp.Regexp = regexp.MustCompile(`(\+\d{1,3}[\s-]?)?\d{2,5}[\s-]?\d{4}[\s-]?\d{4}`)

	return email_verifier.FindString(message_content) != "" || phone_verifier.FindString(message_content) != ""
}
