package models

import (
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       int64  `gorm:"primaryKey" json:"id"`
	Username string `gorm:"type:varchar(255); not null; unique" json:"username"`
	Password string `gorm:"type:varchar(255); not null" json:"password"`
	Books    []Book
}

func (u *User) SaveUser() (*User, error) {

	var err error
	err = DB.Create(&u).Error

	if err != nil {
		return &User{}, err
	}

	return u, nil

}

func (u *User) BeforeSave(tx *gorm.DB) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	u.Username = strings.TrimSpace(u.Username)

	return nil

}
