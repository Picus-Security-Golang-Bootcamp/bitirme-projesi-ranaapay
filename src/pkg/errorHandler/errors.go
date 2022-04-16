package errorHandler

import (
	_type "PicusFinalCase/src/pkg/errorHandler/type"
)

// Panic It throws panic as custom error type.
func Panic(error _type.ErrorType) {

	panic(&_type.ErrorType{
		Code:    error.Code,
		Message: error.Message,
	})
}
