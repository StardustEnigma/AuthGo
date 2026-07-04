package routes

import (
	"github.com/StardustEnigma/AuthGo/db"
	"github.com/StardustEnigma/AuthGo/handler"
	"github.com/StardustEnigma/AuthGo/repository"
	"github.com/StardustEnigma/AuthGo/services"
	"github.com/go-chi/chi/v5"
)

func Routes() *chi.Mux{
	r := chi.NewRouter()
	repo := &repository.Repository{DB :db.Db}
	authService := &services.AuthService{Repo:repo}

	handler := &handler.Handler{AuthService: *authService}
	r.Post("/register",handler.RegisterUser)
	return r
}