package handlers

import (
	"fmt"
	"log"
	"net/http"
	"placepicker/internal/db"
)

type Handler struct {
	Queries *db.Queries
}

func NewHandler(queries *db.Queries) *Handler {
	return &Handler{Queries: queries}
}

func (h *Handler) GetAvailableLocations(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	data, err := h.Queries.ListAvailableLocations(ctx)
	if err != nil {
		log.Printf("Data not found: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	fmt.Println(data)

}
