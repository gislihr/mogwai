package mogwai

import (
	"encoding/json"
	"net/http"
)

//Respond marshals data as json and writes to response writer
func Respond(data interface{}, w http.ResponseWriter) {
	bytes, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}
