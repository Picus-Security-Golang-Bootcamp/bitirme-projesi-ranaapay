package helper

import (
	"PicusFinalCase/src/models"
	"PicusFinalCase/src/pkg/errorHandler"
	"encoding/csv"
	log "github.com/sirupsen/logrus"
	"io"
)

// ReadCSVForCategory reads the given file as csv. Converts the read file to the category slice.
func ReadCSVForCategory(file io.Reader) []models.Category {

	csvReader := csv.NewReader(file)

	records, err := csvReader.ReadAll()
	if err != nil {
		log.Error("Can not read csv file. Err : %v", err)
		errorHandler.Panic(errorHandler.CSVReadError)
	}

	//Converts the read lines to category type.
	var categories []models.Category

	for _, line := range records[1:] {

		categories = append(categories, models.Category{
			CategoryName: line[0],
			Description:  line[1],
		})
	}

	log.Info("Complete Read Csv File")
	return categories
}
