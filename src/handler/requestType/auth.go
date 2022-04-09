package requestType

import (
	"PicusFinalCase/src/models"
	"PicusFinalCase/src/pkg/errorHandler"
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

func (req UserRequestType) ValidateUserRequest() {
	if req.FirstName == "" {
		errorHandler.Panic(errorHandler.FirstNameValidationError)
	}
	if req.LastName == "" {
		errorHandler.Panic(errorHandler.LastNameValidationError)
	}
	if req.Email == "" {
		errorHandler.Panic(errorHandler.EmailValidationError)
	}
	if req.Password == "" {
		errorHandler.Panic(errorHandler.PasswordValidationError)
	}
}

func (req LoginType) ValidateLoginType() {
	if req.FirstName == "" {
		errorHandler.Panic(errorHandler.FirstNameValidationError)
	}
	if req.Password == "" {
		errorHandler.Panic(errorHandler.PasswordValidationError)
	}
}
