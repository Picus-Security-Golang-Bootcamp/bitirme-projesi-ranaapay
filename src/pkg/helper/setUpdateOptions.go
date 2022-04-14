package helper

import "PicusFinalCase/src/models"

const (
	ProductNameVar = "product_name"
	CategoryIdVar  = "category_id"
	PriceVar       = "price"
	StockNumberVar = "stock_number"
	UnitsOnCartVar = "units_on_cart"
	UpdatedAtVar   = "updated_at"
)

func SetProductUpdateOptions(product models.Product) map[string]interface{} {
	updateOptions := map[string]interface{}{
		ProductNameVar: product.GetProductName(),
		CategoryIdVar:  product.GetProductCategoryId(),
		PriceVar:       product.GetProductPrice(),
		StockNumberVar: product.GetProductStockNumber(),
		UnitsOnCartVar: product.GetProductUnitsOnCart(),
		UpdatedAtVar:   product.GetUpdatedAt(),
	}
	return updateOptions
}
