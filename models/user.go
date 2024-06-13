package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Email    string `gorm:"unique" json:"email"`
    Password string `json:"password"`
}

type LoginCredentials struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}
