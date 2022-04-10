package models

import "github.com/shopspring/decimal"

type Product struct {
	Base
	ProductName string          `json:"productName"`
	Price       decimal.Decimal `json:"price"`
	StockNumber int             `json:"stockNumber"`
	UnitsOnCart int             `json:"unitsOnCart"`
	CategoryId  string          `json:"categoryId"`
	Category    Category        `json:"category,omitempty"`
}

func (Product) TableName() string {
	//default table name
	return "products"
}
