package common

import (
	"fmt"
	"runtime"
)

var StatusChannel chan *StatusResponse
var LastStatus StatusResponse

func UpdateAndSendStatus(status StatusResponse) error {
	var err error
	StatusChannel = make(chan *StatusResponse)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				if _, ok := r.(runtime.Error); ok {
					fmt.Print(r.(runtime.Error), "Runtime Error")
				}
			}
		}()

		select {
		case _, ok := <-StatusChannel:
			if ok {
				StatusChannel <- &status
			} else {
				fmt.Println("EmptyChannel")
			}
		default:

			StatusChannel <- &status
		}

	}()
	return err
}
