package responseType

type ResponseType struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}

func NewResponseType(code int, message interface{}) *ResponseType {
	return &ResponseType{
		Code:    code,
		Message: message,
	}
}
