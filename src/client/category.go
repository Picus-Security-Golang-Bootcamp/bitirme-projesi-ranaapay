package client

import (
	"PicusFinalCase/src/handler/responseType"
	"PicusFinalCase/src/pkg/config"
	"PicusFinalCase/src/pkg/errorHandler"
	_type "PicusFinalCase/src/pkg/errorHandler/type"
	"context"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type CategoryClient struct {
	BaseURI    string
	HTTPClient *http.Client
}

func NewCategoryClient(cfg config.ServerConfig, basePath string) *CategoryClient {

	baseURI := fmt.Sprintf("http://localhost:%s/%s%s", cfg.Port, cfg.RoutePrefix, basePath)

	return &CategoryClient{
		BaseURI: baseURI,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *CategoryClient) FindCategoryIsValid(id string) {
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, fmt.Sprintf("%s:%s", c.BaseURI, id), nil)
	if err != nil {
		log.Errorf("NewRequest error : %s", err)
		errorHandler.Panic(errorHandler.ClientError)
	}
	_ = c.sendRequest(req)
}

func (c *CategoryClient) sendRequest(req *http.Request) interface{} {
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
