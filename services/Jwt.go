package services

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var JWTkey = []byte("secret")

type UserClaims struct {
	jwt.StandardClaims
	UserName string
	TeamName string
}

func GenerateToken(userName string, teamName string) (string, error) {
	claims := UserClaims{
		UserName: userName,
		TeamName: teamName,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "DailyEnglish",
			IssuedAt:  time.Now().Unix(), // token will be valid for 1 hour
			ExpiresAt: time.Now().Add(60 * 60 * time.Second).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTkey)

}

func ParseToken(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JWTkey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
