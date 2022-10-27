package authentication

import (
	"encoding/json"
	"net/http"
	"time"
)

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	var userCredentials UserCredentials
	err := json.NewDecoder(r.Body).Decode(&userCredentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if userCredentials.Username == "someUsername" && userCredentials.Password == "somePassword" {
		exp := time.Now().Add(10 * time.Minute)
		jwt, err := GenerateJwt(userCredentials.Username, "admin", exp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = w.Write([]byte(jwt))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusUnauthorized)
}
