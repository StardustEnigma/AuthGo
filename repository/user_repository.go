package repository

import (
	"context"

	"github.com/StardustEnigma/AuthGo/db"
	"github.com/StardustEnigma/AuthGo/models"
)

func CreateUser(ctx context.Context,user models.User)(models.User,error){
	query := `INSERT INTO users
				(username,password,created_at)
				VALUES $1,$2,$3
				RETURNING user_id,username,password,created_at`

	var savedUser models.User
	err := db.Db.QueryRowContext(
		ctx,
		query,
		user,
	).Scan(
		&savedUser.UserId,
		&savedUser.UserName,
		&savedUser.Password,
		&savedUser.CreatedAt,
	)
	if err != nil {
		return models.User{},err
	}
	return savedUser,nil
}

func LoginUser(ctx context.Context, username string)(models.User,error){
	query := `SELECT 
				user_id,
				password,
				created_at
				FROM users
				WHERE username = $1`
	var user models.User	
	err := db.Db.QueryRowContext(
		ctx,
		query,
		username,
	).Scan(
		&user.UserId,
		&user.UserName,
		&user.Password,
		&user.CreatedAt,
	)
	if err!=nil {
		return models.User{},err
	}
	return user,nil
}