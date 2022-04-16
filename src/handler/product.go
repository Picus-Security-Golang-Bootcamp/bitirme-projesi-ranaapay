package handler

import (
	"PicusFinalCase/src/handler/requestType"
	"PicusFinalCase/src/handler/responseType"
	"PicusFinalCase/src/models"
	"PicusFinalCase/src/pkg/config"
	"PicusFinalCase/src/pkg/errorHandler"
	"PicusFinalCase/src/pkg/helper"
	"PicusFinalCase/src/pkg/middleware"
	"PicusFinalCase/src/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type ProductHandler struct {
	productService *service.ProductService
}

func NewProductHandler(r *gin.RouterGroup, config config.JWTConfig, productService *service.ProductService) {
	h := ProductHandler{
		productService: productService,
	}
	r.POST("", middleware.AuthMiddleware(config.SecretKey, models.Admin), h.createProducts)
	r.DELETE("/:id", middleware.AuthMiddleware(config.SecretKey, models.Admin), h.deleteProducts)
	r.GET("", h.listProducts)
	r.GET("/:id", h.findProductById)
	r.PUT("/:id", middleware.AuthMiddleware(config.SecretKey, models.Admin), h.updateProducts)
}

// findProductById
//@Summary       Show a product
// @Description  get product by ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Product ID"
// @Success      200  {object}  responseType.ResponseType
// @Failure		 404 {object} 	_type.ErrorType
// @Failure		 400 {object} 	_type.ErrorType
// @Router       /product/{id} [get]
//Returns the product based on the productId.
func (h *ProductHandler) findProductById(c *gin.Context) {
	pId := c.Param("id")

	res := h.productService.FindByProductId(pId)

	prodRes := responseType.NewProductResponseType(*res)
	c.JSON(http.StatusOK, responseType.NewResponseType(http.StatusOK, prodRes))
}

// listProducts
// @Summary      List products
// @Description  get products
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        sort query string false "sort"
// @Param        page query int false "page"
// @Param        pageSize query int false "page"
// @Param		 product_name query string false "product_name"
// @Param		 category_id query string false "category_id"
// @Param		 price query int false "price"
// @Param		 stock_number query int false "stock_number"
// @Param		 units_on_cart query int false "units_on_cart"
// @Success      200  {object}  responseType.ResponseType
// @Failure		 404 {object} 	_type.ErrorType
// @Failure		 400 {object} 	_type.ErrorType
// @Router       /product 	[get]
//Users can list products without the need for role control. They can
//search according to the parameters they entered.
func (h *ProductHandler) listProducts(c *gin.Context) {

	//request parameters are made available for pagination and search.
	reqQueries := c.Request.URL.Query()
	sortOpt, pageNum, pageSize := helper.SetPaginationOptions(&reqQueries)
	searchFilter := helper.SetSearchFilter(reqQueries)

	total, res := h.productService.FindProducts(searchFilter, sortOpt, pageNum, pageSize)

	productsRes := responseType.NewProductsResponseType(res)
	paginationRes := responseType.NewPaginationType(pageNum, pageSize, total, productsRes)
	c.JSON(http.StatusOK, responseType.NewResponseType(http.StatusOK, paginationRes))
}

// createProducts
//@Summary       Create Product
// @Description  create product in database
// @Tags         products
// @Accept       json
// @Produce      json
// @Param 	     requestType.ProductRequestType body requestType.ProductRequestType true "For create a Product"
//@Success       201  {object}  responseType.ResponseType
// @Failure		 400 {object} 	_type.ErrorType
// @Failure		 500 {object} 	_type.ErrorType
// @Router       /product [post]
//Users in the admin role create individual products for the database.
func (h *ProductHandler) createProducts(c *gin.Context) {

	var reqProduct requestType.ProductRequestType
	if err := c.Bind(&reqProduct); err != nil {
		log.Error("Bind error : %s", err.Error())
		errorHandler.Panic(errorHandler.BindError)
	}

	reqProduct.ValidateProductRequest()

	product := reqProduct.RequestToProductType()

	resId := h.productService.CreateProduct(product)

	c.JSON(http.StatusCreated, responseType.NewResponseType(http.StatusCreated, resId))
	return
}

// deleteProducts
//@Summary       Delete a product
// @Description  delete product by ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Product ID"
// @Success      200 {object}   responseType.ResponseType
//@Failure       400  {object}  _type.ErrorType
// @Failure		 404 {object} 	_type.ErrorType
// @Router       /product/{id} [delete]
//Users in the admin role delete products
func (h *ProductHandler) deleteProducts(c *gin.Context) {

	id := c.Param("id")

	h.productService.DeleteProduct(id)

	c.JSON(http.StatusOK, responseType.NewResponseType(http.StatusOK, true))
	return
}

// updateProducts
//@Summary       Update Product
// @Description  Update product in database
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Product ID"
// @Param 	     requestType.ProductRequestType body requestType.ProductRequestType true "For update a Product"
//@Success       200  {object}  responseType.ResponseType
// @Failure		 400 {object} 	_type.ErrorType
// @Failure		 500 {object} 	_type.ErrorType
// @Router       /product/{id} [put]
//Users in the admin role update products
func (h *ProductHandler) updateProducts(c *gin.Context) {

	reqId := c.Param("id")

	var reqProduct requestType.ProductRequestType
	if err := c.Bind(&reqProduct); err != nil {
		log.Error("Bind error : %s", err.Error())
		errorHandler.Panic(errorHandler.BindError)
	}

	reqProduct.ValidateProductRequest()

	product := reqProduct.RequestToProductType()
	product.SetProductId(reqId)

	res := h.productService.UpdateProduct(product)

	productRes := responseType.NewProductResponseType(res)
	c.JSON(http.StatusOK, responseType.NewResponseType(http.StatusOK, productRes))
	return
}
