package websocket

import (
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
	defer ws.Close()

	user := getUserFromContext(r)
	if user.Id == 0 {
		ws.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Unauthorized"))
		ws.Close()
		return
	}

	id := r.URL.Query().Get("id")
	fmt.Println(id, user)
	/// TAKE TO GIVE CONNECTION

	go wsh.handleConnection(ws, user.Id)
}

func (wsh *WebSocketHandler) handleConnection(conn *websocket.Conn, id int) {
	defer conn.Close()

	for {
		var messages *models.MessangerDTO
		if err := conn.ReadJSON(&messages); err != nil {
			log.Println(err)
			break
		}

		switch messages.Event {
		case "initiateConversation":
			wsh.connectionChat(conn, id, int(messages.Data.RecipientID))
		case "sendMessage":
			wsh.sendMessageW(conn, *messages, 1)
		}
	}
}

func (wsh *WebSocketHandler) connectionChat(conn *websocket.Conn, id int, re_id int) {
	if err := wsh.service.Conversation.ConversationCreateService(&models.Conversations{UserID1: id, UserID2: re_id, CreatedAt: time.Now()}); err != nil {
		log.Println(err)
		return
	}
	// NEED SOME LOGIC
	// BUT WHAT ???? BROO
}

func (wsh *WebSocketHandler) sendMessageW(conn *websocket.Conn, m models.MessangerDTO, sender int) {
	if err := wsh.service.Conversation.SendMessageService(models.Messanger{ConversationID: int(m.Data.ConversationID), UserIDSender: sender, Message: m.Data.Content, CreatedAt: time.Now()}); err != nil {
		log.Println(err)
		return
	}
	// NEED SOME LOGIC
}

func (wsh *WebSocketHandler) broadcasting(conn *websocket.Conn) {
}

func (wsh *WebSocketHandler) Conversation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Println("Method not allowed")
		wsh.renderError(w, http.StatusNotFound)
		return
	}

	user := getUserFromContext(r)
	fmt.Println(user)
	if user.Id == 0 || user == nil {
		wsh.renderError(w, http.StatusBadRequest)
		return
	}

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

	var userId2 int
	if user.Id == chatHistory.Conversation.UserID2 {
		userId2 = chatHistory.Conversation.UserID1
	} else {
		userId2 = chatHistory.Conversation.UserID2
	}
	fmt.Println(userId2, chatHistory)

	wsh.renderPage(w, "chat.html", &models.Chat{
		ConversationID: chatHistory.Conversation.ID,
		User: user,
		UserID2:  userId2,
		Messages: chatHistory.Messages,
	}) /// ???
}

func (wsh *WebSocketHandler) Conversations(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Println("Method not allowed")
		wsh.renderError(w, http.StatusNotFound)
		return
	}

	chats, err := wsh.service.Conversation.ConversationsService()
	if err != nil {
		log.Println(err)
		wsh.renderError(w, http.StatusInternalServerError)
		return
	}

	wsh.renderPage(w, "chats.html", chats)
}
