package models

import "github.com/jinzhu/gorm"

// Answer : structure for holding answers
type Answer struct {
	gorm.Model
	QuestionID uint
	Text       string `gorm:"type:varchar(500)" json:"text"`
	Correct    bool   `gorm:"type:BOOLEAN" json:"correct"`
}

// Answers : slice of answers
type Answers []Answer

// Create : Add a new answer to db
func (answer *Answer) Create() {
	db := GetDB()
	db.Create(answer)
}
