package controllers

import (
	"net/http"
)

func RegistrationHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)

}
