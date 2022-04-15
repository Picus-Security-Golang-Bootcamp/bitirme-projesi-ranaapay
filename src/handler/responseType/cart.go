package responseType

import (
	"PicusFinalCase/src/models"
	log "github.com/sirupsen/logrus"
)

type CartDetailResponseType struct {
	ProductId        string `json:"productId"`
	ProductQuantity  int64  `json:"productQuantity"`
	DetailTotalPrice string `json:"detailTotalPrice"`
}

type CartResponseType struct {
	TotalCartPrice string `json:"totalCartPrice"`
	CartDetails    []CartDetailResponseType
}

func NewCartDetailResponseType(detail models.CartDetails) CartDetailResponseType {
	log.Info("Created CartDetailResponseType according to CartDetails.")
	return CartDetailResponseType{
		ProductId:        detail.ProductId,
		ProductQuantity:  detail.ProductQuantity,
		DetailTotalPrice: detail.DetailTotalPrice.String(),
	}
}

func NewCartResponseType(cart models.Cart) CartResponseType {

	var detailsRes []CartDetailResponseType
	for _, detail := range cart.CartDetails {
		detailsRes = append(detailsRes, NewCartDetailResponseType(detail))
	}

	log.Info("Created CartResponseType according to Cart.")
	return CartResponseType{
		TotalCartPrice: cart.TotalCartPrice.String(),
		CartDetails:    detailsRes,
	}
}
