package app_config

import (
	"encoding/json"
	"fmt"
	"libery_llm_chat_base_service/helpers"
	"os"

	"github.com/Gerardo115pp/patriots_lib/echo"
) // Loads the configuration from the environment variables

var SERVICE_PORT string = os.Getenv("SERVICE_PORT")
var SERVICE_NAME string = os.Getenv("SERVICE_NAME")
var SETTINGS_FILE string = os.Getenv("SETTINGS_FILE")
var JWT_SECRET string = os.Getenv("JWT_SECRET")
var DOMAIN_SECRET string = os.Getenv("DOMAIN_SECRET")
var PRODUCTION_DOMAIN string = os.Getenv("PRODUCTION_DOMAIN")
var DEVELOPMENT_DOMAIN string = os.Getenv("DEVELOPMENT_DOMAIN")
var OPERATION_DATA_PATH string = os.Getenv("OPERATION_DATA_PATH")
var RECAPTCHA_SK string = os.Getenv("RECAPTCHA_SK")

// --------Knowledge System--------

var KNOWLEDGE_QDRANT_REST_API_URL string
var KNOWLEDGE_QDRANT_COLLECTION string
var CONTACT_DETECTION_COLLECTION string = os.Getenv("QDRANT_CONTACT_DETECTION_COLLECTION")
var REMEMBRANCES_STORAGE_FILE string = fmt.Sprintf("%s/remembrances.json", OPERATION_DATA_PATH)
var ADA_EMBEDDING_API_URL string
var ADA_MODEL_NAME string
var ADA_EMBEDDING_DIMENSION int
var EMBEDDING_3_API_URL string
var EMBEDDING_3_MODEL_NAME string
var EMBEDDING_3_DIMENSION int
var VECTOR_DISTANCE string
var MAX_REMEMBRANCES_IN_CONTEXT int = 2

// --------Chat--------

var CHAT_CLAIM_COOKIE_NAME string = "portfolio_libery_chat_token"
var OPEN_AI_API_KEY string = os.Getenv("OPEN_AI_API_KEY")
var CHATS_DATA_PATH string = fmt.Sprintf("%s/chats", OPERATION_DATA_PATH)
var CHAT_INSTRUCTION string
var CHAT_ENABLED bool
var CHAT_TEMPERATURE float64 = 1.2
var CHAT_TOP_P float64 = 1.0
var CHAT_ASSISTANT_MAX_TOKENS int = 100 // Limit of the assistants response
var MAX_USER_TOKENS int = 100           // Limit of tokens per user message
var MAX_CHAT_SIZE int = 10              // Max number of messages in a single chat
var CONTACT_SIMILARITY_THRESHOLD float64 = 0.8
var CONTACT_INFO_EMBEDDING *[]float64

const (
	CHAT_SYSTEM_ROLE = "system"
	CHAT_USER_ROLE   = "user"
	CHAT_BOT_ROLE    = "assistant"
)

// --------Settings--------

var service_settings map[string]any = make(map[string]any)

func VerifyConfig() {

	if SERVICE_PORT == "" {
		panic("SERVICE PORT environment variable is required")
	}

	if SERVICE_NAME == "" {
		panic("SERVICE_NAME environment variable is required")
	}

	if SETTINGS_FILE == "" {
		SETTINGS_FILE = "settings.json"
	}

	if JWT_SECRET == "" {
		panic("JWT_SECRET environment variable is required")
	}

	if DOMAIN_SECRET == "" {
		panic("DOMAIN_SECRET environment variable is required")
	}

	if PRODUCTION_DOMAIN == "" {
		panic("PRODUCTION_DOMAIN environment variable is required")
	}

	if DEVELOPMENT_DOMAIN == "" {
		fmt.Println("Warning: DEVELOPMENT_DOMAIN environment variable is not set. defaulting to PRODUCTION_DOMAIN")
		DEVELOPMENT_DOMAIN = PRODUCTION_DOMAIN
	}

	if RECAPTCHA_SK == "" {
		panic("RECAPTCHA_SK environment variable is required")
	}

	if OPEN_AI_API_KEY == "" {
		panic("OPEN_AI_API_KEY environment variable is required")
	}

	if chats_dir, err := os.Stat(CHATS_DATA_PATH); os.IsNotExist(err) || !chats_dir.IsDir() {
		panic(fmt.Sprintf("Chats data path<%s> not found", CHATS_DATA_PATH))
	}

	if OPERATION_DATA_PATH == "" {
		panic("OPERATION_DATA_PATH environment variable is required")
	}

	err := loadSettings()
	if err != nil {
		panic(fmt.Sprintf("Error loading settings file: %s", err.Error()))
	}
}

