package handler

import (
	"encoding/json"
	"net/http"

	"github.com/StardustEnigma/AuthGo/dto"
	"github.com/StardustEnigma/AuthGo/middleware"
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

func (h *Handler)LoginUser(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	var loginRequest dto.LoginRequest

	json.NewDecoder(r.Body).Decode(&loginRequest)
	token,err := h.AuthService.LoginUser(ctx,loginRequest)

	if err != nil  {
		http.Error(w,"Bad Request",http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	response := struct{
		Token string `json:"token"`
	}{
		Token: token,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler)GetMe(w http.ResponseWriter,r *http.Request){
	userId := r.Context().Value(middleware.UserIdKey).(int)
	w.Header().Set("Content-Type","application/json")

	response := map[string]interface{}{
		"userId" : userId,
		"message" : "You are authenticated",
	}
	json.NewEncoder(w).Encode(response)
}