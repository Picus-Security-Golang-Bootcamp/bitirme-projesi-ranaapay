package handler

import (
	"PicusFinalCase/src/handler/requestType"
	"PicusFinalCase/src/handler/responseType"
	"PicusFinalCase/src/pkg/config"
	"PicusFinalCase/src/pkg/errorHandler"
	"PicusFinalCase/src/pkg/helper"
	"PicusFinalCase/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductHandler struct {
	productService  *service.ProductService
	categoryService *service.CategoryService
}

func NewProductHandler(r *gin.RouterGroup, config config.JWTConfig, productService *service.ProductService, categoryService *service.CategoryService) {
	h := ProductHandler{
		productService:  productService,
		categoryService: categoryService,
	}
	r.POST("", h.createProducts)
	r.DELETE("/:id", h.deleteProducts)
	r.GET("", h.listProducts)
	r.GET("/:id", h.findProductById)
	r.PUT("/:id", h.updateProducts)
}

func (h *ProductHandler) findProductById(c *gin.Context) {
	pId := c.Param("id")
	res := h.productService.FindByProductId(pId)
	prodRes := responseType.NewProductResponseType(*res)
	c.JSON(http.StatusOK, responseType.NewResponseType(http.StatusOK, prodRes))
}

func (h *ProductHandler) listProducts(c *gin.Context) {
	reqQueries := c.Request.URL.Query()
	sortOpt, pageNum, pageSize := helper.SetPaginationOptions(&reqQueries)
	searchFilter := helper.SetSearchFilter(reqQueries)
	total, res := h.productService.FindProducts(searchFilter, sortOpt, pageNum, pageSize)
	productsRes := responseType.NewProductsResponseType(res)
	paginationRes := responseType.NewPaginationType(pageNum, pageSize, total, productsRes)
	c.JSON(http.StatusOK, responseType.NewResponseType(http.StatusOK, paginationRes))
}

func (h *ProductHandler) createProducts(c *gin.Context) {
	/*userRole, ok := c.Get("role")
	if !ok {
		errorHandler.Panic(errorHandler.NotAuthorizedError)
	}
	if userRole != models.Admin {
		errorHandler.Panic(errorHandler.ForbiddenError)
	}*/
	var reqProduct requestType.ProductRequestType
	if err := c.Bind(&reqProduct); err != nil {
		errorHandler.Panic(errorHandler.BindError)
	}
	reqProduct.ValidateProductRequest()
	category := h.categoryService.FindCategory(reqProduct.CategoryId)
	if category == nil {
		errorHandler.Panic(errorHandler.CategoryIdNotValidError)
	}
	product := reqProduct.RequestToProductType()
	resId := h.productService.CreateProduct(product)
	c.JSON(http.StatusCreated, responseType.NewResponseType(http.StatusCreated, resId))
	return
}

func (h *ProductHandler) deleteProducts(c *gin.Context) {
	id := c.Param("id")
	h.productService.DeleteProduct(id)
	c.JSON(http.StatusOK, responseType.NewResponseType(http.StatusOK, true))
	return
}

func (h *ProductHandler) updateProducts(c *gin.Context) {
	/*userRole, ok := c.Get("role")
	if !ok {
		errorHandler.Panic(errorHandler.NotAuthorizedError)
	}
	if userRole != models.Admin {
		errorHandler.Panic(errorHandler.ForbiddenError)
	}*/
	reqId := c.Param("id")
	var reqProduct requestType.ProductRequestType
	if err := c.Bind(&reqProduct); err != nil {
		errorHandler.Panic(errorHandler.BindError)
	}
	reqProduct.ValidateProductRequest()
	category := h.categoryService.FindCategory(reqProduct.CategoryId)
	if category == nil {
		errorHandler.Panic(errorHandler.CategoryIdNotValidError)
	}
	product := reqProduct.RequestToProductType()
	product.SetProductId(reqId)
	res := h.productService.UpdateProduct(product)
	productRes := responseType.NewProductResponseType(res)
	c.JSON(http.StatusOK, responseType.NewResponseType(http.StatusOK, productRes))
	return
}
