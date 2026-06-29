package utils

import (
	"time"

	"github.com/StardustEnigma/AuthGo/models"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserId int `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}
var mySecretKey = []byte("my-Secret-Key")
func GenerateToken(user models.User)(string,error){
	claims := Claims{
		Username: user.UserName,
		UserId: user.UserId,
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