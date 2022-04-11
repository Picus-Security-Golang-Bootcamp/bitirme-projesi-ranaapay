package service

import (
	"PicusFinalCase/src/models"
	"PicusFinalCase/src/pkg/errorHandler"
	"PicusFinalCase/src/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(productRepo *repository.ProductRepository) *ProductService {
	return &ProductService{
		repo: productRepo,
	}
}

func (s *ProductService) CreateProduct(product models.Product) string {
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
