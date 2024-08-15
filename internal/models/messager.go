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

type ConversationDTO struct {
	UserID1   int       `json:"user_id_1"`
	UserID2   int       `json:"user_id_2"`
	CreatedAt time.Time `json:created_at"`
}

type Chat struct {
	ConversationID int
	User           *User
	UserID2        int          `json:"user_id_2"`
	Messages       []*Messanger `json:"chat_history"`
}

type ChatDTO struct {
	Conversation Conversations `json:"conversations"`
	Messages     []*Messanger  `json:"chat_history"`
}

type Chats struct {
	User          *User
	Conversations []*Conversations
}

type MessangerDTO struct {
	Event string `json:"event"`
	Data  struct {
		ConversationID string   `json:"conversationID"`
		RecipientID    string   `json:"recipientID"`
		Content        string `json:"content"`
	} `json:"data"`
}

type Data struct {
	ConversationID string   `json:"conversationID"`
	RecipientID    string   `json:"recipientID"`
	Content        string `json:"content"`
}

type Data1 struct {
	ConversationID int       `json:"conversationID"`
	SenderID       int       `json:"sender_id"`
	Content        string    `json:"content"`
	CreatedAt      time.Time `json:created_at"`
}

type MessangerDTO1 struct {
	Event string `json:"event"`
	Data1 struct {
		ConversationID int       `json:"conversationID"`
		SenderID       int       `json:"sender_id"`
		Content        string    `json:"content"`
		CreatedAt      time.Time `json:created_at"`
	} `json:"data"`
}
