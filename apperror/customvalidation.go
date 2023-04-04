package apperror

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

var (
	customErrors = map[string]error{
		"name.required":   errors.New("is required"),
		"name.min":        errors.New("has to be at-least 5 characters"),
		"name.max":        errors.New("can be at-most 100 characters"),
		"type.required":   errors.New("is required"),
		"type.oneof":      errors.New("has to be one of TypeA, TypeB, TypeC, TypeD"),
		"quantity.min":    errors.New("has to be a positive value"),
		"price.min":       errors.New("has to be a positive value"),
		"description.max": errors.New("can be at-most 500 character"),
		"id.required":     errors.New("is required"),
		"id.uuid":         errors.New("has to be a uuid"),
	}
)

func CustomValidationError(sourceStruct interface{}, err error) []map[string]string {
	errs := make([]map[string]string, 0)
	switch errTypes := err.(type) {
	case validator.ValidationErrors:
		for _, e := range errTypes {
			errorMap := make(map[string]string)

			key := e.Field() + "." + e.Tag()

			if v, ok := customErrors[key]; ok {
				errorMap[e.Field()] = v.Error()
			} else {
				errorMap[e.Field()] = fmt.Sprintf("custom message is not available: %v", err)
			}
			errs = append(errs, errorMap)
		}
		return errs
	case *json.UnmarshalTypeError:
		errs = append(errs, map[string]string{errTypes.Field: fmt.Sprintf("%v cannot be a %v", errTypes.Field, errTypes.Value)})
		return errs
	}
	errs = append(errs, map[string]string{"unknown": fmt.Sprintf("unsupported custom error for: %v", err)})
	return errs
}
