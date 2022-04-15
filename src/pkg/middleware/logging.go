package middleware

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// LoggingMiddleware It logs the information of all incoming requests.
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		log.WithFields(log.Fields{
			"method": c.Request.Method,
			"path":   c.FullPath(),
			"url":    c.Request.RequestURI,
		}).Info("request details")

		c.Next()
	}
}
