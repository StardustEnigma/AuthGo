package repository

import (
	"context"

	"github.com/StardustEnigma/AuthGo/models"
)


type AdminRepository interface{
	CreateAdmin(ctx context.Context,user models.User)(models.User,error)
}

func (r *Repository)CreateAdmin(ctx context.Context,user models.User)(models.User,error){
	query := `INSERT INTO users
				(username,password,created_at,role)
				VALUES ($1,$2,$3,$4)
				RETURNING user_id,username,password,created_at,role`
	var savedAdmin models.User
	
}