package controllers

import (
	"com.github.goscaffold/internal/service"
	"com.github.goscaffold/web/result"
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func NewUserHandler() *UserController {
	return &UserController{}
}

// GET /users/123

// Build Build方法
func (this *UserController) Build(r *gin.Engine) {
	r.GET("/users", UserList)
	r.GET("/user/:id", UserDetail)
}

func UserList(c *gin.Context) {
	ResultWrapper(c)(service.UserServiceGetter.GetUserList(), "")(OK)
}

func UserDetail(c *gin.Context) {
	id := &struct {
		Id int64 `uri:"id" binding:"required"`
	}{}
	result.Result(c.ShouldBindUri(id)).Unwrap()
	ResultWrapper(c)(service.UserServiceGetter.GetUserDetail(id.Id).Unwrap(), "")(OK)
}

//
//func UserSave(c *gin.Context) {
//	u := UserModel.New()
//	result.Result(c.ShouldBindJSON(u)).Unwrap()
//	ResultWrapper(c)("save user", "10086", "true")(OK)
//}
