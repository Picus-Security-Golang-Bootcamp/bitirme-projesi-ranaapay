package repository

import (
	"PicusFinalCase/src/models"
	"PicusFinalCase/src/pkg/errorHandler"
	"gorm.io/gorm"
	"time"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	productRepo := ProductRepository{db: db}
	productRepo.migrations()
	return &productRepo
}

func (r *ProductRepository) CreateProduct(product models.Product) string {
	result := r.db.Create(&product)
	if result.Error != nil {
		errorHandler.Panic(errorHandler.DBCreateError)
	}
	return product.Id
}

func (r *ProductRepository) DeleteProduct(id string) {
	result := r.db.Model(models.Product{}).Where("id = ?", id).Updates(models.Product{
		Base: models.Base{
			DeletedAt: time.Now(),
			IsDeleted: true,
		},
	})
	if result.Error != nil {
		errorHandler.Panic(errorHandler.DBDeleteError)
	}
	if result.RowsAffected == 0 {
		errorHandler.Panic(errorHandler.NotFoundError)
	}
}

func (r *ProductRepository) migrations() {
	if err := r.db.AutoMigrate(&models.Product{}); err != nil {
		errorHandler.Panic(errorHandler.DBMigrateError)
	}
}
