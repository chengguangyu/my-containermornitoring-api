package startserver

import (
	"fmt"
	"github.com/comodo/comodoca-status-api/src/routers"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func StartStatusServer() {
	go func() {
		statusRouter := mux.NewRouter()

		routes.SetupStatus(statusRouter, "")

		statusServer := &http.Server{
			Handler:      statusRouter,
			Addr:         ":8081",
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}

		err := statusServer.ListenAndServe()
		if err != nil {
			fmt.Print(err.Error())
		}
		fmt.Println("status page server is up")
	}()

}
