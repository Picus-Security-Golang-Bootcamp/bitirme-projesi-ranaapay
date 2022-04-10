package repository

import (
	"PicusFinalCase/src/models"
	"PicusFinalCase/src/pkg/errorHandler"
	"errors"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	categoryRepo := CategoryRepository{db: db}
	categoryRepo.migrations()
	return &categoryRepo
}

func (r CategoryRepository) CreateCategories(categories []models.Category) {
	result := r.db.Create(&categories)
	if result.Error != nil {
		errorHandler.Panic(errorHandler.DBCreateError)
	}
}

func (r *CategoryRepository) migrations() {
	if err := r.db.AutoMigrate(&models.Category{}); err != nil {
		errorHandler.Panic(errorHandler.DBMigrateError)
	}
}

func (r *CategoryRepository) FindCategories() *[]models.Category {
	var categories []models.Category
	result := r.db.Find(&categories).Where("is_deleted = ?", "false")
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		errorHandler.Panic(errorHandler.NotFoundError)
	}
	return &categories
}
