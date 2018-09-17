package routes

import (
	"github.com/comodo/comodoca-status-api/src/controller"
	"github.com/gorilla/mux"
)

//comment
func SetupStatus(r *mux.Router, prefix string) {
	s := r.PathPrefix(prefix).Subrouter()
	s.HandleFunc("/status", controllers.StatusHandler).Methods("GET")
}

func SetupService(r *mux.Router, prefix string) {
	s := r.PathPrefix(prefix).Subrouter()
	s.HandleFunc("/helloworld", controllers.HelloWorldHandler).Methods("GET")
	s.HandleFunc("/swagger", controllers.SwaggerHandler).Methods("GET")
}