package service

import (
	"PicusFinalCase/src/models"
	"PicusFinalCase/src/pkg/errorHandler"
	"PicusFinalCase/src/pkg/helper"
	"PicusFinalCase/src/repository"
	log "github.com/sirupsen/logrus"
)

type ProductService struct {
	repo    *repository.ProductRepository
	catRepo *repository.CategoryRepository
}

func NewProductService(productRepo *repository.ProductRepository, catRepo *repository.CategoryRepository) *ProductService {
	return &ProductService{
		repo:    productRepo,
		catRepo: catRepo,
	}
}

func (s *ProductService) FindByProductId(id string) *models.Product {
	res := s.repo.FindProductById(id)
	if res == nil {
		errorHandler.Panic(errorHandler.NotFoundError)
	}
	return res
}

func (s *ProductService) CreateProduct(product models.Product) string {

	//Find the category that its id matches the product CategoryId. If category return
	//nil throws error.
	category := s.catRepo.FindCategory(product.CategoryId)
	if category == nil {
		log.Error("The request categoryId does not exist in the database.")
		errorHandler.Panic(errorHandler.CategoryIdNotValidError)
	}

	productId := s.repo.CreateProduct(product)
	if productId == "" {
		errorHandler.Panic(errorHandler.DBCreateError)
	}
	return productId
}

func (s *ProductService) DeleteProduct(id string) {
	if res := s.repo.DeleteProduct(id); res == 0 {
		errorHandler.Panic(errorHandler.NotFoundError)
	}
}

func (s *ProductService) FindProducts(searchFilter map[string]interface{}, sortOpt string, pageNum int, pageSize int) (int, []models.Product) {
	offset := (pageNum - 1) * pageSize
	total, res := s.repo.FindProducts(searchFilter, sortOpt, offset, pageSize)
	if len(res) == 0 {
		errorHandler.Panic(errorHandler.NotFoundError)
	}
	return total, res
}

func (s *ProductService) UpdateProduct(product models.Product) models.Product {

	//Find the category that its id matches the product CategoryId. If category return
	//nil throws error.
	category := s.catRepo.FindCategory(product.CategoryId)
	if category == nil {
		log.Error("The request categoryId does not exist in the database.")
		errorHandler.Panic(errorHandler.CategoryIdNotValidError)
	}

	product.SetProductUpdatedAt()
	updateOptions := helper.SetProductUpdateOptions(product)
	updatedProduct, err := s.repo.UpdateProduct(product, updateOptions)
	if err != nil {
		errorHandler.Panic(errorHandler.InternalServerError)
	}
	if updatedProduct == nil {
		errorHandler.Panic(errorHandler.NotFoundError)
	}
	return *updatedProduct
}
