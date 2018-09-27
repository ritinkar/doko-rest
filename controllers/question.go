package controllers

import (
	"doko-rest/models"
	"encoding/json"
	"net/http"
)

type request struct {
	Text            string   `json:"text"`
	CorrectAnswerID int      `json:"correctAnswerID"`
	Answers         []string `json:"answers"`
}

// CreateQuestion : Create question
func CreateQuestion(w http.ResponseWriter, r *http.Request) {
	var request request
	var answers []models.Answer
	json.NewDecoder(r.Body).Decode(&request)

	for idx, answer := range request.Answers {
		answers = append(answers, models.Answer{Text: answer, Correct: idx == request.CorrectAnswerID})
	}
	newQuestion := &models.Question{Text: request.Text, Answers: answers}

	newQuestion.Create()

}
