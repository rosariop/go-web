package authentication

import (
	"fmt"
	"testing"
	"time"
)

func TestVerification(t *testing.T) {
	// prepare
	bearer, err := GenerateJwt("someUsername", "admin", time.Now())
	if err != nil {
		fmt.Println("Error generating JWT")
		t.FailNow()
	}

	// action
	actual := validate(bearer)

	//assert

	if !actual {
		fmt.Println("")
		t.FailNow()
	}
}
