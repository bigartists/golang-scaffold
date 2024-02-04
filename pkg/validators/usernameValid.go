package validators

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

var (
	UsernamePattern = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]{5,19}$`)
)

func UsernameValid(fl validator.FieldLevel) bool {
	//return UsernamePattern.MatchString(fl.Field().String())
	validatorError["usernameValid"] = "[用户名]:长度在 6-20 之间，且不能重复，只能包含大小写字母，数字，下划线；第一个字符必须是字母；"
	return UsernamePattern.MatchString(fl.Field().Interface().(string))
}
