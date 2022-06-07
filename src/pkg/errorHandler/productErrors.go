package errorHandler

import (
	_type "PicusFinalCase/src/pkg/errorHandler/type"
	"net/http"
)

var (
	ProductNameValidationError = _type.ErrorType{
		Code:    http.StatusBadRequest,
		Message: "ProductName is required, it can not be empty.",
	}
	CategoryIdValidationError = _type.ErrorType{
		Code:    http.StatusBadRequest,
		Message: "CategoryId is required, it can not be empty.",
	}
	CategoryIdNotValidError = _type.ErrorType{
		Code:    http.StatusBadRequest,
		Message: "A category with the Category Id you entered could not be found. Please enter a valid id.",
	}
	StockNumberValidationError = _type.ErrorType{
		Code:    http.StatusBadRequest,
		Message: "StockNumber is required, it can not be lower than 1.",
	}
	PriceValidationError = _type.ErrorType{
		Code:    http.StatusBadRequest,
		Message: "Price is required, it can not be lower than 1.",
	}
	UnitsOnCartValidationError = _type.ErrorType{
		Code:    http.StatusBadRequest,
		Message: "UnitsOnCart is required, it can not be lower than 0.",
	}
	ProductNotFoundError = _type.ErrorType{
		Code:    http.StatusNotFound,
		Message: "Product not found in database.",
	}
)
