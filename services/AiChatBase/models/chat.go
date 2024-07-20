package models

import (
	"fmt"
	app_config "libery_llm_chat_base_service/Config"
	"time"
)

type ChatMessage struct {
	Order    int    `json:"order"`
	Author   string `json:"author"`
	Content  string `json:"content"`
	SendDate string `json:"send_date"`
}

func CreateChatMessage(order int, author string, content string) *ChatMessage {
	var new_message *ChatMessage = new(ChatMessage)

	new_message.Order = order
	new_message.Author = author
	new_message.Content = content
	new_message.SendDate = time.Now().Format("2006-01-02 15:04:05")

	return new_message
}

func (chat_message *ChatMessage) Write(data []byte) (n int, err error) {
	chat_message.Content += string(data)
	return len(data), nil
}

type ChatRoom struct {
	ID                   string         `json:"id"`
	CreationDate         string         `json:"creation_date"`
	LastMessageDate      string         `json:"last_message_date"`
	InstructionMessage   string         `json:"instruction_message"`
	MaxUserMessageLength int            `json:"max_user_message_length"`
	Messages             []*ChatMessage `json:"messages"`
}

func (chat *ChatRoom) AddMessage(message string, from_user bool) *ChatMessage {
	var new_message *ChatMessage = new(ChatMessage)
	new_message.Order = len(chat.Messages)
	new_message.Content = message
	new_message.Author = app_config.CHAT_USER_ROLE

	if !from_user {
		new_message.Author = app_config.CHAT_BOT_ROLE
	}

	new_message.SendDate = time.Now().Format("2006-01-02 15:04:05")
	chat.Messages = append(chat.Messages, new_message)
	chat.LastMessageDate = new_message.SendDate

	return new_message
}

func (chat *ChatRoom) AppendMessage(message *ChatMessage) error {
	if len(chat.Messages) == app_config.MAX_CHAT_SIZE {
		return fmt.Errorf("Chat room is full")
	}

	chat.Messages = append(chat.Messages, message)
	chat.LastMessageDate = message.SendDate

	return nil
}

func (chat ChatRoom) GetLastMessage() *ChatMessage {
	if len(chat.Messages) == 0 {
		return nil
	}

	return chat.Messages[len(chat.Messages)-1]
}

type ChatRoomPublicResponse struct {
	ID                   string         `json:"id"`
	Messages             []*ChatMessage `json:"messages"`
	MaxUserMessageLength int            `json:"max_user_message_length"`
}

func (chat *ChatRoom) GetPublicResponse() *ChatRoomPublicResponse {
	var public_response *ChatRoomPublicResponse = new(ChatRoomPublicResponse)
	public_response.ID = chat.ID
	public_response.Messages = chat.Messages
	public_response.MaxUserMessageLength = chat.MaxUserMessageLength

	return public_response
}
