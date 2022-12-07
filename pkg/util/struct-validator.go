package util

import (
	"strings"

	"fiber/app/model"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct(obj interface{}) model.ErrorValidationResponse {
	var errors model.ErrorValidationResponse = model.ErrorValidationResponse{}
	err := validate.Struct(obj)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element model.ErrorValidatorMeta
			element.Tag = err.Tag()
			element.Value = err.Param()

			//errors = append(errors, &element)
			errors[strings.ToLower(err.StructField())] = element
		}
	}
	return errors
}
