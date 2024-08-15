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

func (m *MessangerService) ConversationService(conversation_id int) (*models.Conversations, error) {
	return m.Conversation.Conversation1(conversation_id)
}

func (m *MessangerService) ConversationExistService(id1, id2 int) (int, error) {
	return m.Conversation.ConversationExist(id1, id2)
}

func (m *MessangerService) ConversationCreateService(conversation *models.Conversations) error {
	return m.Conversation.ConversationCreate(conversation)
}

func (m *MessangerService) ConversationsService(user_id int) ([]*models.Conversations, error) {
	return m.Conversation.Conversations(user_id)
}

func (m *MessangerService) ConversationHistoryService(conversation_id int) (*models.ChatDTO, error) {
	return m.Conversation.ConversationHistory(conversation_id)
}

func (m *MessangerService) SendMessageService(message models.Messanger) error {
	return m.Conversation.SendMessage(message)
}
