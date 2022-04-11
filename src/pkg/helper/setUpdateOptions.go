package helper

import "PicusFinalCase/src/models"

const (
	ProductNameVar = "product_name"
	CategoryIdVar  = "category_id"
	PriceVar       = "price"
	StockNumberVar = "stock_number"
)

func SetProductUpdateOptions(product models.Product) map[string]interface{} {
	updateOptions := map[string]interface{}{
		ProductNameVar: product.GetProductName(),
		CategoryIdVar:  product.GetProductCategoryId(),
		PriceVar:       product.GetProductPrice(),
		StockNumberVar: product.GetProductStockNumber(),
	}
	return updateOptions
}
