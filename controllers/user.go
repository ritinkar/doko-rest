package controllers

import (
	"doko-rest/models"
	"doko-rest/utils"
	"encoding/json"
	"log"
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
	newUser.Create(request.Password)
	res := utils.Message(true, "User Created")
	utils.Respond(w, res)

}

// AuthenticateUser : authenticates a user
func AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	var request userRequest
	json.NewDecoder(r.Body).Decode(&request)
	existingUser := &models.User{Username: request.Username}
	existingUser.ReadOne()
	log.Println(existingUser)
	if authenticated := utils.ComparePasswords(
		existingUser.HashedPassword,
		request.Password,
		existingUser.Salt); authenticated {
		res := utils.CreateResponse(true, map[string]string{"token": "foobar"})
		utils.Respond(w, res)
	} else {
		res := utils.Message(false, "Wrong username or password")
		utils.Respond(w, res)
	}

}
