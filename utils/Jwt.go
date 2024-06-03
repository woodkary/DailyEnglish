package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var JWTkey = []byte("secret")

type TeamManagerClaims struct {
	jwt.StandardClaims
	ManagerID int
	Team      map[int]string // map[teamID]teamname
}
type UserClaims struct {
	jwt.StandardClaims
	UserID   int
	TeamID   int
	Teamname string
}

func GenerateToken_User(UserID int, TeamID int, Teamname string) (string, error) {
	claims := UserClaims{
		UserID:   UserID,
		TeamID:   TeamID,
		Teamname: Teamname,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "DailyEnglish",
			IssuedAt:  time.Now().Unix(), // token will be valid for 1 hour
			ExpiresAt: time.Now().Add(60 * 60 * time.Second).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTkey)
}
func ParseToken_User(tokenString string) (*UserClaims, error) {
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

func GenerateToken_TeamManager(ManagerID int, team map[int]string) (string, error) {
	temp := time.Now()
	// 使用当前日期和时间创建一个Time对象，时间为凌晨零点
	startOfDay := time.Date(temp.Year(), temp.Month(), temp.Day(), 0, 0, 0, 0, temp.Location())
	// 使用当前日期和时间创建一个Time对象，时间为今天的23:59:59.999999999
	endOfDay := time.Date(temp.Year(), temp.Month(), temp.Day(), 23, 59, 59, 999999999, temp.Location())
	claims := TeamManagerClaims{
		ManagerID: ManagerID,
		Team:      team,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "DailyEnglish",
			IssuedAt:  startOfDay.Unix(), // token will be valid for 1 hour
			ExpiresAt: endOfDay.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTkey)

}

func ParseToken_TeamManager(tokenString string) (*TeamManagerClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TeamManagerClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JWTkey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*TeamManagerClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
