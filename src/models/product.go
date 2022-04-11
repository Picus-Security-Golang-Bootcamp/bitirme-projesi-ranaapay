package models

import (
	"github.com/shopspring/decimal"
	"time"
)

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

func (p *Product) GetProductName() string {
	return p.ProductName
}
func (p *Product) GetProductPrice() decimal.Decimal {
	return p.Price
}
func (p *Product) GetProductStockNumber() int {
	return p.StockNumber
}
func (p *Product) GetProductCategoryId() string {
	return p.CategoryId
}

func (p *Product) SetProductId(id string) {
	p.Id = id
}
func (p *Product) SetProductUpdatedAt() {
	p.UpdatedAt = time.Now()
}
