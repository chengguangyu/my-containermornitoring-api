package common

import (
	"fmt"
	"runtime"
)

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
		if previousStatus := <-StatusChannel; &previousStatus != nil {
			fmt.Print("write into channel1")
			StatusChannel <- &status
		} else {
			//first status
			fmt.Print("write into channel")
			StatusChannel <- &status
		}

	}()
	return err
}
