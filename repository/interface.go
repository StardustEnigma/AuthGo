package repository

import (
	"context"

	"github.com/StardustEnigma/AuthGo/dto"
	"github.com/StardustEnigma/AuthGo/models"
)

type UserRepository interface{
	CreateUser(ctx context.Context,request dto.Register)(models.User,error)
}