package requestType

import (
	"PicusFinalCase/src/models"
	"PicusFinalCase/src/pkg/errorHandler"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
)

type ProductRequestType struct {
	ProductName string  `json:"productName"`
	Price       float64 `json:"price"`
	StockNumber int     `json:"stockNumber"`
	UnitsOnCart int     `json:"unitsOnCart"`
	CategoryId  string  `json:"categoryId"`
}

func (req ProductRequestType) RequestToProductType() models.Product {
	log.Info("Created product type according to request.")
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
		log.Error("Product request ProductName is empty.")
		errorHandler.Panic(errorHandler.ProductNameValidationError)
	}
	if req.Price <= 0 {
		log.Error("Product request price is lower than one.")
		errorHandler.Panic(errorHandler.PriceValidationError)
	}
	if req.StockNumber <= 0 {
		log.Error("Product request StockNumber is lower than one.")
		errorHandler.Panic(errorHandler.StockNumberValidationError)
	}
	if req.CategoryId == "" {
		log.Error("Product request CategoryId is empty.")
		errorHandler.Panic(errorHandler.CategoryIdValidationError)
	}
	if req.UnitsOnCart < 0 {
		log.Error("Product request UnitsOnCart is lower than zero.")
		errorHandler.Panic(errorHandler.UnitsOnCartValidationError)
	}
}
