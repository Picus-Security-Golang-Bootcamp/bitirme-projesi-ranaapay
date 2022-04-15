package requestType

import (
	"PicusFinalCase/src/models"
	"PicusFinalCase/src/pkg/errorHandler"
	log "github.com/sirupsen/logrus"
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
	log.Info("Created user type according to request user type.")
	return models.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
		Role:      1,
	}
}

// ValidateUserRequest It is checked whether the fields of the request comply with the expected values. If not, it throws an error.
func (req UserRequestType) ValidateUserRequest() {
	if req.FirstName == "" {
		log.Error("User request firstname is empty.")
		errorHandler.Panic(errorHandler.FirstNameValidationError)
	}
	if req.LastName == "" {
		log.Error("User request lastname is empty.")
		errorHandler.Panic(errorHandler.LastNameValidationError)
	}
	if req.Email == "" {
		log.Error("User request email is empty.")
		errorHandler.Panic(errorHandler.EmailValidationError)
	}
	if req.Password == "" {
		log.Error("User request password is empty.")
		errorHandler.Panic(errorHandler.PasswordValidationError)
	}
}

// ValidateLoginType It is checked whether the fields of the request comply with the expected values. If not, it throws an error.
func (req LoginType) ValidateLoginType() {
	if req.FirstName == "" {
		log.Error("Login request firstname is empty.")
		errorHandler.Panic(errorHandler.FirstNameValidationError)
	}
	if req.Password == "" {
		log.Error("Login request password is empty.")
		errorHandler.Panic(errorHandler.PasswordValidationError)
	}
}
