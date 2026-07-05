package services

import (
	"context"
	"errors"

	"testing"

	"github.com/StardustEnigma/AuthGo/dto"
	"github.com/StardustEnigma/AuthGo/models"
	"golang.org/x/crypto/bcrypt"
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

func (m *MockAuthRepository) LoginUser(ctx context.Context,username string)(models.User,error){
	if m.shouldErr{
		return models.User{},errors.New("Database error")
	}
	return m.savedUser,nil
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

func TestLoginUser(t *testing.T){
	hashedPassword ,_ := bcrypt.GenerateFromPassword([]byte("password123"),bcrypt.MinCost)
	tests :=[]struct{
		name string
		request dto.LoginRequest
		mockErr bool
		mockUser models.User
		shouldErr bool
	}{
		{
			name: "Valid Login",
			request: dto.LoginRequest{
				Username: "john123",
				Password: "password123",
			},
			mockErr: false,
			mockUser: models.User{
				UserName: "john123",
				Password: string(hashedPassword),
				IsActive: true,
			},
			shouldErr: false,
		},
		{
			name: "Wrong Password",
			request: dto.LoginRequest{
				Username: "john123",
				Password: "WrongPassword",
			},
			mockErr: false,
			mockUser: models.User{
				UserName: "john123",
				Password: string(hashedPassword),
				IsActive: true,
			},
			shouldErr: true,
		},
		{
			name: "Suspendend User",
			request: dto.LoginRequest{
				Username: "badGuy",
				Password: "password123",
			},
			mockErr: false,
			mockUser: models.User{
				UserName: "badGuy",
				Password: string(hashedPassword),
				IsActive: true,
				IsSuspended: true,
			},
			shouldErr: true,
		},
		{
			name: "Database Error",
			request: dto.LoginRequest{
				Username: "ghost",
				Password: "password123",
			},
			mockErr: true,
			mockUser: models.User{},
			shouldErr: true,
		},

	}
	for _,tt := range tests{
		t.Run(tt.name,func(t *testing.T) {
			mockRepo := &MockAuthRepository{
				shouldErr: tt.mockErr,
				savedUser: tt.mockUser,
			}
			service := NewAuthService(mockRepo)

			token,err := service.LoginUser(context.Background(),tt.request)

			if tt.shouldErr {
				if err == nil{
					t.Error("Expected error but got none")
				}
				if token != ""{
					t.Errorf("Expected empty token but got %v",token)
				}
			}else{
				if err != nil {
					t.Errorf("Unexpected Error %v",err)
				}
				if token=="" {
					t.Error("Expected a token but got null")
				}
			}
		})
	}
}