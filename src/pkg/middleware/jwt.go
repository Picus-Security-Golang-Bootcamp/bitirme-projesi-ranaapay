package middleware

import (
	"PicusFinalCase/src/models"
	"PicusFinalCase/src/pkg/errorHandler"
	"PicusFinalCase/src/pkg/helper"
	"github.com/gin-gonic/gin"
	"strings"
)

const AuthorizationType = "Bearer"

func AuthMiddleware(secretKey string, expectedRole models.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") != "" {
			headerVal := strings.Split(c.GetHeader("Authorization"), " ")
			if headerVal[0] != AuthorizationType {
				c.JSON(errorHandler.NotAuthorizedError.Code, errorHandler.NotAuthorizedError)
				c.Abort()
				return
			}
			decodedClaims := helper.VerifyToken(headerVal[1], secretKey)
			if decodedClaims == nil {
				c.JSON(errorHandler.NotAuthorizedError.Code, errorHandler.NotAuthorizedError)
				c.Abort()
				return
			}
			if decodedClaims.Role != expectedRole {
				c.JSON(errorHandler.ForbiddenError.Code, errorHandler.ForbiddenError)
				c.Abort()
				return
			}
			c.Set("id", decodedClaims.UserId)
			return
		} else {
			c.JSON(errorHandler.NotAuthorizedError.Code, errorHandler.NotAuthorizedError)
			c.Abort()
			return
		}
	}
}
