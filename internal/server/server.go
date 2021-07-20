package server

import (
	"computer-control/internal/handle"
	"computer-control/internal/mouse_control"
	"computer-control/internal/volume_control"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/koron/go-ssdp"
)

func StartServerSsdp(port string) {
	myIp := getHostIp().String()
	hosName, err := os.Hostname()
	if err != nil {
		fmt.Println("Error getting hostname: ", err)
	}

	_, err = ssdp.Advertise(
		"urn:schemas-upnp-org:service:computer-control:1",
		"id:"+hosName,
		"http://"+myIp+""+port+"/",
		"ssdp for computer-control",
		3600)
	if err != nil {
		fmt.Println("Error advertising ssdp: ", err)
	}
}

func StartServerRest(port string) {
	r := mux.NewRouter()
	r.HandleFunc("/volume/showVolume", ShowVolume)
	r.HandleFunc("/volume/{newVolume:[0-9]+}", ChangeVolume)
	r.HandleFunc("/volume/mute", Mute)
	r.HandleFunc("/mouse/{x:[0-9]+}/{y:[0-9]+}", SetMousePosition)
	r.HandleFunc("/mouse/right", ClickRightMouse)
	r.Path("/mouse/left").Queries("doubleClick", "{doubleClick}").HandlerFunc(ClickLeftMouse)
	r.PathPrefix("/")

	err := http.ListenAndServe(port, r)
	if err != nil {
		fmt.Println("listening error: ", err)
	}
}

func getHostIp() net.IP {
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)

	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil && ipv4[0] == 192 {
			return ipv4
		}
	}
	return net.IP{}
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

func SetMousePosition(w http.ResponseWriter, r *http.Request) {
	res := mouse_control.SetMousePosition(r)
	handle.ReturnJson(w, res)
}

func ClickLeftMouse(w http.ResponseWriter, r *http.Request) {
	res := mouse_control.ClickLeftMouse(r)
	handle.ReturnJson(w, res)
}

func ClickRightMouse(w http.ResponseWriter, r *http.Request) {
	res := mouse_control.ClickRightMouse(r)
	handle.ReturnJson(w, res)
}
