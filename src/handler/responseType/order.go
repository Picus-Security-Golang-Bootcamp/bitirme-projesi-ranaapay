package responseType

import "PicusFinalCase/src/models"

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
	return OrdersResponseType{
		TotalOrders: len(ordersResList),
		Orders:      ordersResList,
	}
}
