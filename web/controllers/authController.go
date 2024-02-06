package controllers

import (
	"com.github.goscaffold/config"
	"com.github.goscaffold/internal/service"
	"com.github.goscaffold/pkg/utils"
	"com.github.goscaffold/web/dto"
	"com.github.goscaffold/web/result"
	"github.com/gin-gonic/gin"
	"time"
)

type AuthController struct{}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (a *AuthController) Login(c *gin.Context) {
	// 校验输入参数是否合法
	params := &struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}{}
	// 校验参数
	result.Result(c.ShouldBindJSON(params)).Unwrap()

	user, err := service.UserServiceGetter.SignIn(params.Username, params.Password)
	if err != nil {
		ResultWrapper(c)(nil, err.Error())(Error)
		return
	}

	//// 生成 token
	prikey := []byte(config.SysYamlconfig.Jwt.PrivateKey)
	curTime := time.Now().Add(time.Second * 60 * 60 * 24)
	token, _ := utils.GenerateToken(user.Id, prikey, curTime)

	c.Set("token", token)
	ResultWrapper(c)(user, "")(OK)
}

func (a *AuthController) SignUp(c *gin.Context) {
	// 校验输入参数是否合法
	params := &dto.SignupRequest{}
	// 校验参数
	result.Result(c.ShouldBindJSON(params)).Unwrap()

	err := service.UserServiceGetter.SignUp(params.Email, params.Username, params.Password)
	if err != nil {
		ResultWrapper(c)(nil, err.Error())(Error)
		return
	}
	ResultWrapper(c)(true, "")(Created)
}

func (a *AuthController) Build(r *gin.Engine) {
	r.POST("/login", a.Login)
	r.POST("/register", a.SignUp)
}

//func SetUpAuthController(r *gin.Engine) {
//	authController := NewAuthController()
//	r.POST("/login", authController.Login)
//	r.POST("/register", authController.SignUp)
//}
