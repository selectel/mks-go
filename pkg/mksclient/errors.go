package mksclient

import (
	"encoding/json"
	"strings"
)

// MKSError is the custom error type for mks-go API errors.
type MKSError struct {
	StatusCode int
	Message    string
}

// Status returns the HTTP status code of the error.
func (e MKSError) Status() int {
	return e.StatusCode
}

// Error returns the error message.
func (e MKSError) Error() string {
	return e.Message
}

// APIError is the interface for convertable to MKSError types.
type APIError interface {
	ToMKSError(statusCode int) *MKSError
}

// ToMKSError converts to an MKSError with the given status code.
func (e *GenericError) ToMKSError(statusCode int) *MKSError {
	if e == nil {
		return nil
	}

	return &MKSError{
		StatusCode: statusCode,
		Message:    e.Error.Message,
	}
}

// ToMKSError converts to an MKSError with the given status code.
func (e *GenericNotFoundError) ToMKSError(statusCode int) *MKSError {
	if e == nil {
		return nil
	}

	return &MKSError{
		StatusCode: statusCode,
		Message:    e.Error.Message,
	}
}

// HandleAPIErrors processes a list of possible API errors and returns the first one found.
// If no API error is found, it returns a generic MKSError using the provided status code and message.
func HandleAPIErrors(statusCode int, statusMsg string, errors ...APIError) error {
	return HandleAPIErrorsWithBody(statusCode, statusMsg, nil, errors...)
}

// HandleAPIErrorsWithBody is HandleAPIErrors that adds the API message from the response body to the status message.
func HandleAPIErrorsWithBody(statusCode int, statusMsg string, body []byte, errors ...APIError) error {
	for _, err := range errors {
		if err != nil {
			mksErr := err.ToMKSError(statusCode)
			if mksErr == nil {
				continue
			}

			return mksErr
		}
	}

	return &MKSError{
		StatusCode: statusCode,
		Message:    withBodyReason(statusMsg, body),
	}
}

// withBodyReason appends the message of a GenericError-shaped body to the status message.
func withBodyReason(statusMsg string, body []byte) string {
	var genericErr GenericError

	err := json.Unmarshal(body, &genericErr)
	if err != nil {
		return statusMsg
	}

	message := strings.TrimSpace(genericErr.Error.Message)
	if message == "" {
		return statusMsg
	}

	return statusMsg + ": " + message
}
