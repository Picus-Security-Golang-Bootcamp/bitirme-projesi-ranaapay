package models

import (
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"time"
)

type CartDetails struct {
	Base
	ProductId        string          `json:"productId"`
	ProductQuantity  int64           `json:"productQuantity"`
	DetailTotalPrice decimal.Decimal `json:"detailTotalPrice"`
	CartId           string          `json:"cartId"`
}

type Cart struct {
	Base
	UserId      string `json:"userId"`
	IsCompleted bool   `json:"isCompleted"`
	CartDetails []CartDetails
}

func (CartDetails) TableName() string {
	//default table name
	return "cartDetails"
}

func (Cart) TableName() string {
	//default table name
	return "cart"
}

func (c *Cart) SetUserId(userId string) {
	c.UserId = userId
	log.Info("Set Cart UserId field : %s.", c.UserId)
}

func (c *Cart) SetUpdatedAt() {
	c.UpdatedAt = time.Now()
	log.Info("Set Cart UpdatedAt field : %v.", c.UpdatedAt)
}

func (d *CartDetails) SetDetailTotalPrice(total decimal.Decimal) {
	d.DetailTotalPrice = total
	log.Info("Set CartDetails DetailTotalPrice field : %v.", d.DetailTotalPrice)

}

func (d *CartDetails) SetCartId(cartId string) {
	d.CartId = cartId
	log.Info("Set CartDetails CartId field : %v.", d.CartId)
}

func (d *CartDetails) SetUpdatedAt() {
	d.UpdatedAt = time.Now()
	log.Info("Set CartDetails UpdatedAt field : %v.", d.UpdatedAt)
}
