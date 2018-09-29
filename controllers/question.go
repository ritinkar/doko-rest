package controllers

import (
	"doko-rest/models"
	"doko-rest/utils"
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
	var answers models.Answers
	json.NewDecoder(r.Body).Decode(&request)

	for idx, answer := range request.Answers {
		answers = append(answers, models.Answer{Text: answer, Correct: idx == request.CorrectAnswerID})
	}
	newQuestion := &models.Question{Text: request.Text, Answers: answers}
	m := make(map[string]interface{})
	m["data"] = newQuestion
	newQuestion.Create()
	utils.Respond(w, m)

}

// GetQuestions : gets all the questions
func GetQuestions(w http.ResponseWriter, r *http.Request) {
	var questions models.Questions
	questions.ReadAll()
	m := make(map[string]interface{})
	m["data"] = questions
	utils.Respond(w, m)
}
