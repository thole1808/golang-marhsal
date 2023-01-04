package helper

import (
	"encoding/json"
	"net/http"
)

// teruntuk interface menerima semua parameters apa saja
func ResponseJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
