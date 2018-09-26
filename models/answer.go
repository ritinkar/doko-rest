package models

import "github.com/jinzhu/gorm"

// Answer : structure for holding answers
type Answer struct {
	gorm.Model
	QuestionID uint
	text       string `gorm:"type:varchar(500)"`
}
