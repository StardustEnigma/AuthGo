package handler

import (
	"encoding/json"
	"net/http"
	"github.com/StardustEnigma/AuthGo/dto"
	"github.com/StardustEnigma/AuthGo/services"
)

func RegisterUser(w http.ResponseWriter,r *http.Request){
	ctx:=r.Context()

	var register dto.Register
	json.NewDecoder(r.Body).Decode(&register)
	user,err := services.CreateUser(ctx,register)
	if err != nil {
		http.Error(w,"Bad Request",http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(user)
}