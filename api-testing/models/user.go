package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
