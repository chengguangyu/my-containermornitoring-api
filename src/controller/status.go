package controllers

import (
	"encoding/json"
	"github.com/comodo/comodoca-status-api/common"
	"net/http"
)

func StatusHandler(w http.ResponseWriter, _ *http.Request) {

	for response := range common.StatusChannel {
		status := *response
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		statusJson, err := json.Marshal(status)
		if err != nil {
			panic(err)
		}
		//json.NewEncoder(w).Encode(status)
		w.Write(statusJson)
		close(common.StatusChannel)
	}

}
