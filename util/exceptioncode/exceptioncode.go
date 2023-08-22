package exceptioncode

import "errors"

var (
	ErrEmptyResult    = errors.New("empty result")
	ErrInvalidRequest = errors.New("invalid request")
	ErrUnableToLock   = errors.New("unable to lock")

	// spesific postgre error
	ErrForeignKeyViolation = errors.New("foreign key violation")
	ErrUniqueViolation     = errors.New("unique violation")
)

const (
	CodeDataNotFound        = "DATA_NOT_FOUND"
	CodeInvalidRequest      = "INVALID_REQUEST"
	CodeInvalidValidation   = "INVALID_VALIDATION"
	CodeDataLocked          = "DATA_LOCKED"
	CodeBadRequest          = "BAD_REQUEST"
	CodeInternalServerError = "INTERNAL_SERVER_ERROR"
)

type (
	errorType struct {
		ErrorMessage string
	}
	ErrorNotFound            errorType
	ErrorForeignKeyViolation errorType
)
