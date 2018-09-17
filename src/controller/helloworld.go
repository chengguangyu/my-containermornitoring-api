package controllers

import (
	"net/http"
)

func HelloWorldHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello, world! this is Bob from ComodoCA."))
}
