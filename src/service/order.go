package service

import (
	"PicusFinalCase/src/models"
	"PicusFinalCase/src/pkg/errorHandler"
	"PicusFinalCase/src/pkg/helper"
	"PicusFinalCase/src/repository"
	"time"
)

const (
	CreateVar = "order create"
	CancelVar = "order cancel"
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

func (s *OrderService) CreateOrder(userId string) string {

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
	s.updateProduct(CreateVar, cart.CartDetails)

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
	return id
}

func (s *OrderService) ListOrders(userId string) []models.Order {
	res := s.orderRepo.FindUserOrders(userId)
	if len(res) == 0 {
		errorHandler.Panic(errorHandler.NotFoundError)
	}
	return res
}

func (s *OrderService) CancelOrder(orderId string) {

	//Find order by its id
	order := s.orderRepo.FindOrderById(orderId)
	if order == nil {
		errorHandler.Panic(errorHandler.NotFoundError)
	}

	//Check if 14 days have passed since the order was created.
	toAdd := 14 * 24 * time.Hour
	limitTime := order.CreatedAt.Add(toAdd)
	if order.CreatedAt.After(limitTime) {
		errorHandler.Panic(errorHandler.OrderCanNotCancelledError)
	}

	//Cancel order by its id
	if res := s.orderRepo.CancelOrderById(orderId); res == 0 {
		errorHandler.Panic(errorHandler.NotFoundError)
	}

	//Finds users cart by given its. If cannot be found gives an error.
	cart := s.cartRepo.FindUserCartById(order.CartId, true)
	if cart == nil {
		errorHandler.Panic(errorHandler.CartNotFoundError)
	}

	//Update carts isDeleted, and deletedAt fields.
	updateOptions := models.Cart{
		Base: models.Base{DeletedAt: time.Now(), IsDeleted: true},
	}
	rawEffected := s.cartRepo.UpdateUserCart(cart.Id, updateOptions)
	if rawEffected == 0 {
		errorHandler.Panic(errorHandler.DBUpdateError)
	}

	//Updates products unitsOnCart, stockNumber, and updatedAt fields.
	s.updateProduct(CancelVar, cart.CartDetails)

}

func (s *OrderService) updateProduct(options string, cartDetails []models.CartDetails) {

	for _, detail := range cartDetails {

		//Find the product that its id matches the cartDetail productId.
		product := s.productRepo.FindProductById(detail.ProductId)
		if product == nil {
			errorHandler.Panic(errorHandler.ProductDeletedError)
		}

		//Calculating and Updating products unitsOnCart, updatedAt, and stockNumber fields depending on order crud operation.
		var productUnitsOnCart int
		var productStockNum int

		if options == CreateVar {
			productUnitsOnCart = product.UnitsOnCart - int(detail.ProductQuantity)
			productStockNum = product.StockNumber - int(detail.ProductQuantity)
		} else if options == CancelVar {
			productUnitsOnCart = product.GetProductUnitsOnCart()
			productStockNum = productStockNum + int(detail.ProductQuantity)
		}

		product.SetProductUnitsOnCart(productUnitsOnCart)
		product.SetUpdatedAt()
		product.SetProductStockNumber(productStockNum)

		_, err := s.productRepo.UpdateProduct(*product, helper.SetProductUpdateOptions(*product))
		if err != nil {
			errorHandler.Panic(errorHandler.InternalServerError)
		}
	}
}
