package helper

import (
	"PicusFinalCase/src/models"
	"PicusFinalCase/src/pkg/errorHandler"
	"encoding/csv"
	"io"
)

func ReadCSVForCategory(file io.Reader) []models.Category {
	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		errorHandler.Panic(errorHandler.CSVReadError)
	}
	var categories []models.Category
	for _, line := range records[1:] {
		categories = append(categories, models.Category{
			CategoryName: line[0],
			Description:  line[1],
		})
	}
	return categories
}
