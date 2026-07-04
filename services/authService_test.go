package services

import (
	"context"
	"errors"
	"testing"

	"github.com/StardustEnigma/AuthGo/dto"
	"github.com/StardustEnigma/AuthGo/models"
)

type MockAuthRepository struct{
	shouldErr bool
	savedUser models.User
}

func (m *MockAuthRepository) CreateUser(ctx context.Context,user models.User)(models.User,error){
	if m.shouldErr{
		return models.User{},errors.New("Database error")
	}
	user.UserId=1
	return user,nil
}

func TestCreateUser(t *testing.T){
	tests := []struct {
		name string
		request dto.Register
		mockError bool
		shouldErr bool
	}{
		{
		name : "Valid Registeration",
		request :dto.Register{
			Username: "john",
			Password: "password123",
		},
		mockError :false,
		shouldErr:false,
	},
	{
		name : "Database Error",
		request : dto.Register{
			Username :"john",
			Password: "password123",
		},
		mockError: true,
		shouldErr:true,
	},

}
	for _,tt := range tests{
		t.Run(tt.name,func(t *testing.T) {
			mockRepo := &MockAuthRepository{
				shouldErr: tt.mockError,
			}
			service := NewAuthService(mockRepo)
			_,err := service.CreateUser(context.Background(),tt.request)

			if tt.shouldErr {
				if err == nil{
					t.Error("Expected Error but got none")
				}
			}else{
				if err != nil{
					t.Errorf("Unexpected error %v",err)
				}
			}

		})
	}
}