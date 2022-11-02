package authentication

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
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

type realClock struct{}

var now = func() time.Time { return time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC) }

func (realClock) Now() time.Time { return now() }

func TestValidateOK(t *testing.T) {

	// mocking data
	userdata := UserCredentials{Username: "someUsername", Password: "password"}
	var byteBuffer bytes.Buffer
	err := json.NewEncoder(&byteBuffer).Encode(userdata)
	if err != nil {
		fmt.Println("Failing due to bad encoding")
		t.FailNow()
	}

	// prepares request
	req, err := http.NewRequest(http.MethodGet, "/validate", &byteBuffer)
	if err != nil {
		t.Fatal(err)
	}

	// prepares response recorder
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AuthHandler)

	// starts request
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusUnauthorized {
		fmt.Printf("FAIL: Wrong status code: %d", rr.Code)
		t.FailNow()
	}
}
