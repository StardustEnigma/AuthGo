package repository

import (
	"context"

	"github.com/StardustEnigma/AuthGo/models"
)


type AdminRepository interface{
	CreateAdmin(ctx context.Context,user models.User)(models.User,error)
}

func (r *Repository)CreateAdmin(ctx context.Context,user models.User)(models.User,error){
	
}