package requestType

import (
	"PicusFinalCase/src/models"
	"PicusFinalCase/src/pkg/errorHandler"
	"github.com/shopspring/decimal"
)

type ProductRequestType struct {
	ProductName string  `json:"productName"`
	Price       float64 `json:"price"`
	StockNumber int     `json:"stockNumber"`
	UnitsOnCart int     `json:"unitsOnCart"`
	CategoryId  string  `json:"categoryId"`
}

func (req ProductRequestType) RequestToProductType() models.Product {
	return models.Product{
		ProductName: req.ProductName,
		Price:       decimal.NewFromFloat(req.Price),
		StockNumber: req.StockNumber,
		UnitsOnCart: req.UnitsOnCart,
		CategoryId:  req.CategoryId,
	}
}

func (req ProductRequestType) ValidateProductRequest() {
	if req.ProductName == "" {
		errorHandler.Panic(errorHandler.ProductNameValidationError)
	}
	if req.Price <= 0 {
		errorHandler.Panic(errorHandler.PriceValidationError)
	}
	if req.StockNumber <= 0 {
		errorHandler.Panic(errorHandler.StockNumberValidationError)
	}
	if req.CategoryId == "" {
		errorHandler.Panic(errorHandler.CategoryIdValidationError)
	}
	if req.UnitsOnCart < 0 {
		errorHandler.Panic(errorHandler.UnitsOnCartValidationError)
	}
}
