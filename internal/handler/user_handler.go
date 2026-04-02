package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Chintukr2004/auth-service/internal/middleware"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Protected route working",
		"user_id": userID,
	})
}

func (h *UserHandler) AdminOnly(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Welcome Admin",
	})
}
