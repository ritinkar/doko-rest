package controllers

import (
	"encoding/json"
	"net/http"
)

type request struct {
	text          string
	correctAnswer uint
	answers       []string
}

// CreateQuestion : Create question
func CreateQuestion(w http.ResponseWriter, r *http.Request) {
	var request request
	json.NewDecoder(r.Body).Decode(&request)

}
