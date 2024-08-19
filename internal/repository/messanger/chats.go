package messanger

import (
	"database/sql"
	"forum/internal/models"
)

type Chats struct {
	db *sql.DB
}

func NewChatsSqlite(db *sql.DB) *Chats {
	return &Chats{
		db: db,
	}
}

const (
	listOfUsersQuery = "SELECT id, name FROM users"
	addUserChatQuery = ""
	chatNowQuery     = ""
)

func (ch *Chats) ListOfUsersToChat() (*[]models.User, error) {
	var people []models.User
	rows, err := ch.db.Query(listOfUsersQuery)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var person models.User
		if err := rows.Scan(&person.Id, &person.Name); err != nil {
			return nil, err
		}
		people = append(people, person)
	}

	return &people, nil
}

func (ch *Chats) ChatsNow() error {
	return nil
}
