package service

import (
	"PicusFinalCase/src/models"
	"PicusFinalCase/src/pkg/errorHandler"
	"PicusFinalCase/src/pkg/helper"
	"PicusFinalCase/src/repository"
	"time"
)

type OrderService struct {
	orderRepo   *repository.OrderRepository
	productRepo *repository.ProductRepository
	cartRepo    *repository.CartRepository
}

func NewOrderService(orderRepo *repository.OrderRepository, prodRepo *repository.ProductRepository, cartRepo *repository.CartRepository) *OrderService {
	return &OrderService{
		orderRepo:   orderRepo,
		productRepo: prodRepo,
		cartRepo:    cartRepo,
	}
}

func (s *OrderService) CreateOrder(userId string) models.Order {

	//Finds users cart by given userId. If cannot be found gives an error.
	//If cannot be found or the cart has not any cartDetail gives an error.
	cart := s.cartRepo.FindUserCart(userId, false)
	if cart == nil {
		errorHandler.Panic(errorHandler.CartNotFoundError)
	}
	if len(cart.CartDetails) == 0 {
		errorHandler.Panic(errorHandler.CartNotContainCartDetailError)
	}

	//Updates products unitsOnCart, stockNumber, and updatedAt fields.
	s.updateProduct(cart.CartDetails)

	//Update carts isCompleted, and updatedAt fields.
	updateOptions := models.Cart{
		IsCompleted: true,
		Base:        models.Base{UpdatedAt: time.Now()},
	}
	rawEffected := s.cartRepo.UpdateUserCart(cart.Id, updateOptions)
	if rawEffected == 0 {
		errorHandler.Panic(errorHandler.DBUpdateError)
	}

	//Initialized order fields and create it.
	var order models.Order
	order.SetUserId(userId)
	order.SetCartId(cart.Id)

	id := s.orderRepo.CreateOrder(order)
	if id == "" {
		errorHandler.Panic(errorHandler.DBCreateError)
	}
	return order
}

func (s *OrderService) updateProduct(cartDetails []models.CartDetails) {

	for _, detail := range cartDetails {

		//Find the product that its id matches the cartDetail productId.
		product := s.productRepo.FindProductById(detail.ProductId)
		if product == nil {
			errorHandler.Panic(errorHandler.ProductDeletedError)
		}

		//Calculating and Updating products unitsOnCart, updatedAt, and stockNumber fields.
		productUnitsOnCart := product.UnitsOnCart - int(detail.ProductQuantity)
		product.SetProductUnitsOnCart(productUnitsOnCart)

		product.SetUpdatedAt()

		productStockNum := product.StockNumber - int(detail.ProductQuantity)
		product.SetProductStockNumber(productStockNum)

		_, err := s.productRepo.UpdateProduct(*product, helper.SetProductUpdateOptions(*product))
		if err != nil {
			errorHandler.Panic(errorHandler.InternalServerError)
		}
	}
}
