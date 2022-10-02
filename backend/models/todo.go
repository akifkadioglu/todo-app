package models

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	ID          uint `gorm:"primaryKey"`
	Title       string
	Description string
}
