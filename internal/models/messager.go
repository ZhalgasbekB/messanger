package models

import "time"

type Messanger struct {
	ID             int       `json:"id"`
	ConversationID int       `json:"room_id"`
	UserIDSender   int       `json:"user_id_sender"`
	Message        string    `json:"message"`
	CreatedAt      time.Time `json:created_at"`
}

type Conversations struct {
	ID        int       `json:"id"`
	UserID1   int       `json:"user_id_1"`
	UserID2   int       `json:"user_id_2"`
	CreatedAt time.Time `json:created_at"`
}

type MessangerDTO struct {
	Event string `json:"event"`
	Data  struct {
		ConversationID uint   `json:"conversationID"`
		RecipientID    uint   `json:"recipientID"`
		Content        string `json:"content"`
	} `json:"data"`
}

type Chat struct {
	ConversationID int
	User *User
	UserID2  int          `json:"user_id_2"`
	Messages []*Messanger `json:"chat_history"`
}

type ChatDTO struct {
	Conversation Conversations `json:"conversations"`
	Messages     []*Messanger  `json:"chat_history"`
}
