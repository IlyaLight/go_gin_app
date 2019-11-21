package util

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go_gin_app/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// GenerateToken generate tokens used for auth
func GenerateToken(user *models.User) (*string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(viper.GetInt("app.TokenTimeLive")) * time.Hour)
	claims := Claims{
		user.Username,
		user.Password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}
	// Устанавливаем набор параметров для токена
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Подписываем токен нашим секретным ключем
	token, err := tokenClaims.SignedString([]byte(viper.GetString("app.JwtSecret")))
	return &token, errors.WithStack(err)
}

// ParseToken parsing token
func CheckToken(token *string) (bool, error) {
	tokenClaims, err := jwt.ParseWithClaims(*token, &Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(viper.GetString("app.JwtSecret")), nil
		})
	if tokenClaims != nil {
		if _, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return true, nil
		}
	}

	return false, err
}
