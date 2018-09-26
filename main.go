package main

import (
	"context"
	"fmt"
	"github.com/comodo/comodo-logging-lib/startlogger"
	"github.com/comodo/comodo-logging-lib/stoplogger"
	"github.com/comodo/comodoca-status-api/common"
	"github.com/comodo/comodoca-status-api/startserver"
	"github.com/comodo/comodoca-status-api/stopserver"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func HelloWorldHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello, world! this is Bob from ComodoCA."))
}

func SetupService(r *mux.Router, prefix string) {
	s := r.PathPrefix(prefix).Subrouter()
	s.HandleFunc("/helloworld", HelloWorldHandler).Methods("GET")
}

func init() {

}

func main() {
	var err error

	helloWorldRouter := mux.NewRouter()
	serverName := "helloWorld"
	hostname, _ := os.Hostname()
	Log := startlogger.Logger{}

	SetupService(helloWorldRouter, "/v1/comodoca")

	helloWorldServer := &http.Server{
		Handler:      helloWorldRouter,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	startserver.StartStatusServer()
	startlogger.StartLogServer(serverName, hostname, Log)

	go func() {
		err = helloWorldServer.ListenAndServe()
		if err != nil {
			fmt.Print(err.Error())
		}
	}()

	go func() {

		status := common.StatusResponse{
			ServiceName:        "Hello World Example Services",
			ServiceDescription: "A service that exists so documentation can be written for it.",
			Status:             "available",
			SubComponents:      nil,
			VersionNumber:      "2.0",
		}

		err = common.UpdateAndSendStatus(status)
		if err != nil {
			fmt.Print("error")
		}
		Log.Printf("guys, the hello world is damn started")

		time.Sleep(15 * time.Second)
		failStatus := common.StatusResponse{
			ServiceName:        "Database disconnection",
			ServiceDescription: "It is damaged by Bob",
			Status:             "unavailable",
			SubComponents:      nil,
			VersionNumber:      "2.0",
		}
		err = common.UpdateAndSendStatus(failStatus)
		if err != nil {
			fmt.Print("error")
		}
		Log.Warnf("%s", "something bad happens, the database is broken")

		time.Sleep(15 * time.Second)
		recoverStatus := common.StatusResponse{
			ServiceName:        "Database recovered",
			ServiceDescription: "It is fixed by Bob",
			Status:             "available",
			SubComponents:      nil,
			VersionNumber:      "2.0",
		}
		err = common.UpdateAndSendStatus(recoverStatus)
		if err != nil {
			fmt.Print("error")
		}
		Log.Print("Bob is the best, he fixed the database")
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	helloWorldServer.Shutdown(ctx)
	stopserver.StopStatusServer(ctx)
	stoplogger.StopLogger(Log)

	os.Exit(0)
}
