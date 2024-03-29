package handlers

import (
	"goth/.gen/model"
	"goth/internal/logic"
	"goth/internal/templates"
	"log"
	"net/http"
	"sync"
)

func (h *Handler) StancesHandler(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("query")
	query, err := h.store.SaveQuery(text)
	if err != nil {
		log.Fatalf("Error saving query: %v", err)
	}

	stances := []model.Stance{}
	parties, err := h.store.GetParties()
	if err != nil {
		log.Fatalf("Error getting parties: %v", err)
	}

	var wg sync.WaitGroup
	for _, party := range parties {
		wg.Add(1)
		go func(party model.Party) {
			defer wg.Done()

			response, err := logic.CallAgentHook(query.Text, party.Website)
			if err != nil {
				log.Fatalf("Error starting agent: %v", err)
			}

			stance := model.Stance{
				QueryID: *query.ID,
				PartyID: *party.ID,
				Text:    &response,
			}
			err = h.store.SaveStance(stance)
			if err != nil {
				log.Fatalf("Error saving stance: %v", err)
			}
			stances = append(stances, stance)
		}(party)
	}

	wg.Wait()

	c := templates.StanceListResponse(query, parties, stances)
	err = c.Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}

}
