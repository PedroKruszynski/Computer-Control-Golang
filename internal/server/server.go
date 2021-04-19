package server

import (
	"fmt"
	"net/http"
	"volume-control/internal/handle"
	"volume-control/internal/volume_control"

	"github.com/gorilla/mux"
)

func StartServer() {
	r := mux.NewRouter()
	r.HandleFunc("/showVolume", ShowVolume)
	r.HandleFunc("/{newVolume:[0-9]+}", ChangeVolume)
	r.HandleFunc("/mute", Mute)
	http.Handle("/", r)
	fmt.Printf("Server.go")
	http.ListenAndServe(":8888", nil)
}

func ShowVolume(w http.ResponseWriter, r *http.Request) {
	res := volume_control.ShowVolume(r)
	handle.ReturnJson(w, res)
}

func ChangeVolume(w http.ResponseWriter, r *http.Request) {
	res := volume_control.ChangeVolume(r)
	handle.ReturnJson(w, res)
}

func Mute(w http.ResponseWriter, r *http.Request) {
	res := volume_control.Mute(r)
	handle.ReturnJson(w, res)
}
