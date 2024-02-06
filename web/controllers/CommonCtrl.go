package controllers

import (
	"com.github.goscaffold/pkg/globalConstants"
	"fmt"
	"github.com/gin-gonic/gin"
	"sync"
)

type JSONResult struct {
	Message string      `json:"message"`
	Code    int8        `json:"code"`
	Result  interface{} `json:"result"`
	Token   string      `json:"token"`
}

func NewJSONResult(result interface{}) *JSONResult {
	return &JSONResult{
		Message: "",
		Code:    0,
		Result:  result,
		Token:   "",
	}
}

var ResultPool *sync.Pool

func init() {
	ResultPool = &sync.Pool{
		New: func() interface{} {
			return NewJSONResult(nil)
		},
	}
}

type ResultFunc func(result interface{}, message string) func(output Output)
type Output func(c *gin.Context, v interface{})

func ResultWrapper(c *gin.Context) ResultFunc {
	return func(result interface{}, message string) func(output Output) {
		r := ResultPool.Get().(*JSONResult)
		defer ResultPool.Put(r)
		r.Message = message
		//r.Code = code
		token := c.GetString("token")
		r.Token = token
		r.Result = result

		//r.Result = map[string]interface{}{
		//	"data":  result,
		//	"token": token,
		//}

		return func(output Output) {
			output(c, r)
		}
	}
}

func OK(c *gin.Context, v interface{}) {
	// 将v 转成 *JSONResult 类型
	if r, ok := v.(*JSONResult); ok {
		r.Code = globalConstants.HTTPSUCCESS
		r.Message = globalConstants.HTTPMESSAGESUCCESS
	}
	c.JSON(200, v)
}

func Error(c *gin.Context, v interface{}) {
	// 将v 转成 *JSONResult 类型
	if r, ok := v.(*JSONResult); ok {
		r.Code = globalConstants.HTTPFAIL
		if r.Message == "" {
			r.Message = globalConstants.HTTPMESSAGEFAIL
		}
	}
	c.JSON(400, v)
}

func OK2String(c *gin.Context, v interface{}) {
	c.String(200, fmt.Sprintf("%v", v))
}
