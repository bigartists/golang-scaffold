package result

import (
	"com.github.goscaffold/pkg/validators"
	"fmt"
)

type ErrorResult struct {
	data interface{}
	err  error
}

func (this *ErrorResult) Unwrap() interface{} {

	if this.err != nil {
		fmt.Printf("%T\n", this.err) // validator.ValidationErrors T 可以输出报错类型
		// 通过类型判断，捕获 validation错误
		validators.CheckErrors(this.err) // validator.ValidationErrors 类型的数据交给 CheckErrors 处理
		panic(this.err.Error())
	}
	return this.data
}

// 支持多返回值
func Result(vs ...interface{}) *ErrorResult {
	if len(vs) == 1 {
		if vs[0] == nil {
			return &ErrorResult{nil, nil}
		}
		if e, ok := vs[0].(error); ok {
			return &ErrorResult{nil, e}
		}
	}
	if len(vs) == 2 {

		if vs[1] == nil {

			return &ErrorResult{vs[0], nil}
		}
		if e, ok := vs[1].(error); ok {
			return &ErrorResult{vs[0], e}
		}
	}
	return &ErrorResult{nil, fmt.Errorf("error result format")}
}
