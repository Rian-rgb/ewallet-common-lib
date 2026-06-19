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
			msg = fmt.Sprintf("Field %s wajib diisi", fieldName)
		case "email":
			msg = fmt.Sprintf("Format %s tidak valid", fieldName)
		case "min":
			msg = fmt.Sprintf("Field %s minimal berisi %s karakter", fieldName, f.Param())
		case "max":
			msg = fmt.Sprintf("Field %s maksimal berisi %s karakter", fieldName, f.Param())
		case "numeric":
			msg = fmt.Sprintf("Field %s harus berupa angka", fieldName)
		case "oneof":
			msg = fmt.Sprintf(
				"Field %s harus bernilai salah satu dari: %s",
				fieldName,
				f.Param(),
			)
		case "gt":
			msg = fmt.Sprintf(
				"Field %s harus lebih besar dari %s",
				fieldName,
				f.Param(),
			)
		default:
			msg = fmt.Sprintf("Field %s tidak memenuhi aturan %s", fieldName, f.Tag())
		}

		listErrors = append(listErrors, ValidationErrorField{
			Field:   fieldName,
			Message: msg,
		})
	}

	return listErrors
}
