package handler

import (
	"PicusFinalCase/src/handler/requestType"
	"PicusFinalCase/src/handler/responseType"
	"PicusFinalCase/src/pkg/errorHandler"
	"PicusFinalCase/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(r *gin.RouterGroup, authService *service.AuthService) {
	h := &AuthHandler{service: authService}

	r.POST("/create", h.createUser)
	r.POST("/login", h.loginUser)
}

func (h *AuthHandler) createUser(c *gin.Context) {
	var userReq requestType.UserRequestType
	if err := c.Bind(&userReq); err != nil {
		errorHandler.Panic(errorHandler.BindError)
	}
	userReq.ValidateUserRequest()
	user := userReq.RequestToUserType()
	token := h.service.CreateUser(user)
	c.JSON(http.StatusCreated, responseType.NewResponseType(http.StatusCreated, token))
	return
}

func (h *AuthHandler) loginUser(c *gin.Context) {
	var userLogin requestType.LoginType
	if err := c.Bind(&userLogin); err != nil {
		errorHandler.Panic(errorHandler.BindError)
	}
	userLogin.ValidateLoginType()
	token := h.service.LoginUser(userLogin.FirstName, userLogin.Password)
	c.JSON(http.StatusOK, responseType.NewResponseType(http.StatusOK, token))
	return
}
