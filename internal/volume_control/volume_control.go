package volume_control

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/itchyny/volume-go"
)

func ShowVolume(r *http.Request) Response {

	vol, err := volume.GetVolume()
	if err != nil {
		return Response{500, "Not possible get the volume"}
	}

	return Response{200, strconv.Itoa(vol)}
}

func Mute(r *http.Request) Response {

	muted, err := volume.GetMuted()
	if err != nil {
		return Response{500, "Not possible get if volume is muted or not"}
	}

	if muted {
		err = volume.Unmute()
		if err != nil {
			return Response{500, "Not possible unmute the volume"}
		}
		return Response{200, "Volume unmuted"}
	} else {
		err = volume.Mute()
		if err != nil {
			return Response{500, "Not possible mute the volume"}
		}
		return Response{200, "Volume muted"}
	}

}

func ChangeVolume(r *http.Request) Response {

	vars := mux.Vars(r)
	volumeParam, error := vars["newVolume"]
	if error == false {
		return Response{500, "Invalid new volume get in param"}
	}

	newVolume, _ := strconv.Atoi(volumeParam)

	err := volume.SetVolume(newVolume)
	if err != nil {
		return Response{500, "Not possible turn up the volume"}
	}

	actualVolume, _ := volume.GetVolume()
	return Response{200, strconv.Itoa(actualVolume)}
}