func loadSettings() error {
	// Load settings from file
	var settings_path string = fmt.Sprintf("%s/%s", OPERATION_DATA_PATH, SETTINGS_FILE)
	var setting_exists bool
	var setting_value_interface interface{}
	var err error

	if _, err = os.Stat(settings_path); os.IsNotExist(err) {
		return fmt.Errorf("Settings file<%s> not found", settings_path)
	}

	var setting_content []byte

	setting_content, err = os.ReadFile(settings_path)

	var settings map[string]any = make(map[string]any)

	err = json.Unmarshal(setting_content, &settings)
	if err != nil {
		return fmt.Errorf("While unmarshaling settings data, found error <%s>", err.Error())
	}

	var chat_settings map[string]any = settings["chat_settings"].(map[string]any)
	var knowledge_settings map[string]any = settings["knowledge"].(map[string]any)
	var embeddings map[string]any = settings["embeddings"].(map[string]any)

	if embeddings == nil || chat_settings == nil || knowledge_settings == nil {
		return fmt.Errorf("Settings file is missing required sections: 'chat_settings', 'knowledge' or 'embeddings'")
	}

	var chat_instruction_filename string

	chat_instruction_filename = chat_settings["instruction_message_filename"].(string)
	if chat_instruction_filename == "" {
		return fmt.Errorf("Chat instruction message is required")
	}

	chat_instruction_filename = fmt.Sprintf("%s/%s", OPERATION_DATA_PATH, chat_instruction_filename)

	CHAT_INSTRUCTION, err = loadChatInstruction(chat_instruction_filename)
	if err != nil {
		return fmt.Errorf("While loading chat instruction: %s", err.Error())
	}

	CHAT_ENABLED, setting_exists = chat_settings["chat_enabled"].(bool)
	if !setting_exists {
		fmt.Println("Warning: Chat enabled setting not found, defaulting to false.")
		CHAT_ENABLED = false
	}

	if chat_temp, setting_exists := chat_settings["temperature"].(float64); setting_exists {
		CHAT_TEMPERATURE = chat_temp
	}

	if chat_top_p, setting_exists := chat_settings["top_p"].(float64); setting_exists {
		CHAT_TOP_P = chat_top_p
	}

	if chat_max_tokens, setting_exists := chat_settings["assistant_max_tokens"].(int); setting_exists {
		CHAT_ASSISTANT_MAX_TOKENS = chat_max_tokens
	}

	if user_max_tokens, setting_exists := chat_settings["max_user_tokens"].(int); setting_exists {
		MAX_USER_TOKENS = user_max_tokens
	}

	if chat_max_size, setting_exists := chat_settings["max_chat_size"].(int); setting_exists {
		MAX_CHAT_SIZE = chat_max_size
	}

	contact_info_embed, err := loadEmbedding("contact_info_embedding", embeddings)
	if err != nil {
		return fmt.Errorf("While loading contact_info_embedding: %s", err.Error())
	}

	CONTACT_INFO_EMBEDDING = contact_info_embed

	if contact_similarity_threshold, setting_exists := chat_settings["contact_similarity_threshold"].(float64); setting_exists {
		CONTACT_SIMILARITY_THRESHOLD = contact_similarity_threshold
	}

	// Knowledge settings

	KNOWLEDGE_QDRANT_REST_API_URL, setting_exists = knowledge_settings["qdrant_rest_api"].(string)
	if !setting_exists {
		return fmt.Errorf("Knowledge setting 'qdrant_rest_api_url' not found")
	}

	KNOWLEDGE_QDRANT_COLLECTION, setting_exists = knowledge_settings["qdrant_collection_name"].(string)
	if !setting_exists {
		return fmt.Errorf("Knowledge setting 'qdrant_collection_name' not found")
	}

	CONTACT_DETECTION_COLLECTION, setting_exists = embeddings["qdrant_contact_detection_collection_name"].(string)
	if !setting_exists {
		return fmt.Errorf("Embedding setting 'qdrant_contact_detection_collection_name' not found")
	}

	ADA_EMBEDDING_API_URL, setting_exists = embeddings["ada_embedding_api"].(string)
	if !setting_exists {
		return fmt.Errorf("Embedding setting 'ada_embedding_api' not found")
	}

	ADA_MODEL_NAME, setting_exists = embeddings["ada_embedding_api_model_name"].(string)
	if !setting_exists {
		return fmt.Errorf("Embedding setting 'ada_embedding_api_model_name' not found")
	}

	setting_value_interface, setting_exists = embeddings["ada_embedding_dimensions"]
	if !setting_exists {
		echo.Echo(echo.RedFG, fmt.Sprintf("Embedding setting 'ada_embedding_dimensions' not found: %v", embeddings))
		return fmt.Errorf("Embedding setting 'ada_embedding_dimensions' not found")
	} else { // Cast the value to int
		if ada_embedding_dimensions, is_number := setting_value_interface.(float64); is_number {
			ADA_EMBEDDING_DIMENSION = int(ada_embedding_dimensions)
		} else {
			return fmt.Errorf("Error while parsing 'ada_embedding_dimensions': was '%v'", setting_value_interface)
		}
	}

	EMBEDDING_3_API_URL, setting_exists = embeddings["embedding_3_api"].(string)
	if !setting_exists {
		return fmt.Errorf("Embedding setting 'embedding_3_api' not found")
	}

	EMBEDDING_3_MODEL_NAME, setting_exists = embeddings["embedding_3_model_name"].(string)
	if !setting_exists {
		return fmt.Errorf("Embedding setting 'embedding_3_model_name' not found")
	}

	setting_value_interface, setting_exists = embeddings["embedding_3_dimensions"]
	if !setting_exists {
		return fmt.Errorf("Embedding setting 'embedding_3_dimensions' not found")
	} else {
		if embedding_3_dimensions, is_number := setting_value_interface.(float64); is_number {
			EMBEDDING_3_DIMENSION = int(embedding_3_dimensions)
		} else {
			return fmt.Errorf("Error while parsing 'embedding_3_dimensions': was '%v'", setting_value_interface)
		}
	}

	VECTOR_DISTANCE, setting_exists = embeddings["embedding_vector_distance"].(string)
	if !setting_exists {
		return fmt.Errorf("Embedding setting 'embedding_vector_distance' not found")
	}

	setting_value_interface, setting_exists = knowledge_settings["max_context_remembrances"]
	if !setting_exists {
		return fmt.Errorf("Knowledge setting 'max_context_remembrances' not found")
	} else {
		if max_context_remembrance, is_number := setting_value_interface.(float64); is_number {
			MAX_REMEMBRANCES_IN_CONTEXT = int(max_context_remembrance)
		} else {
			return fmt.Errorf("Error while parsing 'max_context_remembrances': was '%v'", setting_value_interface)
		}
	}

	service_settings = settings

	return nil
}

