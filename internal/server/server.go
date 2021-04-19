package server

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"volume-control/internal/handle"
	"volume-control/internal/volume_control"

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
		"urn:schemas-upnp-org:service:volume-control:1",
		"id:"+hosName,
		"http://"+myIp+""+port+"/",
		"ssdp for volume-control",
		3600)
	if err != nil {
		fmt.Println("Error advertising ssdp: ", err)
	}
}

func StartServerRest(port string) {
	r := mux.NewRouter()
	r.HandleFunc("/showVolume", ShowVolume)
	r.HandleFunc("/{newVolume:[0-9]+}", ChangeVolume)
	r.HandleFunc("/mute", Mute)
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
