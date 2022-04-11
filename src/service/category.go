package service

import (
	"PicusFinalCase/src/models"
	"PicusFinalCase/src/pkg/helper"
	"PicusFinalCase/src/repository"
	"io"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) CreateCategories(file io.Reader) {
	categories := helper.ReadCSVForCategory(file)
	s.repo.CreateCategories(categories)
}

func (s *CategoryService) FindCategories() *[]models.Category {
	return s.repo.FindCategories()
}

func (s *CategoryService) FindCategory(id string) *models.Category {
	return s.repo.FindCategory(id)
}
