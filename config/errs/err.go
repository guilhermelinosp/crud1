package errs

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Error struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes,omitempty"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func NewError(message, err string, code int, causes []Causes) *Error {
	return &Error{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewValidationError(message string, causes []Causes) *Error {
	return NewError(message, "validation_error", http.StatusBadRequest, causes)
}

func NewBadRequest(message string) *Error {
	return NewError(message, "bad_request", http.StatusBadRequest, nil)
}

func NewNotFound(message string) *Error {
	return NewError(message, "not_found", http.StatusNotFound, nil)
}

func NewInternalServerError(message string) *Error {
	return NewError(message, "internal_server_error", http.StatusInternalServerError, nil)
}

func NewUnauthorized(message string) *Error {
	return NewError(message, "unauthorized", http.StatusUnauthorized, nil)
}

func NewConflict(message string) *Error {
	return NewError(message, "conflict", http.StatusConflict, nil)
}

func (r *Error) Error() string {
	return fmt.Sprintf("Error [%d - %s]: %s", r.Code, r.Err, r.Message)
}

// ToJSON converts the error to a JSON string
func (r *Error) ToJSON() string {
	jsonData, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(jsonData)
}

// AddCause adds a cause to the error
func (r *Error) AddCause(field, message string) {
	r.Causes = append(r.Causes, Causes{Field: field, Message: message})
}

// IsType checks if the error is of a specific type
func (r *Error) IsType(errType string) bool {
	return r.Err == errType
}
