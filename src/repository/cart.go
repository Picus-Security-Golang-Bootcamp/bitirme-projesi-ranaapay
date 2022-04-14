package repository

import (
	"PicusFinalCase/src/models"
	"PicusFinalCase/src/pkg/errorHandler"
	"gorm.io/gorm"
	"sync"
)

type CartRepository struct {
	db  *gorm.DB
	mux *sync.RWMutex
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	cartRepo := CartRepository{
		db:  db,
		mux: &sync.RWMutex{},
	}
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

func (r *CartRepository) FindUserCart(id string, isCompleted bool) *models.Cart {
	r.mux.Lock()
	defer r.mux.Unlock()

	var cart models.Cart
	result := r.db.Preload("CartDetails").Where("user_id = ? AND is_completed = ?", id, isCompleted).First(&cart)
	if result.Error != nil {
		return nil
	}
	return &cart
}

func (r *CartRepository) CreateUserCart(id string) *models.Cart {
	r.mux.Lock()
	defer r.mux.Unlock()

	var cart models.Cart
	cart.SetUserId(id)
	result := r.db.Create(&cart)
	if result.Error != nil {
		return nil
	}
	return &cart
}

func (r *CartRepository) UpdateUserCart(id string, options models.Cart) int64 {
	r.mux.Lock()
	defer r.mux.Unlock()

	result := r.db.Model(&models.Cart{}).Where("id = ?", id).Updates(options)
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

func (r *CartRepository) UpdateUserCartDetail(id string, options models.CartDetails) int64 {
	result := r.db.Model(&models.CartDetails{}).Where("id = ?", id).Updates(options)
	if result.Error != nil {
		return 0
	}
	return result.RowsAffected
}

func (r *CartRepository) DeleteCartDetails(detail models.CartDetails) bool {
	result := r.db.Unscoped().Delete(&detail)
	if result.Error != nil {
		return false
	}
	return true
}
