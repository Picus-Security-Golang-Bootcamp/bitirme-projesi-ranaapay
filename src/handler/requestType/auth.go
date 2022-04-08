package requestType

import (
	"PicusFinalCase/src/models"
	"errors"
)

type LoginType struct {
	FirstName string `json:"firstName"`
	Password  string `json:"password"`
}

type UserRequestType struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (req UserRequestType) RequestToUserType() models.User {
	return models.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
		Role:      1,
	}
}

func (req UserRequestType) ValidateUserRequest() error {
	if req.FirstName == "" {
		return errors.New("FirstName is required, it can not be empty. ")
	}
	if req.LastName == "" {
		return errors.New("LastName is required, it can not be empty. ")
	}
	if req.Email == "" {
		return errors.New("Email is required, it can not be empty. ")
	}
	if req.Password == "" {
		return errors.New("Password is required, it can not be empty. ")
	}
	return nil
}

func (req LoginType) ValidateLoginType() error {
	if req.FirstName == "" {
		return errors.New("FirstName is required, it can not be empty. ")
	}
	if req.Password == "" {
		return errors.New("Password is required, it can not be empty. ")
	}
	return nil
}
