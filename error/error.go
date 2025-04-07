package liberror

import (
	"fmt"
	"net/http"
)

type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

func New(code int, message string, detail string) *HTTPError {
	return &HTTPError{Code: code, Message: message, Detail: detail}
}

func ErrorBadRequest(message string, detail string) *HTTPError {
	return &HTTPError{Code: http.StatusBadRequest, Message: message, Detail: detail}
}

func ErrorUnauthorized(message string, detail string) *HTTPError {
	return &HTTPError{Code: http.StatusUnauthorized, Message: message, Detail: detail}
}

func ErrorForbidden(message string, detail string) *HTTPError {
	return &HTTPError{Code: http.StatusForbidden, Message: message, Detail: detail}
}

func ErrorNotFound(message string, detail string) *HTTPError {
	return &HTTPError{Code: http.StatusNotFound, Message: message, Detail: detail}
}

func ErrorInternalServerError(message string, detail string) *HTTPError {
	return &HTTPError{Code: http.StatusInternalServerError, Message: message, Detail: detail}
}

func ErrorConflict(message string, detail string) *HTTPError {
	return &HTTPError{Code: http.StatusConflict, Message: message, Detail: detail}
}
