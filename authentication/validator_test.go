package authentication

import (
	"testing"
	"time"
)

func TestVerification(t *testing.T) {
	// prepare
	expected := true
	bearer := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOiIyMDE5LTAxLTAxVDAwOjAwOjAwWiIsInJvbGVzIjpbImFkbWluIl0sInVzZXIiOiJzb21lVXNlcm5hbWUifQ.L5p5s_28DxnyH08l9iGJFAKqMI2KeaI6fFuJe9K2tdY"
	var now = func() time.Time { return time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC) }

	// action
	actual := validate(bearer, now())

	//assert
	if actual != expected {
		t.FailNow()
	}
}
