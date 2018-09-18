package main

import (
	"context"
	"fmt"
	"github.com/comodo/comodoca-status-api/common"
	"github.com/comodo/comodoca-status-api/src/routers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	helloWorldRouter := mux.NewRouter()
	statusRouter := mux.NewRouter()
	common.StatusChannel = make(chan *common.StatusResponse)

	routes.SetupService(helloWorldRouter, "/v1/comodoca")
	routes.SetupStatus(statusRouter, "")

	helloWorldServer := &http.Server{
		Handler:      helloWorldRouter,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	statusServer := &http.Server{
		Handler:      statusRouter,
		Addr:         ":8081",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {

		status := common.StatusResponse{
			ServiceName:        "Example Services",
			ServiceDescription: "A service that exists so documentation can be written for it.",
			Status:             "available",
			SubComponents:      nil,
		}
		fmt.Print(status)

		err := common.UpdateAndSendStatus(status)
		if err != nil {
			fmt.Print("error")
		}
		err = statusServer.ListenAndServe()
		if err != nil {
			fmt.Print(err.Error())
		}
	}()

	err := helloWorldServer.ListenAndServe()
	if err != nil {
		fmt.Print(err.Error())
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	helloWorldServer.Shutdown(ctx)
	statusServer.Shutdown(ctx)
	os.Exit(0)
}
