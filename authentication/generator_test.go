package authentication

import (
	"fmt"
	"testing"
	"time"
)

func TestTokengeneration(t *testing.T) {
	// prepare
	loc := time.UTC
	var now = func() time.Time { return time.Date(2019, 1, 1, 0, 0, 0, 0, loc) }
	expected := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOiIyMDE5LTAxLTAxVDAwOjAwOjAwWiIsInJvbGVzIjpbImFkbWluIl0sInVzZXIiOiJzb21lVXNlcm5hbWUifQ.L5p5s_28DxnyH08l9iGJFAKqMI2KeaI6fFuJe9K2tdY"

	//execute
	actual, _ := GenerateJwt("someUsername", "admin", now())
	fmt.Println(actual)

	//assert
	if actual != expected {
		t.FailNow()
	}

}
