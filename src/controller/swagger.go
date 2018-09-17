package controllers

import (
	"net/http"
)

func SwaggerHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "swagger/swagger.yaml")
}
