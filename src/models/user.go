package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Password  string
	Role      Role
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
