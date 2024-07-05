package handlers

import (
	"libery_llm_chat_base_service/server"
	"net/http"
)

func AliveHandler(service_instance server.Server) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		response.WriteHeader(http.StatusOK)
		return
	}
}
