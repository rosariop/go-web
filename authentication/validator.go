package authentication

import (
	"fmt"
	"strings"
)

func validate(bearer string) bool {
	s := strings.Split(bearer, " ")
	tokenString := s[1]

	fmt.Println(tokenString)

	return true
}
