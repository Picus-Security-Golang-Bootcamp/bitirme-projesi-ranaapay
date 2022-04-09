package errorHandler

import (
	_type "PicusFinalCase/src/pkg/errorHandler/type"
	"net/http"
)

var (
	NotFoundError = _type.ErrorType{
		Code:    http.StatusNotFound,
		Message: "Record not found.",
	}
	BindError = _type.ErrorType{
		Code:    http.StatusBadRequest,
		Message: "Unable to bind the request body.",
	}
	UnmarshalError = _type.ErrorType{
		Code:    http.StatusInternalServerError,
		Message: "Unmarshal Error : Unable to decode into struct",
	}
	GormOpenError = _type.ErrorType{
		Code:    http.StatusInternalServerError,
		Message: "Failed to open db session matching the entered values",
	}
	SqlDBError = _type.ErrorType{
		Code:    http.StatusInternalServerError,
		Message: "SQL DB Error",
	}
	SqlDBPingError = _type.ErrorType{
		Code:    http.StatusInternalServerError,
		Message: "SQL DB Ping Error",
	}
	ConfigNotFoundError = _type.ErrorType{
		Code:    http.StatusInternalServerError,
		Message: "Config File Not Found Error",
	}
	ConvertIdError = _type.ErrorType{
		Code:    http.StatusBadRequest,
		Message: "Id is not valid. Please write valid Id.",
	}
	DBMigrateError = _type.ErrorType{
		Code:    http.StatusInternalServerError,
		Message: "DBMigrate error.",
	}
	DBCreateError = _type.ErrorType{
		Code:    http.StatusInternalServerError,
		Message: "DBCreate error.",
	}
	ConvertError = _type.ErrorType{
		Code:    http.StatusBadRequest,
		Message: "Convert error.",
	}
	QuantityValidationError = _type.ErrorType{
		Code:    http.StatusBadRequest,
		Message: "Quantity is required, it can not be lower than 1.",
	}
	PriceValidationError = _type.ErrorType{
		Code:    http.StatusBadRequest,
		Message: "Price is required, it can not be lower than 1.",
	}
	InternalServerError = _type.ErrorType{
		Code:    http.StatusInternalServerError,
		Message: "Internal server error.",
	}
)
