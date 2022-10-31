package authentication

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func validate(bearer string) bool {
	s := strings.Split(bearer, " ")
	if len(s) != 2 {
		return false
	}
	tokenString := s[1]

	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	})

	if err != nil {
		fmt.Printf("Time now: %s \n", time.Now())
		fmt.Println(err.Error())
		return false
	}

	return true
}
