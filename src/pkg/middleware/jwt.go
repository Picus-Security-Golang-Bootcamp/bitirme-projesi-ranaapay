package middleware

import (
	"PicusFinalCase/src/models"
	"PicusFinalCase/src/pkg/errorHandler"
	"PicusFinalCase/src/pkg/helper"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"strings"
)

const AuthorizationType = "Bearer"

// AuthMiddleware It checks the token in the header according to the secret key it receives.
//If the token is correct, it performs role control according to the role it takes.
func AuthMiddleware(secretKey string, expectedRole models.Role) gin.HandlerFunc {
	return func(c *gin.Context) {

		if c.GetHeader("Authorization") != "" {

			headerVal := strings.Split(c.GetHeader("Authorization"), " ")

			//Check authorization type is equal bearer.
			if headerVal[0] != AuthorizationType {
				log.Error("Authorization type is not Bearer.")
				c.JSON(errorHandler.NotAuthorizedError.Code, errorHandler.NotAuthorizedError)
				c.Abort()
				return
			}

			decodedClaims := helper.VerifyToken(headerVal[1], secretKey)
			if decodedClaims == nil {
				log.Error("The token is not correct.")
				c.JSON(errorHandler.NotAuthorizedError.Code, errorHandler.NotAuthorizedError)
				c.Abort()
				return
			}

			//Check expectedRole is equal to the claim role that is inside the token.
			if decodedClaims.Role != expectedRole {
				log.Error("The user's role is not equal to the expected role.")
				c.JSON(errorHandler.ForbiddenError.Code, errorHandler.ForbiddenError)
				c.Abort()
				return
			}

			//Set the id value in context to claim userId.
			c.Set("id", decodedClaims.UserId)
			log.Info("id field in context is set to : %s", decodedClaims.UserId)
			return

		} else {
			log.Error("Authorization header is empty.")
			c.JSON(errorHandler.NotAuthorizedError.Code, errorHandler.NotAuthorizedError)
			c.Abort()
			return

		}
	}
}
