package utils

import (
	"errors"
	"log"

	"github.com/danielgtaylor/huma/v2"
)

type HTTPError struct {
	Message string
	Status  int
	Cause   error
}

func (e HTTPError) Error() string {
	if e.Cause != nil {
		return e.Message + ": " + e.Cause.Error()
	}

	return e.Message
}

func (e HTTPError) Unwrap() error {
	return e.Cause
}

func HandleHttpError(err error) huma.StatusError {
	// unwrap error to see if it is of our custom type
	if httpError, ok := errors.AsType[HTTPError](err); ok {
		// only log the full details of the error, but response will contain the friendly message
		if httpError.Status/100 == 5 {
			log.Printf("Request error: %s", httpError)
		}

		return huma.NewError(httpError.Status, httpError.Message)
	}

	log.Printf("Unhandled error: %v", err)
	return huma.Error500InternalServerError("Internal Server Error")
}

func NewHTTPError(status int, message string, causes ...error) HTTPError {
	if len(message) == 0 {
		message = "Internal Server Error"
	}

	err := HTTPError{Status: status, Message: message}

	if len(causes) != 0 {
		err.Cause = causes[0]
	}

	return err
}
