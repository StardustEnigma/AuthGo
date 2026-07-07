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
type AuthService interface{
	CreateUser(ctx context.Context, request dto.Register) (models.User, error)
	LoginUser(ctx context.Context, request dto.LoginRequest)(string,error)
}
type authService struct{
	Repo repository.AuthRepository
}
func NewAuthService(repo repository.AuthRepository)AuthService{
	return &authService{
		Repo: repo,
	}
}

func(s *authService) CreateUser(ctx context.Context,request dto.Register)(models.User,error) {
		var user models.User

		user.UserName=request.Username
		hashedPassword,err := bcrypt.GenerateFromPassword([]byte(request.Password),bcrypt.DefaultCost)
		if err != nil {
			return models.User{},err
		}
		user.Password=string(hashedPassword)
		user.CreatedAt=time.Now().UTC()
		user.Role=models.Users
		savedUser,err := s.Repo.CreateUser(ctx,user)
		if err != nil {
			return models.User{},err
		} 
		return savedUser,nil

}

func (s *authService) LoginUser(ctx context.Context,request dto.LoginRequest)(string,error){
	username := request.Username
	user,err :=s.Repo.LoginUser(ctx,username)
	if err!= nil{
		return "",err
	}
	if user.IsSuspended || !user.IsActive {
		return "",errors.New("Plz contact admin for login")
	}
	ok := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(request.Password))
	if ok != nil{
		return "",errors.New("incorrect password")
	}
	token,err := utils.GenerateToken(user)
	if err!= nil {
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

