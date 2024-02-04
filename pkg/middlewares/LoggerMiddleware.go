package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process request
		c.Next()

		// Stop timer
		end := time.Now()

		// Execution time
		latency := end.Sub(start)

		// Access the status we are sending
		status := c.Writer.Status()

		fmt.Printf("[GIN] %v | %3d | %13v | %15s | %-7s %#v\n%s",
			end.Format("2006/01/02 - 15:04:05"),
			status,
			latency,
			c.ClientIP(),
			c.Request.Method,
			c.Request.URL,
			c.Errors.String(),
		)
	}
}
