package server

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-nacelle/nacelle"
)

func GinLogger(nl nacelle.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Time Before Request
		t := time.Now()

		c.Next()
		// After Request

		writerInfo := map[string]interface{}{
			"status":     c.Writer.Status(),
			"method":     c.Request.Method,
			"path":       c.Request.URL.Path,
			"IP":         c.ClientIP(),
			"latency":    time.Since(t),
			"user-agent": c.Request.UserAgent(),
		}

		reqStatus := c.Writer.Status()

		switch {
		case reqStatus >= 400 && reqStatus < 500:
			nl.WarningWithFields(writerInfo, "API-WARN")
		case reqStatus >= 500:
			nl.ErrorWithFields(writerInfo, "API-ERROR")
		default:
			nl.DebugWithFields(writerInfo, "API-DEBUG")
		}
	}
}
