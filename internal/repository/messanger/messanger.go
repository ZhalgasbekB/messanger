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
	conversationCreateQuery = "INSERT INTO conversations (user_id_1, user_id_2, created_at) VALUES (?, ?, ?);"
	conversationAllQuery    = "SELECT * FROM conversations WHERE user_id_1 = ? OR user_id_2 = ?;"
	ALL                     = "SELECT c.id AS conversation_id,  c.user_id_1,  c.user_id_2,   m.message,   m.created_at FROM  conversations c  INNER JOIN      (SELECT  conversation_id,   MAX(created_at) AS last_message_time  FROM     messages  GROUP BY         conversation_id) lm   ON c.id = lm.conversation_id  INNER JOIN      messages m      ON c.id = m.conversation_id AND m.created_at = lm.last_message_time  WHERE    c.user_id_1 = ? OR c.user_id_2 = ? ORDER BY  m.created_at DESC;"

	conversationsHistoryQuery = "SELECT id, conversation_id, user_id_sender, message, created_at FROM messages WHERE conversation_id = ?;"
	sendMessaeegQuery         = "INSERT INTO messages (conversation_id, user_id_sender, message, created_at) VALUES (?, ?, ?, ?);"
	conversationQuery         = "SELECT * FROM conversations WHERE id= ?;"

	findAChatQuery = `SELECT id FROM conversations WHERE (user_id_1 = ? AND user_id_2 = ?) OR (user_id_1 = ? AND user_id_2 = ?);`
)

func (m *MessangerSqlite) ConversationExist(id1, id2 int) (int, error) {
	var conversation_id int
	if err := m.db.QueryRow(findAChatQuery, id1, id2, id2, id1).Scan(&conversation_id); err != nil {
		return -1, err
	}
	return conversation_id, nil
}

func (m *MessangerSqlite) Conversation1(conversation_id int) (*models.Conversations, error) {
	var conversation models.Conversations
	if err := m.db.QueryRow(conversationQuery, conversation_id).Scan(&conversation.ID, &conversation.UserID1, &conversation.UserID2, &conversation.CreatedAt); err != nil {
		return nil, err
	}
	return &conversation, nil
}

func (m *MessangerSqlite) ConversationCreate(conversation *models.Conversations) error {
	if _, err := m.db.Exec(conversationCreateQuery, conversation.UserID1, conversation.UserID2, conversation.CreatedAt); err != nil {
		return err
	}
	return nil
}

func (m *MessangerSqlite) Conversations(user_id int) ([]*models.Conversations, error) {
	var conversations []*models.Conversations
	rows, err := m.db.Query(ALL, user_id, user_id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		conversation := new(models.Conversations)
		if err := rows.Scan(&conversation.ID, &conversation.UserID1, &conversation.UserID2, &conversation.LastMessage, &conversation.CreatedAt); err != nil {
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

	rows, err := m.db.Query(conversationsHistoryQuery, conversation_id)
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
