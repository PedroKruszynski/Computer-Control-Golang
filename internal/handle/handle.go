package handle

import (
	"encoding/json"
	"net/http"

	"volume-control/internal/volume_control"
)

func ReturnJson(w http.ResponseWriter, response volume_control.Response) {
	json.NewEncoder(w).Encode(response)
}
