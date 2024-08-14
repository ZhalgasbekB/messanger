package messanger

import (
	"database/sql"
	"forum/internal/models"
)

type MessangerSqlite struct {
	db *sql.DB
}

func NewMesssangerSqlite(db *sql.DB) *MessangerSqlite {
	return &MessangerSqlite{db: db}
}

const (
	conversationCreateQuery  = ""
	conversationHistoryQuery = ""
	conversationsQuery       = ""
	sendMessaeegQuery        = ""
)

func (m *MessangerSqlite) ConversationCreate(conversation *models.Conversations) error {
	if _, err := m.db.Exec("", conversation.UserID1, conversation.UserID2, conversation.CreatedAt); err != nil {
		return err
	}
	return nil
}

func (m *MessangerSqlite) Conversations() ([]*models.Conversations, error) {
	var conversations []*models.Conversations
	rows, err := m.db.Query("")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		conversation := new(models.Conversations)
		if err := rows.Scan(&conversation.ID, &conversation.UserID1, &conversation.UserID2, &conversation.CreatedAt); err != nil {
			return nil, err
		}
		conversations = append(conversations, conversation)
	}
	return conversations, nil
}

func (m *MessangerSqlite) ConversationHistory() error {
	// >>> con, mees  <<<
	return nil
}

func (m *MessangerSqlite) SendMessage() error {
	return nil
}
