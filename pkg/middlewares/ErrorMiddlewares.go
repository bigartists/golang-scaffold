package middlewares

import "github.com/gin-gonic/gin"

func ErrorHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				context.JSON(500, gin.H{"error": err})
			}
		}()

		context.Next()
	}
}
