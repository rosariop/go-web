package authentication

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func validate(bearer string) bool {
	s := strings.Split(bearer, " ")
	tokenString := s[1]

	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("mysecret"), nil
	})

	if err != nil {
		fmt.Printf("Time now: %s \n", time.Now())
		fmt.Println(err.Error())
		return false
	}
	fmt.Println(token)

	fmt.Println(tokenString)

	return true
}
