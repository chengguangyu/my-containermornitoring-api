package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/comodo/comodoca-status-api/common"
	"net/http"
)

func StatusHandler(w http.ResponseWriter, r *http.Request, statusChan chan *common.StatusResponse) {
	res := <-statusChan

	fmt.Print(&res)

	for res := range statusChan {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&res)
	}

	//w.Write([]byte("received status"))

}
