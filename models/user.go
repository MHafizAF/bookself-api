package models

import (
	"strings"

	"github.com/MHafizAF/bookself-api/utils/token"
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

func VerifyPasword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func SignInCheck(username string, password string) (string, error) {
	var err error

	user := User{}

	err = DB.Model(User{}).Where("username = ?", username).Take(&user).Error

	if err != nil {
		return "", err
	}

	err = VerifyPasword(password, user.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(user.Id)

	if err != nil {
		return "", err
	}

	return token, nil

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
