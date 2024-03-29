package handlers

import (
	"net/http"
	"politics/internal/templates"
)

func (h *Handler) AboutHandler(w http.ResponseWriter, r *http.Request) {
	c := templates.About()
	err := templates.Layout(c, "My website").Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
