package models

import (
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
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
	log.Info("Get Product updated at field : %v.", p.UpdatedAt)
	return p.ProductName
}

func (p *Product) GetProductPrice() decimal.Decimal {
	log.Info("Get Product Price field : %v.", p.Price)
	return p.Price
}

func (p *Product) GetProductStockNumber() int {
	log.Info("Get Product StockNumber field : %v.", p.StockNumber)
	return p.StockNumber
}

func (p *Product) GetProductCategoryId() string {
	log.Info("Get Product CategoryId field : %v.", p.CategoryId)
	return p.CategoryId
}

func (p *Product) GetProductUnitsOnCart() int {
	log.Info("Get Product UnitsOnCart field : %v.", p.UnitsOnCart)
	return p.UnitsOnCart
}

func (p *Product) SetProductId(id string) {
	p.Id = id
	log.Info("Set Product Id field : %v.", p.Id)
}

func (p *Product) SetProductUnitsOnCart(units int) {
	p.UnitsOnCart = units
	log.Info("Set Product UnitsOnCart field : %v.", p.UnitsOnCart)
}

func (p *Product) SetProductStockNumber(stock int) {
	p.StockNumber = stock
	log.Info("Set Product StockNumber field : %v.", p.StockNumber)
}

func (p *Product) SetProductUpdatedAt() {
	p.UpdatedAt = time.Now()
	log.Info("Set Product UpdatedAt field : %v.", p.UpdatedAt)
}
