package handler

import (
	"encoding/json"
	"net/http"
	model "github.com/medvedev-v/radiocontest/internal/model"
	repository "github.com/medvedev-v/radiocontest/internal/repository"
)

type EchoHandler struct {
	repo *repository.UserRepository
}

func NewEchoHandler() *EchoHandler {
	return &EchoHandler{}
}

func (h *EchoHandler) Echo(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	id, err := h.repo.Create(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}
