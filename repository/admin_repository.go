package repository

import (
	"context"
	"errors"

	"github.com/StardustEnigma/AuthGo/db"
	"github.com/StardustEnigma/AuthGo/models"
)



func GetAllUsers(ctx context.Context,offset int, limit int)([]models.User,error){
	query := `SELECT 
				user_id,
				username,
				role,
				is_active,
				is_verified,
				created_at,
				is_suspended
				FROM users
				LIMIT $1
				OFFSET $2`

	rows,err := db.Db.QueryContext(
		ctx,
		query,
		limit,
		offset,
	)
	if err != nil {
		return[]models.User{} ,err
	}
	defer rows.Close()
	var users []models.User
	for rows.Next(){
		var user models.User

		err := rows.Scan(
			&user.UserId,
			&user.UserName,
			&user.Role,
			&user.IsActive,
			&user.IsVerified,
			&user.CreatedAt,
			&user.IsSuspended,
		)
		if err != nil {
			return nil,err
		}
		users = append(users, user)
	}
	if err = rows.Err() ; err != nil {
		return nil,err
	}
	return users,nil
}

func GetUser(ctx context.Context,userId int)(models.User,error){
	query := `SELECT 
			username,
			role,
			password,
			is_active,
			is_verified,
			created_at
			FROM users
			WHERE user_id=$1
			`

	var user models.User
	err := db.Db.QueryRowContext(
		ctx,
		query,
		userId,
	).Scan(
		&user.UserName,
		&user.Role,
		&user.Password,
		&user.IsActive,
		&user.IsVerified,
		&user.CreatedAt,
	)
	if err != nil {
		return models.User{},err
	}
	return user,err
}

func DeActivateUser(ctx context.Context,userid int)error{
	query := `
			UPDATE users
			SET is_active = false
			WHERE user_id = $1`
	
	msg,err := db.Db.ExecContext(
		ctx,
		query,
		userid,
	)
	if err != nil {
		return err
	}
	rowsAffected,err := msg.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected==0 {
		return errors.New("user does not exist")
	}
	return nil
}

func SuspendUser(ctx context.Context,userId int)error{
	query:= `UPDATE users
			SET is_suspended = true
			WHERE user_id = $1
			`
	msg,err := db.Db.ExecContext(
		ctx,
		query,
		userId,
	)
	if err != nil{
		return err
	}
	rowsAffected,err := msg.RowsAffected()
	if err !=nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("User does not exist")
	}
	return nil
}