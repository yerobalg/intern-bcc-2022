package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title     string    `json:"title" gorm:"type:varchar(100);not null"`
	Text      string    `json:"text" gorm:"type:text;not null"`
}

type PostInput struct {
	Title string `json:"title" binding:"required"`
	Text  string `json:"text" binding:"required"`
}
