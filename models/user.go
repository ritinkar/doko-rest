package models

import (
	"doko-rest/utils"
	"fmt"

	"log"

	"github.com/jinzhu/gorm"
)

// User : structure for holding users
type User struct {
	gorm.Model
	Username       string `gorm:"unique" json:"username"`
	HashedPassword string `gorm:"type:varchar(1024)" json:"hashed_password"`
	Role           string `gorm:"type:varchar(64)" json:"role"`
	Salt           string `gorm:"type:varchar(64)" json:"salt"`
}

// JwtToken : structure for holding jwts
type JwtToken struct {
	Token string `json:"token"`
}

// Create : Add a new user to db
func (user *User) Create(password string) {
	user.HashedPassword = utils.HashAndSalt(password + user.Salt)
	db := GetDB()
	fmt.Println(user)
	if err := db.Create(user).Error; err != nil {
		log.Fatal(err)
	}
}
