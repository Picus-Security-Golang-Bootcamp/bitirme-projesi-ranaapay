package responseType

import "PicusFinalCase/src/models"

type CategoryResponseType struct {
	CategoryName string           `json:"categoryName"`
	Description  string           `json:"description"`
	Product      []models.Product `json:"products,omitempty"`
}

func NewCategoryResponseType(category models.Category) CategoryResponseType {
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
	return categoryRes
}
