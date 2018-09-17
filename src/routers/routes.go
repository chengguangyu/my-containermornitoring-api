package routes

import (
	"github.com/comodo/comodoca-status-api/src/controller"
	"github.com/comodo/comodoca-status-api/src/middleware"
	"github.com/gorilla/mux"
)

func Setup(r *mux.Router) {
	middleware.SetupBasicAuth()
	s := r.PathPrefix("/v1/comodoca/registration").Subrouter()
	s.HandleFunc("/@status", middleware.BasicAuth(controllers.StatusHandler)).Methods("GET")
	s.HandleFunc("/events", middleware.BasicAuth(controllers.RegistrationHandler)).Methods("POST")
	s.HandleFunc("/swagger", middleware.BasicAuth(controllers.SwaggerHandler)).Methods("GET")
}
