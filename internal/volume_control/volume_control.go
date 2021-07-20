package volume_control

import (
	"computer-control/internal/handle"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/itchyny/volume-go"
)

func ShowVolume(r *http.Request) handle.Response {

	vol, err := volume.GetVolume()
	if err != nil {
		return handle.Response{HttpCode: 500, Message: "Not possible get the volume"}
	}

	return handle.Response{HttpCode: 200, Message: strconv.Itoa(vol)}
}

func Mute(r *http.Request) handle.Response {

	muted, err := volume.GetMuted()
	if err != nil {
		return handle.Response{HttpCode: 500, Message: "Not possible get if volume is muted or not"}
	}

	if muted {
		err = volume.Unmute()
		if err != nil {
			return handle.Response{HttpCode: 500, Message: "Not possible unmute the volume"}
		}
		return handle.Response{HttpCode: 200, Message: "Volume unmuted"}
	} else {
		err = volume.Mute()
		if err != nil {
			return handle.Response{HttpCode: 500, Message: "Not possible mute the volume"}
		}
		return handle.Response{HttpCode: 200, Message: "Volume muted"}
	}

}

func ChangeVolume(r *http.Request) handle.Response {

	vars := mux.Vars(r)
	volumeParam, error := vars["newVolume"]
	if error == false {
		return handle.Response{HttpCode: 400, Message: "Invalid new volume get in param"}
	}

	newVolume, _ := strconv.Atoi(volumeParam)

	err := volume.SetVolume(newVolume)
	if err != nil {
		return handle.Response{HttpCode: 500, Message: "Not possible turn up the volume"}
	}

	actualVolume, _ := volume.GetVolume()
	return handle.Response{HttpCode: 200, Message: strconv.Itoa(actualVolume)}
}
