package config

import (
	"github.com/go-playground/validator/v10"
)

func ValidateConfig(s interface{}) (err error) {
	validate := validator.New()
	return validate.Struct(s)

}
