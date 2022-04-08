package handler

import (
	"PicusFinalCase/src/handler/requestType"
	"PicusFinalCase/src/models"
	"PicusFinalCase/src/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(r *gin.RouterGroup, authService *service.AuthService) {
	h := &AuthHandler{service: authService}

	r.POST("", h.createUser)
	r.POST("/login", h.loginUser)
}

func (h *AuthHandler) createUser(c *gin.Context) {
	var user models.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
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
	fmt.Println(userLogin.FirstName, userLogin.Password)
	token, err := h.service.LoginUser(userLogin.FirstName, userLogin.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}
	c.JSON(http.StatusOK, token)
	return
}
