package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// LoggingMiddleware logs requests based on the configured log level
func LoggingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Starting time request
		startTime := time.Now()

		// Processing request
		ctx.Next()

		// End Time request
		endTime := time.Now()

		// Execution time
		latencyTime := endTime.Sub(startTime)

		// Request details
		reqMethod := ctx.Request.Method
		reqUri := ctx.Request.RequestURI
		statusCode := ctx.Writer.Status()
		clientIP := ctx.ClientIP()

		fields := log.Fields{
			"METHOD":    reqMethod,
			"URI":       reqUri,
			"STATUS":    statusCode,
			"LATENCY":   latencyTime,
			"CLIENT_IP": clientIP,
		}

		// log levels
		switch log.GetLevel() {
		case log.DebugLevel:
			log.WithFields(fields).Debug("HTTP REQUEST - Debug level")
		case log.InfoLevel:
			log.WithFields(fields).Info("HTTP REQUEST - Info level")
		case log.WarnLevel:
			log.WithFields(fields).Warn("HTTP REQUEST - Warn level")
		case log.ErrorLevel:
			log.WithFields(fields).Error("HTTP REQUEST - Error level")
		case log.FatalLevel:
			log.WithFields(fields).Fatal("HTTP REQUEST - Fatal level")
		default:
			log.WithFields(fields).Info("HTTP REQUEST - Default level")
		}
	}
}
