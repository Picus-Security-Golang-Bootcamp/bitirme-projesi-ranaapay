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
	MarshalError = _type.ErrorType{
		Code:    http.StatusInternalServerError,
		Message: "Marshal Error",
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
	CSVReadError = _type.ErrorType{
		Code:    http.StatusInternalServerError,
		Message: "CSV Read Error",
	}
	FormFileError = _type.ErrorType{
		Code:    http.StatusInternalServerError,
		Message: "Form File Error",
	}
	DBMigrateError = _type.ErrorType{
		Code:    http.StatusInternalServerError,
		Message: "DBMigrate error.",
	}
	DBCreateError = _type.ErrorType{
		Code:    http.StatusInternalServerError,
		Message: "DBCreate error.",
	}
	DBUpdateError = _type.ErrorType{
		Code:    http.StatusInternalServerError,
		Message: "DBUpdate error.",
	}
	DBDeleteError = _type.ErrorType{
		Code:    http.StatusInternalServerError,
		Message: "DBDelete error.",
	}
	NotAuthorizedError = _type.ErrorType{
		Code:    http.StatusUnauthorized,
		Message: "Not Authorized Error",
	}
	ForbiddenError = _type.ErrorType{
		Code:    http.StatusForbidden,
		Message: "You are not authorized to access the requested resource.",
	}
	ConvertError = _type.ErrorType{
		Code:    http.StatusBadRequest,
		Message: "Convert error.",
	}
	ProductDeletedError = _type.ErrorType{
		Code:    http.StatusBadRequest,
		Message: "The product in your cart has been deleted.",
	}
	InternalServerError = _type.ErrorType{
		Code:    http.StatusInternalServerError,
		Message: "Internal server error.",
	}
	ClientError = _type.ErrorType{
		Code:    http.StatusInternalServerError,
		Message: "HttpClient error.",
	}
)
