package common

import (
	"fmt"
	"strings"
)

type APIError struct {
	Message string     `json:"message"`
	Errors  []APIError `json:"errors,omitempty"`
	Field   string     `json:"field,omitempty"`
}

func NewValidationError(fieldErrors ...APIError) APIError {
	return APIError{
		Message: "Some fields weren't valid; see 'errors' for more info.",
		Errors:  fieldErrors,
	}
}

func (e APIError) AddFieldError(fieldName, message string) {
	newError := APIError{
		Field:   fieldName,
		Message: message,
	}
	e.Errors = append(e.Errors, newError)
}

func (e APIError) GetFields() []string {
	fields := []string{}
	for _, fieldError := range e.Errors {
		fields = append(fields, fieldError.Field)
	}
	return fields
}

func (e APIError) GetErrorsByField() map[string]string {
	byField := map[string]string{}
	for _, fieldError := range e.Errors {
		byField[fieldError.Field] = fieldError.Message
	}
	return byField
}

func (e APIError) Error() string {
	if len(e.Errors) > 0 {
		return fmt.Sprintf("%s (fields: %s)", e.Message, strings.Join(e.GetFields(), ", "))
	} else {
		return e.Message
	}
}
