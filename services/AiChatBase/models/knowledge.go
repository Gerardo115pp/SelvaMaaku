package models

import (
	"fmt"

	"github.com/google/uuid"
)

type Remembrance struct {
	ID        string    `json:"id"`
	Document  string    `json:"document"`
	Embedding []float64 `json:"embedding"`
}

type RemembranceSimilarity struct {
	Remembrance Remembrance `json:"remembrance"`
	Similarity  float64     `json:"similarity"`
}

func CreateEmptyRemembrance(document string) (*Remembrance, error) {
	remembrance_uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("Failed to generate new remembrance UUID: %s", err.Error())
	}

	return &Remembrance{
		ID:       remembrance_uuid.String(),
		Document: document,
	}, nil
}
