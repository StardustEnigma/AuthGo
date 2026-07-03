package services

import (
	"context"
	"errors"
	"time"

	"github.com/StardustEnigma/AuthGo/dto"
	"github.com/StardustEnigma/AuthGo/models"
	"github.com/StardustEnigma/AuthGo/repository"
	"github.com/StardustEnigma/AuthGo/utils"
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

func LoginUser(ctx context.Context, request dto.LoginRequest)(string,error){
	user,err := repository.LoginUser(ctx,request.Username)
	if err!= nil {
		return "",err
	}
	if !user.IsActive {
		return "",errors.New("User not active ")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(request.Password))
	if err != nil {
		return "",errors.New("Invalid Password")
	}
	token,err := utils.GenerateToken(user)
	if err != nil {
		return "",err
	}
	return token,nil
}

// AuthService handles authentication and identity management.
//
// Responsibilities:
// - User registration
// - User login
// - Password hashing and verification
// - JWT access token generation
// - Refresh token issuance and rotation
// - Logout and token revocation
// - Email verification workflow
// - Password reset functionality
// - Account activation checks
// - Authentication-related business rules