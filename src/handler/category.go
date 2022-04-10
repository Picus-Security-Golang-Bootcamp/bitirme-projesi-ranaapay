package handler

import (
	"PicusFinalCase/src/handler/responseType"
	"PicusFinalCase/src/models"
	"PicusFinalCase/src/pkg/config"
	"PicusFinalCase/src/pkg/errorHandler"
	"PicusFinalCase/src/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CategoryHandler struct {
	service *service.CategoryService
}

func NewCategoryHandler(r *gin.RouterGroup, config config.JWTConfig, categoryService *service.CategoryService) {
	h := &CategoryHandler{service: categoryService}
	r.POST("/", h.createCategories)
	//r.POST("/", middleware.AuthMiddleware(config.SecretKey), h.createCategories)
	r.GET("/", h.FindCategories)
}

func (h *CategoryHandler) createCategories(c *gin.Context) {
	userRole, ok := c.Get("role")
	if !ok {
		fmt.Println("handler : 28")
		errorHandler.Panic(errorHandler.NotAuthorizedError)
	}
	if userRole != models.Admin {
		errorHandler.Panic(errorHandler.ForbiddenError)
	}
	file, _, err := c.Request.FormFile("csvFile")
	if err != nil {
		errorHandler.Panic(errorHandler.FormFileError)
	}
	h.service.CreateCategories(file)
	c.JSON(http.StatusCreated, responseType.NewResponseType(http.StatusCreated, true))
}

func (h *CategoryHandler) FindCategories(c *gin.Context) {
	categories := h.service.FindCategories()
	categoryRes := responseType.NewCategoriesResponseType(*categories)
	c.JSON(http.StatusOK, responseType.NewResponseType(http.StatusOK, categoryRes))
}
