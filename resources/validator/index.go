package validator

import "github.com/naufalfmm/aquafarm-management-service/utils/validator"

func NewValidator() (validator.Validator, error) {
	return validator.NewV9()
}
