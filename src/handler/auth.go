package handler

import (
	"PicusFinalCase/src/handler/requestType"
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
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := userReq.ValidateUserRequest(); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	user := userReq.RequestToUserType()
	token, err := h.service.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusCreated, token)
	return
}

func (h *AuthHandler) loginUser(c *gin.Context) {
	var userLogin requestType.LoginType
	if err := c.Bind(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := userLogin.ValidateLoginType(); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	token, err := h.service.LoginUser(userLogin.FirstName, userLogin.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}
	c.JSON(http.StatusOK, token)
	return
}
