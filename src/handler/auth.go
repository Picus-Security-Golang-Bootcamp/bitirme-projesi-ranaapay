package handler

import (
	"PicusFinalCase/src/handler/requestType"
	"PicusFinalCase/src/handler/responseType"
	"PicusFinalCase/src/pkg/errorHandler"
	"PicusFinalCase/src/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(r *gin.RouterGroup, authService *service.AuthService) {
	h := &AuthHandler{service: authService}

	r.POST("/register", h.createUser)
	r.POST("/login", h.loginUser)
}

// createUser
//@Summary       Create User
// @Description  create user in database
// @Tags         authentication
// @Accept       json
// @Produce      json
// @Param 	     requestType.UserRequestType body requestType.UserRequestType true "For create a User"
//@Success       201  {object}  responseType.ResponseType
// @Failure		 400 {object} 	_type.ErrorType
// @Failure		 500 {object} 	_type.ErrorType
// @Router       /authentication/register [post]
//Creates a user in the database by obtaining the necessary
//information from the user and JWT token is returned in response.
func (h *AuthHandler) createUser(c *gin.Context) {

	var userReq requestType.UserRequestType
	if err := c.Bind(&userReq); err != nil {
		log.Error("Bind error : %s", err.Error())
		errorHandler.Panic(errorHandler.BindError)
	}

	//Validates the body of the incoming request
	userReq.ValidateUserRequest()

	//Creates a user type based on the body of the request and sends it to the createUser service func.
	user := userReq.RequestToUserType()
	token := h.service.CreateUser(user)

	//Returns the token from the service.
	c.JSON(http.StatusCreated, responseType.NewResponseType(http.StatusCreated, token))
	return
}

// loginUser
//@Summary       Login User
// @Description  login user for app
// @Tags         authentication
// @Accept       json
// @Produce      json
// @Param 	     LoginRequestType body requestType.LoginType true "For login"
//@Success       200  {object}  responseType.ResponseType
// @Failure		 400 {object} 	_type.ErrorType
// @Failure		 500 {object} 	_type.ErrorType
// @Router       /authentication/login [post]
//Users registered in the database log into the system with firstname and password.
//If both information is correct, JWT token is returned.
func (h *AuthHandler) loginUser(c *gin.Context) {

	var userLogin requestType.LoginType
	if err := c.Bind(&userLogin); err != nil {
		log.Error("Bind error : %s", err.Error())
		errorHandler.Panic(errorHandler.BindError)
	}

	userLogin.ValidateLoginType()

	token := h.service.LoginUser(userLogin.FirstName, userLogin.Password)

	c.JSON(http.StatusOK, responseType.NewResponseType(http.StatusOK, token))
	return
}
