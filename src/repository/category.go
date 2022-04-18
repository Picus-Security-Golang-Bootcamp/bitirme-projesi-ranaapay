package repository

import (
	"PicusFinalCase/src/models"
	"PicusFinalCase/src/pkg/errorHandler"
	"errors"
	log "github.com/sirupsen/logrus"
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

// CreateCategories Creates categories in the database. Returns result.
func (r *CategoryRepository) CreateCategories(categories []models.Category) bool {

	result := r.db.Create(&categories)
	if result.Error != nil {
		log.Errorf("Create Category Error : %s", result.Error.Error())
		return false
	}
	return true
}

// FindCategories Find the categories in database. Return category slice.
func (r *CategoryRepository) FindCategories() *[]models.Category {

	var categories []models.Category

	result := r.db.Find(&categories).Where("is_deleted = ?", "false")
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Errorf("Find Categories Error : %s", result.Error.Error())
		return nil
	}

	return &categories
}

// FindCategory Find category by id based on the entered category id parameter. Returns category.
func (r *CategoryRepository) FindCategory(id string) *models.Category {

	var category models.Category
	result := r.db.Where(&models.Category{
		Base: models.Base{
			Id:        id,
			IsDeleted: false,
		},
	}).First(&category)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Errorf("Find Category Error : %s", result.Error.Error())
		return nil
	}

	return &category
}

func (r *CategoryRepository) migrations() {

	if err := r.db.AutoMigrate(&models.Category{}); err != nil {
		log.Errorf("User Migration Error : %v", err)
		errorHandler.Panic(errorHandler.DBMigrateError)
	}
}
