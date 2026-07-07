package repository

import (
	"context"

	"github.com/StardustEnigma/AuthGo/models"
)

type UserRepository interface{
	GetUser(ctx context.Context,userId int)(models.User,error)
}

func (r *Repository) GetUser(ctx context.Context,userId int)(models.User,error){
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

