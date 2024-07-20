package qdrant_db

import (
	"encoding/json"
	"fmt"
	"io"
	app_config "libery_llm_chat_base_service/Config"
	"libery_llm_chat_base_service/models"
)

// <========== create collection ==========>
// Tip: The collection name is a 'path parameter'. e.g PUT /collections/{collection_name}
// and yes, it's PUT not POST. it horrible but that's how they did it.

type QdrantCreateCollectionRequestVectorsParam struct {
	Size     int    `json:"size"`     // dimension for a dense vector
	Distance string `json:"distance"` // the method to calculate the distance between vectors. for cosine similarity, which is the best approach for openai models, use "Cosine"
}

// API method docs: https://api.qdrant.tech/api-reference/collections/create-collection
// Collections docs: https://qdrant.tech/documentation/concepts/collections/
type QdrantCreateCollectionRequest struct {
	Vectors QdrantCreateCollectionRequestVectorsParam `json:"vectors"`
}

func createQdrantCreateCollectionRequest(vector_size int) (*QdrantCreateCollectionRequest, error) {
	if app_config.VECTOR_DISTANCE != "Cosine" && app_config.VECTOR_DISTANCE != "Dot" {
		return nil, fmt.Errorf("Invalid or unsupported vector distance for Qdrant: '%s'. Expected 'Cosine' or 'Dot'", app_config.VECTOR_DISTANCE)
	}

	var vectors_param QdrantCreateCollectionRequestVectorsParam = QdrantCreateCollectionRequestVectorsParam{
		Size:     vector_size,
		Distance: app_config.VECTOR_DISTANCE,
	}

	return &QdrantCreateCollectionRequest{
		Vectors: vectors_param,
	}, nil
}

// Probably good enough to check that Result is true.
type QdrantCreateCollectionResponse struct {
	Result bool    `json:"result"`
	Status string  `json:"status"` // "ok" if successful
	Time   float64 `json:"time"`
}

// <========== List collections ==========>
// A type for list request is not needed as there are no parameters for this request.

type QdrantListCollectionsResponseCollectionItem struct {
	Name string `json:"name"`
}

type QdrantListCollectionsResponseResultParam struct {
	Collections []QdrantListCollectionsResponseCollectionItem `json:"collections"`
}

type QdrantListCollectionsResponse struct {
	Result QdrantListCollectionsResponseResultParam `json:"result"`
	Status string                                   `json:"status"` // "ok" if successful
	Time   float64                                  `json:"time"`
}

// <========== Check collection existence ==========>
// A type for check request is not needed as there are no parameters for this request.

type QdrantCheckCollectionResponseResultParam struct {
	Exists bool `json:"exists"`
}

type QdrantCheckCollectionResponse struct {
	Result QdrantCheckCollectionResponseResultParam `json:"result"`
	Status string                                   `json:"status"` // "ok" regardless of whether it exists or not. Keep that in mind.
	Time   float64                                  `json:"time"`
}

// <========== Create points ==========>
// Tip: The collection name is a 'path parameter'. e.g PUT /collections/{collection_name}/points
type QdrantCreatePointsRequest struct {
	Points []QdrantCreatePointsRequestPointParam `json:"points"`
}

type QdrantCreatePointsRequestPointParam struct {
	Id     string    `json:"id"`     // unique identifier for the point
	Vector []float64 `json:"vector"` // the vector for the point
}

func createPointsRequestFromRemembrance(remembrance models.Remembrance) (*QdrantCreatePointsRequest, error) {
	var points_param []QdrantCreatePointsRequestPointParam = make([]QdrantCreatePointsRequestPointParam, 1)

	points_param[0] = QdrantCreatePointsRequestPointParam{
		Id:     remembrance.ID,
		Vector: remembrance.Embedding,
	}

	return &QdrantCreatePointsRequest{
		Points: points_param,
	}, nil
}

//<========== Search points ==========>
// Tip: The collection name is a 'path parameter'. e.g POST /collections/{collection_name}/points/search

type QdrantSearchPointsRequest struct {
	Vector []float64 `json:"vector"` // the vector to search for
	Limit  int       `json:"limit"`  // the maximum number of results to return
}

type QdrantSearchPointsResponse struct {
	Result []QdrantSearchPointsResponseResultParam `json:"result"`
	Status string                                  `json:"status"` // "ok" if successful
	Time   float64                                 `json:"time"`
}

type QdrantSearchPointsResponseResultParam struct {
	ID      string  `json:"id"`      // the id of the point
	Version int     `json:"version"` // the version of the point
	Score   float64 `json:"score"`
}

type pointSimilarityScore struct {
	ID    string  `json:"id"`
	Score float64 `json:"score"`
}

func createQdrantSearchPointsRequest(vector []float64, limit int) *QdrantSearchPointsRequest {
	return &QdrantSearchPointsRequest{
		Vector: vector,
		Limit:  limit,
	}
}

func parseQdrantSearchPointsResponse(response io.Reader) (*QdrantSearchPointsResponse, error) {
	var response_body []byte
	var err error

	response_body, err = io.ReadAll(response)
	if err != nil {
		return nil, err
	}

	if len(response_body) == 0 {
		return nil, fmt.Errorf("Failed to read response body: %s\nStatus: %s", string(response_body), response_body)
	}

	var search_response QdrantSearchPointsResponse

	err = json.Unmarshal(response_body, &search_response)
	if err != nil {
		return nil, err
	}

	return &search_response, nil
}
