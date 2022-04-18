package service

import (
	"PicusFinalCase/src/models"
	"PicusFinalCase/src/pkg/errorHandler"
	"PicusFinalCase/src/pkg/helper"
	"PicusFinalCase/src/repository"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
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

// CreateOrder Checks according to given userId has any cart in database. Updates product
//and cartDetails fields. Create order based on given userId and its cart.
func (s *OrderService) CreateOrder(userId string) string {

	//Finds users cart by given userId. If cannot be found gives an error.
	//If cannot be found or the cart has not any cartDetail gives an error.
	cart := s.cartRepo.FindUserCart(userId)
	if cart == nil {
		log.Error("The cart with userId does not exist in the database.")
		errorHandler.Panic(errorHandler.CartNotFoundError)
	}
	if len(cart.CartDetails) == 0 {
		log.Error("There are no products in your cart.")
		errorHandler.Panic(errorHandler.CartNotContainCartDetailError)
	}

	//Updates products unitsOnCart, stockNumber, and updatedAt fields. Dependent to the option
	//parameter update cart details and returns total price of order according to founded products.
	totalOrderPrice := s.updateProductandCartDetails(CreateVar, cart.CartDetails)

	//Update carts isCompleted, and updatedAt fields.
	updateOptions := models.Cart{
		IsCompleted: true,
		Base:        models.Base{UpdatedAt: time.Now()},
	}
	rawEffected := s.cartRepo.UpdateUserCart(cart.Id, updateOptions)
	if rawEffected == 0 {
		log.Error("Somethings went wrong when updating cart. ")
		errorHandler.Panic(errorHandler.DBUpdateError)
	}

	//Initialized order fields and create it.
	var order models.Order

	order.SetUserId(userId)
	order.SetCartId(cart.Id)
	order.SetOrderTotalPrice(totalOrderPrice)

	id := s.orderRepo.CreateOrder(order)
	if id == "" {
		log.Error("Something happened when creating order.")
		errorHandler.Panic(errorHandler.DBCreateError)
	}

	return id
}

// ListOrders Find and return all orders based on userId.
func (s *OrderService) ListOrders(userId string) []models.Order {
	res := s.orderRepo.FindUserOrders(userId)
	if len(res) == 0 {
		log.Error("The order with userId : %s does not exist in the database.", userId)
		errorHandler.Panic(errorHandler.NotFoundError)
	}

	return res
}

// CancelOrder Finds order based on given orderId. If 14 days haven't passed since
//the order was created 14, order cancelled. Update models.
func (s *OrderService) CancelOrder(orderId string) {

	//Find order by its id
	order := s.orderRepo.FindOrderById(orderId)
	if order == nil {
		log.Error("The order with orderId does not exist in the database.")
		errorHandler.Panic(errorHandler.NotFoundError)
	}

	//Check if 14 days have passed since the order was created.
	toAdd := 14 * 24 * time.Hour
	limitTime := order.CreatedAt.Add(toAdd)
	if order.CreatedAt.After(limitTime) {
		log.Error("14 days have passed since the order was created, the cancellation failed.")
		errorHandler.Panic(errorHandler.OrderCanNotCancelledError)
	}

	//Cancel order by its id
	if res := s.orderRepo.CancelOrderById(orderId); res == 0 {
		log.Error("Given orderId does not contain in order.")
		errorHandler.Panic(errorHandler.DBDeleteError)
	}

	//Finds users cart by given its. If cannot be found gives an error.
	cart := s.cartRepo.FindUserCartById(order.CartId, true)
	if cart == nil {
		log.Error("Given cartId does not contain in cart.")
		errorHandler.Panic(errorHandler.CartNotFoundError)
	}

	//Update carts isDeleted, and deletedAt fields.
	updateOptions := models.Cart{
		Base: models.Base{DeletedAt: time.Now(), IsDeleted: true},
	}
	rawEffected := s.cartRepo.UpdateUserCart(cart.Id, updateOptions)
	if rawEffected == 0 {
		log.Error("Somethings went wrong when updating cart. ")
		errorHandler.Panic(errorHandler.DBUpdateError)
	}

	//Updates products unitsOnCart, stockNumber, and updatedAt fields.
	s.updateProductandCartDetails(CancelVar, cart.CartDetails)

}

//updateProductandCartDetails Updates products and cartDetails based on given options and cart details. Returns order total price.
func (s *OrderService) updateProductandCartDetails(options string, cartDetails []models.CartDetails) decimal.Decimal {

	var orderTotalPrice decimal.Decimal = decimal.NewFromInt(0)

	for _, detail := range cartDetails {

		//Find the product that its id matches the cartDetail productId.
		product := s.productRepo.FindProductById(detail.ProductId)
		if product == nil {
			log.Error("Somethings went wrong when finding product. ")
			errorHandler.Panic(errorHandler.ProductDeletedError)
		}

		//Calculating and Updating products unitsOnCart, updatedAt, and stockNumber fields depending on order crud operation.
		var productUnitsOnCart int
		var productStockNum int

		if options == CreateVar {

			productUnitsOnCart = product.UnitsOnCart - int(detail.ProductQuantity)
			productStockNum = product.StockNumber - int(detail.ProductQuantity)

			//Updating cart details according to product price number. Calculate total price for order.
			detailPrice := product.Price.Mul(decimal.NewFromInt(detail.ProductQuantity))

			detail.SetDetailTotalPrice(detailPrice)
			detail.SetUpdatedAt()

			updateOptions := models.CartDetails{
				Base: models.Base{
					UpdatedAt: detail.UpdatedAt,
				},
				DetailTotalPrice: detail.DetailTotalPrice,
			}

			s.cartRepo.UpdateUserCartDetail(detail.Id, updateOptions)

			orderTotalPrice = orderTotalPrice.Add(detail.DetailTotalPrice)

		} else if options == CancelVar {

			productUnitsOnCart = product.GetProductUnitsOnCart()
			productStockNum = productStockNum + int(detail.ProductQuantity)

		}

		product.SetProductUnitsOnCart(productUnitsOnCart)
		product.SetUpdatedAt()
		product.SetProductStockNumber(productStockNum)

		_, err := s.productRepo.UpdateProduct(*product, helper.SetProductUpdateOptions(*product))
		if err != nil {
			log.Error("Somethings went wrong when updating product. ")
			errorHandler.Panic(errorHandler.InternalServerError)
		}

	}

	return orderTotalPrice
}
