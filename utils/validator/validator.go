package validator

import "gopkg.in/go-playground/validator.v9"

type Validator interface {
	Validate(i interface{}) error
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
