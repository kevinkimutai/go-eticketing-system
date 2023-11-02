package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique"`
	Role      string `gorm:"type:ENUM('admin', 'user');default:'user'"`
	Password  string `json:"password"`
}
