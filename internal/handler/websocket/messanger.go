package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleConnection(conn *websocket.Conn, id int) {
	defer conn.Close()

	// define a witch type of send or initial ? 
	
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
	/// TAKE TO GIVE CONN
}

func (wsh *WebSocketHandler) Conversation(w http.ResponseWriter, r *http.Request) {
}

func (wsh *WebSocketHandler) Broadcasting(w http.ResponseWriter, r *http.Request) {
}
