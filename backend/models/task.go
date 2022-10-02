package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string `json:"title" gorm:"size:60; NOT NULL;"`
	Description string `json:"description"`
}
