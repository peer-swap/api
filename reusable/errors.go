package reusable

import (
	"fmt"
	"github.com/go-playground/validator"
)

type ValidationErrorsMapper struct {
	validator.ValidationErrors
}

func NewValidationErrorsMapper(validationErrors validator.ValidationErrors) *ValidationErrorsMapper {
	return &ValidationErrorsMapper{ValidationErrors: validationErrors}
}

func (m ValidationErrorsMapper) ErrorMap() map[string][]map[string]string {
	errorMap := map[string][]map[string]string{}

	for _, err := range m.ValidationErrors {
		field := err.Field()
		errData := map[string]string{
			"message": fmt.Sprintf("Validation failed on '%s' with tag '%s'", field, err.Tag()),
			"tag":     err.Tag(),
		}

		if _, ok := errorMap[field]; ok {
			errorMap[field] = append(errorMap[field], errData)
		} else {
			errorMap[field] = []map[string]string{errData}
		}
	}

	return errorMap
}

func (m ValidationErrorsMapper) ErrorMessageMap() map[string]interface{} {
	return map[string]interface{}{
		"message": "The service has encountered an error.",
		"data":    m.ErrorMap(),
	}
}
