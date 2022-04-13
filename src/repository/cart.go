package repository

import (
	"PicusFinalCase/src/models"
	"PicusFinalCase/src/pkg/errorHandler"
	"gorm.io/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	cartRepo := CartRepository{db: db}
	cartRepo.migrations()
	return &cartRepo
}

func (r *CartRepository) migrations() {
	if err := r.db.AutoMigrate(&models.CartDetails{}); err != nil {
		errorHandler.Panic(errorHandler.DBMigrateError)
	}
	if err := r.db.AutoMigrate(&models.Cart{}); err != nil {
		errorHandler.Panic(errorHandler.DBMigrateError)
	}
}

func (r *CartRepository) FindUserCart(id string) *models.Cart {
	var cart models.Cart
	result := r.db.Preload("CartDetails").Where(models.Cart{UserId: id}).First(&cart)
	if result.Error != nil {
		return nil
	}
	return &cart
}

func (r *CartRepository) CreateUserCart(id string) *models.Cart {
	var cart models.Cart
	cart.SetUserId(id)
	result := r.db.Create(&cart)
	if result.Error != nil {
		return nil
	}
	return &cart
}

func (r *CartRepository) UpdateUserCart(id string, options models.Cart) int64 {
	result := r.db.Model(&models.Cart{}).Where("id = ?", id).Updates(options)
	//result := r.db.Save(&cart)
	if result.Error != nil {
		return 0
	}
	return result.RowsAffected
}

func (r *CartRepository) CreateCartDetail(detail models.CartDetails) string {
	result := r.db.Create(&detail)
	if result.Error != nil {
		return ""
	}
	return detail.Id
}
