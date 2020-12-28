package model

import (
	"log"

	validator "github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	id           uint   `gorm:"primaryKey" json:"id"`
	Name         string `gorm:"not null" json:"name"`
	Email        string `gorm:"not null" valid:"email" json:"email,omitempty"`
	PasswordHash string `gorm:"not null" json:"-"`
}

type InputUser struct {
	Name     string `json:"name" valid:"optional"`
	Email    string `json:"email" valid:"email"`
	Password string `json:"password,omitempty"`
}

func (inputUser *InputUser) ValidateInputUser() bool {
	if !validator.IsEmail(inputUser.Email) {
		log.Println("Email format error")
		return false
	}
	if len(inputUser.Password) < 4 {
		log.Printf("Password length < 4, current length is %d", len(inputUser.Password))
		return false
	}
	return true
}
