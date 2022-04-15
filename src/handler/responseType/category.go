package responseType

import (
	"PicusFinalCase/src/models"
	log "github.com/sirupsen/logrus"
)

type CategoryResponseType struct {
	CategoryName string           `json:"categoryName"`
	Description  string           `json:"description"`
	Product      []models.Product `json:"products,omitempty"`
}

func NewCategoryResponseType(category models.Category) CategoryResponseType {
	log.Info("Created CategoryResponseType according to Category.")
	return CategoryResponseType{
		CategoryName: category.CategoryName,
		Description:  category.Description,
		Product:      category.Product,
	}
}

func NewCategoriesResponseType(categories []models.Category) []CategoryResponseType {
	var categoryRes []CategoryResponseType
	for _, category := range categories {
		categoryRes = append(categoryRes, NewCategoryResponseType(category))
	}

	log.Info("Created CategoryResponseType slice according to Category slice.")
	return categoryRes
}
