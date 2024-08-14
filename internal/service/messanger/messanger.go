package messanger

import (
	"forum/internal/models"
	repo "forum/internal/repository"
)

type MessangerService struct {
	Conversation repo.Conversation
}

func NewMessangerService(repo repo.Conversation) *MessangerService {
	return &MessangerService{
		Conversation: repo,
	}
}

func (m *MessangerService) ConversationCreateService(conversation *models.Conversations) error {
	return m.Conversation.ConversationCreate(conversation)
}

func (m *MessangerService) ConversationsService() ([]*models.Conversations, error) {
	return m.Conversation.Conversations()
}

func (m *MessangerService) ConversationHistoryService(conversation_id int) ([]*models.Messanger, error) {
	return m.Conversation.ConversationHistory(conversation_id)
}

func (m *MessangerService) SendMessageService(message models.Messanger) error {
	return m.Conversation.SendMessage(message)
}
