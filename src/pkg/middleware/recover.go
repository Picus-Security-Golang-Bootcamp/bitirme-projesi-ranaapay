package middleware

import (
	_type "PicusFinalCase/src/pkg/errorHandler/type"
	"github.com/gin-gonic/gin"
	"io"
	"log"
)

func Recovery() gin.HandlerFunc {
	return RecoveryWithWriter(gin.DefaultErrorWriter)
}

func RecoveryWithWriter(out io.Writer) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				resErr, cErr := err.(*_type.ErrorType)
				if cErr != true {
					log.Fatalf(resErr.Message)
				}
				c.JSON(resErr.Code, resErr)
			}
		}()
		c.Next() // execute all the handlers
	}
}
