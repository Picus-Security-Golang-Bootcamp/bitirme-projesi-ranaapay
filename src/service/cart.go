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

func (s *CartService) ListCartItems(userId string) *models.Cart {
	return s.findUserCart(userId, false)
}

func (s *CartService) AddToCart(userId string, cartDetail *models.CartDetails) *models.CartDetails {
	product := s.productRepo.FindProductById(cartDetail.ProductId)
	if product == nil {
		errorHandler.Panic(errorHandler.ProductIdNotValidError)
	}
	if product.StockNumber < int(cartDetail.ProductQuantity) {
		errorHandler.Panic(errorHandler.QuantityNotValidError)
	}

	//Find user's cart by userId.
	cart := s.findUserCart(userId, false)

	//It checks if there is any matching productId of the cartDetail from the request in the
	//cartDetails in the user's cart. If it does match, it throws an error.
	isExistProduct := findIfProductExistInCart(cartDetail.ProductId, cart.CartDetails)
	if isExistProduct != nil {
		errorHandler.Panic(errorHandler.ProductExistInCartError)
	}

	//Calculating the price of the cartDetail according to the price of
	//the product found and the quantity of the request. Creates cartDetail
	detailPrice := product.Price.Mul(decimal.NewFromInt(cartDetail.ProductQuantity))
	cartDetail.SetDetailTotalPrice(detailPrice)
	cartDetail.SetCartId(cart.Id)
	detailId := s.cartRepo.CreateCartDetail(*cartDetail)
	if detailId == "" {
		errorHandler.Panic(errorHandler.DBCreateError)
	}

	//Updating cart price according to request cartDetail.
	cartTotal := cart.TotalCartPrice.Add(cartDetail.DetailTotalPrice)
	cart.SetTotalCartPrice(cartTotal)
	s.updateCart(cart)

	//Updating products unitsOnCart field.
	product.SetProductUnitsOnCart(int(cartDetail.ProductQuantity))
	product.SetUpdatedAt()
	_, err := s.productRepo.UpdateProduct(*product, helper.SetProductUpdateOptions(*product))
	if err != nil {
		errorHandler.Panic(errorHandler.InternalServerError)
	}

	return cartDetail
}

func (s *CartService) UpdateCartDetail(userId string, cartDetail *models.CartDetails) *models.CartDetails {
	//Find user's cart by userId.
	cart := s.findUserCart(userId, false)

	//It checks if there is any matching productId of the cartDetail from the request in the
	//cartDetails in the user's cart. If it doesn't match, it throws an error.
	existDetailCart := findIfProductExistInCart(cartDetail.ProductId, cart.CartDetails)
	if existDetailCart == nil {
		errorHandler.Panic(errorHandler.ProductNotExistInCartError)
	}

	//Find the product that its id matches the cartDetail productId.
	product := s.productRepo.FindProductById(cartDetail.ProductId)
	if product == nil {
		errorHandler.Panic(errorHandler.ProductDeletedError)
	}

	//Checking the request product quantity by product's quantity that found.
	validNum := product.StockNumber - product.UnitsOnCart + int(existDetailCart.ProductQuantity)
	if validNum < int(cartDetail.ProductQuantity) {
		errorHandler.Panic(errorHandler.QuantityNotValidError)
	}
	detailPrice := product.Price.Mul(decimal.NewFromInt(cartDetail.ProductQuantity))
	cartDetail.SetDetailTotalPrice(detailPrice)
	cartDetail.SetUpdatedAt()
	updateOptions := models.CartDetails{
		Base: models.Base{
			UpdatedAt: cartDetail.UpdatedAt,
		},
		ProductQuantity:  cartDetail.ProductQuantity,
		DetailTotalPrice: cartDetail.DetailTotalPrice,
	}
	rawEffected := s.cartRepo.UpdateUserCartDetail(existDetailCart.Id, updateOptions)
	if rawEffected == 0 {
		errorHandler.Panic(errorHandler.DBUpdateError)
	}
	cartTotal := cart.TotalCartPrice.Add(existDetailCart.DetailTotalPrice.Neg()).Add(cartDetail.DetailTotalPrice)
	cart.SetTotalCartPrice(cartTotal)
	s.updateCart(cart)
	return cartDetail
}

func (s *CartService) DeleteCartDetail(userId string, productId string) {
	cart := s.findUserCart(userId, false)
	existDetailCart := findIfProductExistInCart(productId, cart.CartDetails)
	if existDetailCart == nil {
		errorHandler.Panic(errorHandler.ProductNotExistInCartError)
	}
	if res := s.cartRepo.DeleteCartDetails(*existDetailCart); res == false {
		errorHandler.Panic(errorHandler.DBDeleteError)
	}
	cartTotal := cart.TotalCartPrice.Add(existDetailCart.DetailTotalPrice.Neg())
	cart.SetTotalCartPrice(cartTotal)
	s.updateCart(cart)
}

func (s *CartService) findUserCart(id string, isCompleted bool) *models.Cart {
	cart := s.cartRepo.FindUserCart(id, isCompleted)
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

func findIfProductExistInCart(productId string, details []models.CartDetails) *models.CartDetails {
	for _, detail := range details {
		if productId == detail.ProductId {
			return &detail
		}
	}
	return nil
}
