package handle

import (
	"encoding/json"
	"net/http"
)

func ReturnJson(w http.ResponseWriter, message string) {
	json.NewEncoder(w).Encode(message)
}
