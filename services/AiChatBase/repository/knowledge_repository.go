package repository

import "libery_llm_chat_base_service/models"

type KnowledgeRepository interface {
	Remember(remembrance models.Remembrance) error
	Recall(memory_vector []float64) ([]models.Remembrance, error)
}

var Knowledge KnowledgeRepository

func SetKnowledgeRepository(repository KnowledgeRepository) {
	Knowledge = repository
}