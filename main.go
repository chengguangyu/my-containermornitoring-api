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
			ServiceName:        "Hello World Example Services",
			ServiceDescription: "A service that exists so documentation can be written for it.",
			Status:             "available",
			SubComponents:      nil,
		}

		err := common.UpdateAndSendStatus(status)
		if err != nil {
			fmt.Print("error")
		}

		//fake database disconnection
		go func() {
			time.Sleep(15 * time.Second)
			failStatus := common.StatusResponse{
				ServiceName:        "Database disconnection",
				ServiceDescription: "It is killed by Bob",
				Status:             "unavailable",
				SubComponents:      nil,
			}
			err := common.UpdateAndSendStatus(failStatus)
			if err != nil {
				fmt.Print("error")
			}

			time.Sleep(15 * time.Second)
			recoverStatus := common.StatusResponse{
				ServiceName:        "Database recovered",
				ServiceDescription: "It is killed by Bob",
				Status:             "unavailable",
				SubComponents:      nil,
			}
			err = common.UpdateAndSendStatus(recoverStatus)
			if err != nil {
				fmt.Print("error")
			}

		}()

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
