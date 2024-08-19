package websocket

import (
	"forum/internal/models"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var upGrader1 = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (wsh *WebSocketHandler) StreamChats(w http.ResponseWriter, r *http.Request) {
	ws2, err := upGrader1.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		wsh.renderError(w, http.StatusInternalServerError)
		return
	}

	user := getUserFromContext(r)
	go wsh.handleLastMessage(ws2, user.Id)
}

func (wsh *WebSocketHandler) handleLastMessage(ws2 *websocket.Conn, user_id int) {
	defer func() {
		ws2.Close()
	}()
	wsh.addChats(user_id, ws2)

	for {
		var stream Stream
		if err := ws2.ReadJSON(&stream); err != nil {
			log.Println(err)
			return
		}

		switch stream.Event {
		case "lastMessage":
			wsh.lastMessage(stream, user_id)
		}
	}
}

func (wsh *WebSocketHandler) lastMessage(stream Stream, user_id int) {
	if err := wsh.service.SendMessageService(models.Messanger{
		ConversationID: stream.StreamData.ConversationID,
		UserIDSender:   user_id,
		Message:        stream.StreamData.LastMessage,
		CreatedAt:      time.Now(),
	}); err != nil {
		log.Println(err)
		return
	}

	wsh.broadcastMessageInChats(stream)
}

func (wsh *WebSocketHandler) broadcastMessageInChats(stream Stream) {
	if conn, ok := wsh.activeChatsConnections[stream.StreamData.UserID]; ok {
		if err := conn.WriteJSON(&RESPQ{
			Event: "lastMessage",
			Stream: SteamResponse{
				Status:         true,
				LastMessage:    stream.StreamData.LastMessage,
				ConversationID: stream.StreamData.ConversationID},
		}); err != nil {
			log.Println(err)
			return
		}
	}
}

type Stream struct {
	Event      string     `json:"event"`
	StreamData StreamData `json:"data"`
}

type StreamData struct {
	ConversationID int    `json:"conversation_id"`
	UserID         int    `json:"user_id"`
	LastMessage    string `json:"last_message"`
}

type SteamResponse struct {
	Status         bool   `json:"status"`
	LastMessage    string `json:"last_message"`
	ConversationID int    `json:"conversation_id"`
}

type RESPQ struct {
	Event  string        `json:"event"`
	Stream SteamResponse `json:"data"`
}
