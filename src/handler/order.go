package handler

import (
	"PicusFinalCase/src/handler/responseType"
	"PicusFinalCase/src/models"
	"PicusFinalCase/src/pkg/config"
	"PicusFinalCase/src/pkg/middleware"
	"PicusFinalCase/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderHandler struct {
	service *service.OrderService
}

func NewOrderHandler(r *gin.RouterGroup, config config.JWTConfig, orderService *service.OrderService) {
	h := &OrderHandler{service: orderService}

	r.POST("", middleware.AuthMiddleware(config.SecretKey, models.Customer), h.completeOrder)
	r.GET("", middleware.AuthMiddleware(config.SecretKey, models.Customer), h.listOrders)
	r.DELETE("/:id", middleware.AuthMiddleware(config.SecretKey, models.Customer), h.cancelOrder)
}

// completeOrder
//@Summary       Create Order
// @Description  create order in database according to users cart
// @Tags         orders
// @Accept       json
// @Produce      json
//@Success       201  {object}  responseType.ResponseType
// @Failure		 400 {object} 	_type.ErrorType
// @Failure		 500 {object} 	_type.ErrorType
// @Router       /order [post]
//Users create an order with the products they add to their cart. Returns orderId.
func (h *OrderHandler) completeOrder(c *gin.Context) {

	userId, _ := c.Get("id")

	res := h.service.CreateOrder(userId.(string))

	c.JSON(http.StatusCreated, responseType.NewResponseType(http.StatusCreated, res))
}

// listOrders
// @Summary      List Orders
// @Description  get users orders
// @Tags         orders
// @Accept       json
// @Produce      json
// @Success      200  {object}  responseType.ResponseType
// @Failure		 404 {object} 	_type.ErrorType
// @Router       /order 	[get]
//Users view their past orders.
func (h *OrderHandler) listOrders(c *gin.Context) {

	userId, _ := c.Get("id")

	res := h.service.ListOrders(userId.(string))

	c.JSON(http.StatusOK, responseType.NewResponseType(http.StatusOK, responseType.NewOrdersResponseType(res)))
}

// cancelOrder
//@Summary       Cancel Order
// @Description  cancel user's order if order created has not passed 14 days
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Order ID"
// @Success      200 {object}   responseType.ResponseType
//@Failure       400  {object}  _type.ErrorType
// @Failure		 404 {object} 	_type.ErrorType
// @Failure		 500 {object} 	_type.ErrorType
// @Router       /order/{id} [delete]
//If the user's order date has not passed 14 days, the user can cancel the order. If 14 days
//have passed after the order creation date, the cancellation request will be invalid.
func (h *OrderHandler) cancelOrder(c *gin.Context) {

	orderId := c.Param("id")

	h.service.CancelOrder(orderId)

	c.JSON(http.StatusOK, responseType.NewResponseType(http.StatusOK, true))
}
