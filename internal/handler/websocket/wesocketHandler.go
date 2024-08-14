package websocket

import (
	"forum/config"
	"forum/internal/service"
	"text/template"
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
