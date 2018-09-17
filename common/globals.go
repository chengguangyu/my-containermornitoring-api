package common

import (
	"fmt"
	"runtime"
)

var StatusChannel chan *StatusResponse

func UpdateAndSendStatus(status StatusResponse, statusChannel chan *StatusResponse) error {
	var err error

	go func() {
		defer func() {
			if r := recover(); r != nil {
				if _, ok := r.(runtime.Error); ok {
					fmt.Print(r.(runtime.Error), "Runtime Error")
				}
			}
		}()
		if previousStatus := <-statusChannel; previousStatus != nil {
			statusChannel <- &status
		} else {
			//first status
			statusChannel <- &status
		}

		return

	}()
	return err
}
