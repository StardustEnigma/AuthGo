package routes

import (
	"github.com/StardustEnigma/AuthGo/handler"
	"github.com/go-chi/chi/v5"
)

func Routes() *chi.Mux{
	r := chi.NewRouter()

	r.Post("/register",handler.RegisterUser)
	return r
}