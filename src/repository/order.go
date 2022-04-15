package repository

import (
	"PicusFinalCase/src/models"
	"PicusFinalCase/src/pkg/errorHandler"
	"errors"
	"gorm.io/gorm"
	"time"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	orderRepo := OrderRepository{db: db}
	orderRepo.migrations()
	return &orderRepo
}

func (r *OrderRepository) migrations() {
	if err := r.db.AutoMigrate(&models.Order{}); err != nil {
		errorHandler.Panic(errorHandler.DBMigrateError)
	}
}

func (r *OrderRepository) CreateOrder(order models.Order) string {
	result := r.db.Create(&order)
	if result.Error != nil {
		return ""
	}
	return order.Id
}

func (r *OrderRepository) FindUserOrders(userId string) []models.Order {
	var orders []models.Order
	result := r.db.Preload("Cart").Where("user_id = ?", userId).Find(&orders)
	if result.Error != nil {
		return nil
	}
	return orders
}

func (r *OrderRepository) FindOrderById(orderId string) *models.Order {
	var order models.Order
	result := r.db.Where("id = ?", orderId).Where("is_cancelled = ?", false).First(&order)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &order
}

func (r *OrderRepository) CancelOrderById(id string) int64 {
	result := r.db.Model(models.Order{}).Where("id = ?", id).Where(IsDeletedFilterVar).Updates(models.Order{
		Base: models.Base{
			DeletedAt: time.Now(),
			IsDeleted: true,
		},
		IsCancelled: true,
	})
	if result.Error != nil {
		return 0
	}
	if result.RowsAffected == 0 {
		return 0
	}
	return result.RowsAffected
}
