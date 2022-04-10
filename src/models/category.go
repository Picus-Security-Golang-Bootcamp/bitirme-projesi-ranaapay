package models

type Category struct {
	Base
	CategoryName string    `json:"categoryName"`
	Description  string    `json:"description"`
	Product      []Product `json:"products,omitempty"`
}

func (Category) TableName() string {
	//default table name
	return "categories"
}
