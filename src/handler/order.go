package handler

import (
	"PicusFinalCase/src/handler/responseType"
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

	r.POST("", middleware.AuthMiddleware(config.SecretKey), h.completeOrder)
}

func (h *OrderHandler) completeOrder(c *gin.Context) {
	userId, _ := c.Get("id")
	res := h.service.CreateOrder(userId.(string))

	c.JSON(http.StatusCreated, responseType.NewResponseType(http.StatusCreated, res))

}
