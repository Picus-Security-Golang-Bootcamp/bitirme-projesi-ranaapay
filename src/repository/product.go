package repository

import (
	"PicusFinalCase/src/models"
	"PicusFinalCase/src/pkg/errorHandler"
	"errors"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"sync"
	"time"
)

var IsDeletedFilterVar = map[string]interface{}{
	"is_deleted": false,
}

type ProductRepository struct {
	db  *gorm.DB
	mux *sync.RWMutex
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	productRepo := ProductRepository{
		db:  db,
		mux: &sync.RWMutex{},
	}
	productRepo.migrations()
	return &productRepo
}

// FindProductById Finds the product based on the entered product id parameter. Returns product.
func (r *ProductRepository) FindProductById(id string) *models.Product {
	r.mux.Lock()
	defer r.mux.Unlock()

	var product models.Product

	result := r.db.Where("id = ?", id).Where(IsDeletedFilterVar).First(&product)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Errorf("Find Cart Error : %s", result.Error.Error())
		return nil
	}

	return &product
}

// CreateProduct Creates a product in the database. Returns product id.
func (r *ProductRepository) CreateProduct(product models.Product) string {

	result := r.db.Create(&product)
	if result.Error != nil {
		log.Errorf("Create Cart Error : %s", result.Error.Error())
		return ""
	}

	return product.Id
}

// DeleteProduct Deletes the product based on the entered product id parameter. Returns raw affected.
func (r *ProductRepository) DeleteProduct(id string) int64 {
	r.mux.Lock()
	defer r.mux.Unlock()

	result := r.db.Model(models.Product{}).Where("id = ?", id).Where(IsDeletedFilterVar).Updates(models.Product{
		Base: models.Base{
			DeletedAt: time.Now(),
			IsDeleted: true,
		},
	})
	if result.Error != nil {
		return 0
	}
	if result.RowsAffected == 0 {
		return 0
	}
	return result.RowsAffected
}

func (r *ProductRepository) migrations() {
	if err := r.db.AutoMigrate(&models.Product{}); err != nil {
		log.Errorf("Product Migration Error : %v", err)
		errorHandler.Panic(errorHandler.DBMigrateError)
	}
}

// FindProducts Finds the products based on the entered searchFilter, sortOpt, offset, pageSize parameter.
//Returns total number of product and product slice.
func (r *ProductRepository) FindProducts(searchFilter map[string]interface{}, sortOpt string, offset int, pageSize int) (int, []models.Product) {

	//Declare two channels for countDocument functions.
	c := make(chan int64)
	errChan := make(chan error)

	//calls countDocument function in goroutine.
	go r.countDocument(searchFilter, c, errChan)

	var products []models.Product
	result := r.db.Preload("Category").Order(sortOpt).Where(searchFilter).Where(IsDeletedFilterVar).Limit(pageSize).Offset(offset).Find(&products)
	if result.Error != nil {
		log.Errorf("Find Product Error : %s", result.Error.Error())
		return 0, nil
	}

	//If countDocument function return error with error channel returns nil values.
	countErr := <-errChan
	if countErr != nil {
		log.Errorf("Find Product Count Error : %s", countErr.Error())
		return 0, nil
	}

	//gets countDocument function response total product number.
	total := <-c

	return int(total), products
}

// UpdateProduct Updates the product based on the entered product and update options parameter. Returns updated product.
func (r *ProductRepository) UpdateProduct(product models.Product, options map[string]interface{}) (*models.Product, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	result := r.db.Model(&product).Where(IsDeletedFilterVar).Updates(options)
	if result.Error != nil {
		log.Errorf("Update Product Error : %s", result.Error.Error())
		return nil, result.Error
	}

	return &product, nil
}

//countDocument Finds the product number based on the entered searchFilter parameter.
////If gorm db count error occurs error result Returns total number of product and product slice.
func (r *ProductRepository) countDocument(filter map[string]interface{}, c chan int64, errChan chan error) {

	var count int64
	result := r.db.Model(&models.Product{}).Where(filter).Where(IsDeletedFilterVar).Count(&count)
	if result.Error != nil {
		log.Errorf("Count Document Error : %s", result.Error.Error())
		errChan <- result.Error
	}

	log.Info("Error channel closing. ")
	close(errChan)

	log.Info("Count result sending to channel. Count value : %d ", count)
	c <- count
}
