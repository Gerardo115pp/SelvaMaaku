package workflows

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"libery_llm_chat_base_service/helpers"
	"libery_llm_chat_base_service/models"
	"libery_llm_chat_base_service/workflows/chat_messages"
	"net/http"
	"os"

	app_config "libery_llm_chat_base_service/Config"

	"github.com/Gerardo115pp/patriots_lib/echo"
)

func DetectContactInfo(message *models.ChatMessage) (bool, error) {
	var looks_like_contact_info bool = false

	looks_like_contact_info = helpers.MessageLooksLikeContactInfo(message.Content)
	if !looks_like_contact_info {
		return false, nil
	}

	echo.Echo(echo.GreenFG, fmt.Sprintf("Message looks like contact info, checking.."))

	var embedding_request *models.AdaEmbedRequest = models.CreateAdaEmbedRequest()

	embedding_request.Input = []string{message.Content}

	response, err := requestAdaEmbedding(embedding_request)
	if err != nil {
		return false, err
	}

	var contact_info_embedding *[]float64 = &response.Data[0].Embedding

	var cosine_similarity float64

	cosine_similarity, err = helpers.ComputeCosineSimilarity(*app_config.CONTACT_INFO_EMBEDDING, *contact_info_embedding)
	if err != nil {
		return false, err
	}

	echo.Echo(echo.GreenFG, fmt.Sprintf("Cosine similarity: %f", cosine_similarity))

	return cosine_similarity >= app_config.CONTACT_SIMILARITY_THRESHOLD, nil
}

func EmbedTextSemantics(text string) ([]float64, error) {
	var embedding_request *models.TextEmbedding3SmallRequest = models.CreateTextEmbedding3SmallRequest(false)

	embedding_request.Input = text

	text_embedding_3_response, err := requestTextEmbedding3Small(embedding_request)
	if err != nil {
		return nil, err
	}

	if len(text_embedding_3_response.Data) == 0 {
		return nil, fmt.Errorf("No embeddings returned")
	}

	return text_embedding_3_response.Data[0].Embedding, nil
}

func getChatInstructionWithContext(chat_room *models.ChatRoom) (*models.ChatMessage, error) {
	var last_message *models.ChatMessage = chat_room.GetLastMessage()

	echo.Echo(echo.GreenFG, fmt.Sprintf("Last message: %s. %s == %s", last_message.Content, last_message.Author, app_config.CHAT_USER_ROLE))

	if last_message == nil || last_message.Author != app_config.CHAT_USER_ROLE {
		return nil, fmt.Errorf("The last message is not from the user. no need to get chat context")
	}

	var user_message_semantics []float64

	user_message_semantics, err := EmbedTextSemantics(last_message.Content)
	if err != nil {
		return nil, err
	}

	var system_message *models.ChatMessage

	system_message, err = chat_messages.CreateChatInstructionMessage(user_message_semantics)
	if err != nil {
		return nil, err
	}

	return system_message, nil
}

func SalesChatWithGPT3Turbo(chat_room *models.ChatRoom) (*models.ChatMessage, error) {
	var chat_request *models.GPT3TurboRequest = models.CreateGPT3TurboRequest()

	chat_request.Temperature = app_config.CHAT_TEMPERATURE
	chat_request.TopP = app_config.CHAT_TOP_P
	chat_request.MaxTokens = app_config.CHAT_ASSISTANT_MAX_TOKENS

	var system_message *models.GPT3TurboMessage = new(models.GPT3TurboMessage)
	var char_instruction_message *models.ChatMessage

	system_message.Role = "system"
	system_message.Content = app_config.CHAT_INSTRUCTION

	// Try to get chat instruction with relevant context to answer the user's question.
	char_instruction_message, err := getChatInstructionWithContext(chat_room)
	if err == nil {
		system_message.Content = char_instruction_message.Content
	} else {
		echo.EchoWarn(fmt.Sprintf("Error getting chat instruction with context: %s.\n\nProceeding with default chat instruction: %s", err.Error(), app_config.CHAT_INSTRUCTION))
	}

	chat_request.Messages = append(chat_request.Messages, system_message)

	var new_message *models.GPT3TurboMessage

	for _, message := range chat_room.Messages {
		new_message = new(models.GPT3TurboMessage)
		new_message.Role = message.Author
		new_message.Content = message.Content

		chat_request.Messages = append(chat_request.Messages, new_message)
	}

	response, err := requestGPT3Turbo(chat_request)
	if err != nil {
		return nil, err
	}

	new_chat_message := chat_room.AddMessage(response.Choices[0].Message.Content, false)

	return new_chat_message, nil
}

