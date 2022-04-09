package models

import (
	"PicusFinalCase/src/pkg/errorHandler"
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

func (u *User) HashPassword() {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		errorHandler.Panic(errorHandler.HashPasswordError)
	}
	u.Password = string(bytes)
}

func (u *User) CheckPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
