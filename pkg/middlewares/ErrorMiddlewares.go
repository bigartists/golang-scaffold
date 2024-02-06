package middlewares

import (
	"com.github.goscaffold/web/controllers"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				//context.JSON(500, gin.H{"error": err})
				controllers.ResultWrapper(context)(nil, err.(string))(controllers.Error)
			}
		}()
		context.Next()
	}
}
