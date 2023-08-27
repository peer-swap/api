package reusable

import "github.com/go-playground/validator"

type Validator struct {
	Input interface{}
}

func NewValidator(input interface{}) *Validator {
	return &Validator{Input: input}
}

func (s Validator) Validate() (bool, *ValidationErrorsMapper) {
	validate := validator.New()
	err := validate.Struct(s.Input)

	if err != nil {
		return true, NewValidationErrorsMapper(err.(validator.ValidationErrors))
	}
	return false, nil
}
