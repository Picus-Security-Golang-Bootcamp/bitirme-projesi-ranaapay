package responseType

import "PicusFinalCase/src/models"

type ProductWithCategoryResponseType struct {
	ProductName string               `json:"productName"`
	Price       string               `json:"price"`
	StockNumber int                  `json:"stockNumber"`
	Category    CategoryResponseType `json:"category"`
}

type ProductResponseType struct {
	ProductName string `json:"productName"`
	Price       string `json:"price"`
	StockNumber int    `json:"stockNumber"`
}

func NewProductWithCategoryResponseType(product models.Product) ProductWithCategoryResponseType {
	return ProductWithCategoryResponseType{
		ProductName: product.ProductName,
		Price:       product.Price.String(),
		StockNumber: product.StockNumber,
		Category:    NewCategoryResponseType(product.Category),
	}
}

func NewProductResponseType(product models.Product) ProductResponseType {
	return ProductResponseType{
		ProductName: product.ProductName,
		Price:       product.Price.String(),
		StockNumber: product.StockNumber,
	}
}

func NewProductsResponseType(products []models.Product) []ProductWithCategoryResponseType {
	var productsRes []ProductWithCategoryResponseType
	for _, product := range products {
		productsRes = append(productsRes, NewProductWithCategoryResponseType(product))
	}
	return productsRes
}
