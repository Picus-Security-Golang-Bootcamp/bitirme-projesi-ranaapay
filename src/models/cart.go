package models

import (
	"github.com/shopspring/decimal"
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
	UserId         string          `json:"userId"`
	TotalCartPrice decimal.Decimal `json:"totalCartPrice"`
	CartDetails    []CartDetails
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
}

func (c *Cart) SetTotalCartPrice(total decimal.Decimal) {
	c.TotalCartPrice = total
}

func (c *Cart) SetUpdatedAt() {
	c.UpdatedAt = time.Now()
}

func (d *CartDetails) SetDetailTotalPrice(total decimal.Decimal) {
	d.DetailTotalPrice = total
}

func (d *CartDetails) SetCartId(cartId string) {
	d.CartId = cartId
}

func (d *CartDetails) SetUpdatedAt() {
	d.UpdatedAt = time.Now()
}
