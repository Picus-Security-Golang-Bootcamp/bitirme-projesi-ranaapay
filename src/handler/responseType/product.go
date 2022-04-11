package responseType

import "PicusFinalCase/src/models"

type ProductResponseType struct {
	ProductName string               `json:"productName"`
	Price       string               `json:"price"`
	StockNumber int                  `json:"stockNumber"`
	Category    CategoryResponseType `json:"category"`
}

func NewProductResponseType(product models.Product) ProductResponseType {
	return ProductResponseType{
		ProductName: product.ProductName,
		Price:       product.Price.String(),
		StockNumber: product.StockNumber,
		Category:    NewCategoryResponseType(product.Category),
	}
}

func NewProductsResponseType(products []models.Product) []ProductResponseType {
	var productsRes []ProductResponseType
	for _, product := range products {
		productsRes = append(productsRes, NewProductResponseType(product))
	}
	return productsRes
}
