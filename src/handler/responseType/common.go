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

type PaginationType struct {
	Page       int         `json:"page"`
	PageSize   int         `json:"pageSize"`
	TotalCount int         `json:"totalCount"`
	Items      interface{} `json:"items"`
}

func NewPaginationType(page int, pageSize int, totalCount int, items interface{}) PaginationType {
	return PaginationType{
		Page:       page,
		PageSize:   pageSize,
		TotalCount: totalCount,
		Items:      items,
	}
}
