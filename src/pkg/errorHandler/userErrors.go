package errorHandler

import (
	_type "PicusFinalCase/src/pkg/errorHandler/type"
	"net/http"
)

var (
	FirstNameValidationError = _type.ErrorType{
		Code:    http.StatusBadRequest,
		Message: "FirstName is required, it can not be empty.",
	}
	LastNameValidationError = _type.ErrorType{
		Code:    http.StatusBadRequest,
		Message: "LastName is required, it can not be empty.",
	}
	EmailValidationError = _type.ErrorType{
		Code:    http.StatusBadRequest,
		Message: "Email is required, it can not be empty.",
	}
	PasswordValidationError = _type.ErrorType{
		Code:    http.StatusBadRequest,
		Message: "Password is required, it can not be empty.",
	}
	HashPasswordError = _type.ErrorType{
		Code:    http.StatusInternalServerError,
		Message: "Hash Password error.",
	}
	GenerateJwtError = _type.ErrorType{
		Code:    http.StatusInternalServerError,
		Message: "Generate Jwt error.",
	}
	PasswordNotTrueError = _type.ErrorType{
		Code:    http.StatusBadRequest,
		Message: "The password you entered is incorrect. Please enter valid password.",
	}
)
