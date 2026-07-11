package services

import (
	"context"
	"time"

	"github.com/StardustEnigma/AuthGo/dto"
	"github.com/StardustEnigma/AuthGo/models"
	"github.com/StardustEnigma/AuthGo/repository"
	"golang.org/x/crypto/bcrypt"
)

// AdminService handles privileged administrative operations.
//
// Responsibilities:
// - Retrieve all users
// - View user details
// - Activate or deactivate accounts
// - Suspend or ban users
// - Change user roles
// - Manage permissions
// - Review system activity
// - Perform moderation tasks
// - Execute administrator-only business operations

type AdminService interface{
	CreateAdmin(ctx context.Context,request dto.Register)(models.User,error)
}

type adminService struct{
	Repo repository.AdminRepository
}

func NewAdminService(repo repository.AdminRepository)AdminService{
	return &adminService{
		Repo: repo,
	}
}

func(s *adminService)CreateAdmin(ctx context.Context,request dto.Register)(models.User,error){
	var user models.User
	user.CreatedAt=time.Now().UTC()
	user.Role=models.Admin
	user.IsVerified=true
	user.IsActive=true
	user.UserName=request.Username

	hashedPassword,err := bcrypt.GenerateFromPassword([]byte(request.Password),bcrypt.DefaultCost)
	if err != nil {
		return models.User{},err
	}
	user.Password=string(hashedPassword)
}