package middleware

import (
	"context"
	app_config "libery_llm_chat_base_service/Config"
	"libery_llm_chat_base_service/models"
	"net/http"
)

func CheckAuthCookie(next func(response http.ResponseWriter, request *http.Request)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		claims_token, err := request.Cookie(app_config.CHAT_CLAIM_COOKIE_NAME)

		if err != nil || claims_token.Value == "" {
			response.WriteHeader(http.StatusUnauthorized)
			return
		}

		client_ip := request.Header.Get("X-Forwarded-For")
		if client_ip == "" {
			client_ip = request.RemoteAddr
		}

		err, claims := models.VerifyChatClaims(claims_token.Value, client_ip)
		if err != nil {
			response.WriteHeader(http.StatusUnauthorized)
			return
		}

		authenticated_request := request.WithContext(context.WithValue(request.Context(), "claims", claims))

		next(response, authenticated_request)
	}
}
