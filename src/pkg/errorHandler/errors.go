package errorHandler

import (
	_type "PicusFinalCase/src/pkg/errorHandler/type"
)

func Panic(error _type.ErrorType) {
	panic(&_type.ErrorType{
		Code:    error.Code,
		Message: error.Message,
	})
}
