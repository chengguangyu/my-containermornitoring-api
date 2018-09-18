package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/comodo/comodoca-status-api/common"
	"net/http"
)

func StatusHandler(w http.ResponseWriter, _ *http.Request) {

	select {
	case response, ok := <-common.StatusChannel:
		if ok {
			status := *response
			w.Header().Set("Content-Type", "application/json")
			//statusJson, err := json.Marshal(status)
			//if err != nil {
			//panic(err)
			//}
			if status.Status == "available" {
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusOK)
			}
			json.NewEncoder(w).Encode(status)
			common.LastStatus = status
			//w.Write(statusJson)
			close(common.StatusChannel)
		} else {
			fmt.Print("called")

			w.Header().Set("Content-Type", "application/json")
			//statusJson, err := json.Marshal(common.LastStatus)

			//if err != nil {
			//	panic(err)
			//}
			if common.LastStatus.Status == "available" {
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusOK)
			}
			json.NewEncoder(w).Encode(common.LastStatus)
			//w.Write(statusJson)
		}

		return

	}

}
