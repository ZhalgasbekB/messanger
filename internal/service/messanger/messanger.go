package messanger

import (
	"forum/internal/models"
	repo "forum/internal/repository"
)

type ConversationService struct {
	Conversation repo.Conversation
}

func NewMessangerService(repo repo.Repository) *ConversationService {
	return &ConversationService{
		Conversation: repo,
	}
}

func (con *ConversationService) ConversationCreateService(conversation *models.Conversations) error {
	return con.Conversation.ConversationCreate(conversation)
}

func (con *ConversationService) ConversationsService() ([]*models.Conversations, error) {
	return con.Conversation.Conversations()
}

func (con *ConversationService) ConversationHistoryService(conversation_id int) ([]*models.Messanger, error) {
	return con.Conversation.ConversationHistory(conversation_id)
}

func (con *ConversationService) SendMessageService(message models.Messanger) error {
	return con.Conversation.SendMessage(message)
}
