package handler

import (
	"encoding/json"
	"net/http"
	"web-analyzer/internal/service"

	"github.com/gorilla/mux"
)

type Handler struct {
	service service.AnalyzerService
	router  *mux.Router
}

func NewHandler(s service.AnalyzerService) *Handler {
	h := &Handler{service: s, router: mux.NewRouter()}
	h.router.HandleFunc("/api/analyze", h.handleAnalyze).Methods("POST")
	return h
}

func (h *Handler) Router() *mux.Router {
	return h.router
}

func (h *Handler) handleAnalyze(w http.ResponseWriter, r *http.Request) {
	var input struct {
		URL string `json:"url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.service.AnalyzeURL(input.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
