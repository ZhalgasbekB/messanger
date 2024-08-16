package http

import (
	"encoding/json"
	"fmt"
	"forum/internal/handler/websocket"
	"forum/internal/models"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

func (h *Handler) getUserFromContext(r *http.Request) *models.User {
	user, ok := r.Context().Value(websocket.KeyUser).(*models.User)
	if !ok {
		return nil
	}
	return user
}

func (h *Handler) getVote(voteStr string) (int, error) {
	rx := regexp.MustCompile(`^[^0,+]{1,}\d*$`)
	if !rx.MatchString(voteStr) {
		return 0, fmt.Errorf("incorrect request vote = %s", voteStr)
	}
	vote, err := strconv.Atoi(voteStr)
	if err != nil {
		return 0, err
	}
	if vote != 1 && vote != -1 {
		return 0, fmt.Errorf("incorrect request vote = %d", vote)
	}
	return vote, nil
}

func (h *Handler) getIntFromForm(value string) (int, error) {
	rx := regexp.MustCompile(`^[^0,+,-]{1,}\d*$`)
	if !rx.MatchString(value) {
		return 0, fmt.Errorf("incorrect request vote = %s", value)
	}

	id, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func getValueFromBody(body []byte, key string) string {
	var data map[string]interface{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		return ""
	}
	res, ok := data[key].(string)
	if !ok {
		return ""
	}
	return res
}

func getValueFromURL(body, key string) string {
	valuer, err := url.ParseQuery(body)
	if err != nil {
		return ""
	}
	return valuer.Get(key)
}
