package routes

import (
	"github.com/comodo/comodoca-status-api/src/controller"
	"github.com/gorilla/mux"
	"net/http"
)

//comment
func SetupStatus(r *mux.Router, prefix string) {
	s := r.PathPrefix(prefix).Subrouter()
	s.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		controllers.StatusHandler(w, r)
	}).Methods("GET")
}
