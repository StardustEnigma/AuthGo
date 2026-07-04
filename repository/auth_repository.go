package repository

import (
	"context"
	"database/sql"

	"github.com/StardustEnigma/AuthGo/db"
	"github.com/StardustEnigma/AuthGo/models"
)
type Repository struct{
	DB *sql.DB
}
type AuthRepository interface{
	CreateUser(ctx context.Context,user models.User)(models.User,error)
	LoginUser(ctx context.Context,username string)(models.User,error)
}

func(r *Repository) CreateUser(ctx context.Context,user models.User)(models.User,error){
	query := `INSERT INTO users
				(username,password,created_at)
				VALUES ($1,$2,$3)
				RETURNING user_id,username,password,created_at`

	var savedUser models.User
	err := r.DB.QueryRowContext(
		ctx,
		query,
		user.UserName,
		user.Password,
		user.CreatedAt,
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

func(r *Repository) LoginUser(ctx context.Context, username string)(models.User,error){
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