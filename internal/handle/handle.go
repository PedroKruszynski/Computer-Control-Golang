package handle

import (
	"encoding/json"
	"net/http"
)

func ReturnJson(w http.ResponseWriter, response Response) {
	json.NewEncoder(w).Encode(response)
}
