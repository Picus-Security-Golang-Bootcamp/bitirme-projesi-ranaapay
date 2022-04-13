package handler

import (
	"PicusFinalCase/src/handler/requestType"
	"PicusFinalCase/src/handler/responseType"
	"PicusFinalCase/src/pkg/config"
	"PicusFinalCase/src/pkg/errorHandler"
	"PicusFinalCase/src/pkg/middleware"
	"PicusFinalCase/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CartHandler struct {
	cartService *service.CartService
}

func NewCartHandler(r *gin.RouterGroup, config config.JWTConfig, cartService *service.CartService) {
	h := &CartHandler{
		cartService: cartService,
	}

	r.POST("", middleware.AuthMiddleware(config.SecretKey), h.AddToCart)
}

func (h *CartHandler) AddToCart(c *gin.Context) {
	userId, ok := c.Get("id")
	if !ok {
		errorHandler.Panic(errorHandler.NotAuthorizedError)
	}
	var reqDetail requestType.CartDetailsRequestType
	if err := c.Bind(&reqDetail); err != nil {
		errorHandler.Panic(errorHandler.BindError)
	}
	reqDetail.ValidateCartDetailsRequest()
	cartDetail := reqDetail.RequestToDetailType()
	res := h.cartService.AddToCart(userId.(string), cartDetail)
	c.JSON(http.StatusOK, responseType.NewCartDetailResponseType(*res))
	return
}
