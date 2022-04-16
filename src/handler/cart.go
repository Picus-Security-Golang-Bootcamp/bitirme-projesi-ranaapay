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

// AddToCart
//@Summary       Add Product To Cart
// @Description  add product to the cart
// @Tags         carts
// @Accept       json
// @Produce      json
// @Param 	     requestType.CartDetailsRequestType body requestType.CartDetailsRequestType true "For add product to the basket"
//@Success       201  {object}  responseType.ResponseType
// @Failure		 400 {object} 	_type.ErrorType
// @Failure		 500 {object} 	_type.ErrorType
// @Router       /cart [post]
// AddToCart Users who are logged into the system and whose
//token has not expired can add their products to the basket.
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

// ListCartItems
//@Summary       Show cart items
// @Description  get cart items by userId
// @Tags         carts
// @Accept       json
// @Produce      json
// @Success      200  {object}  responseType.ResponseType
// @Router       /cart [get]
// ListCartItems Users list the products they add to their cart.
func (h *CartHandler) ListCartItems(c *gin.Context) {
	userId, _ := c.Get("id")
	res := h.cartService.ListCartItems(userId.(string))
	cartRes := responseType.NewCartResponseType(*res)
	c.JSON(http.StatusOK, responseType.NewResponseType(http.StatusOK, cartRes))
}

// UpdateCartItems
//@Summary       Update CartItems
// @Description  Update user carts cartItems in database
// @Tags         carts
// @Accept       json
// @Produce      json
// @Param 	     requestType.CartDetailsRequestType body requestType.CartDetailsRequestType true "For update a cart item"
//@Success       200  {object}  responseType.ResponseType
// @Failure		 400 {object} 	_type.ErrorType
// @Failure		 500 {object} 	_type.ErrorType
// @Router       /cart [put]
// UpdateCartItems Users update the quantity of products added to their cart.
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

// DeleteCartItems
//@Summary       Delete a cart item
// @Description  delete cart item by productId
// @Tags         carts
// @Accept       json
// @Produce      json
// @Param        productId   path      string  true  "Product ID"
// @Success      200 {object}   responseType.ResponseType
//@Failure       400  {object}  _type.ErrorType
// @Failure		 500 {object} 	_type.ErrorType
// @Router       /cart/{productId} [delete]
//Users in the admin role delete products
// DeleteCartItems Users delete products added to their cart.
func (h *CartHandler) DeleteCartItems(c *gin.Context) {
	userId, _ := c.Get("id")
	productId := c.Param("productId")
	h.cartService.DeleteCartDetail(userId.(string), productId)
	c.JSON(http.StatusOK, responseType.NewResponseType(http.StatusOK, true))
	return
}
