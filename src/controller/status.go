package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/comodo/comodoca-status-api/common"
	"net/http"
)

func StatusHandler(w http.ResponseWriter, _ *http.Request) {

	res := <-common.StatusChannel
	fmt.Print(&res)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&res)
	fmt.Println("got res from channel")

}
