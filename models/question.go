package models

import "github.com/jinzhu/gorm"

// Question : structure for holding questions
type Question struct {
	gorm.Model
	Text          string   `gorm:"type:varchar(500)"`
	CorrectAnswer Answer   `gorm:"foreignkey:QuestionID"`
	Answers       []Answer `gorm:"foreignkey:QuestionID"`
}
