package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/StardustEnigma/AuthGo/utils"
)
type contextKey string
const userIdKey contextKey = "userId"

func AuthMiddleWare(next http.Handler)http.Handler{
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader =="" {
				http.Error(w,"Missing Authorization Headers",http.StatusUnauthorized)
				return 
			}

			tokenString := strings.TrimPrefix(
				authHeader,
				"Bearer ",
			)
			if tokenString == authHeader {
				http.Error(w,"Invalid Headers",http.StatusUnauthorized)
				return 
			}
			claims,err := utils.ValidateTokens(
				tokenString,
			)
			if err !=  nil {
				http.Error(w,"Invalid Token",http.StatusUnauthorized)
				return 
			}
			ctx := context.WithValue(
				r.Context(),
				userIdKey,
				claims.UserId,
			)
			r=r.WithContext(
				ctx,
			)
			next.ServeHTTP(
				w,
				r,
			)

		},
	)
}