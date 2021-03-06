package handler

import (
	"PicusFinalCase/src/handler/responseType"
	"PicusFinalCase/src/models"
	"PicusFinalCase/src/pkg/config"
	"PicusFinalCase/src/pkg/errorHandler"
	"PicusFinalCase/src/pkg/middleware"
	"PicusFinalCase/src/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type CategoryHandler struct {
	service *service.CategoryService
}

func NewCategoryHandler(r *gin.RouterGroup, config config.JWTConfig, categoryService *service.CategoryService) {
	h := &CategoryHandler{service: categoryService}

	r.POST("", middleware.AuthMiddleware(config.SecretKey, models.Admin), h.createCategories)
	r.GET("", h.findCategories)
	r.GET("/:id", h.findCategoryById)
}

// createCategories
//@Summary       Create Categories
// @Description  admin can create categories by uploading csv file
// @Tags         category
// @Accept       multipart/form-data
// @Produce      json
// @Param        Authorization  header    string  true  "Authentication header"
// @Param        csvFile  formData  file  true  "category list"
// @Success       200  {object}  responseType.ResponseType
// @Failure		 500 {object} 	_type.ErrorType
// @Router       /category [post]
//Users in the admin role create a new category by uploading a CSV file.
func (h *CategoryHandler) createCategories(c *gin.Context) {

	file, _, err := c.Request.FormFile("csvFile")
	if err != nil {
		log.Error("Form File error : %s", err.Error())
		errorHandler.Panic(errorHandler.FormFileError)
	}

	h.service.CreateCategories(file)

	c.JSON(http.StatusCreated, responseType.NewResponseType(http.StatusCreated, true))
}

// findCategories
// @Summary      List Categories
// @Description  gets all categories in database
// @Tags         category
// @Accept       json
// @Produce      json
// @Success      200  {object}  responseType.ResponseType
// @Failure		 404 {object} 	_type.ErrorType
// @Router       /category 	[get]
//All active and not deleted categories in the database are listed.
func (h *CategoryHandler) findCategories(c *gin.Context) {

	categories := h.service.FindCategories()

	categoryRes := responseType.NewCategoriesResponseType(*categories)
	c.JSON(http.StatusOK, responseType.NewResponseType(http.StatusOK, categoryRes))
}

func (h *CategoryHandler) findCategoryById(c *gin.Context) {
	pId := c.Param("id")

	res := h.service.FindCategory(pId)

	catRes := responseType.NewCategoryResponseType(*res)
	c.JSON(http.StatusOK, responseType.NewResponseType(http.StatusOK, catRes))
}
