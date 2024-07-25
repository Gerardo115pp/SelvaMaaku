package qdrant_db

import (
	"encoding/json"
	"fmt"
	app_config "libery_llm_chat_base_service/Config"
	"libery_llm_chat_base_service/helpers"
	"libery_llm_chat_base_service/models"
	"os"
	"path"

	"github.com/Gerardo115pp/patriots_lib/echo"
)

type ChatKnowledgeRepository struct {
	qdrantAPI string
	Memories  map[string]models.Remembrance
}

func NewChatKnowledgeRepository(qdrant_api string) (*ChatKnowledgeRepository, error) {
	var new_knowledge_db *ChatKnowledgeRepository = new(ChatKnowledgeRepository)

	new_knowledge_db.qdrantAPI = qdrant_api
	new_knowledge_db.Memories = make(map[string]models.Remembrance)

	err := new_knowledge_db.loadMemories()
	if err != nil {
		return nil, err
	}

	// Verify the state of Qdrant's db

	if !new_knowledge_db.qdrantReachable() {
		return nil, fmt.Errorf(fmt.Sprintf("Qdrant API not reachable at <%s>", qdrant_api))
	}

	if exists, err := new_knowledge_db.collectionExists(app_config.KNOWLEDGE_QDRANT_COLLECTION); !exists {
		if err != nil {
			return nil, err
		}

		err = new_knowledge_db.createCollection(app_config.KNOWLEDGE_QDRANT_COLLECTION, app_config.EMBEDDING_3_DIMENSION)
		if err != nil {
			return nil, err
		}
	}

	return new_knowledge_db, nil
}

func (knowledge_db *ChatKnowledgeRepository) createRemembrancesFile() error {
	var ignorance []byte = []byte("{}")

	file_directory := path.Dir(app_config.REMEMBRANCES_STORAGE_FILE)
	if !helpers.FileExists(file_directory) || helpers.FileExists(app_config.REMEMBRANCES_STORAGE_FILE) {
		return nil
	}

	file, err := os.Create(app_config.REMEMBRANCES_STORAGE_FILE)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(ignorance)

	return err
}

func (knowledge_db *ChatKnowledgeRepository) loadMemories() error {
	if !helpers.FileExists(app_config.REMEMBRANCES_STORAGE_FILE) {
		fmt.Printf("Warning: Booting with ignorance, no remembrances file found at <%s>\n", app_config.REMEMBRANCES_STORAGE_FILE)
		return knowledge_db.createRemembrancesFile()
	}

	raw_memories, err := os.ReadFile(app_config.REMEMBRANCES_STORAGE_FILE)
	if err != nil {
		return err
	}

	var found_memories map[string]models.Remembrance

	err = json.Unmarshal(raw_memories, &found_memories)
	if err != nil {
		return err
	}

	knowledge_db.Memories = found_memories

	echo.EchoDebug(fmt.Sprintf("Loaded %d remembrances from file", len(knowledge_db.Memories)))

	return nil
}

func (knowledge_db *ChatKnowledgeRepository) Remember(remembrance models.Remembrance) error {
	knowledge_db.Memories[remembrance.ID] = remembrance

	err := knowledge_db.upsertRemembrance(remembrance) // Save to Qdrant
	if err != nil {
		return err
	}

	knowledge_db.saveMemories()

	return nil
}

func (knowledge_db *ChatKnowledgeRepository) Recall(memory_vector []float64) ([]models.Remembrance, error) {
	var related_memories []models.Remembrance = make([]models.Remembrance, 0)
	var similarity_scores []pointSimilarityScore = make([]pointSimilarityScore, 0)

	similarity_scores, err := knowledge_db.searchRelatedRemembrances(memory_vector, app_config.MAX_REMEMBRANCES_IN_CONTEXT)
	if err != nil {
		return nil, err
	}

	for _, score := range similarity_scores {
		related_remembrance, exists := knowledge_db.Memories[score.ID]

		if !exists {
			echo.EchoWarn(fmt.Sprintf("Recalled remembrance<%s> not found in memory", score.ID))
			continue
		}

		related_memories = append(related_memories, related_remembrance)
	}

	return related_memories, nil
}

func (knowledge_db *ChatKnowledgeRepository) saveMemories() error {
	memories_json, err := json.Marshal(knowledge_db.Memories)
	if err != nil {
		return err
	}

	err = os.WriteFile(app_config.REMEMBRANCES_STORAGE_FILE, memories_json, 0640) // -rw-r-----
	if err != nil {
		return err
	}

	return nil
}
