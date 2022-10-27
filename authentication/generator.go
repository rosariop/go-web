package authentication

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateJwt(username string, role string, exp time.Time) (string, error) {
	secret := []byte("mysecret")

	claims := jwt.MapClaims{}
	claims["exp"] = exp
	claims["authorized"] = true
	claims["user"] = username
	claims["roles"] = []string{role}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := at.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
