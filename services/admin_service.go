package services

import (
	"context"
	"errors"
	"strconv"

	"github.com/StardustEnigma/AuthGo/models"
	"github.com/StardustEnigma/AuthGo/repository"
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

func GetAllUsers(ctx context.Context,pagestr string,limitstr string)([]models.User,error){
	 page,err := strconv.Atoi(pagestr)
	 if err != nil || page <=0{
		
		return nil,errors.New("Error in getting page")
	}
	 limit,err := strconv.Atoi(limitstr)
	if err != nil {
		return nil,errors.New("Error with limit")
	}
	offset := (page-1)*limit
	users,err :=repository.GetAllUsers(ctx,offset,limit) 
	if err != nil{
		return nil,err
	}
	return users,nil

}

func GetUser(ctx context.Context,userId int)(models.User,error){
	user,err := repository.GetUser(ctx,userId)
	if err != nil {
		return models.User{},err
	}
	return user,nil
}

func DeActivate(ctx context.Context,UserId int)error{
	user,err := repository.GetUser(ctx,UserId)
	if err != nil {
		return err
	}
	if(user.IsActive){
		return errors.New("User already active")	
	}
	ok := repository.DeActivateUser(ctx,UserId)

	if ok != nil {
		return ok
	}
	return nil
}

func BanUser(ctx context.Context,userId int)error{
	user,err := repository.GetUser(ctx,userId)
	if err!=nil {
		return err
	}
	if user.IsSuspended{
		return errors.New("User already Suspeneded")
	}
	ok := repository.SuspendUser(ctx,userId)
	if ok != nil {
		return err
	}
	return nil
}

// func ChangeUserRoles(ctx context.Context,userId int)error{
	
// }
