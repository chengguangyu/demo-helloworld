package controllers

import (
	"encoding/json"
	"github.com/comodo/comodoca-status-api/common"
	"net/http"
)

func StatusHandler(w http.ResponseWriter, _ *http.Request) {

	select {
	case response, ok := <-common.StatusChannel:
		if ok {
			status := *response
			w.Header().Set("Content-Type", "application/json")
			if status.Status == "available" {
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
			json.NewEncoder(w).Encode(status)
			common.LastStatus = status
			close(common.StatusChannel)
		} else {

			w.Header().Set("Content-Type", "application/json")

			if common.LastStatus.Status == "available" {
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
			json.NewEncoder(w).Encode(common.LastStatus)
		}

		return

	}

}
