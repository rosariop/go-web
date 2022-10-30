package authentication

import (
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
