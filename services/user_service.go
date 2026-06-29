package services

import (
	"context"
	"time"

	"github.com/StardustEnigma/AuthGo/dto"
	"github.com/StardustEnigma/AuthGo/models"
	"github.com/StardustEnigma/AuthGo/repository"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(ctx context.Context,request dto.Register)(models.User,error) {
		var user models.User
		user.UserName=request.Username
		passwordHash,err :=bcrypt.GenerateFromPassword([]byte(request.Password),bcrypt.DefaultCost)
		user.Password=string(passwordHash)
		if err != nil {
			return models.User{},err
		}
		user.CreatedAt=time.Now().UTC()
		newUser,err := repository.CreateUser(ctx,user)
		if err != nil {
			return models.User{}, err
		}
		return newUser,nil
}