package helper

import (
	"PicusFinalCase/src/pkg/errorHandler"
	"net/url"
	"strconv"
)

func SetSearchFilter(values url.Values) map[string]interface{} {
	var filter = make(map[string]interface{})
	for i, v := range values {
		filter[i] = v[0]
		switch {
		case i == StockNumberVar:
			stockNumber, err := strconv.Atoi(filter[StockNumberVar].(string))
			if err != nil {
				errorHandler.Panic(errorHandler.ConvertError)
			}
			filter[i] = stockNumber
		}
	}
	return filter
}
