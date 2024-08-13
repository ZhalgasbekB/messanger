package handler

import (
	"database/sql"
	"log"
	"net/http"

	"forum/internal/render"
)

// GET
func (h *Handler) filterPostsGET(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/filterposts" {
		log.Printf("filterPostsGET:StatusNotFound:%s\n", r.URL.Path)
		h.renderError(w, http.StatusNotFound) // 404
		return
	}

	if r.Method != http.MethodGet {
		log.Printf("filterPostsGET:StatusMethodNotAllowed:%s\n", r.Method)
		h.renderError(w, http.StatusMethodNotAllowed) // 405
		return
	}
	category := r.URL.Query().Get("category")

	posts, err := h.service.GetByCategory(category)
	if err != nil {
		log.Printf("filterPostsGET:GetByCategory:%s\n", err.Error())
		if err == sql.ErrNoRows {
			h.renderError(w, http.StatusBadRequest) // 400
			return
		}
		h.renderError(w, http.StatusInternalServerError) // 500
		return
	}
	categories, err := h.service.Category.GetAll()
	if err != nil {
		log.Printf("filterPostsGET:GetAll:%s\n", err.Error())
		h.renderError(w, http.StatusInternalServerError) // 500
		return
	}

	user := h.getUserFromContext(r)

	h.renderPage(w, "home.html", &render.MainData{
		User:       user,
		Posts:      posts,
		Categories: categories,
	})
}
