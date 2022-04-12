package middleware

import (
	"PicusFinalCase/src/pkg/helper"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

const AuthorizationType = "Bearer"

func AuthMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") != "" {
			fmt.Println(1)
			headerVal := strings.Split(c.GetHeader("Authorization"), " ")
			if headerVal[0] != AuthorizationType {
				fmt.Println(2)
				return
			}
			decodedClaims := helper.VerifyToken(headerVal[1], secretKey)
			if decodedClaims == nil {
				fmt.Println(3)
				return
			}
			c.Set("id", decodedClaims.UserId)
			c.Set("role", decodedClaims.Role)
			return
		} else {
			return
		}
	}
}
