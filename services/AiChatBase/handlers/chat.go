package handlers

import (
	"encoding/json"
	"fmt"
	app_config "libery_llm_chat_base_service/Config"
	"libery_llm_chat_base_service/helpers"
	"libery_llm_chat_base_service/models"
	"libery_llm_chat_base_service/repository"
	"libery_llm_chat_base_service/server"
	"libery_llm_chat_base_service/workflows"
	"net/http"

	"github.com/Gerardo115pp/patriots_lib/echo"
)

func ChatHandler(portfolio_server server.Server) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			getChatHandler(response, request)
		case http.MethodPost:
			postChatHandler(response, request)
		case http.MethodPatch:
			patchChatHandler(response, request)
		case http.MethodDelete:
			deleteChatHandler(response, request)
		case http.MethodPut:
			putChatHandler(response, request)
		case http.MethodOptions:
			response.WriteHeader(http.StatusOK)
		default:
			response.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func getChatHandler(response http.ResponseWriter, request *http.Request) {
	var claim_data *models.ChatClaims = request.Context().Value("claims").(*models.ChatClaims)

	if claim_data.ChatID == "" {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	var chat_room *models.ChatRoom
	chat_room, err := repository.GetChatByID(claim_data.ChatID, true)
	if err != nil {
		response.WriteHeader(404)
		return
	}

	if len(chat_room.Messages) == 0 {
		chat_room.AddMessage(app_config.CHAT_GREET_MESSAGE, false)
	}
	fmt.Printf("Chat ID: %s\n", chat_room.ID)

	err = repository.SaveChat(chat_room)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	var response_data *models.ChatRoomPublicResponse = chat_room.GetPublicResponse()

	err = json.NewEncoder(response).Encode(response_data)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
	}
}

func postChatHandler(response http.ResponseWriter, request *http.Request) {
	var err error
	var claim_data *models.ChatClaims = request.Context().Value("claims").(*models.ChatClaims)

	echo.Echo(echo.GreenBG, "posting to chat ID: "+claim_data.ChatID)

	// ---- IDENTIFYING CHAT ----

	if claim_data.ChatID == "" {
		echo.Echo(echo.RedFG, fmt.Sprintf("In postChatHandler: Chat ID not found in claims\nclaims: %v", claim_data))
		response.WriteHeader(400)
		return
	}

	var chat_room *models.ChatRoom
	chat_room, err = repository.GetChatByID(claim_data.ChatID, false)
	if err != nil {
		response.WriteHeader(404)
		return
	}

	// ---- READING REQUEST DATA ----
	var request_data map[string]string = make(map[string]string)

	err = json.NewDecoder(request.Body).Decode(&request_data)
	if err != nil {
		echo.Echo(echo.RedFG, fmt.Sprintf("In postChatHandler: Error decoding request data because '%s'", err.Error()))
		response.WriteHeader(400)
		return
	}

	// ---- READING MESSAGE DATA ----
	var message_content string = request_data["content"]
	if helpers.TokenCount(message_content) > float64(app_config.MAX_USER_TOKENS) {
		echo.Echo(echo.RedFG, fmt.Sprintf("Message too long: %f tokens for a limit of %d", helpers.TokenCount(message_content), app_config.MAX_USER_TOKENS))
		helpers.WriteRejection(response, 413, "message-too-long")
		return
	}

	user_message := chat_room.AddMessage(message_content, true)

	// ---- GENERATING ASSISTANT RESPONSE ----
	if len(chat_room.Messages) > app_config.MAX_CHAT_SIZE {
		echo.Echo(echo.RedFG, fmt.Sprintf("Chat message count reached limits: %d messages for a limit of %d", len(chat_room.Messages), app_config.MAX_CHAT_SIZE))
		helpers.WriteRejection(response, 429, "too-many-messages")
		return
	}

	has_contact_info, err := workflows.DetectContactInfo(user_message)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	if has_contact_info {
		handleContactInfo(chat_room, response, request)
		return
	}

	var new_assistant_message *models.ChatMessage
	new_assistant_message, err = workflows.SalesChatWithGPT3Turbo(chat_room) // workflows.SalesChatWithGPT3 uses AddMessage internally, don't duplicate it
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = repository.SaveChat(chat_room)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(200)

	err = json.NewEncoder(response).Encode(new_assistant_message)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
	}

	return
}

func handleContactInfo(chat_room *models.ChatRoom, response http.ResponseWriter, request *http.Request) {
	echo.Echo(echo.GreenFG, "Contact info detected")

	conversation_end_message := chat_room.AddMessage("Thank you for your contact information, Gerardo will get back to you as soon as possible.", false)

	err := repository.SaveChat(chat_room)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(conversation_end_message)

	// TODO: replace telegram message with email or maybe make it so that the service can be configured to enable and disable different notification methods
	// go func() {
	// 	localtime, err := time.LoadLocation("America/Mexico_City")
	// 	if err != nil {
	// 		echo.Echo(echo.RedFG, fmt.Sprintf("Error loading timezone: %s", err.Error()))
	// 		return
	// 	}

	// 	beginning_message := fmt.Sprintf("New contact info received on %s", time.Now().In(localtime).Format("2006-01-02 15:04:05"))
	// 	err = workflows.SendTelegramMessage(beginning_message)
	// 	if err != nil {
	// 		echo.Echo(echo.RedFG, fmt.Sprintf("Error sending telegram message: %s", err.Error()))
	// 		return
	// 	}

	// 	for _, message := range chat_room.Messages {
	// 		telegram_message := fmt.Sprintf("%s:\n\n %s", message.Author, message.Content)
	// 		err = workflows.SendTelegramMessage(telegram_message)
	// 		if err != nil {
	// 			echo.Echo(echo.RedFG, fmt.Sprintf("Error sending telegram message: %s", err.Error()))
	// 			return
	// 		}
	// 	}

	// 	echo.Echo(echo.GreenFG, "Contact info sent to telegram")
	// }()

	return
}

func patchChatHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}
func deleteChatHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}
func putChatHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}
