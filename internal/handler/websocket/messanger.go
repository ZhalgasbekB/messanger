package websocket

import (
	"database/sql"
	"fmt"
	"forum/internal/models"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (wsh *WebSocketHandler) InitialConversation(w http.ResponseWriter, r *http.Request) {
	ws, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	user := getUserFromContext(r)
	go wsh.handleConnection(ws, user.Id)
}

func (wsh *WebSocketHandler) handleConnection(conn *websocket.Conn, id int) {
	defer func() {
		conn.Close()
		wsh.remove(id)
	}()
	wsh.add(id, conn)

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ticker.C:
				if err := conn.WriteJSON(&Event{Event: "Ping"}); err != nil {
					log.Println(err)
					return
				}
			}
		}
	}()

	for {
		var messages *models.MessangerDTO
		if err := conn.ReadJSON(&messages); err != nil {
			log.Println(err)
			break
		}

		switch messages.Event {
		case "closeConnection":
			fmt.Println("Connection Closed")
			break
		case "initiateConversation":
			wsh.connectionChat(conn, id, messages.Data.RecipientID)
		case "sendMessage":
			wsh.sendMessage(*messages, id)
		case "Pong":
			log.Println("Pong received: ", time.Now())
		}
	}
}

func (wsh *WebSocketHandler) connectionChat(conn *websocket.Conn, id int, re_id int) {
	conversation_id, err := wsh.service.ConversationExistService(id, re_id)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
		return
	}

	if conversation_id == -1 {
		if err := wsh.service.Conversation.ConversationCreateService(&models.Conversations{UserID1: id, UserID2: re_id, CreatedAt: time.Now()}); err != nil {
			log.Println(err)
			return
		}
	}

	response := models.MessangerDTO{
		Event: "conversationInitiated",
		Data: &models.Data{
			ConversationID: conversation_id,
		},
	}
	if err := conn.WriteJSON(response); err != nil {
		log.Println("Error sending confirmation:", err)
	}
}

func (wsh *WebSocketHandler) sendMessage(m models.MessangerDTO, sender int) {
	message := models.Messanger{
		ConversationID: m.Data.ConversationID,
		UserIDSender:   sender,
		Message:        m.Data.Content,
		CreatedAt:      time.Now(),
	}

	if err := wsh.service.Conversation.SendMessageService(message); err != nil {
		log.Println(err)
		return
	}

	conversation, err := wsh.service.ConversationService(message.ConversationID)
	if err != nil {
		log.Println("Error retrieving conversation:", err)
		return
	}

	wsh.broadcastingMessages(conversation, &message)
}

func (wsh *WebSocketHandler) broadcastingMessages(mm *models.Conversations, message *models.Messanger) {
	if toUserConn, ok := wsh.activeConnections[mm.UserID1]; ok {
		messageToUser := &models.MessangerDTO1{
			Event: "newMessage",
			Data1: models.Data1{
				ConversationID: message.ConversationID,
				SenderID:       message.UserIDSender,
				Content:        message.Message,
				CreatedAt:      message.CreatedAt,
			},
		}
		if err := toUserConn.WriteJSON(messageToUser); err != nil {
			log.Println(err)
		}
	}

	if fromUserConn, ok := wsh.activeConnections[mm.UserID2]; ok {
		messageToSender := &models.MessangerDTO1{
			Event: "newMessage",
			Data1: models.Data1{
				ConversationID: message.ConversationID,
				SenderID:       message.UserIDSender,
				Content:        message.Message,
				CreatedAt:      message.CreatedAt,
			},
		}
		if err := fromUserConn.WriteJSON(messageToSender); err != nil {
			log.Println(err)
		}
	}
}

func (wsh *WebSocketHandler) Conversation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Println("Method not allowed")
		wsh.renderError(w, http.StatusNotFound)
		return
	}

	user := getUserFromContext(r)

	idConversation, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Println(err)
		wsh.renderError(w, http.StatusBadRequest)
		return
	}

	chatHistory, err := wsh.service.Conversation.ConversationHistoryService(idConversation)
	if err != nil {
		log.Println(err)
		wsh.renderError(w, http.StatusInternalServerError)
		return
	}

	userId2 := chatHistory.Conversation.UserID2
	if user.Id == chatHistory.Conversation.UserID2 {
		userId2 = chatHistory.Conversation.UserID1
	}

	wsh.renderPage(w, "chat.html", &models.Chat{
		ConversationID: chatHistory.Conversation.ID,
		User:           user,
		UserID2:        userId2,
		Messages:       chatHistory.Messages,
	})
}

func (wsh *WebSocketHandler) Conversations(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Println("Method not allowed")
		wsh.renderError(w, http.StatusNotFound)
		return
	}

	user := getUserFromContext(r)

	chats, err := wsh.service.Conversation.ConversationsService(user.Id)
	if err != nil {
		log.Println(err)
		wsh.renderError(w, http.StatusInternalServerError)
		return
	}

	wsh.renderPage(w, "chats.html", &models.Chats{
		User:          user,
		Conversations: chats,
	})
}

type Event struct {
	Event string `json:"event"`
}
