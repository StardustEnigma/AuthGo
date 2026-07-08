package services

import (
	"context"
	"errors"

	"github.com/StardustEnigma/AuthGo/models"
	"github.com/StardustEnigma/AuthGo/repository"
)

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

type UserService interface{
	GetUserProfile(ctx context.Context,userId int)(models.User,error)
}

type userRepository struct{
	Repo repository.UserRepository
}

func NewUserService(repo repository.Repository)UserService{
	return &userRepository{
		Repo: &repo,
	}
}

func (r *userRepository) GetUserProfile(ctx context.Context,userId int)(models.User,error){
	user,err := r.Repo.GetUser(ctx,userId)
	if err != nil {
		return models.User{},err
	}
	if user.IsSuspended || !user.IsActive {
		return  models.User{},errors.New("user is not active or is suspended")
	}

	return  user, nil
}