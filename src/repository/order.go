package repository

import (
	"PicusFinalCase/src/models"
	"PicusFinalCase/src/pkg/errorHandler"
	"gorm.io/gorm"
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
