package handlers

import (
	"fmt"
	app_config "libery_llm_chat_base_service/Config"
	"libery_llm_chat_base_service/models"
	"libery_llm_chat_base_service/server"
	"net/http"
	"time"

	"github.com/Gerardo115pp/patriots_lib/echo"
)

func ChatClaimsHandler(portfolio_server server.Server) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		var current_origin string = request.Host

		echo.Echo(echo.GreenFG, "Request from origin: "+current_origin)

		if current_origin != app_config.PRODUCTION_DOMAIN && current_origin != app_config.DEVELOPMENT_DOMAIN {
			response.WriteHeader(http.StatusForbidden)
			return
		}

		switch request.Method {
		case http.MethodGet:
			getChatClaimsHandler(response, request)
		case http.MethodPost:
			postChatClaimsHandler(response, request)
		case http.MethodPatch:
			patchChatClaimsHandler(response, request)
		case http.MethodDelete:
			deleteChatClaimsHandler(response, request)
		case http.MethodPut:
			putChatClaimsHandler(response, request)
		case http.MethodOptions:
			response.WriteHeader(http.StatusOK)
		default:
			response.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func getChatClaimsHandler(response http.ResponseWriter, request *http.Request) {
	var request_route string = request.URL.Path

	if request_route == "/chat-claims/verify" {
		verifyChatClaims(response, request)
	} else {
		createChatClaims(response, request)
	}

	return
}

func verifyChatClaims(response http.ResponseWriter, request *http.Request) {
	claims_token, err := request.Cookie(app_config.CHAT_CLAIM_COOKIE_NAME)

	if err != nil || claims_token.Value == "" {
		echo.Echo(echo.RedBG, fmt.Sprintf("In verifyChatClaims: Error with decoding cookie or empty value: %s", err.Error()))
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	echo.Echo(echo.GreenBG, "Verifying chat claims: "+claims_token.Value)

	client_ip := request.Header.Get("X-Forwarded-For")
	if client_ip == "" {
		client_ip = request.RemoteAddr
	}

	err, _ = models.VerifyChatClaims(claims_token.Value, client_ip)
	if err != nil {
		echo.Echo(echo.RedBG, fmt.Sprintf("In verifyChatClaims: Error while calling models.VerifyChatClaims: %s", err.Error()))
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.Header().Set("X-Chat-Token", claims_token.Value)
	response.WriteHeader(200)
}

func createChatClaims(response http.ResponseWriter, request *http.Request) {
	client_ip := request.Header.Get("X-Forwarded-For")
	if client_ip == "" {
		client_ip = request.RemoteAddr
	}

	err, token := models.CreateChatClaims(client_ip)
	if err != nil {
		echo.Echo(echo.RedBG, "Error creating chat claims: "+err.Error())
		response.WriteHeader(http.StatusInternalServerError)
	}

	cookie := http.Cookie{
		Name:     app_config.CHAT_CLAIM_COOKIE_NAME,
		Value:    token,
		Expires:  time.Now().Add(time.Minute * 30),
		Domain:   request.Host,
		SameSite: http.SameSiteStrictMode,
		HttpOnly: true,
		Path:     "/",
	}

	http.SetCookie(response, &cookie)

	response.Header().Set("Content-Type", "application/json")
	response.Header().Set("X-Chat-Token", token)
	response.WriteHeader(http.StatusOK)

	return
}

func postChatClaimsHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}
func patchChatClaimsHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}
func deleteChatClaimsHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}
func putChatClaimsHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}
