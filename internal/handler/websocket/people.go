package websocket

import (
	"database/sql"
	"encoding/json"
	"forum/internal/models"
	"log"
	"net/http"
	"time"
)

func (wsh *WebSocketHandler) ListOfUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Println("Method Not Allowed")
		wsh.renderError(w, http.StatusMethodNotAllowed)
		return
	}

	user := getUserFromContext(r)

	con, err := wsh.service.Conversation.ConversationsService(user.Id)
	if err != nil {
		log.Println(err)
		wsh.renderError(w, http.StatusMethodNotAllowed)
		return
	}

	checkEx := map[int]bool{}
	for _, v := range con {
		needUser := v.UserID1
		if user.Id == v.UserID1 {
			needUser = v.UserID2
		}
		checkEx[needUser] = true
	}

	users, err := wsh.service.PeopleService.ListOfUsersToChatService()
	if err != nil {
		log.Println(err)
		wsh.renderError(w, http.StatusInternalServerError)
		return
	}

	var peopleToNew []models.PeopleNotConnect

	for _, v := range *users { // by pointer i can change the number oof pointer
		if !checkEx[v.Id] && v.Id != user.Id {
			peopleToNew = append(peopleToNew, models.PeopleNotConnect{
				UserID: v.Id,
				Name:   v.Name,
			})
		}
	}

	wsh.renderPage(w, "people.html", &models.PeopleNotConnectDTO{
		User:   user,
		People: &peopleToNew,
	})
}

func (wsh *WebSocketHandler) AddUserChats(w http.ResponseWriter, r *http.Request) { // JUST A RELOAD A PAGE
	if r.Method != http.MethodPost {
		log.Println("Method not allowed")
		wsh.renderError(w, http.StatusMethodNotAllowed)
		return
	}

	user := getUserFromContext(r)

	userCh := models.UserChats{}
	if err := json.NewDecoder(r.Body).Decode(&userCh); err != nil {
		log.Println(err)
		wsh.renderError(w, http.StatusMethodNotAllowed)
		return
	}

	if err := wsh.service.Conversation.ConversationCreateService(&models.Conversations{UserID1: user.Id, UserID2: userCh.Data.User2, CreatedAt: time.Now()}); err != nil {
		log.Println(err)
		wsh.renderError(w, http.StatusInternalServerError)
		return
	}

	WriteJSON(w, 200, &RESP{Ok: true})
}
func (wsh *WebSocketHandler) NowChatCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Println("Method Not Allowed")
		wsh.renderError(w, http.StatusMethodNotAllowed)
		return
	}

	user := getUserFromContext(r)

	userCh := models.UserChats{}
	if err := json.NewDecoder(r.Body).Decode(&userCh); err != nil {
		log.Println(err)
		wsh.renderError(w, http.StatusMethodNotAllowed)
		return
	}

	if err := wsh.service.Conversation.ConversationCreateService(&models.Conversations{UserID1: user.Id, UserID2: userCh.Data.User2, CreatedAt: time.Now()}); err != nil {
		log.Println("Method Not Allowed")
		wsh.renderError(w, http.StatusInternalServerError)
		return
	}
	conversations_id, err := wsh.service.Conversation.ConversationExistService(user.Id, userCh.Data.User2)
	if err != nil && err != sql.ErrNoRows {
		log.Println("Method Not Allowed")
		wsh.renderError(w, http.StatusInternalServerError)
		return
	}

	WriteJSON(w, 200, &RESP_WS{
		Ok:             true,
		ConversationID: conversations_id,
	})
}

type RESP_WS struct {
	Ok             bool `json:"ok"`
	ConversationID int  `json:"conversation_id"`
}
