package models

import "github.com/jinzhu/gorm"

// Question : structure for holding questions
type Question struct {
	gorm.Model
	Text    string   `gorm:"type:varchar(500);unique;not null" json:"text"`
	Answers []Answer `gorm:"foreignkey:QuestionID" json:"answers"`
}

// Questions : slice of questions
type Questions []Question

// Create : Add a new question to db
func (question *Question) Create() {
	db := GetDB()
	db.NewRecord(*question)
	db.Create(question)
	db.NewRecord(*question)
}

// ReadAll : ReadAll the questions
func (questions *Questions) ReadAll() {
	db.Find(questions)
}
