package repository

import (
	"context"
	"errors"

	"github.com/StardustEnigma/AuthGo/models"
)

type UserRepository interface{
	GetUserById(ctx context.Context,userId int)(models.User,error)
	GetUserByUsername(ctx context.Context,username string)(models.User,error)
}

func (r *Repository) GetUserById(ctx context.Context,userId int)(models.User,error){
	query :=`SELECT 
			username,
			created_at,
			role,
			is_active,
			is_suspended
			FROM users
			WHERE user_id = $1`

	var user models.User
	err := r.DB.QueryRowContext(
		ctx,
		query,
		userId,
	).Scan(
		&user.UserName,
		&user.CreatedAt,
		&user.Role,
		&user.IsActive,
		&user.IsSuspended,
	)
	if err != nil {
		return models.User{},nil
	}
	return user,nil
}

func (r *Repository)GetUserByUsername(ctx context.Context, username string)(models.User,error){
	query :=`SELECT 
			user_id,
			created_at,
			role,
			is_active,
			is_suspended
			FROM users
			WHERE username = $1`
	var user models.User

	err := r.DB.QueryRowContext(
		ctx,
		query,
		username,
	).Scan(
		&user.UserId,
		&user.CreatedAt,
		&user.Role,
		&user.IsActive,
		&user.IsSuspended,
	)
	if err != nil {
		return models.User{},errors.New("Cannot find user")
	}
	return user,nil
}
