package param_verify

import (
	"gopkg.in/go-playground/validator.v9"
)

func NameValid(fl validator.FieldLevel) bool {
	val := fl.Field().String()
	if val == "admin" {
		return false
	}
	return true
}

func PwdValid(fl validator.FieldLevel) bool {
	val := fl.Field().String()
	if val == "admin" {
		return false
	}
	return true
}

func ValidHandle(bindRes interface{}) error {

	// 参数验证
	validate := validator.New()

	// 注册自定义验证
	validate.RegisterValidation("NameValid", NameValid)
	validate.RegisterValidation("PwdValid", NameValid)
	return validate.Struct(bindRes)
}
