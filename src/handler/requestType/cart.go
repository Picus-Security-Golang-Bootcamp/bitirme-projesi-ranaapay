package requestType

import (
	"PicusFinalCase/src/models"
	"PicusFinalCase/src/pkg/errorHandler"
	log "github.com/sirupsen/logrus"
)

type CartDetailsRequestType struct {
	ProductId       string `json:"productId"`
	ProductQuantity int64  `json:"productQuantity"`
}

// ValidateCartDetailsRequest It is checked whether the fields of the incoming request comply with the expected values. If not, it throws an error.
func (req *CartDetailsRequestType) ValidateCartDetailsRequest() {
	if req.ProductId == "" {
		log.Error("CartDetail request productId is empty.")
		errorHandler.Panic(errorHandler.ProductIdValidationError)
	}
	if req.ProductQuantity <= 0 {
		log.Error("CartDetail request product quantity is lower than one.")
		errorHandler.Panic(errorHandler.QuantityValidationError)
	}
}

func (req *CartDetailsRequestType) RequestToDetailType() *models.CartDetails {

	log.Info("Created CartDetail type according to request.")
	return &models.CartDetails{
		ProductId:       req.ProductId,
		ProductQuantity: req.ProductQuantity,
	}
}
