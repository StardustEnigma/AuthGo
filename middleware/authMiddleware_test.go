package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/StardustEnigma/AuthGo/models"
	"github.com/StardustEnigma/AuthGo/utils"
)



func TestAuthMiddleware( t *testing.T){

	validToken, _ := utils.GenerateToken(models.User{
		UserId:   1,
		UserName: "testuser",
	})
	tests:= []struct{
		name string
		authHeader string
		expectedStatus int
	}{
		{
			name: "Valid header",
			authHeader:  "Bearer " + validToken,
			expectedStatus: http.StatusOK,
		},
		{
			name: "missing auth header",
			authHeader: "",
			expectedStatus: http.StatusUnauthorized,	
		},
		{
			name: "Invalid Bearer format",
			authHeader: "Invalid Format",
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name: "Invalid token",
			authHeader: "Bearer Invalid-Token-xyz",
			expectedStatus: http.StatusUnauthorized,
		},
	}

	for _,tt := range tests{
		t.Run(tt.name,func(t *testing.T) {
			req := httptest.NewRequest("GET","/me",nil)
			if tt.authHeader != ""{
				req.Header.Set("Authorization",tt.authHeader)
			}
			w := httptest.NewRecorder()
			middleware := AuthMiddleWare(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			}))
			middleware.ServeHTTP(w,req)
			if w.Code !=tt.expectedStatus {
				t.Errorf("Expected staus %d got %d",tt.expectedStatus,w.Code)
			}
		})
	}
} 