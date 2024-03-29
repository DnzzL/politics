package handlers

import (
	"goth/internal/templates"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) HistoryHandler(w http.ResponseWriter, r *http.Request) {
	queryID := chi.URLParam(r, "queryID")
	queryIDInt32, err := strconv.ParseInt(queryID, 10, 32)
	if err != nil {
		http.Error(w, "Invalid query ID", http.StatusBadRequest)
		return
	}

	history, err := h.store.GetStancesForQuery(int32(queryIDInt32))
	if err != nil {
		http.Error(w, "Error getting stances", http.StatusInternalServerError)
		return
	}

	c := templates.StanceListFromHistory(history)
	err = c.Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
