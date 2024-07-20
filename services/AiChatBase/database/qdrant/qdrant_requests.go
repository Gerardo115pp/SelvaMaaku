package qdrant_db

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	app_config "libery_llm_chat_base_service/Config"
	"libery_llm_chat_base_service/models"
	"net/http"
)

func (knowledge_db *ChatKnowledgeRepository) createCollection(collection_name string, vector_size int) error {
	create_collection_endpoint := fmt.Sprintf("%s/collections/%s", knowledge_db.qdrantAPI, collection_name)

	create_collection_request, err := createQdrantCreateCollectionRequest(vector_size)
	if err != nil {
		return err
	}

	request_body, err := json.Marshal(create_collection_request)
	if err != nil {
		return err
	}

	body := bytes.NewBuffer(request_body)
	response, err := http.NewRequest(http.MethodPut, create_collection_endpoint, body)
	if err != nil {
		return err
	}

	response.Header.Set("Content-Type", "application/json")
	// response.Header.Set("Authorization", fmt.Sprintf("Bearer %s", app_config.Q)) // The Auth token doesn't seem to be working correctly in qdrant. Requests go through without it.

	client := &http.Client{}
	resp, err := client.Do(response)
	if err != nil {
		return err
	}

	// this is just temporary, qdrant returns an object with a property 'result' that is a boolean. Check that also, cause 200 will be sent even if the collection already exists for example.
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to create collection: %s", resp.Status)
	}

	return nil
}

func (knowledge_db *ChatKnowledgeRepository) collectionExists(collection_name string) (bool, error) {
	list_collections_endpoint := fmt.Sprintf("%s/collections/%s/exists", knowledge_db.qdrantAPI, collection_name)

	response, err := http.NewRequest(http.MethodGet, list_collections_endpoint, nil)
	if err != nil {
		return false, err
	}

	response.Header.Set("Content-Type", "application/json")
	// response.Header.Set("Authorization", fmt.Sprintf("Bearer %s", app_config.Q)) // The Auth token doesn't seem to be working correctly in qdrant. Requests go through without it.

	client := &http.Client{}
	resp, err := client.Do(response)
	if err != nil {
		return false, err
	}

	var response_body []byte

	response_body, err = io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	if len(response_body) == 0 {
		return false, fmt.Errorf("Failed to read response body: %s\nStatus: %s", string(response_body), resp.Status)
	}

	var check_collection_response QdrantCheckCollectionResponse

	err = json.Unmarshal(response_body, &check_collection_response)
	if err != nil {
		return false, err
	}

	return check_collection_response.Result.Exists, nil
}

func (knowledge_db *ChatKnowledgeRepository) qdrantReachable() bool {
	list_collections_endpoint := fmt.Sprintf("%s/healthz", knowledge_db.qdrantAPI)

	response, err := http.NewRequest(http.MethodGet, list_collections_endpoint, nil)
	if err != nil {
		return false
	}

	client := &http.Client{}
	resp, err := client.Do(response)
	if err != nil {
		return false
	}

	return resp.StatusCode == http.StatusOK
}

func (knowledge_db *ChatKnowledgeRepository) upsertRemembrance(remembrance models.Remembrance) error {
	upsert_endpoint := fmt.Sprintf("%s/collections/%s/points", knowledge_db.qdrantAPI, app_config.KNOWLEDGE_QDRANT_COLLECTION)

	create_points_request, err := createPointsRequestFromRemembrance(remembrance)
	if err != nil {
		return err
	}

	request_body, err := json.Marshal(create_points_request)
	if err != nil {
		return err
	}

	body := bytes.NewBuffer(request_body)
	response, err := http.NewRequest(http.MethodPut, upsert_endpoint, body)
	if err != nil {
		return err
	}

	response.Header.Set("Content-Type", "application/json")
	// response.Header.Set("Authorization", fmt.Sprintf("Bearer %s", app_config.Q)) // The Auth token doesn't seem to be working correctly in qdrant. Requests go through without it.

	client := &http.Client{}
	resp, err := client.Do(response)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to upsert remembrance: %s", resp.Status)
	}

	return nil
}

// Given some text encoded as a vector, search for remembrances related to that text and return the remembrance's ids
func (knowledge_db *ChatKnowledgeRepository) searchRelatedRemembrances(subject []float64, limit int) ([]pointSimilarityScore, error) {
	search_endpoint := fmt.Sprintf("%s/collections/%s/points/search", knowledge_db.qdrantAPI, app_config.KNOWLEDGE_QDRANT_COLLECTION)

	points_search_request := createQdrantSearchPointsRequest(subject, limit)

	request_body, err := json.Marshal(points_search_request)
	if err != nil {
		return nil, err
	}

	body := bytes.NewBuffer(request_body)

	request, err := http.NewRequest(http.MethodPost, search_endpoint, body)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	// response.Header.Set("Authorization", fmt.Sprintf("Bearer %s", app_config.Q)) // The Auth token doesn't seem to be working correctly in qdrant. Requests go through without it.

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	var search_response *QdrantSearchPointsResponse

	search_response, err = parseQdrantSearchPointsResponse(response.Body)
	if err != nil {
		return nil, err
	}

	related_memories := make([]pointSimilarityScore, 0)

	for _, result := range search_response.Result {
		related_memories = append(related_memories, pointSimilarityScore{
			ID:    result.ID,
			Score: result.Score,
		})
	}

	return related_memories, nil
}
