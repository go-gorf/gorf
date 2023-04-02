package auth

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email" binding:"required"`
	Password string `json:"Password" binding:"required"`
}
