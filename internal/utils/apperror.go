package utils

import "net/http"


type RestError struct {
	Status int
	Message string
}

func (e RestError) Error() string {
	return e.Message
}

func BadRequestError(message string) RestError {
	return RestError{Status: http.StatusBadRequest, Message: message}
}

func NotFoundError(message string) RestError {
	return RestError{Status: http.StatusNotFound, Message: message}
}

func SystemError(message string) RestError {
	return RestError{Status: http.StatusInternalServerError, Message: message}
}

func Validation(message string) RestError {
	return RestError{Status: http.StatusBadRequest, Message: message}
}