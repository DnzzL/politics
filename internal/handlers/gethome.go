package handlers

import (
	"net/http"
	"politics/internal/templates"
)

func (h *Handler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	history, err := h.store.GetQueries()
	if err != nil {
		http.Error(w, "Error getting stances", http.StatusInternalServerError)
		return
	}

	c := templates.Index(history)
	err = templates.Layout(c, "Politics").Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
