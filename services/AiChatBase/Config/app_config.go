package app_config

import (
	"encoding/json"
	"fmt"
	"os"
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
var CONTACT_INFO_EMBEDDING *[]float64
var CONTACT_SIMILARITY_THRESHOLD float64 = 0.8

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

	CHAT_INSTRUCTION := chat_settings["instruction_message"].(string)
	if CHAT_INSTRUCTION == "" {
		return fmt.Errorf("Chat instruction message is required")
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

	if contact_info_embed, setting_exists := chat_settings["contact_info_embedding"].([]float64); setting_exists {
		CONTACT_INFO_EMBEDDING = &contact_info_embed
	} else {
		return fmt.Errorf("Contact info embedding is required")
	}

	if contact_similarity_threshold, setting_exists := chat_settings["contact_similarity_threshold"].(float64); setting_exists {
		CONTACT_SIMILARITY_THRESHOLD = contact_similarity_threshold
	}

	service_settings = settings

	return nil
}
