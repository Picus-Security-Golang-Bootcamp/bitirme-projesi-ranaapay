package service

import (
	"PicusFinalCase/src/models"
	"PicusFinalCase/src/repository"
)

type ProductService struct {
	productRepo *repository.ProductRepository
}

func NewProductService(productRepo *repository.ProductRepository) *ProductService {
	return &ProductService{
		productRepo: productRepo,
	}
}

func (s *ProductService) CreateProduct(product models.Product) string {
	productId := s.productRepo.CreateProduct(product)
	return productId
}

func (s ProductService) DeleteProduct(id string) {
	s.productRepo.DeleteProduct(id)
}
