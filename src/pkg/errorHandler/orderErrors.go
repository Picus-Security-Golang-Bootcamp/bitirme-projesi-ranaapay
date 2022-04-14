package errorHandler

import (
	_type "PicusFinalCase/src/pkg/errorHandler/type"
	"net/http"
)

var (
	CartNotFoundError = _type.ErrorType{
		Code:    http.StatusBadRequest,
		Message: "There are no products in your cart.",
	}
	CartNotContainCartDetailError = _type.ErrorType{
		Code:    http.StatusBadRequest,
		Message: "There are no products in your cart.",
	}
)
