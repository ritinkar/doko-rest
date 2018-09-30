package controllers

import (
	"doko-rest/models"
	"doko-rest/utils"
	"encoding/json"
	"net/http"
)

type userRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// CreateUser : Create user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var request userRequest
	json.NewDecoder(r.Body).Decode(&request)
	salt := utils.GenerateSalt()
	newUser := &models.User{Username: request.Username, Salt: salt, Role: "admin"}
	res := utils.Message(true, "User Created")
	newUser.Create(request.Password)
	utils.Respond(w, res)

}
