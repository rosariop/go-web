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

func TestVerificationWrongFormat(t *testing.T) {
	// prepare
	bearer := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjAsInJvbGVzIjpbImFkbWluIl0sInVzZXIiOiJzb21lVXNlcm5hbWUifQ.EYi6AVwMoXzj4cJBvD5AQ8qo-AvXEJH2UjFUflyOsMMPo6cbLmINrmDjrBUfcojt8Nzu8uMWKp7V9AfQyf7r2g"

	// action
	actual := validate(bearer)

	//assert

	if actual {
		fmt.Println("")
		t.FailNow()
	}
}

func TestVerificationExpiredToken(t *testing.T) {
	// prepare
	bearer := "Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjAsInJvbGVzIjpbImFkbWluIl0sInVzZXIiOiJzb21lVXNlcm5hbWUifQ.EYi6AVwMoXzj4cJBvD5AQ8qo-AvXEJH2UjFUflyOsMMPo6cbLmINrmDjrBUfcojt8Nzu8uMWKp7V9AfQyf7r2g"

	// action
	actual := validate(bearer)

	//assert

	if actual {
		fmt.Println("")
		t.FailNow()
	}
}
