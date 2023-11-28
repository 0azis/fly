package pkg

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// KEY Слово-секрет, нужен для расшифровки токена
var KEY = []byte("secret")

var TokenTimeAccess int64 = 10000

// CreateAccessToken Метод создания access токена
func CreateAccessToken(userId int) (string, error) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		// Создаем payload структуру
		"userId":      userId,
		"expiredTime": time.Now().Unix() + TokenTimeAccess,
	}).SignedString(KEY)
	return token, err
}

// GetIdentity Расшифровываем токен и получаем из него данные (identity)
func GetIdentity(token string) (int, int64, error) {
	identity, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return KEY, nil
	})

	if err != nil {
		return 0, 0, err
	}

	payload := identity.Claims.(jwt.MapClaims)
	userId := int(payload["userId"].(float64))
	exp := int64(payload["expiredTime"].(float64))

	return userId, exp, nil
}
