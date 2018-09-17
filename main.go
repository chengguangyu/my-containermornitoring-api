package main

import (
	"github.com/hushed/comodoca-status-api/src/routers"
	"net/http"
	"time"
	"os"
	"os/signal"
	"github.com/gorilla/mux"
	"fmt"
	"context"
)

func main() {


	r := mux.NewRouter()

	routes.Setup(r)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			fmt.Print(err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	srv.Shutdown(ctx)
	os.Exit(0)
}