func requestGPT3Turbo(request *models.GPT3TurboRequest) (*models.GPT3TurboResponse, error) {
	var request_body []byte
	var err error

	request_body, err = json.Marshal(request)
	if err != nil {
		echo.Echo(echo.RedFG, fmt.Sprintf("Error marshalling request body: %s", err.Error()))
		return nil, err
	}

	body := bytes.NewBuffer(request_body)
	http_request, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", body)
	if err != nil {
		echo.Echo(echo.RedFG, fmt.Sprintf("Error sending GPT 3-Turbo request: %s", err.Error()))
		return nil, err
	}

	http_request.Header.Add("Content-Type", "application/json")
	http_request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", app_config.OPEN_AI_API_KEY))

	http_client := http.Client{}
	response, err := http_client.Do(http_request)
	if err != nil {
		echo.Echo(echo.RedFG, fmt.Sprintf("Error sending GPT 3-Turbo request: %s", err.Error()))
		return nil, err
	}

	if response.StatusCode < 200 || response.StatusCode > 299 {
		echo.Echo(echo.RedFG, fmt.Sprintf("Error sending GPT 3-Turbo request, status code: %d", response.StatusCode))
		body_content, err := ioutil.ReadAll(response.Body)
		if err != nil {
			echo.Echo(echo.RedFG, fmt.Sprintf("Error reading response body: %s", err.Error()))
			return nil, err
		}
		echo.Echo(echo.RedFG, fmt.Sprintf("Reason: %s", string(body_content)))
		echo.Echo(echo.RedFG, fmt.Sprintf("Request body: %s", string(request_body)))
		return nil, fmt.Errorf("Error sending GPT 3-Turbo request, status code: %d", response.StatusCode)
	}

	return parseGPT3TurboResponse(response)
}

func requestAdaEmbedding(request *models.AdaEmbedRequest) (*models.AdaEmbedResponse, error) {
	var request_body []byte
	var err error

	request_body, err = json.Marshal(request)
	if err != nil {
		echo.Echo(echo.RedFG, fmt.Sprintf("Error marshalling request body: %s", err.Error()))
		return nil, err
	}

	body := bytes.NewBuffer(request_body)
	http_request, err := http.NewRequest("POST", "https://api.openai.com/v1/embeddings", body)

	http_request.Header.Add("Content-Type", "application/json")
	http_request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", app_config.OPEN_AI_API_KEY))

	http_client := http.Client{}
	response, err := http_client.Do(http_request)
	if err != nil {
		echo.Echo(echo.RedFG, fmt.Sprintf("Error sending Ada-002 request: %s", err.Error()))
		return nil, err
	}

	if response.StatusCode < 200 || response.StatusCode > 299 {
		echo.Echo(echo.RedFG, fmt.Sprintf("Error sending Ada-002 request, status code: %d", response.StatusCode))
		body_content, err := ioutil.ReadAll(response.Body)
		if err != nil {
			echo.Echo(echo.RedFG, fmt.Sprintf("Error reading response body: %s", err.Error()))
			return nil, err
		}
		echo.Echo(echo.RedFG, fmt.Sprintf("Reason: %s", string(body_content)))
		echo.Echo(echo.RedFG, fmt.Sprintf("Request body: %s", string(request_body)))
		return nil, fmt.Errorf("Error sending Ada-002 request, status code: %d", response.StatusCode)
	}

	return parseAdaEmbedResponse(response)
}

