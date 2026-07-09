package services

import (
	"context"
	"errors"

	"github.com/StardustEnigma/AuthGo/dto"
	"github.com/StardustEnigma/AuthGo/models"
	"github.com/StardustEnigma/AuthGo/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface{
	GetUserProfile(ctx context.Context,userId int)(models.User,error)
	UpdateUserProfile(ctx context.Context,userId int,request dto.UpdateRequest)(string,error)
}

type userService struct{
	Repo repository.UserRepository
}

func NewUserService(repo repository.Repository)UserService{
	return &userService{
		Repo: &repo,
	}
}

func (r *userService) GetUserProfile(ctx context.Context,userId int)(models.User,error){
	user,err := r.Repo.GetUserById(ctx,userId)
	if err != nil {
		return models.User{},err
	}
	if user.IsSuspended || !user.IsActive {
		return  models.User{},errors.New("user is not active or is suspended")
	}

	return  user, nil
}

func (r *userService) UpdateUserProfile(ctx context.Context,userId int,request dto.UpdateRequest)(string ,error){
	user,err:= r.Repo.GetUserByUsername(ctx,request.OldUsername)
	if err != nil {
		return "",errors.New("Cannot find the user")
	}
	if user.IsSuspended || !user.IsActive {
		return "",errors.New("User is either suspended or not verified")
	}
	passwordCheck := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(request.Password))
	if passwordCheck != nil {
		return "",errors.New("Incorrect Passowrd")
	}
	update := r.Repo.UpdateUsername(ctx,user.UserId,request.NewUsername)
	if update != nil {
		return "",update
	}
	return "Updated Username Successfully",nil
}