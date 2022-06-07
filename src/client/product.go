package client

import (
	"PicusFinalCase/src/handler/requestType"
	"PicusFinalCase/src/handler/responseType"
	"PicusFinalCase/src/models"
	"PicusFinalCase/src/pkg/config"
	"PicusFinalCase/src/pkg/errorHandler"
	_type "PicusFinalCase/src/pkg/errorHandler/type"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type ProductClient struct {
	BaseURI    string
	HTTPClient *http.Client
	Token      string
}

func NewProductClient(cfg config.ServerConfig, basePath string) *ProductClient {

	baseURI := fmt.Sprintf("http://localhost:%s/%s%s", cfg.Port, cfg.RoutePrefix, basePath)

	return &ProductClient{
		BaseURI: baseURI,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *ProductClient) FindProductById(id string) *models.Product {
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, fmt.Sprintf("%s:%s", c.BaseURI, id), nil)
	if err != nil {
		log.Errorf("NewRequest error : %s", err)
		errorHandler.Panic(errorHandler.ClientError)
	}
	res := c.sendRequest(req)
	resProd := res.(*responseType.ProductResponseType)
	price, _ := decimal.NewFromString(resProd.Price)
	product := models.Product{
		Base: models.Base{
			Id: id,
		},
		ProductName: resProd.ProductName,
		Price:       price,
		StockNumber: resProd.StockNumber,
		UnitsOnCart: resProd.UnitsOnCart,
		CategoryId:  resProd.CategoryId,
	}
	return &product
}

func (c *ProductClient) UpdateProduct(product models.Product) {

	price, _ := product.GetProductPrice().Float64()
	prodReq := requestType.ProductRequestType{
		ProductName: product.GetProductName(),
		Price:       price,
		StockNumber: product.GetProductStockNumber(),
		UnitsOnCart: product.GetProductUnitsOnCart(),
		CategoryId:  product.GetProductCategoryId(),
	}

	jsonProd, err := json.Marshal(prodReq)
	if err != nil {
		errorHandler.Panic(errorHandler.MarshalError)
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodPut, fmt.Sprintf("%s:%s", c.BaseURI, product.Id), bytes.NewBuffer(jsonProd))
	if err != nil {
		log.Errorf("NewRequest error : %s", err)
		errorHandler.Panic(errorHandler.ClientError)
	}

	req.Header.Set("Authorization", c.Token)
	_ = c.sendRequest(req)

}

func (c *ProductClient) sendRequest(req *http.Request) interface{} {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	res, doErr := c.HTTPClient.Do(req)
	if doErr != nil {
		log.Errorf("HttpClient Do error : %s", doErr)
		errorHandler.Panic(errorHandler.ClientError)
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes *_type.ErrorType
		if dErr := json.NewDecoder(res.Body).Decode(&errRes); dErr != nil {
			log.Errorf("Send request decode error message error : %s", dErr)
			errorHandler.Panic(errorHandler.ClientError)
		}
		errorHandler.Panic(_type.ErrorType{Code: errRes.Code, Message: errRes.Message})
	}
	var fullResponse responseType.ResponseType
	if err := json.NewDecoder(res.Body).Decode(&fullResponse); err != nil {
		log.Errorf("Send request response decode error : %s", err)
		errorHandler.Panic(errorHandler.ClientError)
	}
	return fullResponse.Message
}
