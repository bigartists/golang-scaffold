package routes

import (
	"com.github.goscaffold/web/controllers"
	"github.com/gin-gonic/gin"
)

func Build(r *gin.Engine) {
	controllers.NewUserHandler().Build(r)
}
