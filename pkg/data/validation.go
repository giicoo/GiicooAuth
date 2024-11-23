package data

import (
	errTools "github.com/giicoo/GiicooAuth/pkg/err_tools"
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
