package validator

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	v9 "gopkg.in/go-playground/validator.v9"
	enTrans "gopkg.in/go-playground/validator.v9/translations/en"
)

func NewV9() (Validator, error) {
	langEn := en.New()
	langId := id.New()
	uni := ut.New(langEn, langEn, langId)
	trans, _ := uni.GetTranslator("en")

	validate := v9.New()
	if err := enTrans.RegisterDefaultTranslations(validate, trans); err != nil {
		return nil, err
	}

	return &CustomValidator{
		validator: validate,
	}, nil
}
