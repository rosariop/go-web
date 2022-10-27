package authentication

import "testing"

func TestVerification(t *testing.T) {
	// prepare
	expected := true
	bearer := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOiIyMDE5LTAxLTAxVDAwOjAwOjAwWiIsInJvbGVzIjpbImFkbWluIl0sInVzZXIiOiJzb21lVXNlcm5hbWUifQ.L5p5s_28DxnyH08l9iGJFAKqMI2KeaI6fFuJe9K2tdY"

	// action
	actual := validate(bearer)

	//assert
	if actual != expected {
		t.FailNow()
	}
}
