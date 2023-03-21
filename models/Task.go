package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model

	Title 		string `gorm:"not null; unique_index"`
	Description string
	Done 		string `gorm:"default:false"`
	UserID 		uint
}		