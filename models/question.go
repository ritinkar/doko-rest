package models

import "github.com/jinzhu/gorm"

// Question : structure for holding questions
type Question struct {
	gorm.Model
	Text    string   `gorm:"type:varchar(500)" json:"text"`
	Answers []Answer `gorm:"foreignkey:QuestionID" json:"answers"`
}

// Create : Add a new question to db
func (question *Question) Create() {
	db := GetDB()
	db.Create(question)
}
