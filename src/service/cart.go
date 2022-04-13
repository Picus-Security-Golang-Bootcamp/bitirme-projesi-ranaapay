package service

import (
	"PicusFinalCase/src/models"
	"PicusFinalCase/src/pkg/errorHandler"
	"PicusFinalCase/src/repository"
	"github.com/shopspring/decimal"
)

type CartService struct {
	cartRepo    *repository.CartRepository
	productRepo *repository.ProductRepository
}

func NewCartService(cartRepo *repository.CartRepository, prodRepo *repository.ProductRepository) *CartService {
	return &CartService{
		cartRepo:    cartRepo,
		productRepo: prodRepo,
	}
}

func (s *CartService) AddToCart(userId string, cartDetail *models.CartDetails) *models.CartDetails {
	product := s.productRepo.FindProductById(cartDetail.ProductId)
	if product == nil {
		errorHandler.Panic(errorHandler.ProductIdNotValidError)
	}
	if product.StockNumber < int(cartDetail.ProductQuantity) {
		errorHandler.Panic(errorHandler.QuantityNotValidError)
	}
	cart := s.findUserCart(userId)
	isExistProduct := findIfProductExistInCart(cartDetail.ProductId, cart.CartDetails)
	if isExistProduct == true {
		errorHandler.Panic(errorHandler.ProductExistInCartError)
	}
	detailPrice := product.Price.Mul(decimal.NewFromInt(cartDetail.ProductQuantity))
	cartDetail.SetDetailTotalPrice(detailPrice)
	cartDetail.SetCartId(cart.Id)
	detailId := s.cartRepo.CreateCartDetail(*cartDetail)
	if detailId == "" {
		errorHandler.Panic(errorHandler.DBCreateError)
	}
	cartTotal := cart.TotalCartPrice.Add(cartDetail.DetailTotalPrice)
	cart.SetTotalCartPrice(cartTotal)
	s.updateCart(cart)
	return cartDetail
}

func (s *CartService) findUserCart(id string) *models.Cart {
	cart := s.cartRepo.FindUserCart(id)
	if cart == nil {
		cart = s.createUserCart(id)
	}
	return cart
}

func (s *CartService) createUserCart(id string) *models.Cart {
	cart := s.cartRepo.CreateUserCart(id)
	if cart == nil {
		errorHandler.Panic(errorHandler.DBCreateError)
	}
	return cart
}

func (s *CartService) updateCart(cart *models.Cart) {
	cart.SetUpdatedAt()
	updateOptions := models.Cart{
		TotalCartPrice: cart.TotalCartPrice,
		Base:           models.Base{UpdatedAt: cart.UpdatedAt},
	}
	rawEffected := s.cartRepo.UpdateUserCart(cart.Id, updateOptions)
	if rawEffected == 0 {
		errorHandler.Panic(errorHandler.DBUpdateError)
	}
}

func findIfProductExistInCart(productId string, details []models.CartDetails) bool {
	for _, detail := range details {
		if productId == detail.ProductId {
			return true
		}
	}
	return false
}