package helper

import (
	"encoding/json"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, code int, payload interface{}) {
	res, _ := json.Marshal(payload)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
