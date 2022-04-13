package errorHandler

import (
	_type "PicusFinalCase/src/pkg/errorHandler/type"
	"net/http"
)

var (
	ProductIdValidationError = _type.ErrorType{
		Code:    http.StatusBadRequest,
		Message: "ProductId is required, it can not be empty.",
	}
	QuantityValidationError = _type.ErrorType{
		Code:    http.StatusBadRequest,
		Message: "Quantity is required, it can not be lower than 1.",
	}
	ProductIdNotValidError = _type.ErrorType{
		Code:    http.StatusBadRequest,
		Message: "A product with the Product Id you entered could not be found. Please enter a valid id.",
	}
	QuantityNotValidError = _type.ErrorType{
		Code:    http.StatusBadRequest,
		Message: "The product is out of stock for the amount entered. Please reduce the number of products you enter.",
	}
	ProductExistInCartError = _type.ErrorType{
		Code:    http.StatusBadRequest,
		Message: "The product you entered is already in your cart.",
	}
	ProductNotExistInCartError = _type.ErrorType{
		Code:    http.StatusBadRequest,
		Message: "The product you entered is not in your cart. Please enter a valid productId.",
	}
)
