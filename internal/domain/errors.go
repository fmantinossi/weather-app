package domain

import "errors"

var (
	ErrNotFound            = errors.New("can not find zipcode")
	ErrUnprocessableEntity = errors.New("invalid zipcode")
)
