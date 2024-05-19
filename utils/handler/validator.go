package handler

import "github.com/go-playground/validator/v10"

func Validate[T any](s T) error {
	validate := validator.New()
	return validate.Struct(s)
}
