package chat_messages

import (
	app_config "libery_llm_chat_base_service/Config"
	"libery_llm_chat_base_service/models"
	"libery_llm_chat_base_service/repository"
	"text/template"

	"github.com/Gerardo115pp/patriots_lib/echo"
)

type ChatInstructionTemplateData struct {
	RelatedKnowledge string `json:"related_knowledge"`
}

func getUserMessageRelatedKnowledge(user_message_semantics []float64) (string, error) {
	remembrances, err := repository.Knowledge.Recall(user_message_semantics)
	if err != nil {
		return "", err
	}

	var related_knowledge string = ""

	for _, remembrance := range remembrances {
		related_knowledge += remembrance.Document + "\n---\n"
	}

	return related_knowledge, nil
}

func CreateChatInstructionMessage(user_message_semantics []float64) (*models.ChatMessage, error) {
	var add_knowledge bool = user_message_semantics != nil && len(user_message_semantics) == app_config.EMBEDDING_3_DIMENSION

	var chat_instruction_data *ChatInstructionTemplateData = new(ChatInstructionTemplateData)

	chat_instruction_data.RelatedKnowledge = ""

	if add_knowledge {
		related_knowledge, err := getUserMessageRelatedKnowledge(user_message_semantics)
		if err != nil {
			return nil, err
		}

		chat_instruction_data.RelatedKnowledge = related_knowledge
	}

	echo.Echo(echo.GreenBG, "Related knowledge:\n"+chat_instruction_data.RelatedKnowledge)

	chat_instruction_message_template, err := template.New("chat_instruction_message").Parse(app_config.CHAT_INSTRUCTION)
	if err != nil {
		return nil, err
	}

	var chat_instruction_message *models.ChatMessage = models.CreateChatMessage(0, app_config.CHAT_SYSTEM_ROLE, "")

	err = chat_instruction_message_template.Execute(chat_instruction_message, chat_instruction_data)
	if err != nil {
		return nil, err
	}

	return chat_instruction_message, nil
}
