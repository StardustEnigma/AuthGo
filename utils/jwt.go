package utils

import (
	"errors"
	"time"

	"github.com/StardustEnigma/AuthGo/models"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserId int `json:"user_id"`
	Username string `json:"username"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}
var mySecretKey = []byte("my-Secret-Key")
func GenerateToken(user models.User)(string,error){
	claims := Claims{
		Username: user.UserName,
		UserId: user.UserId,
		Role: string(user.Role),
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt: jwt.NewNumericDate(
				time.Now(),
			),
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(24*time.Hour),
			),
		},
	}
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)
	tokenString ,err := token.SignedString(
		mySecretKey,
	)
	if err!= nil {
		return "",err
	}
	return tokenString,nil
}

func ValidateTokens(tokenString string)(*Claims,error){

	claims := &Claims{}

	token,err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface {}, error) {
			return mySecretKey,nil
		},
	)
	if err != nil{
		return nil,err
	}
	if _, ok :=token.Method.(*jwt.SigningMethodHMAC); !ok{
		return nil, errors.New("Unexpected Signing Method")
	}
	if !token.Valid{
		return  nil, errors.New("Invalid token")
	}
	return claims,nil
}