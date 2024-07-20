package repository

import "libery_llm_chat_base_service/models"

type ChatRepository interface {
	GetChatByID(chat_id string, create bool) (*models.ChatRoom, error)
	SaveChat(chat *models.ChatRoom) error
}

var chat_repo_implementation ChatRepository

func SetChatRepositoryImplementation(implementation ChatRepository) {
	chat_repo_implementation = implementation
}

func GetChatByID(chat_id string, create bool) (*models.ChatRoom, error) {
	return chat_repo_implementation.GetChatByID(chat_id, create)
}

func SaveChat(chat *models.ChatRoom) error {
	return chat_repo_implementation.SaveChat(chat)
}
