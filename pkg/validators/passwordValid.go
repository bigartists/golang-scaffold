package validators

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

var (
	PasswordPattern = regexp.MustCompile(`^[a-zA-Z0-9_]{6,20}$`)
)

func PasswordValid(fl validator.FieldLevel) bool {
	validatorError["passwordValid"] = "[密码]:长度在 6-20 之间，只能包含字母，数字，下划线；"
	return PasswordPattern.MatchString(fl.Field().String())
}
