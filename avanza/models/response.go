package models

import "fmt"

type BaseResponse struct {
	StatusCode int               `json:"statusCode"`
	Message    string            `json:"message"`
	Time       string            `json:"time"`
	Errors     []string          `json:"errors"`
	Additional map[string]string `json:"additional"`
}

// ErrorResponse represents an API response with an error status code.
type ErrorResponse struct {
	StatusCode int
	BaseResponse
}

// Error returns the details of an error response.
func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("bad status with code '%d': message '%s': errors: '%s'", e.StatusCode, e.Message, e.Errors)
}
