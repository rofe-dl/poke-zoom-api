package utils

import "github.com/danielgtaylor/huma/v2"

type InternalServerError struct{}

func (e InternalServerError) Error() string {
	return "Internal server error"
}

type NotFoundError struct{}

func (e NotFoundError) Error() string {
	return "Entity not found"
}

func HandleHttpError(err error) huma.StatusError {
	switch err.(type) {
	case NotFoundError:
		return huma.Error404NotFound(err.Error())
	default:
		return huma.Error500InternalServerError(err.Error())
	}
}
