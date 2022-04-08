package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
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
