package messanger

import (
	"forum/internal/models"
	"forum/internal/repository"
)

type ChatsService struct {
	Chats repository.People
}

func NewChatsService(people repository.People) *ChatsService {
	return &ChatsService{Chats: people}
}

func (chs *ChatsService) ListOfUsersToChatService() (*[]models.User, error) {
	return chs.Chats.ListOfUsersToChat()
}
