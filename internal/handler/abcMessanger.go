package handler

import (
	"log"
	"net/http"
)

func (h *Handler) Conversations(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Printf("adminGET:StatusNotFound:%s\n", r.URL.Path)
		h.renderError(w, http.StatusNotFound)
		return
	}
}

func (h *Handler) ConversationByID(w http.ResponseWriter, r *http.Request) {
}
