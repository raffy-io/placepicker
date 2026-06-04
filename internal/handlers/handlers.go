package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/raffy-io/placepicker/internal/db"
	"github.com/raffy-io/placepicker/ui"
)

type Handler struct {
	Queries *db.Queries
}

func NewHandler(queries *db.Queries) *Handler {
	return &Handler{Queries: queries}
}

func (h *Handler) Homepage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	pollLocations,err := h.Queries.PollSuggestions(ctx)
	if err != nil {
		log.Printf("Data not found: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	availableLocations,err := h.Queries.ListAvailableLocations(ctx)
	if err != nil {
		log.Printf("Data not found: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	dreamLocations,err := h.Queries.ListDreamLocations(ctx)
	if err != nil {
		log.Printf("Data not found: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	
	component := ui.Homepage(pollLocations,availableLocations,dreamLocations)
	templ.Handler(component).ServeHTTP(w,r)
}

func (h *Handler) Add(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		// Or render a Templ error component
		return
	}

	idStr := r.PostFormValue("id")
	if idStr == "" {
		http.Error(w, "Missing location ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	params := db.SetDreamLocationStatusParams{
		IsDreamLocation: 1,
		ID: id,
	}

	err = h.Queries.SetDreamLocationStatus(ctx,params)
	if err != nil {
		// Handle database errors properly
		http.Error(w, "Failed to update database", http.StatusInternalServerError)
		return
	}

	dreamLocations,err := h.Queries.ListDreamLocations(ctx)
	if err != nil {
		log.Printf("Data not found: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	availableLocations,err := h.Queries.ListAvailableLocations(ctx)
	if err != nil {
		log.Printf("Data not found: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}


	component := ui.AddLocationResponse(availableLocations,dreamLocations)
	templ.Handler(component).ServeHTTP(w,r)
	
	
}

func (h *Handler) Remove(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		// Or render a Templ error component
		return
	}

	idStr := r.PostFormValue("id")
	if idStr == "" {
		http.Error(w, "Missing location ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	params := db.SetDreamLocationStatusParams{
		IsDreamLocation: 0,
		ID: id,
	}

	err = h.Queries.SetDreamLocationStatus(ctx,params)
	if err != nil {
		// Handle database errors properly
		http.Error(w, "Failed to update database", http.StatusInternalServerError)
		return
	}

		dreamLocations,err := h.Queries.ListDreamLocations(ctx)
	if err != nil {
		log.Printf("Data not found: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	availableLocations,err := h.Queries.ListAvailableLocations(ctx)
	if err != nil {
		log.Printf("Data not found: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}


	component := ui.RemoveLocationResponse(availableLocations,dreamLocations)
	templ.Handler(component).ServeHTTP(w,r)
	
}

func (h *Handler) PollLocations(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()

	pollLocations,err := h.Queries.PollSuggestions(ctx)
	if err != nil {
		log.Printf("Data not found: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	availableLocations,err := h.Queries.ListAvailableLocations(ctx)
	if err != nil {
		log.Printf("Data not found: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	dreamLocations,err := h.Queries.ListDreamLocations(ctx)
	if err != nil {
		log.Printf("Data not found: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	component := ui.PollSuggestionResponse(pollLocations,availableLocations,dreamLocations)
	templ.Handler(component).ServeHTTP(w,r)

}