func loadChatInstruction(chat_instruction_filename string) (string, error) {
	if !helpers.FileExists(chat_instruction_filename) {
		return "", fmt.Errorf("Chat instruction file<%s> not found", chat_instruction_filename)
	}

	instruction_content, err := os.ReadFile(chat_instruction_filename)
	if err != nil {
		return "", fmt.Errorf("Error while reading chat instruction file<%s>: %s", chat_instruction_filename, err.Error())
	}

	var looking_at_marker bool = false // If true the content been processed is a marker like '<INS>'

	var chat_instruction []byte = make([]byte, 0)
	var file_length int = len(instruction_content)
	var file_pointer int = 0
	var character byte

	for file_pointer < file_length {
		character = instruction_content[file_pointer]
		file_pointer++

		if looking_at_marker {
			if character == 0x3E {
				looking_at_marker = false

				if instruction_content[file_pointer] == 0x0A { // If the next character after a marker is a new line, skip it
					file_pointer++
				}

			} else if file_pointer == (file_length - 1) {
				return "", fmt.Errorf("Chat instruction file<%s> has an unclosed marker", chat_instruction_filename)
			}

			continue
		}

		if character == 0x3C {
			looking_at_marker = true
			continue
		}

		// echo.Echo(echo.GreenFG, fmt.Sprintf("Character: %c -> %x", character, character))
		chat_instruction = append(chat_instruction, character)
	}

	return string(chat_instruction), nil
}

func loadEmbedding(embedding_key string, embeddings_map map[string]interface{}) (*[]float64, error) {
	var embedding_interface []interface{} = embeddings_map[embedding_key].([]interface{})
	var embedding *[]float64 = new([]float64)

	if _, exists := embeddings_map[embedding_key]; !exists {
		return nil, fmt.Errorf("Embedding<%s> not found in the settings file", embedding_key)
	}

	for h, value := range embedding_interface {
		if float_value, is_float := value.(float64); is_float {
			*embedding = append(*embedding, float_value)
		} else {
			return nil, fmt.Errorf("Error while parsing embedding<%s>[index: %d]: was '%v'", embedding_key, h, value)
		}
	}

	return embedding, nil
}
