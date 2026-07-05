package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/StardustEnigma/AuthGo/dto"
	"github.com/StardustEnigma/AuthGo/models"
)

type MockAuthService struct{
	shouldErr bool
	user models.User
}

func (m *MockAuthService) CreateUser(ctx context.Context,request dto.Register)(models.User,error){
	if m.shouldErr {
		return models.User{},errors.New("service error")
	}
	return models.User{
		UserId: 1,
		UserName: request.Username,
		Role: models.Users,
		CreatedAt: time.Now().UTC(),
	},nil
}


func (m *MockAuthService) LoginUser(ctx context.Context,request dto.LoginRequest)(string,error){
	if m.shouldErr{
		return " ",errors.New("Service error")
	}
	return "mock-jwt-token",nil
}

func TestRegisterUser(t *testing.T){
	tests := []struct {
		name string
		request dto.Register
		mockError bool
		expectedStatus int 
	}{
		{
			name: "Valid registeration",
			request: dto.Register{
				Username: "john",
				Password: "password123",
			},
			mockError: false,
			expectedStatus: http.StatusOK,
		},
		{
			name: "service error",
			request: dto.Register{
				Username: "john",
				Password: "password123",
			},
			mockError: true,
			expectedStatus: http.StatusBadRequest,
		},
	}
	for _,tt := range tests{
		t.Run(tt.name,func (t *testing.T)  {
			body,_:=json.Marshal(tt.request)
			req := httptest.NewRequest("POST","/regsiter",bytes.NewReader(body))
			
			w := httptest.NewRecorder()
			mockService := &MockAuthService{
				shouldErr: tt.mockError,
				user: models.User{
					UserId: 1,
					UserName: tt.request.Username,
				},
			}
			handler := &Handler{AuthService: mockService}
			handler.RegisterUser(w,req)

			if w.Code!= tt.expectedStatus {
				t.Errorf("Expected status %d got %d",tt.expectedStatus,w.Code)
			}
		})
	}
}

func TestLoginUser(t *testing.T){
	tests := []struct{
		name string
		request dto.LoginRequest
		mockError bool
		expectedStatus int

	}{
		{
			name:"Valid login",
			request :dto.LoginRequest{
				Username: "john123",
				Password: "password123",
			},
			mockError : false,
			expectedStatus: http.StatusOK,
		},
		{
			name: "invalid credentials",
			request: dto.LoginRequest{
				Username: "john123",
				Password: "incorrect-password",
			},
			mockError: true,
			expectedStatus: http.StatusBadRequest,
		},
	}
	for _,tt := range tests{
		t.Run(tt.name,func(t *testing.T) {
			body,_:= json.Marshal(tt.request)
			req := httptest.NewRequest("POST","/login",bytes.NewReader(body))
			w := httptest.NewRecorder()
			mockService := &MockAuthService{
				shouldErr: tt.mockError,
			}
			handler := &Handler{AuthService: mockService}
			handler.LoginUser(w,req)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d got %d",tt.expectedStatus,w.Code)
			}
		})
	}
}
