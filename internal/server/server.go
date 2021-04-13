package server

import (
	"fmt"
	"net/http"
	"volume-control/internal/handle"
	"volume-control/internal/volume_control"
)

func StartServer() {
	fmt.Printf("Server.go")
	http.HandleFunc("/", ShowVolumeActual)
	http.ListenAndServe(":8888", nil)
}

func ShowVolumeActual(w http.ResponseWriter, r *http.Request) {
	res := volume_control.ShowVolumeActual(r)
	handle.ReturnJson(w, res)
}
