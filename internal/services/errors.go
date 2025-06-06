package services

import "errors"

var (
	InternalServerError error = errors.New("internal server error")
	AlreadyExists       error = errors.New("already exists")
	NotFound            error = errors.New("not found")
)
