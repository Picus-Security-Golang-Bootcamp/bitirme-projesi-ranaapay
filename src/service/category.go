package service

import (
	"PicusFinalCase/src/models"
	"PicusFinalCase/src/pkg/errorHandler"
	"PicusFinalCase/src/pkg/helper"
	"PicusFinalCase/src/repository"
	log "github.com/sirupsen/logrus"
	"io"
)

type CategoryService struct {
	repo repository.CategoryRepo
}

func NewCategoryService(repo repository.CategoryRepo) *CategoryService {
	return &CategoryService{repo: repo}
}

// CreateCategories Read given file as []models.Category. Creates read categories.
func (s *CategoryService) CreateCategories(file io.Reader) {

	categories := helper.ReadCSVForCategory(file)

	result := s.repo.CreateCategories(categories)
	if result != true {
		log.Error("Something happened when creating categories.")
		errorHandler.Panic(errorHandler.DBCreateError)
	}
}

// FindCategories Finds all categories. If no category found throws error.
// Returns []models.Category.
func (s *CategoryService) FindCategories() *[]models.Category {

	res := s.repo.FindCategories()
	if len(*res) == 0 {
		log.Error("No category found.")
		errorHandler.Panic(errorHandler.NotFoundError)
	}
	return res
}

// FindCategory Finds category given id field. If category can not found
//throws error. Return *models.Category.
func (s *CategoryService) FindCategory(id string) *models.Category {

	category := s.repo.FindCategory(id)
	if category == nil {
		log.Error("The category request id does not exist in the database.")
		errorHandler.Panic(errorHandler.CategoryIdNotValidError)
	}

	return category
}
