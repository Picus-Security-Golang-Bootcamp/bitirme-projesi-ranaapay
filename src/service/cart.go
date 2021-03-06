package service

import (
	"PicusFinalCase/src/client"
	"PicusFinalCase/src/models"
	"PicusFinalCase/src/pkg/errorHandler"
	"PicusFinalCase/src/repository"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
)

type CartService struct {
	cartRepo      *repository.CartRepository
	productClient *client.ProductClient
}

func NewCartService(cartRepo *repository.CartRepository, productClient *client.ProductClient) *CartService {
	return &CartService{
		cartRepo:      cartRepo,
		productClient: productClient,
	}
}

// ListCartItems Find cart with cart details according to userId
func (s *CartService) ListCartItems(userId string) *models.Cart {
	return s.findUserCart(userId)
}

// AddToCart Adds cart detail to cart according to incoming userId, and cartDetail. Updates cart total
//price and product units on cart fields. Returns created cart detail.
func (s *CartService) AddToCart(userId string, cartDetail *models.CartDetails, header string) *models.CartDetails {

	s.productClient.Token = header

	product := s.productClient.FindProductById(cartDetail.ProductId)
	/*
		//Find the product that its id matches the cartDetail productId.
		product := s.productRepo.FindProductById(cartDetail.ProductId)
		if product == nil {
			log.Error("Product that its id matches to incoming cartDetail.ProductId does not contain in database.")
			errorHandler.Panic(errorHandler.ProductIdNotValidError)
		}
	*/
	//Checking the request product quantity by product's quantity that found.
	if (product.StockNumber - product.UnitsOnCart) < int(cartDetail.ProductQuantity) {
		log.Error("Founded product stock numbers arent available for cartDetail product quantity.")
		errorHandler.Panic(errorHandler.QuantityNotValidError)
	}

	//Find user's cart by userId.
	cart := s.findUserCart(userId)

	//It checks if there is any matching productId of the cartDetail from the request in the
	//cartDetails in the user's cart. If it does match, it throws an error.
	isExistProduct := findIfProductExistInCart(cartDetail.ProductId, cart.CartDetails)
	if isExistProduct != nil {
		log.Error("Given cartDetail product already in cart.")
		errorHandler.Panic(errorHandler.ProductExistInCartError)
	}

	//Calculating the price of the cartDetail according to the price of
	//the product found and the quantity of the request. Creates cartDetail
	detailPrice := product.Price.Mul(decimal.NewFromInt(cartDetail.ProductQuantity))
	cartDetail.SetDetailTotalPrice(detailPrice)
	cartDetail.SetCartId(cart.Id)

	detailId := s.cartRepo.CreateCartDetail(*cartDetail)
	if detailId == "" {
		log.Error("Something happened when creating cart detail.")
		errorHandler.Panic(errorHandler.DBCreateError)
	}

	//Updating products unitsOnCart field.
	product.SetProductUnitsOnCart(int(cartDetail.ProductQuantity))

	s.productClient.UpdateProduct(*product)

	/*	_, err := s.productRepo.UpdateProduct(*product, helper.SetProductUpdateOptions(*product))
		if err != nil {
			log.Error("Something happened when updating product.")
			errorHandler.Panic(errorHandler.InternalServerError)
		}
	*/
	return cartDetail
}

// UpdateCartDetail Update cart detail to cart according to incoming userId, and cartDetail. Updates cart total
//price and product units on cart fields. Returns updated cart detail.
func (s *CartService) UpdateCartDetail(userId string, cartDetail *models.CartDetails, header string) *models.CartDetails {

	//Find user's cart by userId.
	cart := s.findUserCart(userId)

	//It checks if there is any matching productId of the cartDetail from the request in the
	//cartDetails in the user's cart. If it doesn't match, it throws an error.
	existDetailCart := findIfProductExistInCart(cartDetail.ProductId, cart.CartDetails)
	if existDetailCart == nil {
		log.Error("Given cartDetail product does not contain in cart.")
		errorHandler.Panic(errorHandler.ProductNotExistInCartError)
	}

	s.productClient.Token = header

	product := s.productClient.FindProductById(cartDetail.ProductId)

	/*
		//Find the product that its id matches the cartDetail productId.
		product := s.productRepo.FindProductById(cartDetail.ProductId)
		if product == nil {
			log.Error("Product that its id matches to incoming cartDetail.ProductId does not contain in database.")
			errorHandler.Panic(errorHandler.ProductDeletedError)
		}
	*/
	//Checking the request product quantity by product's quantity that found.
	validNum := product.StockNumber - product.UnitsOnCart + int(existDetailCart.ProductQuantity)
	if validNum < int(cartDetail.ProductQuantity) {
		log.Error("Founded product stock numbers arent available for cartDetail product quantity.")
		errorHandler.Panic(errorHandler.QuantityNotValidError)
	}

	//Updating the price of the cartDetail according to the price of
	//the product found and the quantity of the request.
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
		log.Error("No raws updated.")
		errorHandler.Panic(errorHandler.DBUpdateError)
	}

	//Calculating and Updating products unitsOnCart field.
	productUnitsOnCart := product.UnitsOnCart - int(existDetailCart.ProductQuantity) + int(cartDetail.ProductQuantity)
	product.SetProductUnitsOnCart(productUnitsOnCart)

	s.productClient.UpdateProduct(*product)

	/*
		_, err := s.productRepo.UpdateProduct(*product, helper.SetProductUpdateOptions(*product))
		if err != nil {
			log.Error("Something happened when updating product.")
			errorHandler.Panic(errorHandler.InternalServerError)
		}
	*/
	return cartDetail
}

// DeleteCartDetail Deletes cart detail according to given userId and productId. Updates cart total price.
func (s *CartService) DeleteCartDetail(userId string, productId string) {

	cart := s.findUserCart(userId)

	existDetailCart := findIfProductExistInCart(productId, cart.CartDetails)
	if existDetailCart == nil {
		log.Error("Given productId does not contain in carts product.")
		errorHandler.Panic(errorHandler.ProductNotExistInCartError)
	}

	if res := s.cartRepo.DeleteCartDetails(*existDetailCart); res == false {
		log.Error("Something happened when deleting cartDetail.")
		errorHandler.Panic(errorHandler.DBDeleteError)
	}
}

//findUserCart Finds user cart according to given id. If cart can not be found creates
//new cart according to given id. Returns *models.Cart
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
		log.Error("Something happened when creating cart.")
		errorHandler.Panic(errorHandler.DBCreateError)
	}

	return cart
}

func findIfProductExistInCart(productId string, details []models.CartDetails) *models.CartDetails {
	for _, detail := range details {
		if productId == detail.ProductId {
			return &detail
		}
	}
	return nil
}
