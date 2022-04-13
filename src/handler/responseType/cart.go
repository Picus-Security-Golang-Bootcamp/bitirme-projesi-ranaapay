package responseType

import (
	"PicusFinalCase/src/models"
)

type CartDetailResponseType struct {
	ProductId        string `json:"productId"`
	ProductQuantity  int64  `json:"productQuantity"`
	DetailTotalPrice string `json:"detailTotalPrice"`
}

func NewCartDetailResponseType(detail models.CartDetails) CartDetailResponseType {
	return CartDetailResponseType{
		ProductId:        detail.ProductId,
		ProductQuantity:  detail.ProductQuantity,
		DetailTotalPrice: detail.DetailTotalPrice.String(),
	}
}
