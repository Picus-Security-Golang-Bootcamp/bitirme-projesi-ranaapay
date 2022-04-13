package requestType

import (
	"PicusFinalCase/src/models"
	"PicusFinalCase/src/pkg/errorHandler"
)

type CartDetailsRequestType struct {
	ProductId       string `json:"productId"`
	ProductQuantity int64  `json:"productQuantity"`
}

func (req *CartDetailsRequestType) ValidateCartDetailsRequest() {
	if req.ProductId == "" {
		errorHandler.Panic(errorHandler.ProductIdValidationError)
	}
	if req.ProductQuantity <= 0 {
		errorHandler.Panic(errorHandler.QuantityValidationError)
	}
}

func (req *CartDetailsRequestType) RequestToDetailType() *models.CartDetails {
	return &models.CartDetails{
		ProductId:       req.ProductId,
		ProductQuantity: req.ProductQuantity,
	}
}
