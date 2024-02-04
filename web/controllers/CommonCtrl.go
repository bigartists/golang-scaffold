package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sync"
)

type JSONResult struct {
	Message string      `json:"message"`
	Code    string      `json:"code"`
	Result  interface{} `json:"result"`
}

func NewJSONResult(message string, code string, result interface{}) *JSONResult {
	return &JSONResult{
		Message: message,
		Code:    code,
		Result:  result,
	}
}

var ResultPool *sync.Pool

func init() {
	ResultPool = &sync.Pool{
		New: func() interface{} {
			return NewJSONResult("", "", nil)
		},
	}
}

type ResultFunc func(message string, code string, result interface{}) func(output Output)
type Output func(c *gin.Context, v interface{})

func ResultWrapper(c *gin.Context) ResultFunc {
	return func(message string, code string, result interface{}) func(output Output) {
		r := ResultPool.Get().(*JSONResult)
		defer ResultPool.Put(r)
		r.Message = message
		r.Code = code
		r.Result = result

		return func(output Output) {
			output(c, r)
		}
	}
}

func OK(c *gin.Context, v interface{}) {

	c.JSON(200, v)
}

func Error(c *gin.Context, v interface{}) {
	c.JSON(400, v)
}

func OK2String(c *gin.Context, v interface{}) {
	c.String(200, fmt.Sprintf("%v", v))
}
