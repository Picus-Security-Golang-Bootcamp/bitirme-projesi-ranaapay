package responseType

import (
	"PicusFinalCase/src/models"
	log "github.com/sirupsen/logrus"
)

type OrderResponseType struct {
	CartId      string           `json:"cartId"`
	UserId      string           `json:"userId"`
	IsCancelled bool             `json:"isCancelled"`
	Cart        CartResponseType `json:"cartResponseType"`
}

type OrdersResponseType struct {
	TotalOrders int                 `json:"totalOrders"`
	Orders      []OrderResponseType `json:"orders"`
}

func NewOrderResponseType(order models.Order) OrderResponseType {
	log.Info("Created OrderResponseType according to Order.")
	return OrderResponseType{
		CartId:      order.CartId,
		UserId:      order.UserId,
		IsCancelled: order.IsCancelled,
		Cart:        NewCartResponseType(order.Cart),
	}
}

func NewOrdersResponseType(orders []models.Order) OrdersResponseType {
	var ordersResList []OrderResponseType
	for _, order := range orders {
		ordersResList = append(ordersResList, NewOrderResponseType(order))
	}

	log.Info("Created OrdersResponseType according to order slice.")
	return OrdersResponseType{
		TotalOrders: len(ordersResList),
		Orders:      ordersResList,
	}
}
