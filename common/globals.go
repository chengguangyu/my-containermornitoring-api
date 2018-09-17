package common

import (
	"fmt"
	"runtime"
)

//global channel to save container status
var StatusChannel chan *StatusResponse

func UpdateAndSendStatus(status StatusResponse) error {
	var err error

	go func() {
		defer func() {
			if r := recover(); r != nil {
				if _, ok := r.(runtime.Error); ok {
					fmt.Print(r.(runtime.Error), "Runtime Error")
				}
			}
		}()
		if previousStatus := <-StatusChannel; previousStatus != nil {
			StatusChannel <- &status
		} else {
			//first status
			StatusChannel <- &status
		}

		return

	}()
	return err
}
