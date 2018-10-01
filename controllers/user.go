package controllers

import (
	"doko-rest/models"
	"doko-rest/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
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
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": existingUser.Username,
			"role":     existingUser.Role,
		})
		tokenString, error := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
		if error != nil {
			fmt.Println(error)
		}

		res := utils.CreateResponse(true, map[string]string{"token": tokenString})
		utils.Respond(w, res)
	} else {
		res := utils.Message(false, "Wrong username or password")
		utils.Respond(w, res)
	}

}
