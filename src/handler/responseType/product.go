package responseType

import (
	"PicusFinalCase/src/models"
	log "github.com/sirupsen/logrus"
)

type ProductWithCategoryResponseType struct {
	ProductName string               `json:"productName"`
	Price       string               `json:"price"`
	StockNumber int                  `json:"stockNumber"`
	UnitsOnCart int                  `json:"unitsOnCart"`
	Category    CategoryResponseType `json:"category"`
}

type ProductResponseType struct {
	ProductName string `json:"productName"`
	Price       string `json:"price"`
	StockNumber int    `json:"stockNumber"`
	UnitsOnCart int    `json:"unitsOnCart"`
	CategoryId  string `json:"categoryId"`
}

func NewProductWithCategoryResponseType(product models.Product) ProductWithCategoryResponseType {
	log.Info("Created ProductWithCategoryResponseType according to Product.")
	return ProductWithCategoryResponseType{
		ProductName: product.ProductName,
		Price:       product.Price.String(),
		StockNumber: product.StockNumber,
		UnitsOnCart: product.UnitsOnCart,
		Category:    NewCategoryResponseType(product.Category),
	}
}

func NewProductResponseType(product models.Product) ProductResponseType {
	log.Info("Created ProductResponseType according to Product.")
	return ProductResponseType{
		ProductName: product.ProductName,
		Price:       product.Price.String(),
		UnitsOnCart: product.UnitsOnCart,
		StockNumber: product.StockNumber,
		CategoryId:  product.CategoryId,
	}
}

func NewProductsResponseType(products []models.Product) []ProductWithCategoryResponseType {
	var productsRes []ProductWithCategoryResponseType
	for _, product := range products {
		productsRes = append(productsRes, NewProductWithCategoryResponseType(product))
	}

	log.Info("Created ProductWithCategoryResponseType slice according to product slice.")
	return productsRes
}
