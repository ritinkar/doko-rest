package utils

import (
	"encoding/json"
	"net/http"
)

// Message : Generates a response object from a string
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

// Respond : Adds content-type before writing out response
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// CreateResponse : Generates a response object
func CreateResponse(status bool, data interface{}) map[string]interface{} {
	return map[string]interface{}{"status": status, "data": data}
}
