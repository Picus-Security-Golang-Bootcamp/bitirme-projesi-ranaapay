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

// FindByProductId Finds product given id field. If product can not found
// throws error. Return *models.Product.
func (s *ProductService) FindByProductId(id string) *models.Product {
	res := s.repo.FindProductById(id)
	if res == nil {
		log.Error("The request id does not exist in the database.")
		errorHandler.Panic(errorHandler.NotFoundError)
	}
	return res
}

// CreateProduct Checks given product.categoryId is existed in database. If not throws
//error. Create product based on given product. Return created product.Id.
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
		log.Error("Something happened when creating product.")
		errorHandler.Panic(errorHandler.DBCreateError)
	}
	return productId
}

// DeleteProduct Deletes product according to given productId.
// If product not found throws error.
func (s *ProductService) DeleteProduct(id string) {

	if res := s.repo.DeleteProduct(id); res == 0 {
		log.Error("Given productId does not contain in product.")
		errorHandler.Panic(errorHandler.NotFoundError)
	}
}

// FindProducts Find products according to given searchFilter, sortOpt, pageSize, pageNum. And
//finds total number of products according to searchFilter. If no products found throws error.
func (s *ProductService) FindProducts(searchFilter map[string]interface{}, sortOpt string, pageNum int, pageSize int) (int, []models.Product) {
	offset := (pageNum - 1) * pageSize
	total, res := s.repo.FindProducts(searchFilter, sortOpt, offset, pageSize)
	if len(res) == 0 {
		log.Error("The request id does not exist in the database.")
		errorHandler.Panic(errorHandler.NotFoundError)
	}
	return total, res
}

// UpdateProduct Checks given product.categoryId is existed in database. If not throws
//error. Update product based on given product. Return updated product.
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
		log.Error("Somethings went wrong when updating product. ")
		errorHandler.Panic(errorHandler.InternalServerError)
	}
	if updatedProduct == nil {
		log.Error("The request product does not exist in the database.")
		errorHandler.Panic(errorHandler.NotFoundError)
	}
	return *updatedProduct
}
