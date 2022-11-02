package authentication

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"
)

func TestTokengeneration(t *testing.T) {
	// prepare
	loc := time.UTC
	var now = func() time.Time { return time.Date(2019, 1, 1, 0, 0, 0, 0, loc) }
	expected := "Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE1NDYzMDE0MDAsInJvbGVzIjpbImFkbWluIl0sInVzZXIiOiJzb21lVXNlcm5hbWUifQ.OMsojasNz6tDUjnEwJOxfGLw2fEZ5B1DkSvVpl0JcPwhX0BYQW-0vkywZRIq_BtZwOpMZtkqNDQOOv4DnNV9pw"

	//execute
	actual, _ := GenerateJwt("someUsername", "admin", now())

	//assert
	if actual != expected {
		t.FailNow()
	}
}

func TestGenerateHandlerWithReturnToken(t *testing.T) {

	// mocking data
	userdata := UserCredentials{Username: "someUsername", Password: "somePassword"}
	var byteBuffer bytes.Buffer
	err := json.NewEncoder(&byteBuffer).Encode(userdata)
	if err != nil {
		fmt.Println("Failing due to bad encoding")
		t.FailNow()
	}

	tokenPattern := "Bearer [A-Za-z0-9-_]*\\.[A-Za-z0-9-_]*\\.[A-Za-z0-9-_]*$"

	// prepares request
	req, err := http.NewRequest(http.MethodGet, "/authenticate", &byteBuffer)
	if err != nil {
		t.Fatal(err)
	}

	// prepares response recorder
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AuthHandler)

	// starts request
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		fmt.Printf("FAIL: Wrong status code: %d", rr.Code)
		t.FailNow()
	}

	responseToken := rr.Body.String()

	isJWT, err := regexp.Match(tokenPattern, []byte(responseToken))
	if err != nil {
		log.Fatalln(err)
	}

	if !isJWT {
		t.FailNow()
	}
}

func TestGenerateHandlerWithWrongUsername(t *testing.T) {

	// mocking data
	userdata := UserCredentials{Username: "username", Password: "somePassword"}
	var byteBuffer bytes.Buffer
	err := json.NewEncoder(&byteBuffer).Encode(userdata)
	if err != nil {
		fmt.Println("Failing due to bad encoding")
		t.FailNow()
	}

	// prepares request
	req, err := http.NewRequest(http.MethodGet, "/authenticate", &byteBuffer)
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

func TestGenerateHandlerWithWrongPassword(t *testing.T) {

	// mocking data
	userdata := UserCredentials{Username: "someUsername", Password: "password"}
	var byteBuffer bytes.Buffer
	err := json.NewEncoder(&byteBuffer).Encode(userdata)
	if err != nil {
		fmt.Println("Failing due to bad encoding")
		t.FailNow()
	}

	// prepares request
	req, err := http.NewRequest(http.MethodGet, "/authenticate", &byteBuffer)
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
