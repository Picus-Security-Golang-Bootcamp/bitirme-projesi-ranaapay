package repository

import (
	"PicusFinalCase/src/models"
	"PicusFinalCase/src/pkg/errorHandler"
	log "github.com/sirupsen/logrus"
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
		log.Errorf("User Migration Error : %v", err)
		errorHandler.Panic(errorHandler.DBMigrateError)
	}
	if err := r.db.AutoMigrate(&models.Cart{}); err != nil {
		log.Errorf("User Migration Error : %v", err)
		errorHandler.Panic(errorHandler.DBMigrateError)
	}
}

// FindUserCart Finds the cart based on the entered userId parameter. Returns cart.
func (r *CartRepository) FindUserCart(userId string) *models.Cart {
	r.mux.Lock()
	defer r.mux.Unlock()

	var cart models.Cart
	result := r.db.Preload("CartDetails").Where("user_id = ?", userId).First(&cart)
	if result.Error != nil {
		log.Errorf("Find Cart Error : %s", result.Error.Error())
		return nil
	}
	return &cart
}

// FindUserCartById Finds the cart based on the entered cartId parameter. Returns cart.
func (r *CartRepository) FindUserCartById(cartId string, isCompleted bool) *models.Cart {
	r.mux.Lock()
	defer r.mux.Unlock()

	var cart models.Cart

	result := r.db.Preload("CartDetails").Where("id = ? AND is_completed = ?", cartId, isCompleted).First(&cart)
	if result.Error != nil {
		log.Errorf("Find Cart Error : %s", result.Error.Error())
		return nil
	}

	return &cart
}

// CreateUserCart Creates a cart in the database. Returns cart.
func (r *CartRepository) CreateUserCart(id string) *models.Cart {
	r.mux.Lock()
	defer r.mux.Unlock()

	var cart models.Cart
	cart.SetUserId(id)

	result := r.db.Create(&cart)
	if result.Error != nil {
		log.Errorf("Create Cart Error : %s", result.Error.Error())
		return nil
	}

	return &cart
}

// UpdateUserCart Updates the cart based on the entered cart id parameter. Returns updated cart rows affected
func (r *CartRepository) UpdateUserCart(id string, options models.Cart) int64 {
	r.mux.Lock()
	defer r.mux.Unlock()

	result := r.db.Model(&models.Cart{}).Where("id = ?", id).Updates(options)
	if result.Error != nil {
		log.Errorf("Update Cart Error : %s", result.Error.Error())
		return 0
	}

	return result.RowsAffected
}

// CreateCartDetail Creates a cart details in the database. Returns created detail id.
func (r *CartRepository) CreateCartDetail(detail models.CartDetails) string {

	result := r.db.Create(&detail)
	if result.Error != nil {
		log.Errorf("Create Cart Error : %s", result.Error.Error())
		return ""
	}

	return detail.Id
}

// UpdateUserCartDetail Updates the cart detail based on the entered cart detail id parameter. Returns updated cart rows affected
func (r *CartRepository) UpdateUserCartDetail(id string, options models.CartDetails) int64 {

	result := r.db.Model(&models.CartDetails{}).Where("id = ?", id).Updates(options)
	if result.Error != nil {
		log.Errorf("Update Cart Error : %s", result.Error.Error())
		return 0
	}

	return result.RowsAffected
}

// DeleteCartDetails Deletes the cart detail based on the entered cart detail parameter. Returns is deleted.
func (r *CartRepository) DeleteCartDetails(detail models.CartDetails) bool {
	result := r.db.Unscoped().Delete(&detail)
	if result.Error != nil {
		log.Errorf("Delete Cart Error : %s", result.Error.Error())
		return false
	}
	return true
}
