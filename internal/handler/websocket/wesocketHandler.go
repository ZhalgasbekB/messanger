package websocket

import (
	"fmt"
	"forum/config"
	"forum/internal/models"
	"forum/internal/service"
	"html/template"
	"io"
	"log"
	"net/http"
)

type WebSocketHandler struct {
	service      *service.Service
	template     *template.Template
	googleConfig config.GoogleConfig
	githubConfig config.GithubConfig
}

func NewWebHandler(service *service.Service, tpl *template.Template, googleCfg config.GoogleConfig, githubCfg config.GithubConfig) *WebSocketHandler {
	return &WebSocketHandler{
		service:      service,
		template:     tpl,
		googleConfig: googleCfg,
		githubConfig: githubCfg,
	}
}

func (wsh *WebSocketHandler) renderError(w http.ResponseWriter, code int) {
	w.WriteHeader(code)

	err := wsh.template.ExecuteTemplate(w, "error.html", struct {
		Code int
		Text string
	}{
		Code: code,
		Text: http.StatusText(code),
	})
	if err != nil {
		log.Printf("ExecuteTemplate:%s\n", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError) // 500
	}
}

func (wsh *WebSocketHandler) renderPage(w http.ResponseWriter, file string, data interface{}) {
	err := wsh.template.ExecuteTemplate(w, file, data)
	if err != nil {
		log.Printf("ExecuteTemplate:%s\n", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError) // 500
	}
}

func (wsh *WebSocketHandler) getUserInfoFromApi(accessToken string, userInfoURL string) ([]byte, error) {
	req, err := http.NewRequest("GET", userInfoURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type conKay string

var KeyUser = conKay("user")

func getUserFromContext(r *http.Request) *models.User {
	user, ok := r.Context().Value(KeyUser).(*models.User) // Use keyUser here
	if !ok {
		return nil
	}
	return user
}
