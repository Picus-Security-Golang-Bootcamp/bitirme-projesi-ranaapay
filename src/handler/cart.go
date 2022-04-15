package handler

import (
	"PicusFinalCase/src/handler/requestType"
	"PicusFinalCase/src/handler/responseType"
	"PicusFinalCase/src/models"
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

	r.POST("", middleware.AuthMiddleware(config.SecretKey, models.Customer), h.AddToCart)
	r.GET("", middleware.AuthMiddleware(config.SecretKey, models.Customer), h.ListCartItems)
	r.PUT("", middleware.AuthMiddleware(config.SecretKey, models.Customer), h.UpdateCartItems)
	r.DELETE("/:productId", middleware.AuthMiddleware(config.SecretKey, models.Customer), h.DeleteCartItems)
}

func (h *CartHandler) AddToCart(c *gin.Context) {
	userId, _ := c.Get("id")
	var reqDetail requestType.CartDetailsRequestType
	if err := c.Bind(&reqDetail); err != nil {
		errorHandler.Panic(errorHandler.BindError)
	}
	reqDetail.ValidateCartDetailsRequest()
	cartDetail := reqDetail.RequestToDetailType()
	res := h.cartService.AddToCart(userId.(string), cartDetail)
	detailRes := responseType.NewCartDetailResponseType(*res)
	c.JSON(http.StatusCreated, responseType.NewResponseType(http.StatusCreated, detailRes))
	return
}

func (h *CartHandler) ListCartItems(c *gin.Context) {
	userId, _ := c.Get("id")
	res := h.cartService.ListCartItems(userId.(string))
	cartRes := responseType.NewCartResponseType(*res)
	c.JSON(http.StatusOK, responseType.NewResponseType(http.StatusOK, cartRes))
}

func (h CartHandler) UpdateCartItems(c *gin.Context) {
	userId, _ := c.Get("id")
	var reqDetail requestType.CartDetailsRequestType
	if err := c.Bind(&reqDetail); err != nil {
		errorHandler.Panic(errorHandler.BindError)
	}
	reqDetail.ValidateCartDetailsRequest()
	cartDetail := reqDetail.RequestToDetailType()
	res := h.cartService.UpdateCartDetail(userId.(string), cartDetail)
	detailRes := responseType.NewCartDetailResponseType(*res)
	c.JSON(http.StatusOK, responseType.NewResponseType(http.StatusOK, detailRes))
	return
}

func (h *CartHandler) DeleteCartItems(c *gin.Context) {
	userId, _ := c.Get("id")
	productId := c.Param("productId")
	h.cartService.DeleteCartDetail(userId.(string), productId)
	c.JSON(http.StatusOK, responseType.NewResponseType(http.StatusOK, true))
	return
}
