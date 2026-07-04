package handler

import (
	"encoding/json"
	"net/http"
	"github.com/StardustEnigma/AuthGo/dto"
	"github.com/StardustEnigma/AuthGo/services"
)
type Handler struct{
	AuthService services.AuthService
}

func (h *Handler)RegisterUser(w http.ResponseWriter,r *http.Request){
	ctx:=r.Context()

	var register dto.Register
	json.NewDecoder(r.Body).Decode(&register)
	user,err := h.AuthService.CreateUser(ctx,register)
	if err != nil {
		http.Error(w,"Bad Request",http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(user)
}