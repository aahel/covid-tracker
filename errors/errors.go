package errors

import (
	"errors"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

// AppError struct holds the value of HTTP status code and custom error message.
type AppError struct {
	Status  int    `json:"status"`
	Message string `json:"error_message,omitempty"`
	Debug   error  `json:"-"`
}

var ner = errors.New

// IsRequiredErr returns new error with custom error message
func IsRequiredErr(key string) error {
	return ner(key + " is required")
}

// IsInvalidErr returns new error with custom error message
func IsInvalidErr(key string) error {
	return ner(key + " is invalid")
}

func (err *AppError) Error() string {
	return err.Message
}

// AddDebug method is used to add a debug error which will be printed
// during the error execution if it is not nil. This is purely for developers'
// debugging purposes
func (err *AppError) AddDebug(erx error) *AppError {
	if err != nil {
		err.Debug = erx
	}

	return err
}

// NewAppError returns the new apperror object
func NewAppError(status int, message string) *AppError {
	return &AppError{
		Status:  status,
		Message: message,
	}
}

// BadRequest will return `http.StatusBadRequest` with custom message.
func BadRequest(message string) *AppError { // 400
	return NewAppError(http.StatusBadRequest, message)
}

// NotFound will return `http.StatusNotFound` with custom message.
func NotFound(message string) *AppError { // 404
	return NewAppError(http.StatusNotFound, message)
}

// InternalServerStd will return `http.StatusInternalServerError` with static message.
func InternalServerStd() *AppError { // 500
	return NewAppError(http.StatusInternalServerError, "Something went wrong")
}

// IsMongoNoDocErr should return true if the err is redis: nil
func IsMongoNoDocErr(err error) bool {
	return err == mongo.ErrNoDocuments
}
