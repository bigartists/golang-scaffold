package validators

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"log"
)

var myvalidate *validator.Validate

// 解析错误信息
var validatorError map[string]string

func init() {
	validatorError = make(map[string]string)
}

// 处理错误信息，将官方错误信息处理的更加人性化一点；
// 例如：Required -> 不能为空

func CheckErrors(errors error) {
	if errs, ok := errors.(validator.ValidationErrors); ok {
		for _, err := range errs {
			// Tag() 返回的是错误的类型，例如：Required、Min、Max、Email 等等

			if v, exists := validatorError[err.Tag()]; exists {
				fmt.Println("v = ", v)
				log.Println("err.Tag() = ", err.Tag()) // 这里既然能答应出Tag，就能通过查找记录的map把自定义错误信息返回出去；
				// panic 会被 中间件捕获，最终400输出；
				panic(v)
			}
		}
	}
}

func Build() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		myvalidate = v
	} else {
		log.Fatal("error validator")
	}
	if err := myvalidate.RegisterValidation("usernameValid", UsernameValid); err != nil {
		log.Fatal("usernameValid error")
	}
	if err2 := myvalidate.RegisterValidation("passwordValid", PasswordValid); err2 != nil {
		log.Fatal("passwordValid error")
	}
}
