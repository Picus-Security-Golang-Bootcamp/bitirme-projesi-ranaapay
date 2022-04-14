package repository

import (
	"PicusFinalCase/src/models"
	"PicusFinalCase/src/pkg/errorHandler"
	"errors"
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

func (r *ProductRepository) FindProductById(id string) *models.Product {
	r.mux.Lock()
	defer r.mux.Unlock()
	var product models.Product
	result := r.db.Where("id = ?", id).Where(IsDeletedFilterVar).First(&product)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &product
}

func (r *ProductRepository) CreateProduct(product models.Product) string {
	result := r.db.Create(&product)
	if result.Error != nil {
		return ""
	}
	return product.Id
}

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
		errorHandler.Panic(errorHandler.DBMigrateError)
	}
}

func (r *ProductRepository) FindProducts(searchFilter map[string]interface{}, sortOpt string, offset int, pageSize int) (int, []models.Product) {
	c := make(chan int64)
	errChan := make(chan error)
	go r.countDocument(searchFilter, c, errChan)

	var products []models.Product
	result := r.db.Preload("Category").Order(sortOpt).Where(searchFilter).Where(IsDeletedFilterVar).Limit(pageSize).Offset(offset).Find(&products)
	if result.Error != nil {
		return 0, nil
	}
	countErr := <-errChan
	if countErr != nil {
		return 0, nil
	}
	total := <-c
	return int(total), products
}

func (r *ProductRepository) UpdateProduct(product models.Product, options map[string]interface{}) (*models.Product, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	result := r.db.Model(&product).Where(IsDeletedFilterVar).Updates(options)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &product, nil
}

func (r *ProductRepository) countDocument(filter map[string]interface{}, c chan int64, errChan chan error) {
	var count int64
	result := r.db.Model(&models.Product{}).Where(filter).Where(IsDeletedFilterVar).Count(&count)
	if result.Error != nil {
		errChan <- result.Error
	}
	close(errChan)
	c <- count
}
