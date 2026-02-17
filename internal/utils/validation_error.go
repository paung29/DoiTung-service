package utils

import "github.com/go-playground/validator/v10"

type FieldError struct {
	Field   string `json:"field"`
	Rule    string `json:"rule"`
	Message string `json:"message"`
}

func FormatValidationErrors(err error) []FieldError {
	verrs, ok := err.(validator.ValidationErrors)
	if !ok {
		return []FieldError{{Field: "", Rule: "invalid", Message: "Invalid input"}}
	}

	out := make([]FieldError, 0, len(verrs))
	for _, e := range verrs {
		field := e.Field()
		rule := e.Tag()

		msg := field + " is invalid"
		if rule == "required" {
			msg = field + " is required"
		}

		out = append(out, FieldError{
			Field:   field,
			Rule:    rule,
			Message: msg,
		})
	}
	return out
}