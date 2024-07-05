package models

import "gorm.io/gorm"

type TodoItem struct {
	gorm.Model
	// ID          uint   `gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
