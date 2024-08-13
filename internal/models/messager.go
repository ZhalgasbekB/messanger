package models

import "time"

type Messanger struct {
	ID             int       `json:"id"`
	ConversationID int       `json:"room_id"`
	UserID         int       `json:"user_id"`
	Message        string    `json:"message"`
	CreatedAt      time.Time `json:created_at"`
}

type Conversations struct {
	ID int
}

type Members struct{}
