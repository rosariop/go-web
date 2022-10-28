package authentication

import (
	"fmt"
	"strings"
	"time"
)

func validate(bearer string, t time.Time) bool {
	s := strings.Split(bearer, " ")
	tokenString := s[1]

	fmt.Println(tokenString)

	return true
}
