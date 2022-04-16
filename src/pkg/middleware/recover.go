package middleware

import (
	_type "PicusFinalCase/src/pkg/errorHandler/type"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io"
)

func Recovery() gin.HandlerFunc {
	return RecoveryWithWriter(gin.DefaultErrorWriter)
}

// RecoveryWithWriter Recovers panics thrown in the application. The incoming error is returned to the user as a response.
func RecoveryWithWriter(out io.Writer) gin.HandlerFunc {
	return func(c *gin.Context) {

		defer func() {

			if err := recover(); err != nil {

				//It is converted to custom error type.
				resErr, cErr := err.(*_type.ErrorType)
				if cErr != true {
					log.Fatal(resErr.Message)
				}

				c.JSON(resErr.Code, resErr)
			}
		}()

		c.Next()
	}
}
