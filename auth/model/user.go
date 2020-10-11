package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	id           uint   `gorm:"primaryKey" json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
}
