package models

import log "github.com/sirupsen/logrus"

type Order struct {
	Base
	CartId      string `json:"cartId"`
	UserId      string `json:"userId"`
	IsCancelled bool   `json:"isCancelled"`
	Cart        Cart
}

func (Order) TableName() string {
	//default table name
	return "orders"
}

func (o *Order) SetCartId(id string) {
	o.CartId = id
	log.Info("Set Order CartId field : %v.", o.CartId)
}
func (o *Order) SetUserId(id string) {
	o.UserId = id
	log.Info("Set Order UserId field : %v.", o.UserId)
}
