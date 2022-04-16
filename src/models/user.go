package models

import (
	"PicusFinalCase/src/pkg/errorHandler"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Base
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      Role   `json:"role"`
}

func (User) TableName() string {
	//default table name
	return "users"
}

type Role int

const (
	Admin Role = iota
	Customer
)

// HashPassword Encrypts the user's password and assigns it to the user password.
func (u *User) HashPassword() {

	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		log.Error("Hash Password Error : %v", err)
		errorHandler.Panic(errorHandler.HashPasswordError)
	}

	u.Password = string(bytes)
}

// CheckPasswordHash Compares the user encrypted password with the incoming encrypted password. If it's not the same, it throws an error.
func (u *User) CheckPasswordHash(password string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
