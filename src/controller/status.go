package controllers

import (
	"net/http"
)

func StatusHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	/*
duration := time.Now().Sub(started)
if duration.Seconds() > 10 {
	w.WriteHeader(500)
	w.Write([]byte(fmt.Sprintf("error: %v", duration.Seconds())))
} else {
	w.WriteHeader(200)
	w.Write([]byte("ok"))
}
*/
}