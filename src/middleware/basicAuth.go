package middleware

import (
	"net/http"
)

type ApiKey struct {
	Api      string `bson:"api" json:"api"`
	ClientId string `bson:"clientId" json:"clientId"`
	Token    string `bson:"token" json:"token"`
	Status   bool   `bson:"status" json:"status"`
}

var apiKeys []ApiKey

func SetupBasicAuth() {

	//fetch api keys from database
	apiKeys[0] = ApiKey{
		ClientId: "comodoca",
		Token:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ",
		Status:   true,
		Api:      "registration",
	}
}

func BasicAuth(h func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		username, password, authOK := r.BasicAuth()
		if !authOK {
			w.WriteHeader(http.StatusUnauthorized)
		}

		for _, key := range apiKeys {
			if key.ClientId == username {
				if password != key.Token {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}

				h(w, r)
				return
			}
		}

		w.WriteHeader(http.StatusUnauthorized)
		return

	}
}