func requestTextEmbedding3Small(request *models.TextEmbedding3SmallRequest) (*models.TextEmbedding3SmallResponse, error) {
	var request_body []byte
	var err error

	request_body, err = json.Marshal(request)
	if err != nil {
		echo.Echo(echo.RedFG, fmt.Sprintf("Error marshalling request body: %s", err.Error()))
		return nil, err
	}

	body := bytes.NewBuffer(request_body)
	http_request, err := http.NewRequest("POST", app_config.EMBEDDING_3_API_URL, body)

	http_request.Header.Add("Content-Type", "application/json")
	http_request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", app_config.OPEN_AI_API_KEY))

	http_client := http.Client{}
	response, err := http_client.Do(http_request)
	if err != nil {
		echo.Echo(echo.RedFG, fmt.Sprintf("Error sending Text Embedding 3-Small request: %s", err.Error()))
		return nil, err
	}

	if response.StatusCode < 200 || response.StatusCode > 299 {
		echo.Echo(echo.RedFG, fmt.Sprintf("Error sending Text Embedding 3-Small request, status code: %d", response.StatusCode))
		body_content, err := io.ReadAll(response.Body)
		if err != nil {
			echo.Echo(echo.RedFG, fmt.Sprintf("Error reading response body: %s", err.Error()))
			return nil, err
		}

		echo.Echo(echo.RedFG, fmt.Sprintf("Reason: %s", string(body_content)))
		echo.Echo(echo.RedFG, fmt.Sprintf("Request body: %s", string(request_body)))
		return nil, fmt.Errorf("Error sending Text Embedding 3-Small request, status code: %d", response.StatusCode)
	}

	return parseTextEmbedding3SmallResponse(response)
}

func requestGPT3TurboIdea(prompt string) (string, error) {
	request_data := map[string]interface{}{
		"max_tokens":  40,
		"temperature": 1.26,
		"model":       "gpt-3.5-turbo",
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": prompt,
			},
		},
	}

	request_body, err := json.Marshal(request_data)
	if err != nil {
		echo.Echo(echo.RedFG, fmt.Sprintf("Error marshalling request body: %s", err.Error()))
		return "", err
	}

	body := bytes.NewBuffer(request_body)
	request, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", body)
	if err != nil {
		echo.Echo(echo.RedFG, fmt.Sprintf("Error creating new project idea: %s", err.Error()))
		return "", err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("OPENAI_API_KEY")))

	http_client := http.Client{}
	response, err := http_client.Do(request)
	if err != nil {
		echo.Echo(echo.RedFG, fmt.Sprintf("Error creating new project idea: %s", err.Error()))
		return "", err
	}

	if response.StatusCode != 200 {
		echo.Echo(echo.RedFG, fmt.Sprintf("Error creating new project idea, status code: %d", response.StatusCode))
		body_content, err := ioutil.ReadAll(response.Body)
		if err != nil {
			echo.Echo(echo.RedFG, fmt.Sprintf("Error reading response body: %s", err.Error()))
			return "", err
		}
		echo.Echo(echo.RedFG, fmt.Sprintf("Reason: %s", string(body_content)))
		echo.Echo(echo.RedFG, fmt.Sprintf("Request body: %s", string(request_body)))
		return "", err
	}

	response_data, err := parseGPT3TurboResponse(response)
	if err != nil {
		echo.Echo(echo.RedFG, fmt.Sprintf("Error parsing response: %s", err.Error()))
		return "", err
	}

	return response_data.Choices[0].Message.Content, nil
}

func parseGPT3TurboResponse(response *http.Response) (*models.GPT3TurboResponse, error) {
	var response_data *models.GPT3TurboResponse = new(models.GPT3TurboResponse)
	err := json.NewDecoder(response.Body).Decode(&response_data)
	if err != nil {
		echo.Echo(echo.RedFG, fmt.Sprintf("Error decoding response: %s", err.Error()))
		return nil, err
	}

	return response_data, nil
}

func parseAdaEmbedResponse(response *http.Response) (*models.AdaEmbedResponse, error) {
	var response_data *models.AdaEmbedResponse = new(models.AdaEmbedResponse)
	err := json.NewDecoder(response.Body).Decode(&response_data)
	if err != nil {
		echo.Echo(echo.RedFG, fmt.Sprintf("Error decoding response: %s", err.Error()))
		return nil, err
	}

	return response_data, nil
}

func parseTextEmbedding3SmallResponse(response *http.Response) (*models.TextEmbedding3SmallResponse, error) {
	var response_data *models.TextEmbedding3SmallResponse = new(models.TextEmbedding3SmallResponse)
	err := json.NewDecoder(response.Body).Decode(&response_data)
	if err != nil {
		echo.Echo(echo.RedFG, fmt.Sprintf("Error decoding response: %s", err.Error()))
		return nil, err
	}

	return response_data, nil
}
