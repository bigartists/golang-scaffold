package controllers

import (
	"com.github.goscaffold/internal/service"
	"com.github.goscaffold/web/result"
	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
	// 校验输入参数是否合法
	params := &struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}{}
	// 校验参数
	result.Result(c.ShouldBindJSON(params)).Unwrap()
	ResultWrapper(c)("get user success", "100001", service.UserServiceGetter.SignIn(params.Username, params.Password).Unwrap())(OK)
}

func SignUp(c *gin.Context) {

}
