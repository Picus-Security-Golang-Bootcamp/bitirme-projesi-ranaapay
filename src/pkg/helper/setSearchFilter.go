package helper

import (
	"PicusFinalCase/src/pkg/errorHandler"
	log "github.com/sirupsen/logrus"
	"net/url"
	"strconv"
)

// SetSearchFilter It takes the url.values and converts them to the map required for filtering.
func SetSearchFilter(values url.Values) map[string]interface{} {

	var filter = make(map[string]interface{})

	for i, v := range values {

		filter[i] = v[0]

		switch {
		case i == StockNumberVar:
			stockNumber, err := strconv.Atoi(filter[StockNumberVar].(string))
			if err != nil {
				log.Error("Stock number could not convert. Err : %v", err)
				errorHandler.Panic(errorHandler.ConvertError)
			}
			filter[i] = stockNumber
		}
	}

	log.Info("Created search filter for pagination.")
	return filter
}
