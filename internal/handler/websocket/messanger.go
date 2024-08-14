package websocket

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleConnection(conn *websocket.Conn, id int) {
	defer conn.Close()

	// define a witch type of send or initial ???
}

func (wsh *WebSocketHandler) InitialConversation(w http.ResponseWriter, r *http.Request) {
	ws, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer ws.Close()

	id := r.URL.Query().Get("id")
	fmt.Println(id)
	/// TAKE TO GIVE CONNECTION
}

func (wsh *WebSocketHandler) Conversation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Println("Method not allowed")
		wsh.renderError(w, http.StatusNotFound)
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

	wsh.renderPage(w, "", chatHistory) /// ADD SOME LOGIC
}

func (wsh *WebSocketHandler) Conversations(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Println("Method not allowed")
		wsh.renderError(w, http.StatusNotFound)
		return
	}
}

func (wsh *WebSocketHandler) Broadcasting(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Println("Method not allowed")
		wsh.renderError(w, http.StatusNotFound)
		return
	}
}
