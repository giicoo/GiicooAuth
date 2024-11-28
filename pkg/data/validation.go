package data

import (
	"github.com/giicoo/GiicooAuth/pkg/errTools"
	"github.com/go-playground/validator"
)

func ValidateStructure(s interface{}) error {
	validate := validator.New()
	err := validate.Struct(s)
	if err != nil {
		return errTools.WrapError(err, errTools.ErrInvalidJSON)
	}
	return nil
}
