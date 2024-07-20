package handlers

import (
	"encoding/json"
	"fmt"
	"libery_llm_chat_base_service/helpers"
	"libery_llm_chat_base_service/models"
	"libery_llm_chat_base_service/repository"
	"libery_llm_chat_base_service/server"
	"libery_llm_chat_base_service/workflows"
	"net/http"
)

func KnowledgeHandler(service_instance server.Server) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			getKnowledgeHandler(response, request)
		case http.MethodPost:
			postKnowledgeHandler(response, request)
		case http.MethodPatch:
			patchKnowledgeHandler(response, request)
		case http.MethodDelete:
			deleteKnowledgeHandler(response, request)
		case http.MethodPut:
			putKnowledgeHandler(response, request)
		case http.MethodOptions:
			response.WriteHeader(http.StatusOK)
		default:
			response.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func getKnowledgeHandler(response http.ResponseWriter, request *http.Request) {
	knowledge_request := &struct {
		Subject string `json:"subject"`
	}{}

	err := json.NewDecoder(request.Body).Decode(knowledge_request)
	if err != nil {
		helpers.WriteRejection(response, 400, "Invalid JSON format. Expected single field 'subject' of type string")
		return
	}

	encoded_subject, err := workflows.EmbedTextSemantics(knowledge_request.Subject)
	if err != nil {
		helpers.WriteRejection(response, 500, "Failed to encode knowledge text")
		return
	}

	remembrances, err := repository.Knowledge.Recall(encoded_subject)
	if err != nil {
		helpers.WriteRejection(response, 500, fmt.Sprintf("Failed to recall knowledge: %s", err.Error()))
		return
	}

	subject_knowledge := &struct {
		Remembrances []string `json:"remembrances"`
	}{
		Remembrances: make([]string, len(remembrances)),
	}

	for i, remembrance := range remembrances {
		subject_knowledge.Remembrances[i] = remembrance.Document
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(200)

	json.NewEncoder(response).Encode(subject_knowledge)
	return
}

func postKnowledgeHandler(response http.ResponseWriter, request *http.Request) {
	new_knowledge_data := &struct {
		Knowledge string `json:"knowledge"`
	}{}

	err := json.NewDecoder(request.Body).Decode(new_knowledge_data)
	if err != nil {
		helpers.WriteRejection(response, 400, "Invalid JSON format. Expected single field 'knowledge' of type string")
		return
	} else if new_knowledge_data.Knowledge == "" {
		helpers.WriteRejection(response, 422, "Knowledge field is required and cannot be an empty string")
		return
	}

	new_remembrance, err := models.CreateEmptyRemembrance(new_knowledge_data.Knowledge)
	if err != nil {
		helpers.WriteRejection(response, 500, fmt.Sprintf("Failed to create new remembrance: %s", err.Error()))
		return
	}

	encoded_knowledge, err := workflows.EmbedTextSemantics(new_knowledge_data.Knowledge)
	if err != nil {
		helpers.WriteRejection(response, 500, "Failed to encode knowledge text")
		return
	}

	new_remembrance.Embedding = encoded_knowledge

	err = repository.Knowledge.Remember(*new_remembrance)
	if err != nil {
		helpers.WriteRejection(response, 500, fmt.Sprintf("Failed to acquire this new knowledge: %s", err.Error()))
		return
	}

	helpers.WriteBooleanResponse(response, true)
	return
}

func patchKnowledgeHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}
func deleteKnowledgeHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}
func putKnowledgeHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}
