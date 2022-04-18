package repository

import (
	"PicusFinalCase/src/models"
	"PicusFinalCase/src/pkg/errorHandler"
	"errors"
	log "github.com/sirupsen/logrus"
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
		log.Errorf("User Migration Error : %v", err)
		errorHandler.Panic(errorHandler.DBMigrateError)
	}
}

// CreateOrder Creates an order in the database. Returns created order id.
func (r *OrderRepository) CreateOrder(order models.Order) string {

	result := r.db.Create(&order)
	if result.Error != nil {
		log.Errorf("Create Cart Error : %s", result.Error.Error())
		return ""
	}

	return order.Id
}

// FindUserOrders Finds the order based on the entered userId parameter. Return orders.
func (r *OrderRepository) FindUserOrders(userId string) []models.Order {
	var orders []models.Order

	result := r.db.Preload("Cart.CartDetails").Where("user_id = ?", userId).Find(&orders)
	if result.Error != nil {
		log.Errorf("Find Order Error : %s", result.Error.Error())
		return nil
	}

	return orders
}

// FindOrderById Finds the order based on the entered orderId parameter. Return order.
func (r *OrderRepository) FindOrderById(orderId string) *models.Order {

	var order models.Order

	result := r.db.Where("id = ?", orderId).Where("is_cancelled = ?", false).First(&order)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Errorf("Find Order Error : %s", result.Error.Error())
		return nil
	}

	return &order
}

// CancelOrderById Updates the order is deleted field based on the entered order id parameter. Returns updated order rows affected
func (r *OrderRepository) CancelOrderById(id string) int64 {

	result := r.db.Model(models.Order{}).Where("id = ?", id).Where(IsDeletedFilterVar).Updates(models.Order{
		Base: models.Base{
			DeletedAt: time.Now(),
			IsDeleted: true,
		},
		IsCancelled: true,
	})
	if result.Error != nil {
		log.Errorf("Update Order Error : %s", result.Error.Error())
		return 0
	}
	if result.RowsAffected == 0 {
		log.Errorf("No Rows Affected : %s", result.Error.Error())
		return 0
	}

	return result.RowsAffected
}
