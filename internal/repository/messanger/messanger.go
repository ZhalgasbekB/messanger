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
	conversationCreateQuery  = "INSERT INTO conversations (user_id_1, user_id_2, create_at) VALUES (?, ?, ?);"
	conversationHistoryQuery = "SELECT id, user_id_1, user_id_2, create_at FROM conversations;"
	conversationsQuery       = "SELECT id, conversation_id, user_id_sender, messages, create_at FROM messages WHERE conversation_id = ?;"
	sendMessaeegQuery        = "INSERT INTO messages (conversation_id, user_id_sender, messages, create_at) VALUES (?, ?, ?, ?);"
	conversationQuery        = "SELECT * FROM conversations WHERE id= ?;"
)

func (m *MessangerSqlite) ConversationCreate(conversation *models.Conversations) error {
	if _, err := m.db.Exec(conversationCreateQuery, conversation.UserID1, conversation.UserID2, conversation.CreatedAt); err != nil {
		return err
	}
	return nil
}

func (m *MessangerSqlite) Conversations() ([]*models.Conversations, error) {
	var conversations []*models.Conversations
	rows, err := m.db.Query(conversationsQuery)
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

func (m *MessangerSqlite) ConversationHistory(conversation_id int) (*models.ChatDTO, error) {
	var convers models.Conversations
	if err := m.db.QueryRow(conversationQuery, conversation_id).Scan(&convers.ID, &convers.UserID1, &convers.UserID2, &convers.CreatedAt); err != nil {
		return nil, err
	}

	var messages []*models.Messanger

	rows, err := m.db.Query(conversationHistoryQuery)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		message := new(models.Messanger)
		if err := rows.Scan(&message.ID, &message.ConversationID, &message.UserIDSender, &message.Message, &message.CreatedAt); err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	return &models.ChatDTO{
		Conversation: convers,
		Messages:     messages,
	}, nil
}

func (m *MessangerSqlite) SendMessage(message models.Messanger) error {
	if _, err := m.db.Exec(sendMessaeegQuery, message.ConversationID, message.UserIDSender, message.Message, message.CreatedAt); err != nil {
		return err
	}
	return nil
}
