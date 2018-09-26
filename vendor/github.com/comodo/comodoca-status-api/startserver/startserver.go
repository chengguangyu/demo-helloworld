package startserver

import (
	"fmt"
	"github.com/comodo/comodoca-status-api/src/routers"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

var StatusServer *http.Server

func StartStatusServer() {
	go func() {
		statusRouter := mux.NewRouter()

		routes.SetupStatus(statusRouter, "")

		StatusServer = &http.Server{
			Handler:      statusRouter,
			Addr:         ":8081",
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}

		err := StatusServer.ListenAndServe()
		if err != nil {
			fmt.Print(err.Error())
		}
		fmt.Println("status page server is up")
	}()

}
