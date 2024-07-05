package models

import "time"

type ChatMessage struct {
	Order    int    `json:"order"`
	Author   string `json:"author"`
	Content  string `json:"content"`
	SendDate string `json:"send_date"`
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
	new_message.Author = "user"

	if !from_user {
		new_message.Author = "assistant"
	}

	new_message.SendDate = time.Now().Format("2006-01-02 15:04:05")
	chat.Messages = append(chat.Messages, new_message)
	chat.LastMessageDate = new_message.SendDate

	return new_message
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
