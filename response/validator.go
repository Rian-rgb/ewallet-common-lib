package response

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
)

func MapValidationErrors(err error) []ValidationErrorField {
	if err == nil {
		return nil
	}

	var validationErrors validator.ValidationErrors
	ok := errors.As(err, &validationErrors)
	if !ok {
		return nil
	}

	var listErrors []ValidationErrorField

	for _, f := range validationErrors {
		var msg string

		fieldName := f.Field()

		switch f.Tag() {
		case "required":
			msg = fmt.Sprintf("The %s field is required", fieldName)

		case "email":
			msg = fmt.Sprintf("The %s field must be a valid email address", fieldName)

		case "min":
			msg = fmt.Sprintf(
				"The %s field must be at least %s characters long",
				fieldName,
				f.Param(),
			)

		case "max":
			msg = fmt.Sprintf(
				"The %s field must not exceed %s characters",
				fieldName,
				f.Param(),
			)

		case "numeric":
			msg = fmt.Sprintf("The %s field must be numeric", fieldName)

		case "oneof":
			msg = fmt.Sprintf(
				"The %s field must be one of: %s",
				fieldName,
				f.Param(),
			)

		case "gt":
			msg = fmt.Sprintf(
				"The %s field must be greater than %s",
				fieldName,
				f.Param(),
			)

		default:
			msg = fmt.Sprintf(
				"The %s field failed validation for rule '%s'",
				fieldName,
				f.Tag(),
			)
		}

		listErrors = append(listErrors, ValidationErrorField{
			Field:   fieldName,
			Message: msg,
		})
	}

	return listErrors
}
