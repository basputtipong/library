package liberror

import "fmt"

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
