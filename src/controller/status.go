package controllers

import (
	"github.com/comodo/comodoca-status-api/common"
	"net/http"
)

func StatusHandler(w http.ResponseWriter, _ *http.Request, statusChan chan *common.StatusResponse) {

	_, ok := <-statusChan
	if !ok {
		w.Write([]byte("didn't receive"))
	}
	//w.WriteHeader(http.StatusOK)
	//w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(&res)
	w.Write([]byte("received status"))

}
