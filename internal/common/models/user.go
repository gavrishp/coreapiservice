package models

import "gorm.io/gorm"

type User struct {
	gorm.Model         // adds ID, created_at etc.
	Name        string `json:"name"`
	Role        string `json:"role"`
	Description string `json:"description"`
}
