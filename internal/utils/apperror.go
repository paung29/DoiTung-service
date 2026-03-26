package utils

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)


type RestError struct {
	Status int	`json:"-"`
	Message string `json:"message"`
	Errors  []FieldError `json:"errors,omitempty"`
}

func (e RestError) Error() string {
	return e.Message
}

func BadRequestError(message string) RestError {
	return RestError{Status: http.StatusBadRequest, Message: message}
}

func UnauthorizedError(message string) RestError {
	return RestError{Status: http.StatusUnauthorized, Message: message}
}

func ForbiddenError(message string) RestError {
	return RestError{Status: http.StatusForbidden, Message: message}
}

func NotFoundError(message string) RestError {
	return RestError{Status: http.StatusNotFound, Message: message}
}

func SystemError(message string) RestError {
	return RestError{Status: http.StatusInternalServerError, Message: message}
}

func ValidationError(message string, errs []FieldError) RestError {
	return RestError{
		Status:  http.StatusBadRequest,
		Message: message,
		Errors:  errs,
	}
}

func HandleError(context *fiber.Ctx, err error) error {
	if restErr, ok := err.(RestError); ok {
		return context.Status(restErr.Status).JSON(fiber.Map{
			"success": false,
			"message": restErr.Message,
			"errors":  restErr.Errors,
		})
	}

	return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"success": false,
		"message": "internal server error",
	})
}

func ParseAndValidate[T any](context *fiber.Ctx, form *T) error {
	if err := context.BodyParser(form); err != nil {
		return BadRequestError("invalid json format")
	}

	if err := Validate.Struct(form); err != nil {
		return ValidationError("validation error", FormatValidationErrors(err))
	}

	return nil
}